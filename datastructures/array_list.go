package datastructures

type ArrayList struct {
	_array []interface{}
}

func (array ArrayList) Get(index int) interface{} {
	return array._array[index]
}

func (array ArrayList) Size() int {
	return len(array._array)
}

func (array *ArrayList) Append(item interface{}) {
	array._array = append(array._array, item)
}

func (array *ArrayList) Set(index int, item interface{}) {
	array._array[index] = item
}

func (array ArrayList) Contains(item interface{}) bool {
	for _, value := range array._array {
		if value == item {
			return true
		}
	}
	return false
}

func (array ArrayList) Array() []interface{} {
	return array._array
}
