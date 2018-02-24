package serv

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

// setHTMLTemplate 设置网页模板
func setHTMLTemplate(router *gin.Engine) {
	templ := template.New("")
	// to do
	router.SetHTMLTemplate(templ)
}
