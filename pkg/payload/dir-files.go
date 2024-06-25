package payload

type DirFiles struct {
	Kind string
	Name string
	Sub  []DirFiles
}
