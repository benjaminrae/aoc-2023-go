package pkg

import (
	"os"
)

func GetFileContents(filePath string) string {
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}
