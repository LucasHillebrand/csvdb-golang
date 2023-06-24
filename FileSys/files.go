package filesys

import (
	"os"
)

func GetFile(FileName string) string {
	out := ""
	data, err := os.ReadFile(FileName)
	if err == nil {
		for i := 0; i < len(data); i++ {
			out += string(data[i])
		}
	}
	return out
}

func GetCsvFile(FileName string) [][]string {
	lines := GetLines(FileName)
	out := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		out[i] = Split(lines[i], ",")
	}
	return out
}

func GetLines(FileName string) []string {
	data := GetFile(FileName)
	out := Split(data, "\n")
	return out
}

func CsvToString(val [][]string) string {
	out := ""
	for i := 0; i < len(val); i++ {
		for j := 0; j < len(val[i]); j++ {
			out += val[i][j]
			if j+1 < len(val[i]) {
				out += ","
			}
		}
		if i+1 < len(val) {
			out += "\n"
		}
	}
	return out
}

func LinesToString(org []string) string {
	out := ""
	for i := 0; i < len(org); i++ {
		if i+1 < len(org) {
			out += org[i] + "\n"
		} else {
			out += org[i]
		}
	}
	return out
}
