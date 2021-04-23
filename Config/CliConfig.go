package Config

import "flag"

type CliConfig struct{
	Tag *string
	Version *int64
	OutputDirectory *string
	Tech *string
	Tail []string
}


func LoadCliConfig() CliConfig{
	tag := flag.String("tag", "", "Tag of the File you want to download")
	version := flag.Int64("version", 0, "The Version of the File you want to download")
	outputDir := flag.String("outputDir", ".", "The Directory the File should be downloaded to")
	tech := flag.String("tech", "docker", "The Technology that the File is for")

	flag.Parse()


	return CliConfig{
		Tag:             	tag,
		Version:         	version,
		OutputDirectory: 	outputDir,
		Tech:				tech,
		Tail:				flag.Args(),
	}
}

