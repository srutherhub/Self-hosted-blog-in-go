package models

// @Type TPage
// Represent one markdown page
type TPage struct {
	Name     string
	Path     string
	IsFolder bool
}

func CreatePage() TPage {
	return TPage{Name: "", IsFolder: false, Path: ""}
}

func (inst *TPage) AddName(name string) *TPage {
	inst.Name = name
	return inst
}

func (inst *TPage) AddPath(path string) *TPage {
	inst.Path = path
	return inst
}

func (inst *TPage) AddIsFolder(is bool) *TPage {
	inst.IsFolder = is
	return inst
}

