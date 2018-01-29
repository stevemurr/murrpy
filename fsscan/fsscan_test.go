package fsscan

import "testing"

func TestScan(t *testing.T) {
	files := Scan("../local")
	for _, file := range files {
		t.Log(file)
	}
}
