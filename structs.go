package main

import "encoding/xml"

// Config represents the config data.
type Config struct {
	Address string `json:"address"`
}

// RSS holds our RSS data in a simple struct.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel defines the RSS data for the channel itself.
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Copyright   string `xml:"copyright"`
	Image       Image  `xml:"image"`
	Item        Item   `xml:"item"`
}

// Image gives the Channel a logo.
type Image struct {
	URL   string `xml:"url"`
	Title string `xml:"title"`
}

// Item is the actual data for our Channel.
type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	PubDate   string `xml:"pubdate"`
	Enclosure struct {
		URL  string `xml:"url,attr"`
		Type string `xml:"type,attr"`
	} `xml:"enclosure"`
}
