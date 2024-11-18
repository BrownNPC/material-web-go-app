package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Icons can be specified by name, unicode code point, or have an <svg> child element.

You must load the appropriate icon font to use icons. `LinkIconFontRounded()`, `LinkIconFontSharp()` and `LinkIconFontOutlined()` are available.

https://github.com/material-components/material-web/blob/main/docs/components/icon.md#mdicon-md-icon
*/
func Icon() app.HTMLElem {
	return app.Elem("md-icon")
}

/*
This icon font is the default one used in Material Design 3.

use the append() function to add this to your RawHeaders app property

	func main() {
		handler := &app.Handler{
			Name:       "Material Design App",
			RawHeaders: append(md.RawHeaders(), md.LinkIconFontOutlined()),
		}
*/
func LinkIconFontRounded() string {
	return `<link href="https://fonts.googleapis.com/icon?family=Material+Symbols+Rounded" rel="stylesheet">`
}

/*
This icon font is similar to the rounded one, but with sharp corners.

use the append() function to add this to your RawHeaders app property

	func main() {
		handler := &app.Handler{
			Name:       "Material Design App",
			RawHeaders: append(md.RawHeaders(), md.LinkIconFontOutlined()),
		}
*/
func LinkIconFontSharp() string {
	return `<link href="https://fonts.googleapis.com/icon?family=Material+Symbols+Sharp" rel="stylesheet">`
}

/*
This icon font features outlined symbols, providing a distinct style.

use the append() function to add this to your RawHeaders app property

	func main() {
		handler := &app.Handler{
			Name:       "Material Design App",
			RawHeaders: append(md.RawHeaders(), md.LinkIconFontOutlined()),
		}
*/
func LinkIconFontOutlined() string {
	return `<link href="https://fonts.googleapis.com/icon?family=Material+Symbols+Outlined" rel="stylesheet">`
}
