package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Switches toggle the state of an item on or off.

Switches are similar to checkboxes, and can be unselected or selected.

https://github.com/material-components/material-web/blob/main/docs/components/switch.md#mdswitch-md-switch
*/
func Switch() app.HTMLElem {
	return app.Elem("md-switch")
}
