package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Dialogs provide important prompts in a user flow.

https://github.com/material-components/material-web/blob/main/docs/components/dialog.md#mddialog-md-dialog
*/
func Dialog() app.HTMLElem {
	return app.Elem("md-dialog")
}
