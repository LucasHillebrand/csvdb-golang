package filesys

type FileTree struct {
	Name  string
	IsDir bool
}

func ListDirs(base []FileTree) []string {
	out := make([]string, 0)
	for i := 0; i < len(base); i++ {
		if base[i].IsDir {
			out = growList(out, base[i].Name)
		}
	}
	return out
}

func ListFiles(base []FileTree) []string {
	out := make([]string, 0)
	for i := 0; i < len(base); i++ {
		if !base[i].IsDir {
			out = growList(out, base[i].Name)
		}
	}
	return out
}
