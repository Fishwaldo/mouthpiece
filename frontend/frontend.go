package frontend

import (
	"embed"
)

//go:generate npm run build

//go:embed dist
var FrontEndFiles embed.FS
