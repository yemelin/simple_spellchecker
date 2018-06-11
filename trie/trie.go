package trie

// Trie - words converted to tree
type Trie struct {
	Letter byte
	Next   []*Trie
	Final  bool //last letter in the word
}

func create(letter byte) *Trie {
	return &Trie{letter, []*Trie{}, false}
}

// Create - build a trie from a list of words sorted alphabetically
func Create(text [][]byte) *Trie {
	rootNode := create(0)
	t := &task{
		bucket:   text,
		rootNode: rootNode,
	}
	t.perform()
	return rootNode
}

type task struct {
	bucket   [][]byte
	rootNode *Trie
	n        int
}

var counter int

func (t *task) perform() {
	var c byte
	var from int
	var node *Trie
	for to, word := range t.bucket {
		if len(word) <= t.n {
			continue
		}
		if word[t.n] != c {
			if node != nil {
				tcur := &task{
					bucket:   t.bucket[from:to],
					rootNode: node,
					n:        t.n + 1,
				}
				tcur.perform()
			}
			from = to //new bucket start
			c = word[t.n]
			node = create(c)
			t.rootNode.Next = append(t.rootNode.Next, node)
		}
		if len(word) == t.n+1 {
			node.Final = true
		}
	}
	if node != nil {
		tcur := &task{
			bucket:   t.bucket[from:len(t.bucket)],
			rootNode: node,
			n:        t.n + 1,
		}
		tcur.perform()
	}
}

type buffer struct {
	b []string
}

func (buffer *buffer) append(s string) {
	buffer.b = append(buffer.b, s)
}

func (t *Trie) Restore() []string {
	ret := &buffer{[]string{}}
	t._restore(make([]byte, 0, 40), ret)
	return ret.b
}

func (t *Trie) _restore(buf []byte, ret *buffer) {
	b := buf
	if t.Letter != 0 {
		b = append(buf, t.Letter)
		if t.Final {
			ret.append(string(b))
		}
	}
	if len(t.Next) != 0 {
		for _, subTrie := range t.Next {
			subTrie._restore(b, ret)
		}
	}
}
