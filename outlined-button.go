package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type outlinedButton struct {
	app.Compo
	//Whether or not the button is disabled.

	disabled bool
	// 	Whether or not the button is "soft-disabled" (disabled but still focusable).
	// Use this when a button needs increased visibility when disabled. See https://www.w3.org/WAI/ARIA/apg/practices/keyboard-interface/#kbd_disabled_controls for more guidance on when this is needed.
	softDisabled bool
	//The URL that the link button points to.
	href string
	//Where to display the linked href URL for a link button. Common options include _blank to open in a new tab.
	target string
	// Whether to render the icon at the inline end of the label rather than the inline start.
	// Note: Link buttons cannot have trailing icons.
	trailingIcon bool
	// The default behavior of the button. May be "button", "reset", or "submit" (default).
	type_ string
	// The value added to a form with the button's name when the button submits a form.
	value string
	name  string
	// what is inside the button
	body   app.UI
	events map[string]app.EventHandler
}

func OutlinedButton() *outlinedButton {
	return &outlinedButton{}
}

func (o *outlinedButton) Disabled(v bool) *outlinedButton {
	o.disabled = v
	return o
}

func (o *outlinedButton) SoftDisabled(v bool) *outlinedButton {
	o.softDisabled = v
	return o
}

func (o *outlinedButton) Href(v string) *outlinedButton {
	o.href = v
	return o
}

func (o *outlinedButton) Target(v string) *outlinedButton {
	o.target = v
	return o
}

func (o *outlinedButton) TrailingIcon(v bool) *outlinedButton {
	o.trailingIcon = v
	return o
}

func (o *outlinedButton) Type(v string) *outlinedButton {
	o.type_ = v
	return o
}

func (o *outlinedButton) Value(v string) *outlinedButton {
	o.value = v
	return o
}

func (o *outlinedButton) Name(v string) *outlinedButton {
	o.name = v
	return o
}

func (o *outlinedButton) Body(v app.UI) *outlinedButton {
	o.body = v
	return o
}

func (o *outlinedButton) On(name string, handler app.EventHandler) *outlinedButton {
	if o.events == nil {
		o.events = make(map[string]app.EventHandler)
	}
	o.events[name] = handler
	return o
}

func (o *outlinedButton) Render() app.UI {
	button := app.Elem("md-outlined-button")

	// Apply boolean attributes
	if o.disabled {
		button.Attr("disabled", "true")
	}
	if o.softDisabled {
		button.Attr("soft-disabled", "true")
	}
	if o.trailingIcon {
		button.Attr("trailing-icon", "true")
	}

	// Apply string attributes
	attrs := map[string]string{
		"type":   o.type_,
		"value":  o.value,
		"name":   o.name,
		"href":   o.href,
		"target": o.target,
	}
	for key, value := range attrs {
		if value != "" {
			button.Attr(key, value)
		}
	}

	// Attach events
	for name, handler := range o.events {
		button.On(name, handler)
	}

	// Set the body
	if o.body != nil {
		button.Body(o.body)
	}

	return button
}
