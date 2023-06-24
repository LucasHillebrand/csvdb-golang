package filesys

import (
	"os"
)

func Copy(orgPath, dstPath string) {
	if IsDir(orgPath) {
		CopyDir(orgPath, dstPath)
	} else {
		CopyFile(orgPath, dstPath)
	}
}

func CopyFile(orgPath string, dstPath string) (err string) {
	orgFile, fail := os.ReadFile(orgPath)
	if fail != nil {
		err = "you don't have the permission to read the orig file"
	} else {
		os.WriteFile(dstPath, orgFile, os.FileMode(0666))
	}
	return err
}

func MkDir(path string) (err string) {
	fail := os.Mkdir(path, os.FileMode(0777))
	if fail != nil {
		err = "error directory can't be created"
	}
	return err
}

func Remove(path string) (err string) {
	fail := os.RemoveAll(path)
	if fail != nil {
		err = "file: \"" + path + "\" can't be removed"
	}
	return
}

func ScanDirTree(dir string) []FileTree {
	if byte(dir[len(dir)-1]) != '/' {
		dir += "/"
	}
	que := make([]string, 1)
	que[0] = dir
	out := make([]FileTree, 1)
	out[0] = FileTree{Name: shrinkLeft(dir, len(dir)-1), IsDir: true}
	for j := 0; j < len(que); j++ {
		directory, _ := os.ReadDir(que[j])
		for i := 0; i < len(directory); i++ {
			if directory[i].IsDir() {
				que = growList(que, que[j]+directory[i].Name()+"/")
			}
			out = growFileTree(out, FileTree{Name: shrinkLeft(que[j]+directory[i].Name(), len(dir)-1), IsDir: directory[i].IsDir()})
			//fmt.Printf("%d %s < Name | IsDir > %t \n", len(out), out[len(out)-1].Name, out[len(out)-1].IsDir)
		}
	}
	return out
}

func CopyDir(origPath, dstPath string) (err string) {
	err = ""
	if !IsDir(origPath) {
		err = "selectet File/directory isn't a directory or is existent"
		return
	}
	if origPath[len(origPath)-1] == '/' {
		nP := ""
		for i := 0; i < len(origPath)-1; i++ {
			nP += string(origPath[i])
		}
		origPath = nP
	}

	if dstPath[len(dstPath)-1] == '/' {
		dP := ""
		for i := 0; i < len(dstPath)-1; i++ {
			dP += string(dstPath[i])
		}
		dstPath = dP
	}

	Tree := ScanDirTree(origPath)

	dirs := ListDirs(Tree)
	for i := 0; i < len(dirs); i++ {
		MkDir(dstPath + dirs[i])
	}

	files := ListFiles(Tree)
	for i := 0; i < len(files); i++ {
		CopyFile(origPath+files[i], dstPath+files[i])
	}
	return
}

func IsDir(path string) bool {
	basePath := path
	if byte(path[len(path)-1]) == '/' {
		basePath = ""
		for i := 0; i < len(path)-1; i++ {
			basePath += string(path[i])
		}
	}

	DirPath, FileName := getPath(basePath)
	fls, _ := os.ReadDir(DirPath)
	out := false

	for i := 0; i < len(fls); i++ {
		if fls[i].Name() == FileName {
			out = fls[i].IsDir()
		}
	}
	return out
}
