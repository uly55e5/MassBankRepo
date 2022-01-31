package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

const (
	mbPath string = "/home/david/Projekte/MassBank-data"
	output string = "mb2_tag_list.txt"
)

var tags = map[string]map[string]bool{}

var filesN = 0

var withSubtags = map[string]bool{
	"CH$LINK":              true,
	"SP$LINK":              true,
	"AC$MASS_SPECTROMETRY": true,
	"AC$CHROMATOGRAPHY":    true,
	"AC$GENERAL":           true,
	"MS$FOCUSED_ION":       true,
	"MS$DATA_PROCESSING":   true,
}

func main() {
	err := filepath.Walk(mbPath, ReadTags)
	if err != nil {
		println(err.Error())
	}
	writeList()
	println(filesN)
}

func writeList() {
	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		println(err.Error())
	}
	for k, st := range tags {
		for kk, _ := range st {
			file.WriteString(k + " " + kk + "\n")
			println(k, kk)
		}
	}
}

func ReadTags(path string, info os.FileInfo, err error) error {
	if strings.HasSuffix(info.Name(), ".txt") {
		filesN += 1
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {

			line := scanner.Text()
			parts := strings.Split(line, " ")
			if len(parts) > 0 && len(parts[0]) > 0 && !strings.HasPrefix(line, "//") {
				tag := strings.Trim(parts[0], ": ")
				subtag := ""
				if len(parts) > 1 {
					if withSubtags[tag] && err == nil {
						subtag = parts[1]
					}
				}
				subtags := tags[tag]
				if subtags == nil {
					subtags = map[string]bool{subtag: true}
				} else {
					subtags[subtag] = true
				}
				tags[tag] = subtags
			}
		}
	}
	return nil
}
