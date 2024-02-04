package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file specified")
		return
	}

	filename := os.Args[1]

	swapColorsInFile(filename)
}

func swapColorsInFile(filename string) {
	fileBody, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	hexColorRegexPattern := `#([A-Fa-f0-9]{8}|[A-Fa-f0-9]{6}|[A-Fa-f0-9]{4}|[A-Fa-f0-9]{3})`
	hexColorRegex, err := regexp.Compile(hexColorRegexPattern)
	if err != nil {
		panic(err)
	}

	strFileBody := string(fileBody)

	// extract all hex codes from file with regex
	matchedHexColors := hexColorRegex.FindAllString(string(fileBody), -1)

	// convert hex codes to rgb
	for _, hexColor := range matchedHexColors {
		r, g, b, alpha := hexToRGB(hexColor)
		// swap colors
		formattedColor := formatCssProperty(r, g, b, alpha)
		strFileBody = strings.ReplaceAll(strFileBody, hexColor, formattedColor)
	}

	// write the new file with the swapped colors
	os.WriteFile(fmt.Sprintf("%s_result.css", filename), []byte(strFileBody), 0644)
}

func hexToRGB(hex string) (int, int, int, float64) {
	if string(hex[0]) == "#" {
		hex = hex[1:]
	}
	// check if hex is length 3 or 6
	var r, g, b, alpha string
	switch len(hex) {
	case 3:
		r = string(hex[0]) + string(hex[0])
		g = string(hex[1]) + string(hex[1])
		b = string(hex[2]) + string(hex[2])
	case 4:
		r = string(hex[0]) + string(hex[0])
		g = string(hex[1]) + string(hex[1])
		b = string(hex[2]) + string(hex[2])
		alpha = string(hex[3]) + string(hex[3])
	case 6:
		r = string(hex[0]) + string(hex[1])
		g = string(hex[2]) + string(hex[3])
		b = string(hex[4]) + string(hex[5])
	case 8:
		r = string(hex[0]) + string(hex[1])
		g = string(hex[2]) + string(hex[3])
		b = string(hex[4]) + string(hex[5])
		alpha = string(hex[6]) + string(hex[7])
	default:
		panic("invalid hex code")
	}

	var parsedAlpha float64
	if alpha != "" {
		parsedAlpha = float64(hexToDecimal(alpha)) / 255
	} else {
		parsedAlpha = 1
	}

	return hexToDecimal(r), hexToDecimal(g), hexToDecimal(b), parsedAlpha
}

var hexMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"a": 10,
	"b": 11,
	"c": 12,
	"d": 13,
	"e": 14,
	"f": 15,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}

func hexToDecimal(hex string) int {
	var res int

	fmt.Println(hex)

	if len(hex) == 1 {
		res = hexMap[hex]
	} else {
		firstVal := hexMap[string(hex[0])]
		shiftedVal := firstVal << 4
		res = shiftedVal + hexMap[string(hex[1])]
	}

	return res
}

func formatCssProperty(r, g, b int, alpha float64) string {
	if alpha == 1 {
		return fmt.Sprintf("rgb(%d %d %d)", r, g, b)
	} else {
		return fmt.Sprintf("rgba(%d %d %d / %v)", r, g, b, strconv.FormatFloat(alpha, 'f', 5, 64))
	}
}
