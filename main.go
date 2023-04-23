package main
 
import (
	"fmt"
)

type Person struct {
	firstname, lastname string
}

type SpecialPerson struct {
	firstname, middlename, lastname string
}

type NameGetter interface {
	FullName() string
}

func (p Person) FullName() string {
	return p.firstname + " " + p.lastname
}

func (sp SpecialPerson) FullName() string {
	return sp.firstname + " " + sp.middlename + " " + sp.lastname
}

func main() {
	var p = Person { 
		"Adam",
		"McDowell",
	}
	var sp = SpecialPerson {
		"Adam",
		"P",
		"McDowell",
	}
	ng := []NameGetter{sp, p}

	var i int = 0
	for range ng {
		fmt.Println(ng[i].FullName())
		i ++
	}
}