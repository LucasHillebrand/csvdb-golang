package main

import (
	filesys "csvDB/FileSys"
	"strconv"
)

func Runtime(orgpath string, cmd []string, main *Master, upd *Update) []string {
	out := make([]string, 2) // out[0] == output, out[1] == clientCommands/errors

	SPath := filesys.Split(orgpath, ":")
	GPath, _ := getPath(SPath[0], SPath[1], main)
	switch cmd[0] {
	case "ChangePath":
		if len(cmd) >= 2 {
			Path := make([]int, 2)
			Path[0] = 0
			if len(cmd) >= 3 {
				Path, _ = getPath(cmd[1], cmd[2], main)
			} else {
				Path, _ = getPath(cmd[1], "", main)
			}
			SPath := make([]string, 2)
			if Path[0] != -1 {
				SPath[0] = main.MasterIndex[Path[0]]
			}
			if Path[1] != -1 {
				SPath[1] = main.DBs[Path[0]].DBIndex[Path[1]]
			}
			for i := 0; i < len(SPath); i++ {
				if SPath[i] == "" {
					SPath[i] = "-1"
				}
			}

			out[1] = "NEWP:" + SPath[0] + ":" + SPath[1] + ""
			out[0] = "succes"
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "Search":
		if len(cmd) >= 4 {
			recursive := true
			if len(cmd) >= 5 {
				if cmd[4] == "false" {
					recursive = false
				}
			}
			if cmd[1] == "col" {
				line, _ := strconv.Atoi(cmd[2])
				if line >= 0 {
					outVal := "lines"
					col := main.DBs[GPath[0]].Tables[GPath[1]].GetCol(line)
					res, _ := searchArr(col, cmd[3], recursive)
					for i := 0; i < len(res); i++ {
						outVal += ":" + strconv.Itoa(res[i])
					}

					out[0] = outVal
					out[1] = "FORMAT"
				} else {
					out[1] = "ERROR:Nothing found"
				}
			} else if cmd[1] == "row" {
				line, _ := strconv.Atoi(cmd[2])
				var lines []int
				lines, _ = main.DBs[GPath[0]].Tables[GPath[1]].SearchRow(line, cmd[3], recursive)
				if len(lines) >= 1 {
					outVal := "cols:"
					for i := 0; i < len(lines); i++ {
						outVal += strconv.Itoa(lines[i])
						if i+1 < len(lines) {
							outVal += ":"
						}
					}
					out[0] = outVal
					out[1] = "FORMAT"
				} else {
					out[1] = "ERROR:Nothing found"
				}
			} else {
				out[1] = "ERROR:to few arguments"
			}
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "Add":
		if len(cmd) >= 3 {
			if cmd[1] == "row" {
				arr := filesys.Split((cmd[2]), ":")

				main.DBs[GPath[0]].Tables[GPath[1]].AddRow(arr)
				*main.DBs[GPath[0]].UpdPerms[GPath[1]] += 1
				out[0] = "succes"
			} else if cmd[1] == "col" {
				arr := filesys.Split((cmd[2]), ":")

				if len(arr) >= 2 {
					main.DBs[GPath[0]].Tables[GPath[1]].AddCol(arr[0], arr[1])
					out[0] = "succes"
					*main.DBs[GPath[0]].UpdPerms[GPath[1]] += 1
				} else {
					out[1] = "ERROR:to few seperated values"
				}
			} else {
				out[1] = "ERROR:no valid col or row value"
			}
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "Remove":
		if len(cmd) >= 3 {
			item, _ := strconv.Atoi(cmd[2])
			if cmd[1] == "col" {
				main.DBs[GPath[0]].Tables[GPath[1]].RemoveCol(item)
				*main.DBs[GPath[0]].UpdPerms[GPath[1]] += 1
				out[0] = "Succes"
			} else if cmd[1] == "row" {
				main.DBs[GPath[0]].Tables[GPath[1]].RemoveRow(item)
				*main.DBs[GPath[0]].UpdPerms[GPath[1]] += 1
				out[0] = "Succes"
			} else {
				out[1] = "ERROR:no valid col or row value"
			}
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "Get":
		if len(cmd) >= 2 {
			if cmd[1] == "all" {
				out[0] = filesys.CsvToString(main.DBs[GPath[0]].Tables[GPath[1]].Database)
				out[1] = "CSV"
			} else if cmd[1] == "coords" && len(cmd) >= 3 {
				coords := filesys.Split((cmd[2]), ":")
				if len(coords) >= 2 {
					row, _ := strconv.Atoi(coords[1])
					col, _ := strconv.Atoi(coords[0])
					out[0] = main.DBs[GPath[0]].Tables[GPath[1]].Database[row][col]
				} else {
					out[1] = "ERROR:coords false formated"
				}
			} else if cmd[1] == "row" && len(cmd) >= 3 {
				rowNum, _ := strconv.Atoi(cmd[2])
				Row := main.DBs[GPath[0]].Tables[GPath[1]].GetRow(rowNum)
				outVal := "row:"
				for i := 0; i < len(Row); i++ {
					outVal += Row[i]
					if i+1 < len(Row) {
						outVal += ":"
					}
				}
				out[0] = outVal
				out[1] = "FORMAT"
			} else if cmd[1] == "col" && len(cmd) >= 3 {
				colNum, _ := strconv.Atoi(cmd[2])
				Col := main.DBs[GPath[0]].Tables[GPath[1]].GetCol(colNum)
				outVal := "column:"
				for i := 0; i < len(Col); i++ {
					outVal += Col[i]
					if i+1 < len(Col) {
						outVal += ":"
					}
				}
				out[0] = outVal
				out[1] = "FORMAT"
			} else {
				out[1] = "ERROR:no valid mode"
			}
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "Replace":
		if len(cmd) >= 3 {
			cmd[1] = (cmd[1])
			SCoords := filesys.Split(cmd[1], ":")
			Coords := make([]int, 2)
			for i := 0; i < len(Coords); i++ {
				Coords[i], _ = strconv.Atoi(SCoords[i])
			}
			main.DBs[GPath[0]].Tables[GPath[1]].Replace(Coords[1], Coords[0], cmd[2])
			*main.DBs[GPath[0]].UpdPerms[GPath[1]] += 1
			out[0] = "Succes"
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "List":
		if len(cmd) >= 2 {
			if cmd[1] == "dbs" {
				val := main.MasterIndex
				outVal := "Databases:"
				for i := 0; i < len(val); i++ {
					outVal += val[i]
					if i+1 < len(val) {
						outVal += ":"
					}
				}
				out[0] = outVal
				out[1] = "FORMAT"
			} else if cmd[1] == "tables" {
				val := main.DBs[GPath[0]].DBIndex
				outVal := "Tables:"
				for i := 0; i < len(val); i++ {
					outVal += val[i]
					if i+1 < len(val) {
						outVal += ":"
					}
				}
				out[0] = outVal
				out[1] = "FORMAT"
			} else {
				out[1] = "ERROR:no valid mode"
			}
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "AddDB":
		if len(cmd) >= 4 {
			main.AddDB(cmd[1], cmd[2], filesys.Split((cmd[3]), ":"), upd)
			out[0] = "succes"
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "AddTable":
		if len(cmd) >= 3 {
			main.DBs[GPath[0]].CreateTable(cmd[1], filesys.Split((cmd[2]), ":"), upd)
			out[0] = "succes"
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "RemoveDB":
		if len(cmd) >= 2 {
			main.RemoveDB(cmd[1])
			out[0] = "succes"
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "RemoveTable":
		if len(cmd) >= 2 {
			main.DBs[GPath[0]].DeleteTable(cmd[1])
			out[0] = "succes"
		} else {
			out[1] = "ERROR:to few arguments"
		}
	case "Stop":
		upd.Stop()
		//time.Sleep(time.Millisecond * 5)
	case "Init":
		Path := []int{0, 0}
		SPath := make([]string, 2)
		if Path[0] != -1 {
			SPath[0] = main.MasterIndex[Path[0]]
		}
		if Path[1] != -1 {
			SPath[1] = main.DBs[Path[0]].DBIndex[Path[1]]
		}
		for i := 0; i < len(SPath); i++ {
			if SPath[i] == "" {
				SPath[i] = "-1"
			}
		}

		out[1] = "NEWP:" + SPath[0] + ":" + SPath[1]
		out[0] = "Succesfull login"
	default:
		out[0] = "Command Not Valid help might be:\n" + filesys.GetFile("commands.txt")
	}

	return out
}

func getPath(DB, Table string, main *Master) (out []int, err string) {
	out = make([]int, 2)

	for i := 0; i < len(out); i++ {
		out[i] = -1
	}

	item, newErr := searchArr(main.MasterIndex, DB, false)
	if newErr == "" && len(item) > 0 {
		out[0] = item[0]
		item, _ = searchArr(main.DBs[out[0]].DBIndex, Table, false)
		if len(item) > 0 {
			out[1] = item[0]
		}
	} else {
		err = "database not found"
	}
	return
}
