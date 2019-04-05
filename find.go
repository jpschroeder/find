package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type options struct {
	name      string
	followSym bool
	empty     bool
}

func find(opt options, root string) []string {
	var ret []string
	if len(root) < 1 {
		return ret
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("failure accessing a path %q: %v\n", path, err)
			return err
		}

		skip := false

		if opt.empty {
			if info.IsDir() {
				isEmpty, _ := isDirEmpty(path)
				if isEmpty {
					skip = true
				}
			} else {
				if info.Size() > 0 {
					skip = true
				}
			}
		}

		if opt.name != "" {
			match, _ := filepath.Match(opt.name, info.Name())
			if !match {
				skip = true
			}
		}

		if opt.followSym {
			sympath, _ := filepath.EvalSymlinks(path)
			if filepath.Clean(path) != sympath {
				skip = true
				symret := find(opt, sympath)
				for _, s := range symret {
					actual := strings.Replace(s, sympath, path, 1)
					ret = append(ret, addPrefix(root, actual))
				}
			}
		}

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

func isDirEmpty(dirname string) (bool, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return false, err
	}
	names, err := f.Readdirnames(-1)
	f.Close()
	if err != nil {
		return false, err
	}
	return len(names) > 0, nil
}

func main() {
	// todo
	// error handling
	// refactoring
	// testing

	// In the linux version of the find command the name and empty flags come after the starting directory
	// The prompt listed them before the starting directory.  That is how it is implemented here.
	name := flag.String("name", "", "Base of file name (the path with the  leading  directories  removed)  matches  shell  pattern  pattern.")
	followSym := flag.Bool("L", false, "Follow symbolic links.")
	empty := flag.Bool("empty", false, "File is empty and is either a regular file or a directory.")

	flag.Parse()

	opt := options{
		name:      *name,
		followSym: *followSym,
		empty:     *empty,
	}

	args := flag.Args()

	root := filepath.FromSlash("./")
	if len(args) > 0 {
		root = args[0]
	}

	output := find(opt, root)
	for _, entry := range output {
		fmt.Println(entry)
	}
}
