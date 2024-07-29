package swagger

import (
	"embed"
	"io/fs"
	"main/router"
	"net/http"
)

//go:embed ui
var embeddedUi embed.FS

func Init(mainRouter *router.MainRouter) {
	sub, _ := fs.Sub(embeddedUi, "ui")
	fsHandler := http.FileServer(http.FS(sub))
	
	mainRouter.Handle("/doc", http.StripPrefix("/doc/", fsHandler))
}