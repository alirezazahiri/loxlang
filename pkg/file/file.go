package file

type File struct {
	Path string
}

func New(path string) *File {
	return &File{
		Path: path,
	}
}
