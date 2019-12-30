package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsbfa8d115ce0617d89507412d5393a462f8e9b003 = "<!doctype html>\r\n<body>\r\n  <p>Can you see this? → {{.Bar}}</p>\r\n  <p>Can you see this 2? → {{.Bar | hAge}}</p>\r\n</body>\r\n"
var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\r\n<body>\r\n  <p>Hello, {{.Foo}}</p>\r\n</body>\r\n"
var _Assets2563e667e12083ecf0c5292d70a6c4de3840819c = "{{define \"goods/good.html\"}}\r\n<html>\r\n<body>\r\n<h1>Goods</h1>\r\n<span>{{.good  }}</span>\r\n</body>\r\n</html>\r\n{{end}}"
var _Assets76b8f1208fed471ae72580184d170a3791bbeb1f = "{{define \"users/user.html\"}}\r\n<html>\r\n<body>\r\n<h1>Goods</h1>\r\n<span>{{.age | hAge}}</span>\r\n</body>\r\n</html>\r\n{{end}}"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html", "views"}, "/html": []string{"bar.tmpl", "index.tmpl"}, "/views": []string{}, "/views/goods": []string{"good.html"}, "/views/users": []string{"user.html"}}, map[string]*assets.File{
	"/views/users/user.html": &assets.File{
		Path:     "/views/users/user.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1577697215, 1577697215667814100),
		Data:     []byte(_Assets76b8f1208fed471ae72580184d170a3791bbeb1f),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1577697117, 1577697117399451100),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1577696993, 1577696993916432800),
		Data:     nil,
	}, "/views/goods": &assets.File{
		Path:     "/views/goods",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1577694270, 1577694270846547100),
		Data:     nil,
	}, "/views/goods/good.html": &assets.File{
		Path:     "/views/goods/good.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1577694270, 1577694270841667400),
		Data:     []byte(_Assets2563e667e12083ecf0c5292d70a6c4de3840819c),
	}, "/views/users": &assets.File{
		Path:     "/views/users",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1577697215, 1577697215671693300),
		Data:     nil,
	}, "/html/bar.tmpl": &assets.File{
		Path:     "/html/bar.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1577696993, 1577696993916432800),
		Data:     []byte(_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003),
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1574233719, 1574233719973302100),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}, "/views": &assets.File{
		Path:     "/views",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1577692575, 1577692575079824100),
		Data:     nil,
	}}, "")
