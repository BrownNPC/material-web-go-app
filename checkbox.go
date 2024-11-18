package md

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

/*
Checkboxes allow users to select one or more items from a set. Checkboxes can turn an option on or off.

There's one type of checkbox in Material. Use this selection control when the user needs to select one or more options from a list.

https://github.com/material-components/material-web/blob/main/docs/components/checkbox.md#mdcheckbox-md-checkbox
*/
func Checkbox() app.HTMLElem {
	return app.Elem("md-checkbox").Attr("touch-target", "wrapper")
}
