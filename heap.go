package graph

type nextVertex struct {
	distance int
	current  string
}

type vertexesQueue []nextVertex

func (h vertexesQueue) Len() int           { return len(h) }
func (h vertexesQueue) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h vertexesQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *vertexesQueue) Push(x any) {
	*h = append(*h, x.(nextVertex))
}

func (h *vertexesQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
