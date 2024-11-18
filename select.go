package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Select menus display a list of choices on temporary surfaces and display the currently selected menu item above the menu.

Select (also referred to as a dropdown menu) allows choosing a value from a fixed list of available options.

https://github.com/material-components/material-web/blob/main/docs/components/select.md#mdfilledselect-md-filled-select
*/
func FilledSelect() app.HTMLElem {
	return app.Elem("md-filled-select")
}

/*
Select menus display a list of choices on temporary surfaces and display the currently selected menu item above the menu.

Select (also referred to as a dropdown menu) allows choosing a value from a fixed list of available options.

https://github.com/material-components/material-web/blob/main/docs/components/select.md#mdoutlinedselect-md-outlined-select
*/
func OutlinedSelect() app.HTMLElem {
	return app.Elem("md-outlined-select")
}

/*
Select menus display a list of choices on temporary surfaces and display the currently selected menu item above the menu.

Select (also referred to as a dropdown menu) allows choosing a value from a fixed list of available options.

https://github.com/material-components/material-web/blob/main/docs/components/select.md#mdselectoption-md-select-option
*/
func SelectOption() app.HTMLElem {
	return app.Elem("md-select-option")
}
