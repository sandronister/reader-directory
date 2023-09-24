package managerdirectory

import (
	"fmt"
)

type Printer struct{}

func (p *Printer) printDir(list []DirFiles, path string) {
	for _, item := range list {
		if item.Kind == "dir" {
			fmt.Println("+" + path + item.Name + "/")
			p.PrintList(item.Sub, path+item.Name+"/")
		}
	}
}

func (p *Printer) PrintList(list []DirFiles, path string) {
	for _, item := range list {
		if item.Kind == "file" {
			fmt.Println("\t" + path + item.Name)
		}
	}
	for _, item := range list {
		if item.Kind == "dir" {
			p.printDir(item.Sub, path+item.Name+"/")
		}
	}
}
