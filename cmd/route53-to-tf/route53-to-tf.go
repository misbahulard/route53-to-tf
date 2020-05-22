package main

import (
	"fmt"
	"os"
	"route53-to-tf/pkg/route53-to-tf/cmd"

	flag "github.com/spf13/pflag"
)

func main() {
	var source *string = flag.StringP("source", "s", "", "source file path")
	var output *string = flag.StringP("output", "o", "", "output file path")

	flag.Parse()

	if *source == "" {
		fmt.Println("please add source file")
		os.Exit(0)
	}

	cmd.Run(*source, *output)
}
