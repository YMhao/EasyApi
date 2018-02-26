package web

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// SetHTMLTemplate 设置网页模板
func SetHTMLTemplate(router *gin.Engine) {
	templ := template.New("")
	template.Must(templ.New("Index").Parse(IndexPage))
	router.SetHTMLTemplate(templ)
}

func SetHtmlRouter(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "main", nil)
	})
}
