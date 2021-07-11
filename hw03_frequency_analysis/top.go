package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordWithCounter struct {
	word    string
	counter int
}

type wordsList []wordWithCounter

var space *regexp.Regexp = regexp.MustCompile(`[[:space:]]+`)

func (w wordsList) Len() int {
	return len(w)
}

func (w wordsList) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w wordsList) Less(i, j int) bool {
	if w[i].counter == w[j].counter {
		s := sort.StringSlice{w[i].word, w[j].word}
		return s.Less(0, 1)
	}
	return w[i].counter > w[j].counter // reverse sorting
}

func sortMapIntoSlice(m map[string]int) []wordWithCounter {
	res := make(wordsList, len(m))
	i := 0
	for k, v := range m {
		res[i] = wordWithCounter{word: k, counter: v}
		i++
	}
	sort.Sort(res)
	return res
}

func Top10(s string) []string {
	res := make([]string, 0)
	wordsCount := make(map[string]int)
	if s == "" {
		return res
	}

	s = space.ReplaceAllString(s, ` `)
	separatedStrings := strings.Split(s, ` `)

	for _, word := range separatedStrings {
		_, found := wordsCount[word]
		if !found {
			wordsCount[word] = 1
			continue
		}
		wordsCount[word]++
	}

	sorted := sortMapIntoSlice(wordsCount)
	if len(sorted) > 10 {
		sorted = sorted[:10]
	}
	for _, v := range sorted {
		res = append(res, v.word)
	}
	return res
}
