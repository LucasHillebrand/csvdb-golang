package main

import (
	filesys "csvDB/FileSys"
	"os"
)

type DBTable struct {
	Database [][]string
	FileName string
}

func LoadDBTable(filename string, upd *Update, updPerm *int) *DBTable {
	out := DBTable{}
	out.Load(filename)
	upd.Add(&out, updPerm)
	return &out
}

func CreateDBTable(filename string, header []string, upd *Update, updPerm *int) *DBTable {
	out := DBTable{}
	out.FileName = filename
	out.Database = make([][]string, 1)
	out.Database[0] = header
	os.WriteFile(filename, toBytes(filesys.CsvToString(out.Database)), os.FileMode(0666))
	upd.Add(&out, updPerm)
	return &out
}

func (db *DBTable) Load(fileName string) {
	db.Database = filesys.GetCsvFile(fileName)
	db.FileName = fileName
}

// Datasheet changeVal
func (itSelf *DBTable) AddRow(addon []string) {
	if len(addon) != len(itSelf.Database[0]) {
		tmparr := make([]string, len(itSelf.Database[0]))
		for i := 0; i < len(itSelf.Database[0]); i++ {
			if len(addon) > i {
				tmparr[i] = addon[i]
			} else {
				tmparr[i] = ""
			}
		}
		addon = tmparr
	}

	itSelf.Database = grow2DArrDown(itSelf.Database, addon)
}

func (itSelf *DBTable) RemoveRow(index int) (err string) {
	err = ""
	if index != 0 {
		itSelf.Database = remove2DArrDown(itSelf.Database, index)
	} else {
		err = "ERROR: row 0 can't be removed"
	}
	return err
}

func (itSelf *DBTable) AddCol(header, defVal string) {
	itSelf.Database = grow2DArrLeft(itSelf.Database, header, defVal)
}

func (itSelf *DBTable) RemoveCol(item int) (err string) {
	if len(itSelf.Database[0]) > 1 && item < len(itSelf.Database[0]) {
		itSelf.Database = remove2DArrLeft(itSelf.Database, item)
	} else if !(len(itSelf.Database[0]) > 1) {
		err = "ERROR: collumn can't be removed due to \"to few cols\""
	} else {
		err = "ERROR: collumn can't be removed due to \"selected item isn't available\""
	}
	return
}

func (itSelf *DBTable) Replace(row, col int, value string) (err string) {
	if row < len(itSelf.Database) && col < len(itSelf.Database[0]) {
		itSelf.Database[row][col] = value
	} else {
		err = "ERROR: collumn or/and row coords are wrong"
	}
	return
}

func (itSelf *DBTable) Get(row, col int) string {
	return itSelf.Database[row][col]
}
func (itSelf *DBTable) GetCol(col int) []string {
	tmpArr := make([]string, len(itSelf.Database))
	for i := 0; i < len(itSelf.Database); i++ {
		tmpArr[i] = itSelf.Database[i][col]
	}
	return tmpArr
}

func (itSelf *DBTable) GetRow(row int) []string {
	return itSelf.Database[row]
}

func (itSelf *DBTable) SearchCol(col int, keyword string, recursive bool) (result []int, err string) {
	result, err = searchArr(itSelf.GetCol(col), keyword, recursive)

	return
}

func (itSelf *DBTable) SearchRow(row int, keyword string, recursive bool) (result []int, err string) {
	result, err = searchArr(itSelf.Database[row], keyword, recursive)

	return
}
