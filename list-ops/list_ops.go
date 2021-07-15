package listops

type binFunc = func(a, b int) int
type predFunc = func(a int) bool
type unaryFunc = func(a int) int

type IntList []int

// Foldr reduces the IntList into int using the binFunc from the right.
func (intList IntList) Foldr(binFunc binFunc, init int) int {
	acc := init
	for i := len(intList) - 1; i >= 0; i-- {
		acc = binFunc(intList[i], acc)
	}
	return acc
}

// Foldl reduces the IntList into int using the binFunc from the right.
func (intList IntList) Foldl(binFunc binFunc, init int) int {
	acc := init
	for _, v := range intList {
		acc = binFunc(acc, v)
	}
	return acc
}

// Filter returns a new IntList with all elements that pass the predicate.
func (intList IntList) Filter(predFunc predFunc) IntList {
	newList := make(IntList, 0, len(intList))
	for _, v := range intList {
		if predFunc(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// Length returns the length of the IntList.
func (intList IntList) Length() int {
	return len(intList)
}

// Map maps each element of the IntList with the unaryFunc.
func (intList IntList) Map(unaryFunc unaryFunc) IntList {
	newList := make(IntList, 0, len(intList))
	for _, v := range intList {
		newList = append(newList, unaryFunc(v))
	}
	return newList
}

// Reverse reverses the IntList.
func (intList IntList) Reverse() IntList {
	newList := make(IntList, 0, len(intList))
	for i := len(intList) - 1; i >= 0; i-- {
		newList = append(newList, intList[i])
	}
	return newList
}

// Append appends the elements of the IntList to the end of the IntList.
func (intList IntList) Append(intList2 IntList) IntList {
	return append(intList, intList2...)
}

// Concat concatenates the IntList with the IntList passed in.
func (intList IntList) Concat(intList2 []IntList) IntList {
	for _, v := range intList2 {
		intList = append(intList, v...)
	}
	return intList
}
