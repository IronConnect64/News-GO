//		Copyright (c) 2019 IronConnect64
//
//		News-GO, a RSS Feed server for the PSP News Channel, which will make it very easy to create and serve PSP-friendly RSS files.
//
//	Permission is hereby granted, free of charge, to any person obtaining a copy
//	of this software and associated documentation files (the "Software"), to deal
//	in the Software without restriction, including without limitation the rights
//	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//	copies of the Software, and to permit persons to whom the Software is
//	furnished to do so, subject to the following conditions:
//
//	The above copyright notice and this permission notice shall be included in all
//	copies or substantial portions of the Software.
//
//	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//	SOFTWARE.

package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Our URL list, held very tight, uwu.
var (
	urls map[string]string
	port string
)

// Our RSS Data, extra friendly for your PSP.
type rss struct {
	XMLName xml.Name `xml:"rss"`

	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Copyright   string `xml:"copyright"`

		Image struct {
			URL   string `xml:"url"`
			Title string `xml:"title"`
		} `xml:"image"`

		Item []struct {
			Title     string `xml:"title"`
			Link      string `xml:"link"`
			PubDate   string `xml:"pubdate"`
			Enclosure struct {
				URL  string `xml:"url,attr"`
				Type string `xml:"type,attr"`
			} `xml:"enclosure"`
		} `xml:"item"`
	} `xml:"channel"`
}

func main() {
	log.Println("News-GO - 1.2.1")

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Println("No port providen; using port 50052.")
		port = "50052"
	} else {
		port = os.Args[1]
	}

	// Let's make an gin engine for our server, and fetch all URLs.
	log.Println("Initializing server...")
	server := gin.Default()

	// URL list.
	urls = make(map[string]string)

	// All RSS URLs from Reuters, held in a nice and simple array.
	urls["arts"] = "http://feeds.reuters.com/news/artsculture"
	urls["business"] = "http://feeds.reuters.com/reuters/businessNews"
	urls["company"] = "http://feeds.reuters.com/reuters/companyNews"
	urls["entertainment"] = "http://feeds.reuters.com/reuters/entertainment"
	urls["environment"] = "http://feeds.reuters.com/reuters/environment"
	urls["health"] = "http://feeds.reuters.com/reuters/healthNews"
	urls["lifestyle"] = "http://feeds.reuters.com/reuters/lifestyle"
	urls["money"] = "http://feeds.reuters.com/reuters/wealth"
	urls["oddlyEnough"] = "http://feeds.reuters.com/reuters/oddlyEnoughNews"
	urls["pictures"] = "http://feeds.reuters.com/reuters/ReutersPictures"
	urls["people"] = "http://feeds.reuters.com/reuters/peopleNews"
	urls["politics"] = "http://feeds.reuters.com/reuters/PoliticsNews"
	urls["science"] = "http://feeds.reuters.com/reuters/scienceNews"
	urls["sports"] = "http://feeds.reuters.com/reuters/sportsNews"
	urls["technology"] = "http://feeds.reuters.com/reuters/technologyNews"
	urls["top"] = "http://feeds.reuters.com/reuters/topNews"
	urls["us"] = "http://feeds.reuters.com/reuters/domesticNews"
	urls["world"] = "http://feeds.reuters.com/reuters/worldNews"

	// Our request handlers.
	// This is just so you feel nicer.
	server.GET("/", func(ctx *gin.Context) {
		fmt.Fprintln(ctx.Writer, "Hey, I'm News-GO, how are you today?")
	})

	server.POST("/", func(ctx *gin.Context) {
		// Simple check if you're using a PSP, or CURL, you silly user ;3
		if ctx.GetHeader("HTTP_X_PSP_BROWSER") == "" {
			ctx.AbortWithStatus(http.StatusUnavailableForLegalReasons)

			// Actual stuff.
		} else if getURL(ctx.PostForm("topic")) {
			resp, err := http.Get(urls[ctx.PostForm("topic")])
			if err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}

			var xmlData rss
			xml.NewDecoder(resp.Body).Decode(&xmlData)

			// Converts and shortened it, I think.
			xml.NewEncoder(ctx.Writer).Encode(xmlData)

			// Other stuff; Bad request then.
		} else {
			ctx.AbortWithStatus(http.StatusBadRequest)
		}
	})

	// If possible, we can replace this with RunTLS() in the future.
	log.Println("Starting server...")
	server.Run(":" + port)
}

// Simple check if we have a proper topic.
func getURL(s string) bool {
	if s == "" {
		return false
	}

	for _, v := range urls {
		if v == s {
			return true
		}
	}
	return false
}
