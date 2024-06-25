package main

import "lab.directory/pkg/utils"

func main() {
	manager := utils.Reader{}
	printer := utils.Printer{}
	list := manager.GetFilesList("../")
	printer.PrintList(list, "../")
}
