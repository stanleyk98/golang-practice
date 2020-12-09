package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
)

func findIP(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"            // IP 값은 255.249.199.?99 이내에서 해결
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP // IP 각 4부분
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			//fmt.Printf("%d\t%s\n", n, line)
			ip := findIP(line)
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}
			fmt.Println(ip)

		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
