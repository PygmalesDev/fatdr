package app

type FatDir struct {
	name       string
	path       string
	fullFath   string
	size       float32
	sizePrefix string
}

func newFatDir(path, name string) *FatDir {
	return &FatDir{
		name: name,
		path: path,
	}
}
