package ResourceLoader

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"rabadaz.at/defaultfiles/Config"
	"rabadaz.at/defaultfiles/FileParamCLI"
	"rabadaz.at/defaultfiles/Utils"
	"strings"
)




func DownloadResource(resource Config.FileConfig, outputDir string, params []FileParamCLI.ReplacementParameters) error{
	resp, err := http.Get(Utils.ConstructURL(resource.Path))

	if err != nil{
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200{
		return errors.New("Status Code of resource was not 200-OK ")
	}

	rawResource, convertErr := ioutil.ReadAll(resp.Body)

	if convertErr != nil{
		return convertErr
	}
	sResource := string(rawResource)

	sResource = replaceParams(sResource, params)

	path := constructNewFilePath(outputDir, resource)

	nFile, fErr := os.Create(path)
	if fErr != nil{
		return fErr
	}
	defer nFile.Close()

	_, err = nFile.WriteString(sResource)
	if err != nil {
		return err
	}

	return nil
}

func constructNewFilePath(outputDir string, resource Config.FileConfig) string{
	fName := filepath.Base(resource.Path)

	if resource.RenameTo != nil{
		fName = *resource.RenameTo
	}

	path := fmt.Sprintf("%s%s", outputDir, fName )

	if !strings.HasSuffix(outputDir, "/"){
		path = fmt.Sprintf("%s/%s", outputDir, fName )
	}
	return path
}

func replaceParams(resource string, params []FileParamCLI.ReplacementParameters,)string{
	for _, param := range params{
		fullKey := fmt.Sprintf("$$%s$$", param.Parameter.Key)

		resource = strings.Replace(resource, fullKey, param.Replacement,-1)
	}
	return resource
}
