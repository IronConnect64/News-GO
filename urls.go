package main

// GetURLs returns all Reuters RSS URLs.
func GetURLs() map[string]string {
	urls := make(map[string]string)

	// All RSS URLs from Reuters, held in a nice, simple array.
	urls["Arts"] = "http://feeds.reuters.com/news/artsculture"
	urls["Business"] = "http://feeds.reuters.com/reuters/businessNews"
	urls["Company News"] = "http://feeds.reuters.com/reuters/companyNews"
	urls["Entertainment"] = "http://feeds.reuters.com/reuters/entertainment"
	urls["Environment"] = "http://feeds.reuters.com/reuters/environment"
	urls["Health News"] = "http://feeds.reuters.com/reuters/healthNews"
	urls["Lifestyle"] = "http://feeds.reuters.com/reuters/lifestyle"
	urls["Money"] = "http://feeds.reuters.com/reuters/wealth"
	urls["Oddly Enough"] = "http://feeds.reuters.com/reuters/oddlyEnoughNews"
	urls["Pictures"] = "http://feeds.reuters.com/reuters/ReutersPictures"
	urls["People"] = "http://feeds.reuters.com/reuters/peopleNews"
	urls["Politics"] = "http://feeds.reuters.com/reuters/PoliticsNews"
	urls["Science"] = "http://feeds.reuters.com/reuters/scienceNews"
	urls["Sports"] = "http://feeds.reuters.com/reuters/sportsNews"
	urls["Technology"] = "http://feeds.reuters.com/reuters/technologyNews"
	urls["Top News"] = "http://feeds.reuters.com/reuters/topNews"
	urls["US News"] = "http://feeds.reuters.com/reuters/domesticNews"
	urls["World"] = "http://feeds.reuters.com/reuters/worldNews"

	return urls
}
