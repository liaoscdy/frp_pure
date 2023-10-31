package frpc

import (
	"embed"

	"frp_pure/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
