package main

import "fmt"

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	fmt.Println("KB:", KB, " bytes")
	fmt.Println("MB:", MB, " bytes")
	fmt.Println("GB:", GB, " bytes")
	fmt.Println("TB:", TB, " bytes")
	fmt.Println("PB:", PB, " bytes")

}
