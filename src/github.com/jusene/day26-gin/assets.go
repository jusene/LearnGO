package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsbfa8d115ce0617d89507412d5393a462f8e9b003 = "<!doctype html>\r\n<body>\r\n  <p>Can you see this? → {{.Bar}}</p>\r\n</body>"
var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\r\n<body>\r\n  <p>Hello, {{.Foo}}</p>\r\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{".\\html"}, "/html": []string{"bar.tmpl", "index.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1592791609, 1592791609649929800),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1592791669, 1592791669877880000),
		Data:     nil,
	}, "/html/bar.tmpl": &assets.File{
		Path:     "/html/bar.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1592791669, 1592791669876892100),
		Data:     []byte(_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003),
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1592791641, 1592791641333471000),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}}, "")
