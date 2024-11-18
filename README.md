# GO bindings for [Material-Web](https://github.com/material-components/material-web) to be used with [GO-APP](https://go-app.dev/)

example
```go
package main

import (
	"log"
	"net/http"

	md "github.com/BrownNPC/material-web-go-app"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Example struct {
	app.Compo
}

func (h *Example) Render() app.UI {
	return app.Div().
		//center using grid layout
		Body(

			md.LinearProgress().
				Attr("indeterminate", true),
			md.Icon().Text("settings"),
			app.H1().Text("Hello World").Class("md-typescale-display-large"),
			md.FilledTextField().Attr("value", "Hello World"),
			md.FilledIconButton().
				Body(md.Icon().
					Text("settings"),
				),
			md.Slider(),
			md.ChipSet().
				Body(
					md.AssistChip().Text("assist chip"),
					md.FilterChip().Text("filter chip"),
				),
		)
}
func main() {
	handler := &app.Handler{
		Name:       "Material Design App",
		RawHeaders: append(md.RawHeaders(), md.LinkIconFontOutlined()),
	}
	app.Route("/", func() app.Composer { return &Example{} })
	app.RunWhenOnBrowser()
	log.Fatal(http.ListenAndServe(":8000", handler))
}
```
To run this
```bash
go mod init app
go mod tidy
GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o  web/app.wasm
go run .
```