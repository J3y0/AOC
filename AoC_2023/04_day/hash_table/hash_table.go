package hash_table

type HashNode struct {
	Key   int
	Value int
}

// Open addressing Hash Table
type HashTable struct {
	Capacity int
	Array    []HashNode
}

func NewHashTable(capacity int) *HashTable {
	newHT := new(HashTable)
	newHT.Capacity = capacity
	newHT.Array = make([]HashNode, capacity)
	for i := range newHT.Array {
		newHT.Array[i].Key = -1
		newHT.Array[i].Value = -1
	}
	return newHT
}

func (ht *HashTable) Hash(key int) int {
	return (key % ht.Capacity)
}

// Return true if insertion is ok, else no free space available
// and couldn't insert element
func (ht *HashTable) Insert(key int, value int) bool {
	index := ht.Hash(key)
	// Find index for next free space
	var counter int
	for ht.Array[index].Key != -1 && ht.Array[index].Key != key {
		index = (index + 1) % ht.Capacity // Linear probing
		counter += 1
		if counter > ht.Capacity {
			// No free space available
			return false
		}
	}

	// Insert new node
	ht.Array[index].Key = key
	ht.Array[index].Value = value
	return true
}

// Receive the value and ok to true if we found the key, 0 and false else
func (ht *HashTable) Get(key int) (value int, ok bool) {
	index := ht.Hash(key)
	var counter int
	for ht.Array[index].Key != key {
		index = (index + 1) % ht.Capacity // Linear probing
		counter += 1
		if counter > ht.Capacity {
			// Not found
			return 0, false
		}
	}
	value = ht.Array[index].Value
	ok = true
	return
}
