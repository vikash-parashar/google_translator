package cli

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const (
	// url  = "https://translate.google.co.in/?sl=en&tl=fr&text=sdfhdfsdsdhsd&op=translate"
	// url2 = "https://translate.google.co.in/?sl=en&tl=fr&op=translate"
	// url3 = "https://translate.googleapis.com"
	url4 = "https://translate.googleapis.com/translate_a/single"
)

func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup) {
	client := http.Client{}

	req, err := http.NewRequest("GET", url4, nil)

	query := req.URL.Query()

	query.Add("client", "gtx")
	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)
	req.URL.RawQuery = query.Encode()

	if err != nil {
		log.Fatalf(" error while making request to url : %s", err)
	}
	// client.Do(req)
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Println("failed to read the response body")
	// }
	// strChan <- string(data)
	defer res.Body.Close()

	if res.StatusCode == http.StatusTooManyRequests {
		str <- "You Have Been Rate Limited , Try Again"
		wg.Done()
		return
	}

	parsedJSON, err := gabs.ParseJSONBuffer(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	nestOne, err := parsedJSON.ArrayElement(0)
	if err != nil {
		log.Fatalln(err)
	}
	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatalln(err)
	}
	translatedString, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatalln(err)
	}
	str <- translatedString.Data().(string)
}
