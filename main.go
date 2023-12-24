package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/logrusorgru/aurora"
)

// DELIMITER holds a string to delimition.
const DELIMITER = " "

func main() {
	delimFlag := flag.String(
		"delim",
		",",
		"-delim='|'",
	)

	flag.Parse()

	fp := readStdin(flag.Args())

	delim := []rune(*delimFlag)
	if len(delim) > 1 {
		panic("delimeter must be one character")
	}

	reader := csv.NewReader(fp)
	reader.Comma = delim[0]
	reader.LazyQuotes = true

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for i, col := range record {
			rotateColor(i, col)
		}
		fmt.Println()
	}
}

// rotateColor exports the colored columns with circuration.
func rotateColor(i int, col interface{}) {
	switch i % 7 {
	case 0:
		fmt.Print(aurora.Magenta(col), DELIMITER)
	case 1:
		fmt.Print(aurora.Blue(col), DELIMITER)
	case 2:
		fmt.Print(aurora.Brown(col), DELIMITER)
	case 3:
		fmt.Print(aurora.Green(col), DELIMITER)
	case 4:
		fmt.Print(aurora.White(col), DELIMITER)
	case 5:
		fmt.Print(aurora.Cyan(col), DELIMITER)
	case 6:
		fmt.Print(aurora.Red(col), DELIMITER)
	}
	return
}

// readStdin returns the CSV from stdin.
func readStdin(args []string) *os.File {
	var fp *os.File
	if len(args) < 2 {
		fp = os.Stdin
	} else {
		var err error
		fp, err = os.Open(args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}
	return fp
}
