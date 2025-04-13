package main

import (
	"html/template"
	"self-hosted-blog/models"

	"github.com/gin-gonic/gin"
)

// @global variables
var GFiles = models.GetFilePaths()

func main() {
	//load assets
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	//Routes
	index := pageData{Nav: GFiles, Article: template.HTML([]byte{})}
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", index)
	})
	createPageRoutes(router, GFiles)

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/")
	})
	router.Run()
}

type pageData struct {
	Nav     []map[string][]models.TPage
	Article template.HTML
}

func createPageRoutes(router *gin.Engine, pagelist []map[string][]models.TPage) {
	for _, folder := range pagelist {
		for dir, pages := range folder {
			for _, page := range pages {
				sendData := pageData{Nav: GFiles, Article: template.HTML(page.GetData())}
				if !page.IsDir() {
					if page.GetNameNoExt() != dir {
						router.GET(page.GetUrl(), func(c *gin.Context) {
							c.HTML(200, "index.html", sendData)
						})
					} else {
						router.GET(dir, func(c *gin.Context) {
							c.HTML(200, "index.html", sendData)
						})
					}
				}
			}
		}
	}
}
