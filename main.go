package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"

	"github.com/yemelin/simple_corrector/corrector"
	"github.com/yemelin/simple_corrector/trie"
)

type fixer interface {
	FixErrors(text []string) ([]string, int)
}

func readWords(fileName, sepRe string) ([]string, error) {
	s, _ := ioutil.ReadFile(fileName)
	sep := regexp.MustCompile(sepRe)
	ret := sep.Split(string(s), -1)
	return ret, nil
}

func main() {
	start := time.Now().UnixNano()
	dictWords, _ := readWords("resources/vocabulary.txt", "\n")
	// TODO: lowercase all bytes before splitting may be faster
	vocabulary := make([]string, len(dictWords))
	for i := 0; i < len(vocabulary); i++ {
		vocabulary[i] = strings.ToLower(dictWords[i])
	}
	// fmt.Println(len(vocabulary))
	// c := corrector.New(vocabulary)
	words, _ := readWords("resources/187.txt", " +")
	readEnd := time.Now().UnixNano()
	// correctWords, numErrors := c.FixErrors(words)
	// fmt.Println("Total corrections: ", numErrors)
	// fmt.Println("Corrected text:")
	// fmt.Println(correctWords)

	// v := []string{"aaab", "aa", "aacb", "bccc", "bcdd", "cabb", "cdef"}
	// fmt.Println(v)
	t := trie.Create(vocabulary)
	trieEnd := time.Now().UnixNano()
	// t := trie.Create(v)
	// words := t.Restore()
	var count byte
	for _, word := range words {
		// fmt.Println(word)
		corrector.NewTask(word, t).Perform()
		// fmt.Println(corrector.Min, corrector.Candidate)
		count += corrector.Min
	}
	done := time.Now().UnixNano()
	totalTime := done - start
	fmt.Println(count)
	fmt.Println("time: ", totalTime/int64(time.Millisecond))
	fmt.Println("read time: ", (readEnd-start)/int64(time.Millisecond))
	fmt.Println("trie time: ", (trieEnd-readEnd)/int64(time.Millisecond))
	fmt.Println("corrector time: ", (done-trieEnd)/int64(time.Millisecond))
}
