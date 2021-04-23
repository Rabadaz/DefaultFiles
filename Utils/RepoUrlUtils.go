package Utils

import (
	"fmt"
	"strings"
)

const MainRepositoryURL = "https://raw.githubusercontent.com/Rabadaz/DefaultFilesResource/main/"

func ConstructURL(relativePath string) string{
	if strings.HasPrefix(relativePath, "./" ) {
		relativePath = relativePath[2:]
	}
	if strings.HasPrefix(relativePath, "/"){
		relativePath = relativePath[1:]
	}

	return fmt.Sprintf("%s%s", MainRepositoryURL, relativePath)
}