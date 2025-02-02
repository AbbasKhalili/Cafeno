package Cafeno_Domain

type Category struct {
	Name string
}

func CreateCategory(name string) *Category {
	return &Category{name}
}
