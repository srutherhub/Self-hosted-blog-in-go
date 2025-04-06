package main

import (
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"self-hosted-blog/constants"
	"self-hosted-blog/models"
)

// @Function main
// @Description: main entry point into application
// @Params
// @Return: None
func main() {
	http.HandleFunc("/", handler)
	createRoutesFromList(getPageList())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// @Function mdHandler
// @Description: callback fundtion that handles what should be served when hitting a specific endpoint
// @Params
//
//	w http.ResponseWriter
//	r http.Request
//
// @Return: None

func mdHandler(md []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, innerHtml(string(mdToHTML(md))+constants.HtmlStyle, "article", "class=article"))
	}
}

// @Function handler
// @Description: callback fundtion that handles what should be served when hitting a specific endpoint
// @Params
//
//	w http.ResponseWriter
//	r http.Request
//
// @Return: None

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, renderPageList(getPageList())+constants.HtmlStyle)
}

// @Function createRoutesFromList
// @Description
func createRoutesFromList(list map[string][]models.TPage) {
	for _, i := range list {
		for _, j := range i {
			http.HandleFunc(j.Path, mdHandler(readMdFile("./home"+j.Path)))
		}

	}
}

// @Function getPageList
// @Description: Returns a list of pages and folders from the src directory
// @Params
// @Return: map[string][]models.TPage
func getPageList() map[string][]models.TPage {
	var output = make(map[string][]models.TPage)
	err := filepath.Walk("./home/", func(path string, info os.FileInfo, err error) error {
		check(err)
		data := models.CreatePage()
		data.AddName(info.Name()).AddIsFolder(info.IsDir()).AddPath(getFilePathLink(path))
		var a, _ = parsePathAndFile(path, data.IsFolder)
		if data.IsFolder {
			if _, ok := output[path]; !ok {
				output[a] = []models.TPage{}
			}
		}
		if !data.IsFolder {
			output[a] = append(output[a], data)
		}
		return err
	})
	check(err)
	return output
}

// @Function parsePathAndFile
// @Description: take a path to extract and format the path string and file string
func parsePathAndFile(path string, isDir bool) (string, string) {
	var split = strings.Split(path, "/")
	pathOutput := ""
	fileOutput := ""
	if split[0] == "." {
		pathOutput += "/" + split[1]
	} else {
		if isDir {
			for i := 0; i < len(split); i++ {
				pathOutput += "/" + split[i]
			}
		} else {
			for i := 0; i < len(split)-1; i++ {
				pathOutput += "/" + split[i]
			}
			fileOutput += split[len(split)-1]
		}
	}

	return pathOutput, fileOutput
}

// @Function renderPageList
// @Description: Renders the page list in an html list
// @Params
//
//	list []TPage
//
// @Return string
func renderPageList(list map[string][]models.TPage) string {
	output := ""
	for in, i := range list {
		output += innerHtml(in, "h3", "")
		for _, j := range i {
			var linkPath string = "href='" + j.Path + "'"
			output += innerHtml(getHtmlSpaces(5)+j.Name, "a", linkPath)
		}
	}
	return innerHtml(output,"div","class=links-container")
}

func getHtmlSpaces(num int) string {
	output:=""
	i:=0
	for i < num {
		output+="&nbsp;"
		i++
	}
	return output
}

// @Function innerHtml
// @Description: Takes in a value and wraps it around an html tag
// @Params
//
//	val string
//	tag string
//
// @Return: string
func innerHtml(val, tag, props string) string {
	var open = "<" + tag + " " + props + ">"
	var close = "</" + tag + ">"
	return open + val + close
}

// @Function getFilePath
// @Description: This function takes a path as a string and returns the link where the file will be display
// @Params
//
//	path string
//
// @ Return
//
//	string
func getFilePathLink(path string) string {
	return path[4:]
}

// @Function readMdFile
// Description: Takes in a file path and returns the byte value of the file
// @Params
//
//	file string
//
// @Return []byte
func readMdFile(file string) []byte {
	data, err := os.ReadFile(file)
	check(err)
	return data
}

// @Function mdToHTML
// @Description: Function copied from above git repo, convert markdown to html
func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

// @Function check
// Description: Takes in an error and throws it
// Inputs:
//
//	e error
//
// Return: None
func check(e error) {
	if e != nil {
		panic(e)
	}
}
