package main

import (
	filesys "csvDB/FileSys"
	"os"
)

type Database struct {
	Path     string
	DBIndex  []string
	Tables   []*DBTable
	UpdPerms []*int
}

func OpenDB(root, name string, upd *Update) *Database {
	if name[len(name)-1] != '/' {
		name += "/"
	}
	if root[len(root)-1] != '/' {
		root += "/"
	}
	path := root + name

	lines := filesys.GetLines(path + "index.txt")

	out := Database{
		DBIndex: lines,
		Path:    path,
	}

	updPerms := make([]int, len(lines))
	Table := make([]*DBTable, len(lines))

	for i := 0; i < len(lines); i++ {
		Table[i] = LoadDBTable(path+lines[i]+".csv", upd, &updPerms[i])
	}
	out.UpdPerms = make([]*int, len(updPerms))
	for i := 0; i < len(updPerms); i++ {
		out.UpdPerms[i] = &updPerms[i]
	}

	out.Tables = Table

	return &out
}

func CreateDB(path, name string, TableName string, TableHeader []string, upd *Update) *Database {
	if path[len(path)-1] != '/' {
		path += "/"
	}
	out := Database{
		Path: path + name + "/",
	}
	filesys.MkDir(out.Path)
	out.CreateTable(TableName, TableHeader, upd)

	return &out
}

func (itSelf *Database) CreateTable(name string, header []string, upd *Update) string {
	c, _ := searchArr(itSelf.DBIndex, name, false)
	err := ""
	if len(c) == 0 {
		updPerm := 0
		itSelf.Tables = growDBTable(itSelf.Tables, CreateDBTable(itSelf.Path+name+".csv", header, upd, &updPerm))
		itSelf.UpdPerms = growIntPointLst(itSelf.UpdPerms, &updPerm)
		itSelf.DBIndex = growStringArr(itSelf.DBIndex, name)

		os.WriteFile(itSelf.Path+"index.txt", toBytes(filesys.LinesToString(itSelf.DBIndex)), os.FileMode(0666))
	} else {
		err = name + " already exists"
	}
	return err
}

func (itSelf *Database) DeleteTable(name string) string {
	err := ""
	if len(itSelf.DBIndex) > 1 {
		item, _ := searchArr(itSelf.DBIndex, name, false)
		if len(item) > 0 {
			itSelf.DBIndex = removeStringArr(itSelf.DBIndex, item[0])
			itSelf.Tables = removePointTable(itSelf.Tables, item[0])
			itSelf.UpdPerms = removePointInt(itSelf.UpdPerms, item[0])
			os.WriteFile(itSelf.Path+"index.txt", toBytes(filesys.LinesToString(itSelf.DBIndex)), os.FileMode(0666))
			filesys.Remove(itSelf.Path + name + ".csv")
		} else {
			err = name + " doesn't exist"
		}
	} else {
		err = "to few tables"
	}
	return err
}
