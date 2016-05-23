package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	ch <- fmt.Sprintf("%s", body)
}

func main() {

	ch := make(chan string)

	for _, arg := range os.Args[1:] {
		go fetch(arg, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

}
