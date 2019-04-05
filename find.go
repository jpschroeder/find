package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func find(empty bool, root string) []string {
	var ret []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("failure accessing a path %q: %v\n", path, err)
			return err
		}

		skip := empty && info.Size() > 0
		if skip {
			return nil
		}
		ret = append(ret, addPrefix(root, path))

		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", root, err)
		return ret
	}
	return ret
}

func addPrefix(root string, path string) string {
	if strings.HasPrefix(root, "./") && !strings.HasPrefix(path, "./") {
		return "./" + path
	}
	if path == "./" {
		return "."
	}
	return path
}

func main() {
	// In the linux version of the find command the -empty flag comes after the starting directory
	// The prompt listed it before the starting directory.  That is how it is implemented here.
	empty := flag.Bool("empty", false, "File is empty and is either a regular file or a directory.")

	flag.Parse()

	args := flag.Args()

	root := filepath.FromSlash("./")
	if len(args) > 0 {
		root = args[0]
	}

	output := find(*empty, root)
	for _, entry := range output {
		fmt.Println(entry)
	}
}
