// Package ui (this file) contains the source code for both the Response Console and the Response Admin apps. In addition,
// these apps are embedded in separate embed.FS instances which are served via the static file server in-binary.
package ui

import "embed"

//go:embed dist/*
var UIAssetsFS embed.FS
