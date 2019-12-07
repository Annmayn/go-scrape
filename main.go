package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	visitURL := "https://www.bea-brak.de/bravsearch/search.brak"

	c := colly.NewCollector()

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		if e.Attr("id") == "searchForm:imgCaptcha" {
			fullURL := e.Request.AbsoluteURL(e.Attr("src"))
			fmt.Println("Full url: ", fullURL)

			reader := bufio.NewReader(os.Stdin)
			captchaSolved, _ := reader.ReadString('\n')

			data := url.Values{
				"searchForm:txtPostal":  {"01099"},
				"searchForm:txtCaptcha": {captchaSolved},
			}

			fmt.Println("1")
			client := http.Client{
				Timeout: 5 * time.Second,
			}
			res, e := client.PostForm(visitURL, data)
			fmt.Println("2")
			if e != nil {
				fmt.Println("Error: ", e)
				return
			}
			var result map[string]interface{}
			json.NewDecoder(res.Body).Decode(&result)
			fmt.Println("Form: ", result["form"])
			fmt.Println("Result: ", result)
			// fmt.Println(res.Status)
			// fmt.Println(res.Header)
			// fmt.Println(res.Body)

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting...", r.URL)
	})

	c.Visit(visitURL)
}
