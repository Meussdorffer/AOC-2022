package main

type Queue [][2]int

func (q *Queue) Enqueue(item [2]int) {
	*q = append(*q, item)
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Dequeue() ([2]int, bool) {
	if !q.isEmpty() {
		item := (*q)[0]
		*q = (*q)[1:]
		return item, true
	}
	return [2]int{}, false
}
