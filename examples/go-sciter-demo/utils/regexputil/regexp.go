package regexputil

import "regexp"

// input:待解析string pattern:表达式 duplRemove：是否去重
func RegexMatchAll(input string, pattern string, duplRemove bool) []string {
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)
	// 去重
	seen := make(map[string]bool)
	results := []string{}
	for _, url := range matches {
		if duplRemove && !seen[url] {
			seen[url] = true
			results = append(results, url)
		} else {
			results = append(results, url)
		}
	}
	return results
}
