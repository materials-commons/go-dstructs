package trie

type TrieChildren map[string]*Trie
type Path []string

type Trie struct {
	ValueSet bool
	Value interface{}
	Children TrieChildren
}

func NewTrie() *Trie {
	return &Trie{
		Children: make(TrieChildren),
	}
}

func (t *Trie) Retrieve(path Path) (value interface{}, hasValue, pathExists bool) {
	node := t.retrieve(path)
	if node != nil {
		value = node.Value
		hasValue = node.ValueSet
		pathExists = true
	} else {
		value = nil
		hasValue, pathExists = false, false
	}

	return 
}

func (t *Trie) PathExists(path Path) bool {
	_, _, pathExists := t.Retrieve(path)
	return pathExists
}

func (t *Trie) retrieve(path Path) (node *Trie) {
	if zeroPath(path) {
		return t
	}

	key, rest := keyRest(path)
	if res, ok := t.Children[key]; ok {
		return res.retrieve(rest)
	}

	return nil
}

func (t *Trie) Add(path Path, value interface{}) bool {
	return t.add(path, value, true)
}

func (t *Trie) add(path Path, value interface{}, existing bool) bool {
	if zeroPath(path) {
		if existing {
			return false
		} else {
			t.ValueSet = true
			t.Value = value
			return true
		}
	}

	key, rest := keyRest(path)
	if _, ok := t.Children[key]; !ok {
		entry := NewTrie()
		t.Children[key] = entry
		existing = false
	}
	return t.add(rest, value, existing)
}

func (t *Trie) Modify(path Path, value interface{}) (found bool) {
	found = false
	node := t.retrieve(path)
	if node != nil {
		node.ValueSet = true
		node.Value = value
		found = true
	}

	return
}

func (t *Trie) Set(path Path, value interface{}) (newNode bool) {
	return t.set(path, value, false)
}

func (t *Trie) set(path Path, value interface{}, existing bool) (newNode bool) {
	if zeroPath(path) {
		t.ValueSet = true
		t.Value = value
		return !existing
	}

	key, rest := keyRest(path)
	if _, ok := t.Children[key]; !ok {
		entry := NewTrie()
		t.Children[key] = entry
		existing = false
	}
	return t.set(rest, value, existing)
}

func zeroPath(path Path) bool {
	return len(path) == 0
}

func keyRest(path Path) (key string, rest Path) {
	key = path[0]
	rest = path[1:]
	return
}
