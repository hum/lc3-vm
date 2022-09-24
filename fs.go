package main

import "os"

func ReadObjFile(fp string) ([]byte, int, error) {
	data, err := os.ReadFile(fp)
	if err != nil {
		return data, -1, err
	}
	// Size of the input
	var n int = len(data)
	return data, n, nil
}
