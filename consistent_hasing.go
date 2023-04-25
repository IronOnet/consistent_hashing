package main 

import (
	"fmt" 
	"hash/fnv" 
	"sort" 
	"strconv"
)

// Hash is a function that returns the hash of a string 
type Hash func(string) uint32 

// Consistent hash is a type that represents a hash ring 
type ConsistentHash struct{
	hash Hash  // the hash function 
	replicas int 

	keys []uint32 // sorted list of cached keys 
	hashMap map[uint32]string // map of hash values to node names
}

func NewConsistentHash(replicas int, nodes []string) *ConsistentHash{
	cHash := &ConsistentHash{
		hash: hashFnv32a, 
		replicas: replicas, 
		hashMap: make(map[uint32]string),
	}

	// Add the nodes to the hash ring 
	for _, node := range nodes{
		for i:= 0; i < cHash.replicas; i++{
			key := cHash.hash(fmt.Sprintf("%s:%d", node, i))
			cHash.keys = append(cHash.keys, key) 
			cHash.hashMap[key] = node 
		}
	}

	// Sort the keys  
	sort.Slice(cHash.keys, func(i, j int) bool{
		return cHash.keys[i] < cHash.keys[j]
	})

	return cHash 
}

// hashFnv32a is  a HashFunction that returns the FNV-1a hash of 
// a string 
func hashFnv32a(s string) uint32{
	h := fnv.New32a() 
	h.Write([]byte(s))
	return h.Sum32()
}

// GetNode returns the name of the node a key belongs to 
func (chash *ConsistentHash) GetNode(key string) string{
	if len(chash.keys) == 0{
		return "" 
	}

	// Hash the keys 
	hash := chash.hash(key) 

	// Binary search for the firs key >= hash 
	idx := sort.Search(len(chash.keys), func(i int) bool{
		return chash.keys[i] >= hash 
	})

	// Wrap around to the begining of the ring if necessary 
	if idx == len(chash.keys){
		idx = 0 
	}

	// Return the node name 
	return chash.hashMap[chash.keys[idx]]
}

func main(){
	// Create a new consistent hash with 3 new replicas and 3 nodes 
	chash := NewConsistentHash(3, []string{"node1", "node2", "node3"}) 

	// Print the node that each key belongs to 
	for i:= 0; i < 30; i++{
		key := "key" + strconv.Itoa(i) 
		node := chash.GetNode(key) 
		fmt.Printf("key=%s node=%s\n", key, node)
	}
}