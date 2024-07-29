package specs

import (
	"embed"
	"io/fs"
	"main/router"
	"net/http"
)

//go:embed swaggerui
var embeddedSwaggerui embed.FS

func Init(mainRouter *router.MainRouter) {
	// NOTE: Makes content of swaggerui folder the root (the swaggerui folder is embedded)
	subFs, _ := fs.Sub(embeddedSwaggerui, "swaggerui")
	fsHandler := http.StripPrefix(
		"/specs",
		http.FileServer(http.FS(subFs)),
	)

	mainRouter.Handle("/specs/*", fsHandler)
}
