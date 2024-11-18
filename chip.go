package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Chips should always appear in a set. Chip sets are toolbars that can display any type of chip or other toolbar items.

https://github.com/material-components/material-web/blob/main/docs/components/chip.md#mdchipset-md-chip-set
*/
func ChipSet() app.HTMLElem {
	return app.Elem("md-chip-set")

}

/*
Assist chips represent smart or automated actions that can span multiple apps, such as opening a calendar event from the home screen.

Assist chips function as though the user asked an assistant to complete the action. They should appear dynamically and contextually in a UI.

https://github.com/material-components/material-web/blob/main/docs/components/chip.md#mdassistchip-md-assist-chip
*/
func AssistChip() app.HTMLElem {
	return app.Elem("md-assist-chip")
}

/*
Filter chips use tags or descriptive words to filter content. They can be a good alternative to toggle buttons or checkboxes.

https://github.com/material-components/material-web/blob/main/docs/components/chip.md#mdfilterchip-md-filter-chip
*/
func FilterChip() app.HTMLElem {
	return app.Elem("md-filter-chip")
}

/*
Input chips represent discrete pieces of information entered by a user, such as Gmail contacts or filter options within a search field.

Input chips whose icons are user images may add the avatar attribute to display the image in a larger circle.

https://github.com/material-components/material-web/blob/main/docs/components/chip.md#mdinputchip-md-input-chip
*/
func InputChip() app.HTMLElem {
	return app.Elem("md-input-chip")
}

/*
Suggestion chips help narrow a user’s intent by presenting dynamically generated suggestions, such as possible responses or search filters.

https://github.com/material-components/material-web/blob/main/docs/components/chip.md#mdsuggestionchip-md-suggestion-chip
*/
func SuggestionChip() app.HTMLElem {
	return app.Elem("md-suggestion-chip")
}
