package main

import (
    "fmt"
    "math/rand"
    "time"
)

type RandomizedSet struct {
    nums       []int
    valToIndex map[int]int
}

func Constructor() RandomizedSet {
    return RandomizedSet{
        nums:       []int{},
        valToIndex: make(map[int]int),
    }
}

func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.valToIndex[val]; ok {
        return false  // Value already exists
    }
    this.valToIndex[val] = len(this.nums)
    this.nums = append(this.nums, val)
    return true
}

func (this *RandomizedSet) Remove(val int) bool {
    if _, ok := this.valToIndex[val]; !ok {
        return false  // Value does not exist
    }
    index := this.valToIndex[val]
    lastElement := this.nums[len(this.nums)-1]

    this.nums[index] = lastElement
    this.valToIndex[lastElement] = index
    this.nums = this.nums[:len(this.nums)-1]
    delete(this.valToIndex, val)
    return true
}

func (this *RandomizedSet) GetRandom() int {
    rand.Seed(time.Now().UnixNano())
    return this.nums[rand.Intn(len(this.nums))]
}

func main() {
    obj := Constructor()
    fmt.Println(obj.Insert(1))   // Output: true
    fmt.Println(obj.Insert(2))   // Output: true
    fmt.Println(obj.Remove(1))   // Output: true
    fmt.Println(obj.Insert(2))   // Output: false (2 already present)
    fmt.Println(obj.GetRandom()) // Output: 2 (since 2 is the only element)
}
