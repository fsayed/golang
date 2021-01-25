package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

/*
func reportPanic() {
	p := recover()

	if p == nil {
		return
	}
	
	err, ok := p.(error)
	
	if err != nil {
		fmt.Println(err)
	} else {
		panic(err)
	}
}

*/

type Page struct {
	URL string
	Size int
}

func coolFunc(channel chan Page, urls []string) {
	var currentPage Page
	for _, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		size := len(body)
		currentPage.URL = url
		currentPage.Size = size
		
		channel <- currentPage
		//channel <- Page{URL: url, Size: size}
	}


}


func main() {
	channel := make(chan Page)
	urls := []string {"http://www.google.ca", "https://www.tsn.ca", "https://en.wikipedia.com" }
	
	
	go coolFunc(channel, urls)
	//fmt.Println(<- channel)
	page := <-channel
	fmt.Println("page type is: ", reflect.TypeOf(page))
	fmt.Printf("Page: %s, Size: %d\n", page.URL, page.Size)
	fmt.Println(<- channel)
	fmt.Println(<- channel)
}
