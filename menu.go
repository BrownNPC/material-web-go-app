package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
When opened, menus position themselves to an anchor.

Thus, either anchor or anchorElement must be supplied to md-menu before opening.

Additionally, a shared parent of position:relative should be present around the menu and it's anchor, 	because the default menu is positioned relative to the anchor element.

https://github.com/material-components/material-web/blob/main/docs/components/menu.md#mdmenu-md-menu
*/
func Menu() app.HTMLElem {
	return app.Elem("md-menu")
}

/*
Menus also render menu items such as md-menu-item and handle keyboard navigation between md-menu-items as well as typeahead functionality. Additionally, md-menu interacts with md-menu-items to help you determine how a menu was closed.

Listen for and inspect the close-menu custom event's details to determine what action and items closed the menu.

https://github.com/material-components/material-web/blob/main/docs/components/menu.md#mdmenuitem-md-menu-item
*/
func MenuItem() app.HTMLElem {
	return app.Elem("md-menu-item")
}

/*
You can compose <md-menu>s inside of an <md-sub-menu>'s menu slot,

but first the has-overflow attribute must be set on the root <md-menu> to disable overflow scrolling and display the nested submenus.

https://github.com/material-components/material-web/blob/main/docs/components/menu.md#mdsubmenu-md-sub-menu
*/
func SubMenu() app.HTMLElem {
	return app.Elem("md-sub-menu")
}
