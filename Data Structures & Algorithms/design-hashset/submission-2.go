type MyHashSet struct {
	values []bool
}

func Constructor() MyHashSet {
	return MyHashSet{
		values: make([]bool, 0),
	}
}

func (this *MyHashSet) Add(key int) {
	if key > len(this.values)-1 {
		newArr := make([]bool, key+1)

		for index, val := range this.values {
			if val {
				newArr[index] = true
			}
		}

		newArr[key] = true

		this.values = newArr
	} else {
		this.values[key] = true
	}
}

func (this *MyHashSet) Remove(key int) {
	if !(key > len(this.values)-1) {
		this.values[key] = false
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if !(key > len(this.values)-1) {
		return this.values[key]
	} else {
		return false
	}
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 