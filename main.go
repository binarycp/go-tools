package main

import (
	"github.com/gin-gonic/gin"
	"go-tools/resources"
	"html/template"
	"net/http"
)

func main() {
	engine := gin.Default()
	template := template.Must(template.New("demo").ParseFS(resources.Resource, "dist/index.html", "dist/favicon.ico"))
	engine.SetHTMLTemplate(template)

	engine.StaticFS("/public", http.FS(resources.Resource))

	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})

	engine.Run(":7752")
}
