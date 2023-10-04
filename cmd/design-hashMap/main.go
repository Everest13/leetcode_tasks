package main

import "fmt"

// Task #706 Design HashMap.
// https://leetcode.com/problems/design-hashmap/description/
// Runtime: 89 ms, beats 81.16%; memory usage: 7.8mb, beats 83,5%.
func main() {
	myHashMap := Constructor()

	myHashMap.Put(1, 1)
	fmt.Println("The map is now ", myHashMap.keyVal)

	myHashMap.Put(2, 2)
	fmt.Println("The map is now ", myHashMap.keyVal)

	myHashMap.Get(1)
	fmt.Println("The map is now ", myHashMap.keyVal)

	hM := myHashMap.Get(3)
	fmt.Println("The result is ", hM)

	myHashMap.Put(2, 1)
	fmt.Println("The map is now ", myHashMap.keyVal)

	myHashMap.Get(2)
	fmt.Println("The map is now ", myHashMap.keyVal)

	myHashMap.Remove(2)
	fmt.Println("The map is now ", myHashMap.keyVal)

	myHashMap.Get(2)
	fmt.Println("The map is now ", myHashMap.keyVal)
}

type MyHashMap struct {
	keyVal map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{map[int]int{}}
}

func (hM *MyHashMap) Put(key int, value int)  {
	hM.keyVal[key] = value
}

func (hM *MyHashMap) Get(key int) int {
	if val, ok := hM.keyVal[key]; ok {
		return val
	}
	return -1
}

func (hM *MyHashMap) Remove(key int)  {
	delete(hM.keyVal, key)
}
