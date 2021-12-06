package main

import "fmt"

func main() {
	testString := "vV"
	fmt.Println(letterCounter(testString, "a"))
	//fmt.Println(letterCounter("Go is awesome", "a"))
}

func letterCounter(str string, letter string) int {
	alphabet := map[string]int{
		"a":     0,
		"b":     0,
		"c":     0,
		"d":     0,
		"e":     0,
		"f":     0,
		"g":     0,
		"h":     0,
		"i":     0,
		"j":     0,
		"k":     0,
		"l":     0,
		"m":     0,
		"n":     0,
		"o":     0,
		"p":     0,
		"q":     0,
		"r":     0,
		"s":     0,
		"t":     0,
		"u":     0,
		"v":     0,
		"w":     0,
		"x":     0,
		"y":     0,
		"z":     0,
		"space": 0,
	}
	fmt.Println(alphabet["a"])
	bs := []byte(str)
	fmt.Println(bs)
	for _, v := range bs {
		switch {
		case v == 65 || v == 97:
			alphabet["a"]++

		case v == 66 || v == 98:
			alphabet["b"]++

		case v == 67 || v == 99:
			alphabet["c"]++

		case v == 68 || v == 100:
			alphabet["d"]++

		case v == 69 || v == 101:
			alphabet["e"]++

		case v == 70 || v == 102:
			alphabet["f"]++

		case v == 71 || v == 103:
			alphabet["g"]++

		case v == 72 || v == 104:
			alphabet["h"]++

		case v == 73 || v == 105:
			alphabet["i"]++

		case v == 74 || v == 106:
			alphabet["j"]++

		case v == 75 || v == 107:
			alphabet["k"]++

		case v == 76 || v == 108:
			alphabet["l"]++

		case v == 77 || v == 109:
			alphabet["m"]++

		case v == 78 || v == 110:
			alphabet["n"]++

		case v == 79 || v == 111:
			alphabet["o"]++

		case v == 80 || v == 112:
			alphabet["p"]++

		case v == 81 || v == 113:
			alphabet["q"]++

		case v == 82 || v == 114:
			alphabet["r"]++

		case v == 83 || v == 115:
			alphabet["s"]++

		case v == 84 || v == 116:
			alphabet["t"]++

		case v == 85 || v == 117:
			alphabet["u"]++

		case v == 86 || v == 118:
			alphabet["v"]++

		case v == 87 || v == 119:
			alphabet["w"]++

		case v == 88 || v == 120:
			alphabet["x"]++

		case v == 89 || v == 121:
			alphabet["y"]++

		case v == 90 || v == 122:
			alphabet["z"]++

		case v == 32:
			alphabet["space"]++

		}
	}
	fmt.Println(alphabet)
	return alphabet[letter]
}
