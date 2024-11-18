package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
A divider is a thin line that groups content in lists and containers.

Dividers can reinforce tapability, such as when used to separate list items or define tappable regions in an accordion.

https://github.com/material-components/material-web/blob/main/docs/components/divider.md#mddivider-md-divider
*/
func Divider() app.HTMLElem {
	return app.Elem("md-divider")
}
