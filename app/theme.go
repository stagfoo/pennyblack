package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct {
	font fyne.Resource
	size fyne.Size
}

func (t *CustomTheme) Font(style fyne.TextStyle) fyne.Resource {
	return t.font
}

func (t *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (t *CustomTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *CustomTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 24.0 // Increase from default (usually ~14)
	case theme.SizeNameHeadingText:
		return 26.0 // Increase heading size
	case theme.SizeNameSubHeadingText:
		return 22.0 // Increase subheading size
	case theme.SizeNameCaptionText:
		return 18.0 // Increase caption size
	default:
		return theme.DefaultTheme().Size(name)
	}
}
