package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Standard icon buttons do not have a background or outline, and have the lowest emphasis of the icon buttons.

https://github.com/material-components/material-web/blob/main/docs/components/icon-button.md#mdiconbutton-md-icon-button
*/
func IconButton() app.HTMLElem {
	return app.Elem("md-icon-button")
}

/*
Filled icon buttons have higher visual impact and are best for high emphasis actions.

https://github.com/material-components/material-web/blob/main/docs/components/icon-button.md#mdfillediconbutton-md-filled-icon-button
*/
func FilledIconButton() app.HTMLElem {
	return app.Elem("md-filled-icon-button")
}

/*
Filled tonal icon buttons are a middle ground between filled and outlined icon buttons.

They're useful in contexts where the button requires slightly more emphasis than an outline would give,
such as a secondary action paired with a high emphasis action.

https://github.com/material-components/material-web/blob/main/docs/components/icon-button.md#mdfilledtonaliconbutton-md-filled-tonal-icon-button
*/
func FilledTonalIconButton() app.HTMLElem {
	return app.Elem("md-filled-tonal-icon-button")
}

/*
Outlined icon buttons are medium-emphasis icon buttons.

They contain actions that are important, but aren’t the primary action in an app.

https://github.com/material-components/material-web/blob/main/docs/components/icon-button.md#mdoutlinediconbutton-md-outlined-icon-button
*/
func OutlinedIconButton() app.HTMLElem {
	return app.Elem("md-outlined-icon-button")
}
