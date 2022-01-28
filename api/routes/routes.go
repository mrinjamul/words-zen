package routes

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrinjamul/words-zen/api/handler"
)

// ViewsFs for static files
var ViewsFs embed.FS

func InitRoutes(routes *gin.Engine) {
	// Serve the frontend
	// This will ensure that the angular files are served correctly
	fsRoot, err := fs.Sub(ViewsFs, "views")
	if err != nil {
		log.Println(err)
	}
	routes.NoRoute(gin.WrapH(http.FileServer(http.FS(fsRoot))))
	// routes.StaticFS("/", http.FS(fsRoot))

	// routes.NoRoute(func(ctx *gin.Context) {
	// 	dir, file := path.Split(ctx.Request.RequestURI)
	// 	ext := filepath.Ext(file)
	// 	if file == "" || ext == "" {
	// 		ctx.File("./views" + "/index.html")
	// 	} else {
	// 		ctx.File("./views" + "/" + path.Join(dir, file))
	// 	}
	// })

	// health check
	routes.GET("/api/health", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"health": "ok",
			},
		)
	})

	// Backend API
	routes.GET("/api/:phrase", handler.GetHandler)
	routes.POST("/api", handler.PostHandler)
}
