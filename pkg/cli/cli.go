package cli

import (
	"flag"
	"fmt"
	"os"
)

func GetArgs() (string, int) {
	density := flag.Int("density", 5, "Specify density")
	flag.Parse()

	fmt.Println(*density)

	if len(flag.Args()) < 1 {
		fmt.Println("Please provide an image name")
		os.Exit(1)
	}
	imageName := flag.Arg(0)

	imageDirPath, err := os.Getwd()
	if err != nil {
		panic("Error accured while getting image path")
	}

	imageFullPath := fmt.Sprintf("%s/%s", imageDirPath, imageName)

	return imageFullPath, *density
}
