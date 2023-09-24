package main

import "lab.directory/managerdirectory"

func main() {
	manager := managerdirectory.Reader{}
	printer := managerdirectory.Printer{}
	list := manager.GetFilesList("../")
	printer.PrintList(list, "../")
}
