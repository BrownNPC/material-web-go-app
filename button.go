package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Elevated buttons are essentially filled tonal buttons with a shadow.

To prevent shadow creep, only use them when absolutely necessary, such as when the button requires visual separation from a patterned background.

https://github.com/material-components/material-web/blob/main/docs/components/button.md#mdelevatedbutton-md-elevated-button
*/
func ElevatedButton() app.HTMLElem {
	return app.Elem("md-elevated-button")
}

/*
Filled buttons have the most visual impact after the FAB, and should be used for important, final actions that complete a flow, like Save, Join now, or Confirm.

https://github.com/material-components/material-web/blob/main/docs/components/button.md#mdfilledbutton-md-filled-button
*/
func FilledButton() app.HTMLElem {
	return app.Elem("md-filled-button")
}

/*
A filled tonal button is an alternative middle ground between filled and outlined buttons.

They're useful in contexts where a lower-priority button requires slightly more emphasis than an outline would give, such as "Next" in an onboarding flow.

https://github.com/material-components/material-web/blob/main/docs/components/button.md#mdfilledtonalbutton-md-filled-tonal-button
*/
func FilledTonalButton() app.HTMLElem {
	return app.Elem("md-filled-tonal-button")
}

/*
Outlined buttons are medium-emphasis buttons. They contain actions that are important, but aren’t the primary action in an app.

https://github.com/material-components/material-web/blob/main/docs/components/button.md#mdoutlinedbutton-md-outlined-button
*/
func OutlinedButton() app.HTMLElem {
	return app.Elem("md-outlined-button")
}

/*
Text buttons are used for the lowest priority actions, especially when presenting multiple options.

https://github.com/material-components/material-web/blob/main/docs/components/button.md#mdtextbutton-md-text-button
*/
func TextButton() app.HTMLElem {
	return app.Elem("md-text-button")
}
