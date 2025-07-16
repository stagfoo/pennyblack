package files

import (
	"fmt"
	"io"
	"log"
	"strings"

	"fyne.io/fyne/v2/widget"
	"github.com/taylorskalyo/goreader/epub"
	"golang.org/x/net/html"
)

func ReadEPUB(epubPath string) ([]epub.Itemref, *epub.ReadCloser) {
	// Open the EPUB file
	book, err := epub.OpenReader(epubPath)
	if err != nil {
		log.Fatal(err)
	}
	// defer book.Close()

	// The Reader embeds Container, so you can access Container fields directly
	fmt.Printf("Found %d rootfiles\n", len(book.Rootfiles))

	// Access the first rootfile (there's usually only one)
	if len(book.Rootfiles) == 0 {
		log.Fatal("No rootfiles found")
	}

	rootfile := book.Rootfiles[0]
	fmt.Printf("Rootfile path: %s\n", rootfile.FullPath)
	fmt.Printf("Media type: %s\n", rootfile.Identifier)

	// The Package field contains the parsed OPF data
	pkg := book.Rootfiles[0].Package

	// Access metadata
	fmt.Printf("Title: %s\n", pkg.Metadata.Title)
	fmt.Printf("Coverage: %s\n", pkg.Metadata.Coverage)
	fmt.Printf("Creator: %s\n", pkg.Metadata.Creator)
	fmt.Printf("Language: %s\n", pkg.Metadata.Language)
	fmt.Printf("Description: %s\n", pkg.Metadata.Description)
	fmt.Printf("Publisher: %s\n", pkg.Metadata.Publisher)

	// Access manifest items
	// fmt.Printf("\nManifest items (%d):\n", len(pkg.Manifest.Items))
	// for _, item := range pkg.Manifest.Items {
	// 	fmt.Printf("- ID: %s, HREF: %s, MediaType: %s\n",
	// 		item.ID, item.HREF, item.MediaType)
	// }

	// Access spine (reading order)
	// fmt.Printf("\nSpine items (%d):\n", len(pkg.Spine.Itemrefs))
	// for _, itemref := range pkg.Spine.Itemrefs {
	// 	fmt.Printf("- IDREF: %s\n", itemref.IDREF)
	// }
	return pkg.Spine.Itemrefs, book
}
func ReadItem(item epub.Item) string {
	// Open the content of the first item
	reader, err := item.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// Read and print the content
	content, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	return string(content)
}

// xhtmlToRichText parses a simple XHTML string and converts it to RichText segments.
// Note: This is a simplified parser for demonstration.
func XhtmlToRichText(xhtml string) (*widget.RichText, error) {
	tokenizer := html.NewTokenizer(strings.NewReader(xhtml))
	var segments []widget.RichTextSegment

	// Basic state tracking for styles
	var isBold, isItalic bool

	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				// End of the document, we're done
				return widget.NewRichText(segments...), nil
			}
			// Some other error
			return nil, err

		case html.TextToken:
			text := string(tokenizer.Text())
			segment := &widget.TextSegment{Text: text}

			// Apply current style
			if isBold {
				segment.Style = widget.RichTextStyleStrong
			} else if isItalic {
				segment.Style = widget.RichTextStyleEmphasis
			}
			segments = append(segments, segment)

		case html.StartTagToken, html.EndTagToken:
			tn, _ := tokenizer.TagName()
			tagName := string(tn)

			var styleState bool
			if tt == html.StartTagToken {
				styleState = true // Entering a tag
			} else {
				styleState = false // Exiting a tag
			}

			switch tagName {
			case "strong", "b":
				isBold = styleState
			case "em", "i":
				isItalic = styleState
			case "p", "h1", "h2", "h3", "br":
				// Add a newline for block elements or line breaks
				if tt == html.EndTagToken || tagName == "br" {
					segments = append(segments, &widget.TextSegment{Text: "\n", Style: widget.RichTextStyleParagraph})
				}
			}
		}
	}
}
