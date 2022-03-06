package url_titles

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

func Titles(urls ...string) <-chan string {
	ch := make(chan string)

	for _, url := range urls {
		go func (url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			ch <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return ch
}
