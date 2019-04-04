package main

import "testing"

func TestFind(t *testing.T) {
	expected := []string{
		"testcase",
		"testcase/file1",
		"testcase/file2",
		"testcase/folder1",
		"testcase/folder1/file3",
		"testcase/folder1/file4",
		"testcase/folder1/folder2",
		"testcase/folder1/folder2/file5",
		"testcase/folder1/folder2/file6",
		"testcase/folder3",
		"testcase/folder3/file7",
		"testcase/folder3/file8",
	}
	actual := find("testcase")
	if len(expected) != len(actual) {
		t.Errorf("Invalid actual length: %v", actual)
		return
	}
	for i := 0; i < len(actual); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Invalid match at %d : %s - %v", i, expected[i], actual[i])
			return
		}
	}
}
