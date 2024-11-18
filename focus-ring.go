package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Focus rings are accessible outlines for components to show keyboard focus.

Focus rings follow the same heuristics as :focus-visible to determine when they are visible.

https://github.com/material-components/material-web/blob/main/docs/components/focus-ring.md#mdfocusring-md-focus-ring
*/
func FocusRing() app.HTMLElem {
	return app.Elem("md-focus-ring")
}
