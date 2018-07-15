package main

import (
	"flag"
	"os"
	"bufio"
	"github.com/bountylabs/service/api_common/language/unicode"
	"strings"
	"fmt"
)

var in = flag.String("in", "data/en", "--path path to file to normalize")
var out = flag.String("out", "data/en.out", "--path path to file to normalize")

func main() {

	flag.Parse()

	file, err := os.Open(*in)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	out, err := os.Create(*out)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		parts := strings.Split(strings.TrimSpace(scanner.Text()), "\t")
		if len(parts) != 2 {
			panic("malformed")
		}

		base := unicode.MustNormalizeUnicode(strings.ToLower(parts[0]))
		form := unicode.MustNormalizeUnicode(strings.ToLower(parts[1]))

		_, err := out.WriteString(fmt.Sprintf("%s\t%s\n", base, form))
		if err != nil {
			panic(err)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}


}
