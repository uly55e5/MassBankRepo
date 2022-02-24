package main

import (
	"flag"
	"github.com/uly55e5/MassBankRepo/helper/lib"
	"path/filepath"
)

func main() {
	//defer profile.Start(profile.ProfilePath(".")).Stop()
	mbPath := flag.String("f", "", "path to massbank data")
	outFile := flag.String("o", "mbParser-out.txt", "name of the output file")
	stats := flag.Bool("s", false, "output statistics")
	flag.Parse()
	if *stats {
		lib.SetMode(lib.Stats)
	}
	err := filepath.Walk(*mbPath, lib.ReadTags)
	if err != nil {
		println(err.Error())
	}
	lib.WriteList(*outFile)
}
