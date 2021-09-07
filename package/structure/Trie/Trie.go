package Trie

type Trie struct {
	prefix byte
	result []string
	suffix map[byte]*Trie
}

const (
	MaxWorker = 32
)

var (
	resultChan = make(chan string) //结果通道
	downChan   = make(chan struct{})
	worksChan  = make(chan *Trie)
	nowWorker  = 0 //目前工人
)

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{suffix: map[byte]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word, value string) {
	root := this
	for i := 0; i < len(word); i++ {
		if root.suffix[word[i]] == nil {
			root.suffix[word[i]] = &Trie{
				prefix: word[i],
				suffix: map[byte]*Trie{},
			}
		}
		root = root.suffix[word[i]]
	}
	root.result = append(root.result, value)
}

func searchCenter() (ret []string) {
	for {
		select {
		case result := <-resultChan:
			ret = append(ret, result)
		case <-downChan:
			nowWorker--
			if nowWorker == 0 {
				return
			}
		case works := <-worksChan:
			nowWorker++
			go search(works, true)
		}
	}
}

func search(tire *Trie, master bool) {
	if master {
		defer func() {
			downChan <- struct{}{}
		}()
	}
	if tire == nil {
		return
	}
	for _, v := range tire.suffix {
		if nowWorker < MaxWorker {
			worksChan <- v
		} else {
			search(v, false)
		}
	}
	for _, v := range tire.result {
		resultChan <- v
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) (ret []string) {
	nowWorker = 1
	go func() {
		root := this
		for i := 0; i < len(word); i++ {
			if root.suffix[word[i]] != nil {
				root = root.suffix[word[i]]
			} else {
				return
			}
		}
		search(root, false)
		downChan <- struct{}{} //结束任务
	}()
	return searchCenter()
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
