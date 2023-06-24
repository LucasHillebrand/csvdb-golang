package main

import (
	filesys "csvDB/FileSys"
	"net"
)

func growDataPoints(inp []*[][]string, addon *[][]string) []*[][]string {
	out := make([]*[][]string, len(inp)+1)
	for i := 0; i < len(inp)+1; i++ {
		if i < len(inp) {
			out[i] = inp[i]
		} else {
			out[i] = addon
		}
	}
	return out
}

func growStringArr(inp []string, addon string) []string {
	out := make([]string, len(inp)+1)
	for i := 0; i < len(inp)+1; i++ {
		if i < len(inp) {
			out[i] = inp[i]
		} else {
			out[i] = addon
		}
	}
	return out
}

func growIntPointLst(inp []*int, addon *int) []*int {
	out := make([]*int, len(inp)+1)
	for i := 0; i < len(inp)+1; i++ {
		if i < len(inp) {
			out[i] = inp[i]
		} else {
			out[i] = addon
		}
	}
	return out
}

func grow2DArrDown(org [][]string, addon []string) [][]string {
	out := make([][]string, len(org)+1)

	for i := 0; i < len(org)+1; i++ {
		if i < len(org) {
			out[i] = org[i]
		} else {
			out[i] = addon
		}
	}

	return out
}

func remove2DArrDown(org [][]string, item int) [][]string {
	org[item] = org[len(org)-1]

	out := make([][]string, len(org)-1)
	for i := 0; i < len(org)-1; i++ {
		if i < len(org) {
			out[i] = org[i]
		}
	}

	return out
}

func grow2DArrLeft(org [][]string, header, defaultValue string) [][]string {
	out := make([][]string, len(org))
	for i := 0; i < len(org); i++ {
		if i == 0 {
			out[i] = growStringArr(org[i], header)
		} else {
			out[i] = growStringArr(org[i], defaultValue)
		}
	}
	return out
}

func removeStringArr(org []string, item int) []string {
	out := make([]string, len(org)-1)

	org[item] = org[len(org)-1]

	for i := 0; i < len(org)-1; i++ {
		out[i] = org[i]
	}
	return out
}

func remove2DArrLeft(org [][]string, item int) [][]string {
	out := make([][]string, len(org))
	for i := 0; i < len(org); i++ {
		//fmt.Println(org[i])
		out[i] = removeStringArr(org[i], item)
	}
	return out
}

func growIntArr(org []int, addon int) []int {
	out := make([]int, len(org)+1)
	for i := 0; i < len(org)+1; i++ {
		if i < len(org) {
			out[i] = org[i]
		} else {
			out[i] = addon
		}
	}

	return out
}

func searchArr(array []string, keyword string, recursive bool) (result []int, err string) {
	result = make([]int, 0)
	if !recursive {
		for i := 0; i < len(array); i++ {
			if array[i] == keyword {
				result = growIntArr(result, i)
			}
		}
	}

	if recursive {
		for i := 0; i < len(array); i++ {
			if filesys.Count(array[i], keyword) > 0 {
				result = growIntArr(result, i)
			}
		}
	}

	if len(result) == 0 {
		err = "value not found"
	}

	return
}

func growDBTable(org []*DBTable, addon *DBTable) []*DBTable {
	out := make([]*DBTable, len(org)+1)
	for i := 0; i < len(org)+1; i++ {
		if i < len(org) {
			out[i] = org[i]
		} else {
			out[i] = addon
		}
	}
	return out
}

func growDBPointArr(org []*Database, new *Database) []*Database {
	out := make([]*Database, len(org)+1)
	for i := 0; i < len(org)+1; i++ {
		if i < len(org) {
			out[i] = org[i]
		} else {
			out[i] = new
		}
	}
	return out
}

func removeDBPointArr(org []*Database, index int) []*Database {
	out := make([]*Database, len(org)-1)
	org[index] = org[len(org)-1]

	for i := 0; i < len(out); i++ {
		out[i] = org[i]
	}

	return out
}

func removePointTable(org []*DBTable, index int) []*DBTable {
	out := make([]*DBTable, len(org)-1)
	org[index] = org[len(org)-1]
	for i := 0; i < len(out); i++ {
		out[i] = org[i]
	}
	return out
}

func removePointInt(org []*int, index int) []*int {
	out := make([]*int, len(org)-1)
	org[index] = org[len(org)-1]
	for i := 0; i < len(out); i++ {
		out[i] = org[i]
	}
	return out
}

func growConnArr(org []*net.Conn, addon *net.Conn) []*net.Conn {
	out := make([]*net.Conn, len(org)+1)
	for i := 0; i < len(org)+1; i++ {
		if i < len(org) {
			out[i] = org[i]
		} else {
			out[i] = addon
		}
	}

	return out
}

func removeConnArr(org []*net.Conn, item int) []*net.Conn {
	out := make([]*net.Conn, len(org)-1)
	org[item] = org[len(org)-1]

	for i := 0; i < len(out); i++ {
		out[i] = org[i]
	}

	return out
}
