type DynamicArray struct {
    array []int
    capacity int
    length int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{
        array: make([]int, capacity),
        capacity: capacity,
        length: 0,
    }
}

func (da *DynamicArray) Get(i int) int {
    return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.array[i] = n 
}

func (da *DynamicArray) Pushback(n int) {
    if (da.length == da.capacity) {
        da.resize()
    }

    da.array[da.length] = n
    da.length++
}

func (da *DynamicArray) Popback() int {
    if da.length > 0 {
        da.length--
    }
    return da.array[da.length]
}

func (da *DynamicArray) resize() {
    da.capacity = 2 * da.capacity
    newArray := make([]int, da.capacity)

    for i, n := range da.array {
        newArray[i] = n
    }
    da.array = newArray
}

func (da *DynamicArray) GetSize() int {
    return da.length
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
