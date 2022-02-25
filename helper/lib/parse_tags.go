package lib

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var tags = map[string]map[string]bool{}

type statValues struct {
	uniqueVals map[string]bool
	frequency  int
}

var stats = map[string]map[string]statValues{}

type Mode int

const (
	Tags Mode = iota
	Stats
	UniqueVals
)

var mode Mode = Tags
var withSubtags = map[string]bool{
	"CH$LINK":              true,
	"SP$LINK":              true,
	"AC$MASS_SPECTROMETRY": true,
	"AC$CHROMATOGRAPHY":    true,
	"AC$GENERAL":           true,
	"MS$FOCUSED_ION":       true,
	"MS$DATA_PROCESSING":   true,
}

var filesN = 0

func SetMode(m Mode) {
	mode = m
}

func WriteList(outFileName string) {
	file, err := os.OpenFile(outFileName, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		println(err.Error())
	}
	switch mode {
	case Tags:
		outputTags(file)
	case Stats:
		outputStats(file)
	}
}

func outputTags(file *os.File) {
	for k, st := range tags {
		for kk, _ := range st {
			file.WriteString(k + " " + kk + "\n")
			println(k, kk)
		}
	}
	println(filesN)
}

func outputStats(file *os.File) {
	for t, st := range stats {
		for stt, v := range st {
			s := t + " " + stt + ":  " + strconv.Itoa(v.frequency) + " / " + strconv.Itoa(len(v.uniqueVals)) + "\n"
			file.WriteString(s)
			println(s)
		}
	}
}

type processFunc func(tag string, subtag string, value string)

func ReadTags(path string, info os.FileInfo, err error) error {
	var processTags processFunc = addToTag
	switch mode {
	case Tags:
		processTags = addToTag
	case Stats, UniqueVals:
		processTags = buildStatistics
	}
	if info != nil && strings.HasSuffix(info.Name(), ".txt") {
		filesN += 1
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		scanner := bufio.NewScanner(file)
		var parts = make([]string, 3)
		for scanner.Scan() {
			line := scanner.Text()
			parts = strings.SplitN(line, " ", 3)
			if len(parts) > 0 && len(parts[0]) > 0 && !strings.HasPrefix(line, "//") {
				tag := strings.Trim(parts[0], ": ")
				subtag := ""
				value := ""
				if len(parts) > 1 {
					if withSubtags[tag] && err == nil {
						subtag = parts[1]
					} else {
						value = parts[1]
					}
					if len(parts) > 2 {
						value += " " + parts[2]
					}
				}
				processTags(tag, subtag, value)
			}
		}
	}
	return nil
}

func addToTag(tag string, subtag string, value string) {
	subtags := tags[tag]
	if subtags == nil {
		subtags = map[string]bool{subtag: true}
	} else {
		subtags[subtag] = true
	}
	tags[tag] = subtags
}

func buildStatistics(tag string, subtag string, value string) {
	subtags := stats[tag]
	if subtags == nil {
		subtags = map[string]statValues{}
	}
	var statVals = subtags[subtag]
	if statVals.uniqueVals == nil {
		statVals.uniqueVals = map[string]bool{value: true}
	} else if _, ok := statVals.uniqueVals[value]; !ok {
		statVals.uniqueVals[value] = true
	}
	statVals.frequency += 1
	subtags[subtag] = statVals
	stats[tag] = subtags
}
