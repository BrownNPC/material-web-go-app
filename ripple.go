package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Ripples are state layers used to communicate the status of a component or interactive element.

A state layer is a semi-transparent covering on an element that indicates its state.
A layer can be applied to an entire element or in a circular shape.

https://github.com/material-components/material-web/blob/main/docs/components/ripple.md#mdripple-md-ripple
*/
func Ripple() app.HTMLElem {
	return app.Elem("md-ripple")
}
