package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Floating Action Button
FAB represents the most important action on a screen. It puts key actions within reach.

Extended FABs help people take primary actions. They're wider than FABs to accommodate a text label and larger target area.

https://github.com/material-components/material-web/blob/main/docs/components/fab.md#mdfab-md-floating-action-button
*/
func Fab() app.HTMLElem {
	return app.Elem("md-fab")
}

/*
Branded FABs use a brightly colored logo for their icon. Unlike FAB, branded FABs do not have color variants.

https://github.com/material-components/material-web/blob/main/docs/components/fab.md#branded-fab
*/
func BrandedFab() app.HTMLElem {
	return app.Elem("md-branded-fab")
}
