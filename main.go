package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func readFile(fn string) []string {
	r := make([]string, 0)
	file, err := os.Open(fn)
	if err != nil {
		return r
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		r = append(r, line)
	}

	return r
}

func writeFile(fn string, lines []string) {
	file, err := os.Create(fn)
	if err != nil {
		return
	}

	defer file.Close()

	for _, line := range lines {
		file.WriteString(line + "\n")
	}
}

func parseStringNames(lines []string) []string {
	r := make([]string, 0)
	for _, line := range lines {
		datas := strings.Split(line, "\">")
		s := strings.Replace(datas[0], "<string name=\"", "", 1)
		r = append(r, s)
	}
	return r
}

func parseNamesStrings(lines []string) map[string]string {
	r := make(map[string]string, 0)
	for _, line := range lines {
		datas := strings.Split(line, "\">")
		s := strings.Replace(datas[0], "<string name=\"", "", 1)
		r[s] = line
	}
	return r
}

func main() {
	fn1 := flag.String("fn1", "fn1", "first file name")
	fn2 := flag.String("fn2", "fn2", "second file name")
	of := flag.String("of", "of", "output file name")

	flag.Parse()

	data1 := readFile(*fn1)
	data2 := readFile(*fn2)

	names := parseStringNames(data1)
	for index, name := range names {
		fmt.Println(index, name)
	}

	names_strings := parseNamesStrings(data2)

	rdata := make([]string, 0)
	for _, name := range names {
		v, ok := names_strings[name]
		if ok {
			rdata = append(rdata, v)
		} else {
			fmt.Println("not found", name)
		}

	}

	writeFile(*of, rdata)
}
