package sort

type Comparable interface {
	Less()
	Swap(i, j IntComp)
}

// Comparable interface types for method attachments
type StringComp string
type IntComp []int

func (s StringComp) Less() {

}

func (d IntComp) Swap(i, j int) {
	temp := d[i]
	d[i] = d[j]
	d[j] = temp
}

func min(numbers []int) int {
	var mindex int
	mvalue := numbers[0]
	for i, d := range numbers {
		if d < mvalue {
			mindex = i
			mvalue = numbers[i]
		}
	}
	return mindex
}

func max(numbers []int) int {
	var maxindex, maxvalue int
	for i, d := range numbers {
		if d > maxvalue {
			maxindex = i
			maxvalue = numbers[i]
		}
	}
	return maxindex
}

func Insertion(items []Comparable) {
	for i, item := range items {
		for d, compare := range items[i:] {
			if compare < item {
				items.Swap(i, d)
			}
		}
	}
}
