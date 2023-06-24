package filesys

func growFileTree(inp []FileTree, newValue FileTree) []FileTree {
	out := make([]FileTree, len(inp)+1)
	for i := 0; i < len(inp); i++ {
		out[i] = inp[i]
	}
	out[len(inp)] = newValue
	return out
}

func growList(inp []string, newValue string) []string {
	out := make([]string, len(inp)+1)
	for i := 0; i < len(inp); i++ {
		out[i] = inp[i]
	}
	out[len(inp)] = newValue
	return out
}

func shrinkLeft(orig string, items int) string {
	out := ""
	for i := items; i < len(orig); i++ {
		out += string(orig[i])
	}
	return out
}

func getPath(path string) (fileRoot, fileName string) {
	fileRoot = ""
	fileName = ""

	fileNameRev := ""
	fileRootRev := ""

	current := "Name"
	for i := len(path) - 1; i >= 0; i-- {
		if byte(path[i]) == '/' && current == "Name" {
			current = "Root"
		}
		if current == "Name" {
			fileNameRev += string(path[i])
		} else {
			fileRootRev += string(path[i])
		}
	}

	for i := len(fileNameRev) - 1; i >= 0; i-- {
		fileName += string(fileNameRev[i])
	}

	for i := len(fileRootRev) - 1; i >= 0; i-- {
		fileRoot += string(fileRootRev[i])
	}

	return
}
