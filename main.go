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
	d, _ := ioutil.ReadFile("resources/vocabulary.txt")
	w, _ := ioutil.ReadFile("resources/187.txt")
	readEnd := time.Now().UnixNano()

	dictWords := regexp.MustCompile("\n").Split(string(d), -1)
	words := regexp.MustCompile(" +").Split(string(w), -1)
	vocabulary := make([]string, len(dictWords))
	for i := 0; i < len(vocabulary); i++ {
		vocabulary[i] = strings.ToLower(dictWords[i])
	}
	splitEnd := time.Now().UnixNano()

	t := trie.Create(vocabulary)
	trieEnd := time.Now().UnixNano()

	var count byte
	for _, word := range words {
		corrector.NewTask(word, t).Perform()
		count += corrector.Min
	}
	done := time.Now().UnixNano()

	totalTime := done - start
	fmt.Println(count)
	fmt.Println("time: ", totalTime/int64(time.Millisecond))
	fmt.Println("read time: ", (readEnd-start)/int64(time.Millisecond))
	fmt.Println("split time: ", (splitEnd-readEnd)/int64(time.Millisecond))
	fmt.Println("trie time: ", (trieEnd-readEnd)/int64(time.Millisecond))
	fmt.Println("corrector time: ", (done-trieEnd)/int64(time.Millisecond))
}
