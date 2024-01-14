package assets

import (
	_ "embed"
)

//go:embed html.base64
var HtmlBase64 string

//go:embed jpeg.base64
var JpegBase64 string
