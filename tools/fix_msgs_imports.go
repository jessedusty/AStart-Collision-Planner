package main

import (
	"fmt"
	"github.com/dave/dst/decorator"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func contains(s []string, str string) bool {
	for _, this := range s {
		if this == str {
			return true
		}
	}
	return false
}

func fixFile(fileName string, packages []string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	f, err := decorator.Parse(file)
	if err != nil {
		panic(err)
	}

	for _, n := range f.Imports {
		if n.Path.Kind == token.STRING {
			if contains(packages, n.Path.Value) {
				fmt.Println("   ", n.Path.Value)
				n.Path.Value = fmt.Sprintf("\"ee631_midterm/msgs/%s\"", n.Path.Value[1:len(n.Path.Value)-1])
			}
		}
	}

	outFile, err := os.OpenFile(fileName, os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	decorator.Fprint(outFile, f)

}

func main() {

	dir, err := ioutil.ReadDir("msgs")
	if err != nil {
		panic(err)
	}

	var packages []string
	for _, file := range dir {
		if file.IsDir() {
			packages = append(packages, "\""+file.Name()+"\"")
		}
	}

	filepath.Walk("msgs", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".go") {
			fmt.Println(path)
			fixFile(path, packages)
		}

		return nil
	})
}
