package main

import (
	"html/template"
	"self-hosted-blog/models"

	"github.com/gin-gonic/gin"
)

// @global variables
var GFiles = sortPageList(models.GetFilePaths())

// @description: main entry point of application
// @return: void
func main() {
	//load assets
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	index := pageData{Nav: GFiles, Article: template.HTML([]byte{})}

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", index)
	})
	createPageRoutes(router, GFiles)

	router.Run()
}

func createPageRoutes(router *gin.Engine, pagelist []map[string][]models.TPage) {
	for _, folder := range pagelist {
		for _, pages := range folder {
			for _, page := range pages {
				sendData := pageData{Nav: GFiles, Article: template.HTML(page.GetData())}
				if !page.IsDir() {
					router.GET(page.GetUrl(), func(c *gin.Context) {
						c.HTML(200, "index.html", sendData)
					})
				}
			}
		}
	}
}

func sortPageList(pagelist map[string][]models.TPage) []map[string][]models.TPage {
	output := make([]map[string][]models.TPage, 0)

	for key, item := range pagelist {
		elem := map[string][]models.TPage{key: item}
		if key == "home" {
			c := make([]map[string][]models.TPage, 0)
			c = append(c, elem)
			output = append(output, c...)
		} else {
			output = append(output, elem)
		}
	}
	return output
}

type pageData struct {
	Nav     []map[string][]models.TPage
	Article template.HTML
}
