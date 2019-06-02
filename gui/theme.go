package gui

import (
	"fyne.io/fyne"
	"image/color"
)

type MaterialTheme struct {
	regular, bold, italic, bolditalic, monospace fyne.Resource
}

func (MaterialTheme) BackgroundColor() color.Color {
	// #303030
	return color.RGBA{R: 0x30, G: 0x30, B: 0x30, A: 0xFF}
}

func (MaterialTheme) ButtonColor() color.Color {
	// #2196F3
	return color.RGBA{R: 0x21, G: 0x96, B: 0xF3, A: 0xFF}
}

func (MaterialTheme) HyperlinkColor() color.Color {
	// #2196F3
	return color.RGBA{R: 0x21, G: 0x96, B: 0xF3, A: 0xFF}
}

func (t *MaterialTheme) TextColor() color.Color {
	return color.White
}

func (MaterialTheme) PlaceHolderColor() color.Color {
	panic("implement me")
}

func (MaterialTheme) PrimaryColor() color.Color {
	// #2196F3
	return color.RGBA{R: 0x21, G: 0x96, B: 0xF3, A: 0xFF}
}

func (MaterialTheme) FocusColor() color.Color {
	// #FF5722
	return color.RGBA{R: 0xFF, G: 0x57, B: 0x22, A: 0xFF}
}

func (MaterialTheme) ScrollBarColor() color.Color {
	panic("implement me")
}

func (MaterialTheme) TextSize() int {
	return 14
}

func (t *MaterialTheme) TextFont() fyne.Resource {
	return nil
}

func (MaterialTheme) TextBoldFont() fyne.Resource {
	return nil
}

func (MaterialTheme) TextItalicFont() fyne.Resource {
	return nil
}

func (MaterialTheme) TextBoldItalicFont() fyne.Resource {
	panic("implement me")
}

func (MaterialTheme) TextMonospaceFont() fyne.Resource {
	panic("implement me")
}

func (MaterialTheme) Padding() int {
	return 4
}

func (MaterialTheme) IconInlineSize() int {
	panic("implement me")
}

func (MaterialTheme) ScrollBarSize() int {
	panic("implement me")
}
