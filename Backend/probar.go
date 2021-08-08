package main

import (
	"fmt"
)

func Atoi(s string) int {
	var (
		n uint64
		i int
		v byte
	)
	for ; i < len(s); i++ {
		d := s[i]
		if '0' <= d && d <= '9' {
			v = d - '0'
		} else if 'a' <= d && d <= 'z' {
			v = d - 'a' + 10
		} else if 'A' <= d && d <= 'Z' {
			v = d - 'A' + 10
		} else {
			n = 0
			break
		}
		n *= uint64(10)
		n += uint64(v)
	}
	return int(n)
}

func main() {
	IniV := "2021-02-12"
	finV := "2021-03-12"
	IniV := "2021-02-12"
	finV := "2021-03-12"

	dia1 := Atoi(IniV[8:])
	mes1 := Atoi(IniV[5:7])
	dia1f := Atoi(finV[8:])
	mes1f := Atoi(finV[5:7])
	dia2 := Atoi(Ini[8:])
	mes2 := Atoi(Ini[5:7])
	dia2f := Atoi(fin[8:])
	mes2f := Atoi(fin[5:7])

	fmt.Print(dia)

	fmt.Print(mes)

}
