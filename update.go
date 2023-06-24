package main

import (
	filesys "csvDB/FileSys"
	"fmt"
	"os"
	"time"
)

type Update struct {
	fileNames  []string
	dataPoints []*[][]string
	updPerm    []*int
	stop       bool
}

func NewUpdate() Update {
	out := Update{
		stop: false,
	}
	return out
}

func (itSelf *Update) Stop() {
	itSelf.stop = true
}

func (itSelf *Update) Add(db *DBTable, updPerm *int) {
	itSelf.fileNames = growStringArr(itSelf.fileNames, db.FileName)
	itSelf.dataPoints = growDataPoints(itSelf.dataPoints, &db.Database)
	itSelf.updPerm = growIntPointLst(itSelf.updPerm, updPerm)
}

func (itSelf *Update) Upd(kA *KeepAlive) {
	for !itSelf.stop {
		//time.Sleep(time.Minute * 1)
		time.Sleep(time.Millisecond * 500)
		kA.Add(1)
		for i := 0; i < len(itSelf.fileNames); i++ {
			if *itSelf.updPerm[i] > 0 {
				os.WriteFile(itSelf.fileNames[i], toBytes(filesys.CsvToString(*itSelf.dataPoints[i])), os.FileMode(0666))
				fmt.Printf("updated file: %s\n", itSelf.fileNames[i])
				*itSelf.updPerm[i]--
			}
		}
		kA.Done()
	}
}
