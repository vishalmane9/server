package main

type Person struct {
	Name `json:"Name"`
	Age  `json:"Age"`
	Year `json:"Year"`
}

type Name string
type Age float64
type Year int

type Info interface {
	getInfo() string
}

func (p Name) getInfo() Name {
	return p
}
func (p Age) getInfo() Age {
	return p
}
func (p Year) getInfo() Year {
	return p
}
