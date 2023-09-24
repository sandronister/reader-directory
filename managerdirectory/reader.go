package managerdirectory

import (
	"fmt"
	"io/fs"
	"os"
)

type Reader struct{}

func (m *Reader) check(err error) {
	if err != nil {
		fmt.Print((err))
		panic(err)
	}
}

func (m *Reader) getDir(path string) []fs.DirEntry {
	list, err := os.ReadDir(path)
	m.check(err)
	return list
}

func (m *Reader) GetFilesList(path string) []DirFiles {
	list := m.getDir(path)
	result := make([]DirFiles, 0)

	for _, item := range list {
		if item.IsDir() {
			if item.Name() != ".git" {
				subDir := m.GetFilesList(path + item.Name() + "/")
				subItem := DirFiles{Kind: "dir", Name: item.Name(), Sub: subDir}
				result = append(result, subItem)
			}
			continue
		}
		subFile := DirFiles{Kind: "file", Name: item.Name()}
		result = append(result, subFile)

	}

	return result
}
