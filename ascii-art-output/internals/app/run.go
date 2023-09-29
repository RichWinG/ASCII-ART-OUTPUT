package app

import (
	"ascii-art/internals/check"
	"fmt"
	"os"
	"strings"
)

func Run() {
	if len(os.Args) > 4 || len(os.Args) <= 1 {
		return
	}
	if len(os.Args) == 2 {
		return
	}
	input, output := os.Args[2], os.Args[1]
	if len(os.Args) == 3 {
		input = os.Args[1]
	}
	if output[len(output)-4:] != ".txt" {
		fmt.Println("err: files have to be on .txt format")
		return
	}
	if output == "--output=./banners/standard.txt" {
		fmt.Println("this file already exists")
		return
	} else if output == "--output=./banners/shadow.txt" {
		fmt.Println("this file already exists")
		return
	} else if output == "--output=./banners/thinkertoy.txt" {
		fmt.Println("this file already exists")
		return
	}

	banner := "./banners/standard.txt"
	if len(os.Args) > 3 {
		banner = os.Args[3]
	}
	if !check.Ascii(input) {
		return
	}

	elemsMap := make(map[rune][]string)
	if len(os.Args) > 3 && os.Args[3] == banner {
		switch banner {
		case "standard":
			banner = "./banners/standard.txt"
		case "shadow":
			banner = "./banners/shadow.txt"
		case "thinkertoy":
			banner = "./banners/thinkertoy.txt"
		default:
			// fmt.Println("Type a right banner name")
			return
		}
	}
	data, err := os.ReadFile(banner)
	if err != nil {
		return
	}

	changed := strings.ReplaceAll(string(data), "\r\n", "\n")
	sliceData := strings.Split(changed, "\n")              // splits standart.txt into multiple substrings by enters
	input = strings.ReplaceAll(input, "\\n", string('\n')) // replace occurrences of the "\\n" with the newline character '\n'
	splittedArr := strings.Split(input, string('\n'))

	for i := ' '; i <= '~'; i++ {
		var strs []string
		for j := 0; j < 8; j++ {
			res := (int(i-' ') * 9) + j + 1
			strs = append(strs, sliceData[res])
		}
		elemsMap[i] = strs
	}

	res := ""
	if check.Valid(splittedArr) {
		for _, el := range splittedArr {
			if len(el) > 0 {
				line := getAsciiArtLine(el, elemsMap)
				res += line
			} else {
				res += string('\n')
			}
		}
	} else {
		for i := 0; i < len(splittedArr)-1; i++ { // handling empty input
			res = "\n" + res
		}
	}

	fmt.Print(res)
	fileName := strings.TrimLeft(os.Args[1], "--output=")
	file, err := os.Create(fileName)
	if err != nil {
		return
	}

	file.WriteString(res)
}

func getAsciiArtLine(str string, mapOfEl map[rune][]string) string {
	res := ""
	for i := 0; i < 8; i++ {
		for _, elem := range str {
			res += mapOfEl[elem][i]
		}
		res += "\n"
	}
	return res
}
