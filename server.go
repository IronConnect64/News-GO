//	Copyright (c) 2019 PSPConnect64
//
//	News-GO, a RSS Feed server for the PSP News Channel, which will make it very easy to create and serve PSP-friendly RSS files.
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

// If I'm honest - this software is nothing but unnecessary.
// The PSP could do this all by itself and that'll be it, but apparently, I wanted to create this a long time ago.
// I, sometimes, really hate my earlier decisions. Quite "hilarious" to hear that from me as well, right?

package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/firstrow/tcp_server"
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
	log.Println("News-GO - 2.0.0")

	var urls = make(map[string]string)

	urls["arts"] = "http://feeds.reuters.com/news/artsculture"
	urls["business"] = "http://feeds.reuters.com/reuters/businessNews"
	urls["company"] = "http://feeds.reuters.com/reuters/companyNews"
	urls["entertainment"] = "http://feeds.reuters.com/reuters/entertainment"
	urls["environment"] = "http://feeds.reuters.com/reuters/environment"
	urls["health"] = "http://feeds.reuters.com/reuters/healthNews"
	urls["lifestyle"] = "http://feeds.reuters.com/reuters/lifestyle"
	urls["money"] = "http://feeds.reuters.com/reuters/wealth"
	urls["oddly"] = "http://feeds.reuters.com/reuters/oddlyEnoughNews"
	urls["pictures"] = "http://feeds.reuters.com/reuters/ReutersPictures"
	urls["people"] = "http://feeds.reuters.com/reuters/peopleNews"
	urls["politics"] = "http://feeds.reuters.com/reuters/PoliticsNews"
	urls["science"] = "http://feeds.reuters.com/reuters/scienceNews"
	urls["sports"] = "http://feeds.reuters.com/reuters/sportsNews"
	urls["technology"] = "http://feeds.reuters.com/reuters/technologyNews"
	urls["top"] = "http://feeds.reuters.com/reuters/topNews"
	urls["us"] = "http://feeds.reuters.com/reuters/domesticNews"
	urls["world"] = "http://feeds.reuters.com/reuters/worldNews"

	server := tcp_server.New(":50052")

	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		fmt.Println(message)

		check := func(s string) bool { // Is this impressive work?
			for _, v := range urls {
				if v == s {
					return true
				}
			}

			return false
		}

		if !check(message) {
			c.Send("INVALID_TOPIC")
			return
		}

		resp, err := http.Get(urls[message])
		if err != nil {
			c.Send("SERVER_ERROR_1")
			log.Printf("The server experienced an error, as we tried to fetch the data: %s\n", err.Error())
			return
		}

		var xmlData rss
		if err := xml.NewDecoder(resp.Body).Decode(&xmlData); err != nil {
			c.Send("SERVER_ERROR_2")
			log.Printf("The server experienced an error, as we tried to decode the received data: %s\n", err.Error())
			return
		}

		if err := xml.NewEncoder(c.Conn()).Encode(xmlData); err != nil {
			c.Send("SERVER_ERROR_2")
			log.Printf("The server experienced an error, as we tried to encode the received data: %s\n", err.Error())
			return
		}
	})

	server.Listen()
}
