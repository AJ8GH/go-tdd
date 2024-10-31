package webcheck

type WebsiteChecker func(string) bool

func CheckWebsites(w WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = w(url)
	}

	return results
}
