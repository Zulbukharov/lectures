package healthchecker

type WebsiteChecker func(string) bool

type response struct {
	url   string
	valid bool
}

func CheckWebsitesHealth(wc WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}

	resultChannel := make(chan *response)

	for i, url := range urls {
		go func(u string, i int) {
			resultChannel <- &response{url: u, valid: wc(u)}
		}(url, i)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.url] = result.valid
	}

	return results
}
