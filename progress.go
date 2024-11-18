package md

import "github.com/maxence-charriere/go-app/v10/pkg/app"

/*
Circular progress indicators display progress by animating along an invisible circular track in a clockwise direction.

https://github.com/material-components/material-web/blob/main/docs/components/progress.md#mdcircularprogress-md-circular-progress
*/
func CircularProgress() app.HTMLElem {
	return app.Elem("md-circular-progress")
}

/*
Linear progress indicators display progress by animating along the length of a fixed, visible track.

https://github.com/material-components/material-web/blob/main/docs/components/progress.md#mdlinearprogress-md-linear-progress
*/
func LinearProgress() app.HTMLElem {
	return app.Elem("md-linear-progress")
}
