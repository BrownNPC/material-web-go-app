package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Elevation is the relative distance between two surfaces along the z-axis.

Material's elevation system is deliberately limited to just a handful of levels. This creative constraint means you need to make thoughtful decisions about your UI’s elevation story.

https://github.com/material-components/material-web/blob/main/docs/components/elevation.md#mdelevation-md-elevation
*/

func Elevation() app.HTMLElem {
	return app.Elem("md-elevation")
}
