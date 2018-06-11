package corrector

import (
	"github.com/yemelin/simple_corrector/trie"
)

type task struct {
	distances []byte
	node      *trie.Trie
	word      []byte
}

func NewTask(word []byte, root *trie.Trie) *task {
	Min = 100
	Candidate = ""
	distances := make([]byte, len(word)+1)
	for i := 1; i < len(distances); i++ {
		distances[i] = byte(i)
	}
	return &task{
		distances: distances,
		node:      root,
		word:      word,
	}
}

var Min byte = 100
var Candidate string

func (t *task) Perform() {
	t._perform([]byte{})
}

func (t *task) _perform(buf []byte) {
	if len(t.node.Next) == 0 {
		return
	}
	for _, child := range t.node.Next {
		d := distances(t.distances, child.Letter, t.word)

		k := d[0]
		if d[0] > byte(len(t.word)) {
			k = byte(len(t.word))
		}
		if d[k] > Min {
			continue
		}

		if child.Final && d[len(d)-1] < Min {
			Min = d[len(d)-1]
			Candidate = string(buf) + string(child.Letter)
		}
		nextTask := &task{
			distances: d,
			node:      child,
			word:      t.word,
		}
		nextTask._perform(append(buf, child.Letter))
	}
}

func distances(d []byte, letter byte, word []byte) []byte {
	ret := make([]byte, len(d))
	ret[0] = d[0] + 1 //next level
	for i := 1; i < len(ret); i++ {
		var cost byte
		if word[i-1] != letter {
			cost = 1
		}
		ret[i] = minVal(d[i-1]+cost, d[i]+1, ret[i-1]+1)
	}
	return ret
}

func minVal(r, i, d byte) byte {
	m := r
	if i < m {
		m = i
	}
	if d < m {
		return d
	}
	return m
}
