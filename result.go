package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Result struct {
	TotalFiles     int
	TotalLines     int
	TotalChars     int
	FileExtensions map[string]bool
}

func ResultFromPaths(path string) *Result {
	r := &Result{FileExtensions: make(map[string]bool)}

	err := filepath.Walk(path, walker(r))
	if err != nil {
		log.Fatal(err)
	}

	return r
}

func (r Result) Print() {
	fmt.Println()
	fmt.Printf("Files: %d.\n", r.TotalFiles)
	fmt.Printf("Lines of code: %d.\n", r.TotalLines)
	fmt.Printf("Characters: %d.\n", r.TotalChars)
	fmt.Printf("Average characters per line: %d.\n", r.averageCharacters())
	if len(r.FileExtensions) == 0 {
		fmt.Println("No extensions.")
		return
	}
	fmt.Println("File extensions:")
	for k := range r.FileExtensions {
		fmt.Printf("\t* %s\n", k)
	}
}

func (r Result) averageCharacters() int {
	average := 0
	if r.TotalChars > 0 && r.TotalLines > 0 {
		average = r.TotalChars / r.TotalLines
	}
	return average
}

func walker(r *Result) func(path string, info os.FileInfo, err error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if skipPath(path) {
			return nil
		}
		if skipFile(info) {
			return nil
		}
		fmt.Printf("\t%v\n", path)

		r.TotalFiles++
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		lines := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")
		r.TotalLines += len(lines)
		for _, line := range lines {
			r.TotalChars += len(line)
		}

		parts := strings.Split(path, ".")
		r.FileExtensions[parts[len(parts)-1]] = true
		return nil
	}
}

func skipPath(path string) bool {
	for _, s := range skips {
		if strings.Contains(path, s) {
			return true
		}
	}
	return false
}

func skipFile(info os.FileInfo) bool {
	if info.IsDir() {
		return true
	}
	for _, ext := range extensions {
		parts := strings.Split(info.Name(), ".")
		if len(parts) < 2 {
			return true
		}
		fileExt := parts[len(parts)-1]
		if fileExt == ext {
			return false
		}
	}
	return true
}
