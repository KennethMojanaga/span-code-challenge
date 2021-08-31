package datastructures

type SetInt struct {
	_set         map[int]bool
	_initialized bool
}

func (set *SetInt) Size() int {
	if !set._initialized {
		set.Init([]int{})
		set._initialized = true
	}

	return len(set._set)
}

func (set *SetInt) Init(items []int) []int {
	if !set._initialized {
		set._set = make(map[int]bool)
		set._initialized = true
	}

	for _, item := range items {
		set._set[item] = true
	}

	return set.Array()
}

func (set *SetInt) Add(item int) {
	if !set._initialized {
		set.Init([]int{})
		set._initialized = true
	}

	set._set[item] = true
}

func (set *SetInt) Delete(item int) {
	if !set._initialized {
		set.Init([]int{})
		set._initialized = true
	}

	delete(set._set, item)
}

func (set *SetInt) Exists(item int) bool {
	if !set._initialized {
		set.Init([]int{})
		set._initialized = true
	}

	_, exists := set._set[item]
	return exists
}

func (set *SetInt) Array() []int {
	if !set._initialized {
		set.Init([]int{})
		set._initialized = true
	}

	results := []int{}
	for key := range set._set {
		results = append(results, key)
	}
	return results
}
