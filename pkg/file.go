package pkg

import (
	"os"
	"path"
)

func GetInputFileContents(filename string) string {
	var inputDirectory = "../input"

	var path = path.Join(inputDirectory, filename)

	bytes, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}
