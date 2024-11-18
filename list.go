package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Lists are continuous, vertical indexes of text and images
<md-list> is a container composed of <md-list-item>s of different types.

https://github.com/material-components/material-web/blob/main/docs/components/list.md
*/
func List() app.HTMLElem {
	return app.Elem("md-list")
}

/*
Lists are continuous, vertical indexes of text and images
<md-list> is a container composed of <md-list-item>s of different types.

https://github.com/material-components/material-web/blob/main/docs/components/list.md
*/
func ListItem() app.HTMLElem {
	return app.Elem("md-list-item")
}
