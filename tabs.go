package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Tabs organize groups of related content that are at the same level of hierarchy.

Tabs contain multiple primary or secondary tab children. Use the same type of tab in a tab bar.

https://github.com/material-components/material-web/blob/main/docs/components/tabs.md#mdtabs-md-tabs
*/
func Tabs() app.HTMLElem {
	return app.Elem("md-tabs")
}

/*
Primary tabs are placed at the top of the content pane under a top app bar.

They display the main content destinations.

https://github.com/material-components/material-web/blob/main/docs/components/tabs.md#mdprimarytab-md-primary-tab
*/
func PrimaryTab() app.HTMLElem {
	return app.Elem("md-primary-tab")
}

/*
Secondary tabs are used within a content area to further separate related content and establish hierarchy.

https://github.com/material-components/material-web/blob/main/docs/components/tabs.md#mdsecondarytab-md-secondary-tab
*/
func SecondaryTab() app.HTMLElem {
	return app.Elem("md-secondary-tab")

}
