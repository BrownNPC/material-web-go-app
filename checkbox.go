package md

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// A material design checkbox
// https://github.com/material-components/material-web/blob/main/docs/components/checkbox.md
type checkbox struct {
	app.Compo
	checked       bool
	indeterminate bool //Whether or not the checkbox is indeterminate. apparently should not be called from html tag https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox#indeterminate_state_checkboxes
	required      bool
	value         string
	disabled      bool
	name          string
	class         string
	onChange      app.EventHandler
}

// A material design checkbox
// https://github.com/material-components/material-web/blob/main/docs/components/checkbox.md
func Checkbox() *checkbox {
	return &checkbox{}
}

// Whether or not the checkbox is selected.
func (c *checkbox) Checked(v bool) *checkbox {
	c.checked = v
	return c
}

// Whether or not the checkbox is indeterminate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox#indeterminate_state_checkboxes
// func (c *checkbox) Indeterminate(v bool) *checkbox {
// 	c.indeterminate = v
// 	return c
// }

// When true, require the checkbox to be selected when participating in form submission.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox#validation
func (c *checkbox) Required(v bool) *checkbox {
	c.required = v
	return c
}

func (c *checkbox) Disabled(v bool) *checkbox {
	c.disabled = v
	return c
}

func (c *checkbox) Name(v string) *checkbox {
	c.name = v
	return c
}

// The value of the checkbox that is submitted with a form when selected.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input/checkbox#value
func (c *checkbox) Value(v string) *checkbox {
	c.value = v
	return c
}

// The native change event on <input>https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/change_event
func (c *checkbox) OnChange(handler app.EventHandler) *checkbox {
	c.onChange = handler
	return c
}

func (c *checkbox) Indeterminate(v bool) *checkbox {
	c.indeterminate = v
	return c
}

// Class sets the class of the checkbox element.
func (c *checkbox) Class(v string) *checkbox {
	c.class = v
	return c
}

func (c *checkbox) Render() app.UI {
	e := app.Elem("md-checkbox")
	if c.checked {
		e.Attr("checked", "true")
	}
	if c.indeterminate {
		e.Attr("indeterminate", "true")
	}
	if c.required {
		e.Attr("required", "true")
	}
	if c.disabled {
		e.Attr("disabled", "true")
	}
	if c.name != "" {
		e.Attr("name", c.name)
	}
	if c.value != "" {
		e.Attr("value", c.value)
	}
	if c.onChange != nil {
		e.OnChange(c.onChange)
	}
	if c.class != "" {
		e.Class(c.class)
	}
	return e

}
