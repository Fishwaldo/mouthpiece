package frontend

import (
	"embed"
	_ "embed"
)

//go:generate npm run build

//go:embed dist/*
var FrontEndFiles embed.FS