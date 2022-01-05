package combination

import "fmt"

type Component interface {
	Display()
}

type File struct {
	Name string
}

func NewFile(name string) File {
	return File{
		Name: name,
	}
}

func (f *File) Display() {
	fmt.Printf("File: %s \n", f.Name)
}

type Folder struct {
	Name       string
	Components []Component
}

func NewFolder(name string) Folder {
	return Folder{
		Name:name,
	}
}

func (f *Folder) Add(c Component) {
	f.Components = append(f.Components, c)
}

func (f *Folder) Display() {
	fmt.Printf("folder: %s \n", f.Name)
	for _, c := range f.Components {
		c.Display()
	}
}
