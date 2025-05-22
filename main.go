package main

import (
	"fmt"
	"os"
	"strings"

	"bufio"
)

func AppendFileExtension(path string) string {
	if !strings.Contains(path, "xlsx") {
		path += ".xlsx"
	}

	return path
}

func FileNamePrompt(path string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, path+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}

	return AppendFileExtension(strings.TrimSpace(s))
}

func main() {
	filename := FileNamePrompt("What is the filename for this Excel file?")

	if err := GenerateExcel(filename); err != nil {
		fmt.Fprintln(os.Stderr, "%w", err)
		os.Exit(1)
	}

	fmt.Println("Excel file created:", filename)
}
