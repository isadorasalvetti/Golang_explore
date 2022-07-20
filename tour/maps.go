package tour

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	ret_map := make(map[string]int)
	for _, word := range(fields) {
		count, ok := ret_map[word]
		if ok {  // If not necessary could just do ret_map[word] = count+1 for both cases.
			ret_map[word] = count+1 
		} else { 
			ret_map[word] = 1
		}
	}
	return ret_map
}

func main() {
	wc.Test(WordCount)
}
