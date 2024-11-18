package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Text fields let users enter text into a UI.

Filled and outlined text fields are functionally identical.

https://github.com/material-components/material-web/blob/main/docs/components/text-field.md#mdoutlinedtextfield-md-outlined-text-field
*/
func OutlinedTextField() app.HTMLElem {
	return app.Elem("md-outlined-text-field")
}

/*
Text fields let users enter text into a UI.

Filled and outlined text fields are functionally identical.

https://github.com/material-components/material-web/blob/main/docs/components/text-field.md#mdfilledtextfield-md-filled-text-field
*/
func FilledTextField() app.HTMLElem {
	return app.Elem("md-filled-text-field")
}
