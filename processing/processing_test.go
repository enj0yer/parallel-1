package processing

import "testing"

func compareArrays[T comparable](elems1 []T, elems2 []T) bool {
	if len(elems1) != len(elems2) {
		return false
	}
	for i := 0; i < len(elems1); i++ {
		if elems1[i] != elems2[i] {
			return false
		}
	}
	return true
}

func TestProcessing(t *testing.T) {
	initItems := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	applier := func(i int) (int, error) { return i + 1, nil }
	converter := NewConverter(initItems, applier)

	processedItems := make([]int, len(initItems))
	for i := 0; i < len(initItems); i++ {
		value, _ := applier(initItems[i])
		processedItems[i] = value
	}

	result, err := converter.ProcessSequentially()
	if err != nil {
		t.Fatalf("error while sequentially processing: %v", err)
	}

	if !compareArrays(processedItems, result) {
		t.Errorf("arrays are not equal")
	}

	result, err = converter.ProcessSimultaneously(4)
	if err != nil {
		t.Fatalf("error while simultaneous processing: %v", err)
	}
	if !compareArrays(processedItems, result) {
		t.Errorf("arrays are not equal")
	}

}
