package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func find(root string) []string {
	var ret []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !strings.HasPrefix(path, root) {
			path = root + path
		}
		ret = append(ret, path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", root, err)
		return ret
	}
	return ret
}

func main() {
	output := find("./")
	for _, entry := range output {
		fmt.Println(entry)
	}
}
