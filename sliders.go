package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Sliders allow users to view and select a value (or range) along a track. They're ideal for adjusting settings such as volume and brightness, or for applying image filters.

Sliders can use icons or labels to represent a numeric or relative scale.

https://github.com/material-components/material-web/blob/main/docs/components/slider.md#mdslider-md-slider
*/
func Slider() app.HTMLElem {
	return app.Elem("md-slider")
}
