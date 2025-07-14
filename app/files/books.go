package files

import (
	"fmt"
	"log"

	"github.com/taylorskalyo/goreader/epub"
)

func ReadEBUB(epubPath string) {
	// Open the EPUB file
	book, err := epub.OpenReader(epubPath)
	if err != nil {
		log.Fatal(err)
	}
	defer book.Close()

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
	fmt.Printf("Creator: %s\n", pkg.Metadata.Creator)
	fmt.Printf("Language: %s\n", pkg.Metadata.Language)
	fmt.Printf("Description: %s\n", pkg.Metadata.Description)
	fmt.Printf("Publisher: %s\n", pkg.Metadata.Publisher)

	// Access manifest items
	fmt.Printf("\nManifest items (%d):\n", len(pkg.Manifest.Items))
	for _, item := range pkg.Manifest.Items {
		fmt.Printf("- ID: %s, HREF: %s, MediaType: %s\n",
			item.ID, item.HREF, item.MediaType)
	}

	// Access spine (reading order)
	fmt.Printf("\nSpine items (%d):\n", len(pkg.Spine.Itemrefs))
	for _, itemref := range pkg.Spine.Itemrefs {
		fmt.Printf("- IDREF: %s\n", itemref.IDREF)
	}
}
