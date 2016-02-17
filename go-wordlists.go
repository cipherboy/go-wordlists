package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage:", os.Args[0], " /path/to/wordlist\n\nDescription:\nBasic word list generator taking the original input and adding numbers of\nlengths 1-4 (0-9, 00-99, 000-999, and 0000-9999) and writing to STDOUT using\nbuffio (performance reasons).")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	for scanner.Scan() {
		line := scanner.Text()
		f.Write([]byte(fmt.Sprintf("%s\n", line)))
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 10; k++ {
					for l := 0; l < 10; l++ {
						f.Write([]byte(fmt.Sprintf("%s%d%d%d%d\n", line, i, j, k, l)))
					}
					f.Write([]byte(fmt.Sprintf("%s%d%d%d\n", line, i, j, k)))
				}
				f.Write([]byte(fmt.Sprintf("%s%d%d\n", line, i, j)))
			}
			f.Write([]byte(fmt.Sprintf("%s%d\n", line, i)))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}
}
