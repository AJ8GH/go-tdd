package webcheck

type (
	WebsiteChecker func(string) bool
	result         struct {
		string
		bool
	}
)

func CheckWebsites(w WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultsChannel <- result{u, w(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}
