package main

import (
	filesys "csvDB/FileSys"
	"os"
)

type Master struct {
	DBs         []*Database
	MasterIndex []string
	Path        string
}

func CreateMaster(path string, upd *Update, DBName, DBTableName string, header []string) *Master {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	filesys.MkDir(path)
	out := Master{
		Path: path,
		DBs:  make([]*Database, 0),
	}
	out.AddDB(DBName, DBTableName, header, upd)
	os.WriteFile(path+"index.txt", toBytes(filesys.LinesToString(out.MasterIndex)), os.FileMode(0666))

	return &out
}

func LoadMaster(path string, upd *Update) *Master {
	if path[len(path)-1] != '/' {
		path += "/"
	}

	out := Master{
		MasterIndex: filesys.GetLines(path + "index.txt"),
		Path:        path,
	}

	dbs := make([]*Database, len(out.MasterIndex))

	for i := 0; i < len(out.MasterIndex); i++ {
		dbs[i] = OpenDB(out.Path, out.MasterIndex[i], upd)
	}

	out.DBs = dbs
	return &out
}

func (itSelf *Master) AddDB(name string, TableName string, TableHeader []string, upd *Update) string {
	index, _ := searchArr(itSelf.MasterIndex, name, false)
	err := ""
	if len(index) == 0 {
		db := CreateDB(itSelf.Path, name, TableName, TableHeader, upd)
		itSelf.DBs = growDBPointArr(itSelf.DBs, db)
		itSelf.MasterIndex = growStringArr(itSelf.MasterIndex, name)
		os.WriteFile(itSelf.Path+"index.txt", toBytes(filesys.LinesToString(itSelf.MasterIndex)), os.FileMode(0666))
	} else {
		err = name + " already exists"
	}
	return err
}

func (itSelf *Master) RemoveDB(name string) string {
	err := ""
	if len(itSelf.MasterIndex) > 1 {
		index, err := searchArr(itSelf.MasterIndex, name, false)
		if err == "" {
			itSelf.DBs = removeDBPointArr(itSelf.DBs, index[0])
			itSelf.MasterIndex = removeStringArr(itSelf.MasterIndex, index[0])
			os.WriteFile(itSelf.Path+"index.txt", toBytes(filesys.LinesToString(itSelf.MasterIndex)), os.FileMode(0666))
			filesys.Remove(itSelf.Path + name)
		} else {
			err = "name not found"
		}
	} else {
		err = "you need 2 or more databases to remove one"
	}
	return err
}
