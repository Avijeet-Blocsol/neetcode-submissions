type MyHashMap struct {
	values []int
}

func Constructor() MyHashMap {
	return MyHashMap{
		values: make([]int, 0),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	if key > len(this.values)-1 {
		newArr := make([]int, key+1)

		for i := range newArr {
			newArr[i] = -1
		}

		for index, val := range this.values {
			newArr[index] = val
		}

		newArr[key] = value

		this.values = newArr
	} else {
		this.values[key] = value
	}
}

func (this *MyHashMap) Get(key int) int {
	if !(key > len(this.values)-1) {
		return this.values[key]
	} else {
		return -1
	}
}

func (this *MyHashMap) Remove(key int) {
	if !(key > len(this.values)-1) {
		this.values[key] = -1
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */