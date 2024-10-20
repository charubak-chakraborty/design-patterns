package composite

type Folder struct {
	components []Component
	Name       string
}

func (f *Folder) Search(s string) {
	for _, c := range f.components {
		c.Search(s)
	}
}
func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}
