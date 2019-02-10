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
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("News-GO - 1.1.0")

	// TODO: implementing RSS parsing

	// Let's make an gin engine for our server, and fetch all URLs.
	log.Println("Initializing server...")
	server := gin.Default()
	getURLs()

	// Our request handlers.
	server.POST("/news", newsHandler)

	// If possible, we can replace this with RunTLS() in the future.
	log.Println("Starting server...")
	server.Run()
}

func newsHandler(ctx *gin.Context) {

	if getURL(ctx.PostForm("topic")) {
		parseToPSP(urls[ctx.PostForm("topic")], ctx)
	} else {
		ctx.Redirect(http.StatusBadRequest, "What are you even doing here?") // maybe I'll add some weird HTML data
	}
}

func parseToPSP(url string, ctx *gin.Context) {
	resp, err := http.Get(url)
	if err != nil {
		ctx.Redirect(http.StatusInternalServerError, "Couldn't fetch XML data.")
	}

	var xmlData RSS
	xml.NewDecoder(resp.Body).Decode(&xmlData)

	// Converts and shortened it, I think
	xml.NewEncoder(ctx.Writer).Encode(xmlData)
}
