package models

import (
	"os"
	"path/filepath"
	"self-hosted-blog/utils"
	"strings"
)

func GetFilePaths() map[string][]TPage {
	output := make(map[string][]TPage)
	root := filepath.Join("home")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		utils.Check(err)
		page := createPage()
		page.addName(filepath.Base(path))
		page.addPath(splitFilePath(filepath.Dir(path)))
		page.addIsFolder(info.IsDir())
		page.getFileData(page.GetName(), page.GetPath())
		page.addUrl(page.GetPath(), page.GetName())
		if info.IsDir() {
			if _, ok := output[page.GetName()]; !ok {
				output[page.GetName()] = []TPage{}
			}
		} else {
			lastDir := page.GetPath()[len(page.GetPath())-1]
			output[lastDir] = append(output[lastDir], page)
		}
		return err
	})
	utils.Check(err)
	return output
}

// @type TPage
// @description Represent one markdown page
type TPage struct {
	name     string
	path     []string
	isFolder bool
	data     []byte
	url      string
}

func createPage() TPage {
	return TPage{name: "", isFolder: false, path: []string{}}
}

// @method
func (inst *TPage) GetName() string {
	return inst.name
}

// @method
func (inst *TPage) GetPath() []string {
	return inst.path
}

// @method
func (inst *TPage) GetUrl() string {
	return inst.url
}

// @method
func (inst *TPage) IsDir() bool {
	return inst.isFolder
}

// @method
func (inst *TPage) GetData() []byte {
	return inst.data
}

// @method
func (inst *TPage) addName(name string) *TPage {
	inst.name = name
	return inst
}

// @method
func (inst *TPage) addPath(path []string) *TPage {
	inst.path = path
	return inst
}

// @method
func (inst *TPage) addUrl(path []string, name string) *TPage {
	inst.url = path[len(path)-1] + "/" + name
	return inst
}

// @method
func (inst *TPage) addIsFolder(is bool) *TPage {
	inst.isFolder = is
	return inst
}

// @method
func (inst *TPage) getFileData(name string, path []string) *TPage {
	inst.data = readFile(name, path)
	return inst
}

func splitFilePath(path string) []string {
	return strings.Split(filepath.Clean(path), string(filepath.Separator))
}

func readFile(name string, path []string) []byte {
	format := "." + string(filepath.Separator) + filepath.Join(path...) + string(filepath.Separator) + name
	data, err := os.ReadFile(format)
	if err != nil {
		return []byte{}
	} else {
		return utils.MdToHTML(data)
	}

}
