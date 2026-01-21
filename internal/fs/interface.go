package fs

type FileSystem interface {
	ExpandPath(path string) (string, error)
	DirSize(path string) (int64, error)
}
