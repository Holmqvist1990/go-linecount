package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
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

	extensions = readExtensions()
	skips = readSkips()

	for _, s := range logo {
		fmt.Println(s)
	}

	fmt.Println("\nRecursively reading all folders:")

	r := ResultFromPaths(*path)
	r.Print()

	return nil
}

func readExtensions() []string {
	b, err := os.ReadFile("./extensions.txt")
	if err != nil {
		panic(fmt.Sprintf("readExtensions: %v", err))
	}
	return strings.Split(string(b), "\r")
}

func readSkips() []string {
	b, err := os.ReadFile("./skip.txt")
	if err != nil {
		panic(fmt.Sprintf("readSkips: %v", err))
	}
	return strings.Split(string(b), "\r")
}

var (
	extensions = []string{}
	skips      = []string{}
)

var logo = []string{
	"",
	" ██████   ██████        ██      ██ ███    ██ ███████  ██████  ██████  ██    ██ ███    ██ ████████ ██ ",
	"██       ██    ██       ██      ██ ████   ██ ██      ██      ██    ██ ██    ██ ████   ██    ██    ██ ",
	"██   ███ ██    ██ █████ ██      ██ ██ ██  ██ █████   ██      ██    ██ ██    ██ ██ ██  ██    ██    ██ ",
	"██    ██ ██    ██       ██      ██ ██  ██ ██ ██      ██      ██    ██ ██    ██ ██  ██ ██    ██       ",
	" ██████   ██████        ███████ ██ ██   ████ ███████  ██████  ██████   ██████  ██   ████    ██    ██ ",
}
