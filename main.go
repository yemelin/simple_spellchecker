package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/yemelin/simple_corrector/corrector"
	"github.com/yemelin/simple_corrector/trie"
)

func readWords(fileName string, sep byte) ([][]byte, error) {
	s, _ := ioutil.ReadFile(fileName)
	lims := fastSplit(s, sep)

	ret := make([][]byte, 0, len(lims)/2)
	for i := 0; i < len(lims); i += 2 {
		ret = append(ret, s[lims[i]:lims[i+1]])
	}
	return ret, nil
}

func fastSplit(s []byte, sep byte) []int {
	in := false
	ret := make([]int, 0, 100000)
	for i := 0; i < len(s); i++ {
		if in == (s[i] == sep) {
			ret = append(ret, i)
			in = !in
		}
	}
	if in {
		ret = append(ret, len(s))
	}
	return ret
}

func main() {
	start := time.Now().UnixNano()

	vocabulary, _ := readWords("resources/vocabulary.txt", '\n')
	words, _ := readWords("resources/187.txt", ' ')
	readEnd := time.Now().UnixNano()

	for i := 0; i < len(vocabulary); i++ {
		for j := 0; j < len(vocabulary[i]); j++ {
			vocabulary[i][j] = vocabulary[i][j] + 32 //lowercase
		}
	}
	upCaseEnd := time.Now().UnixNano()

	t := trie.Create(vocabulary)
	trieEnd := time.Now().UnixNano()

	l := len(words)
	var wg sync.WaitGroup
	wg.Add(l)
	counts := make([]byte, len(words))
	for i := 0; i < len(words); i++ {
		i1 := i
		go func() {
			counts[i1], _ = corrector.NewTask(words[i1], t).Perform()
			wg.Done()
		}()
	}
	wg.Wait()
	count := 0
	for i := 0; i < l; i++ {
		count += int(counts[i])
	}
	done := time.Now().UnixNano()

	totalTime := done - start
	fmt.Println(count)
	fmt.Println("time: ", totalTime/int64(time.Millisecond))
	fmt.Println("read time: ", (readEnd-start)/int64(time.Millisecond))
	fmt.Println("upCase time: ", (upCaseEnd-readEnd)/int64(time.Millisecond))
	fmt.Println("trie time: ", (trieEnd-readEnd)/int64(time.Millisecond))
	fmt.Println("corrector time: ", (done-trieEnd)/int64(time.Millisecond))
}
