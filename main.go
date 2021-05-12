package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	path := flag.String("path", "", "root project folder")
	flag.Parse()
	if *path == "" {
		return fmt.Errorf("path cannot be empty")
	}

	for _, s := range logo {
		fmt.Println(s)
	}

	fmt.Println("\nRecursively reading all folders:")

	r := ResultFromPaths(*path)
	r.Print()

	return nil
}

var extensions = []string{
	"go",
	"cs",
	"js",
	"swift",
	"java",
	"json",
	"ruby",
	"md",
	"strings",
	"html",
	"css",
}

var skips = []string{
	".gitignore",
	"node_modules",
	"build",
	"package.json",
	"package-lock.json",
	"netcoreapp3.1",
	".min.js",
}

var logo = []string{
	"",
	" ██████   ██████        ██      ██ ███    ██ ███████  ██████  ██████  ██    ██ ███    ██ ████████ ██ ",
	"██       ██    ██       ██      ██ ████   ██ ██      ██      ██    ██ ██    ██ ████   ██    ██    ██ ",
	"██   ███ ██    ██ █████ ██      ██ ██ ██  ██ █████   ██      ██    ██ ██    ██ ██ ██  ██    ██    ██ ",
	"██    ██ ██    ██       ██      ██ ██  ██ ██ ██      ██      ██    ██ ██    ██ ██  ██ ██    ██       ",
	" ██████   ██████        ███████ ██ ██   ████ ███████  ██████  ██████   ██████  ██   ████    ██    ██ ",
}
