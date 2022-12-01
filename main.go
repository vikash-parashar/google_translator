package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/vikashparashar/golang_project_03/cli"
)

var wg = sync.WaitGroup{}
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
	strChan := make(chan string)

	wg.Add(1)
	var reqBody = &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	go cli.RequestTranslate(reqBody, strChan, &wg)
	processedString := strings.ReplaceAll(<-strChan, "+", " ")
	// res := <-strChan
	// fmt.Println(res)
	fmt.Printf("%s", processedString)
	close(strChan)
	wg.Wait()
}


// terminal cmd to run the program

/*
go run -s en -st hello -t fn

// Output : Bonjour
*/