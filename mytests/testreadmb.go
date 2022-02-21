package main

import "MassBankRepo/lib/massbank"

func main() {
	var mb = massbank.Massbank{}
	mb.ParseFile("TestFiles/MassBank Files/XX000201.txt")
	println("finished reading")
	res, err := mb.Validate()
	println(res, err)
}
