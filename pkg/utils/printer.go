package utils

import (
	"fmt"

	"lab.directory/pkg/payload"
)

type Printer struct{}

func (p *Printer) printDir(list []payload.DirFiles, path string) {
	for _, item := range list {
		if item.Kind == "dir" {
			fmt.Println("+" + path + item.Name + "/")
			p.PrintList(item.Sub, path+item.Name+"/")
		}
	}
}

func (p *Printer) PrintList(list []payload.DirFiles, path string) {
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
