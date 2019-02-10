package main

var urls map[string]string

func getURLs() {
	urls = make(map[string]string)

	// All RSS URLs from Reuters, held in a nice, simple array.
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
}

func getURL(s string) bool {
	for _, v := range urls {
		if v == s {
			return true
		}
	}
	return false
}
