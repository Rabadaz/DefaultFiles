package main

import (
	"fmt"
	"log"
	"os"
	Config "rabadaz.at/defaultfiles/Config"
	"rabadaz.at/defaultfiles/FileParamCLI"
	"rabadaz.at/defaultfiles/ResourceLoader"
)

func main() {
	//Create main Logger for Application
	mainLogger := log.New(os.Stderr, "[Defaultfiles::MainLog]\t", 0)

	//Load Remote Config from Github
	remoteConfig, err := Config.LoadRemoteConfig()

	if err != nil {
		mainLogger.Println("Cannot download RemoteConfig")
		os.Exit(-1)
	}

	//Load Config from CLI Arguments
	cliConfig := Config.LoadCliConfig()

	//Check if Tag is specified
	if *cliConfig.Tag == ""{
		mainLogger.Println("Tag not given")
		os.Exit(-1)
	}

	//Finding File in Remote Repo
	file := Config.FindFileInRemoteConfig(remoteConfig, *cliConfig.Tag, *cliConfig.Tech)
	if file == nil{
		mainLogger.Println("Tag not found in Remote Config")
		os.Exit(-1)
	}
	fmt.Println("Using File", file.Tag)

	pa := FileParamCLI.DecideParams(file.Params)

	downloadError := ResourceLoader.DownloadResource(*file,*cliConfig.OutputDirectory,pa)

	if downloadError != nil{
		mainLogger.Fatal("Could not download Resource")
	}


}


