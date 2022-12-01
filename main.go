package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vikashparashar/golang_project_03/cli"
)

var (
	sourceLang string
	targetLang string
	sourceText string
)

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Convert From [english]")
	flag.StringVar(&targetLang, "t", "fr", "Convert Into [french]")
	flag.StringVar(&sourceText, "st", "", "Text That Need To Be Converted Into Target Language")

}

func main() {
	flag.Parse()

	// is user dones not enter any flag , then show them options available for flags
	// and exit the program immediatly
	if flag.NArg() == 0 {
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var reqBody = &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	cli.RequestTranslate(reqBody)
}
