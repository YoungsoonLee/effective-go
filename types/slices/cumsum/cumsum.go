package main

func cumsum(a []int) []int {
	//var cs []int
	// tune up
	cs := make([]int, 0, len(a))

	sum := 0
	for _, v := range a {
		sum += v
		cs = append(cs, sum)
	}
	return cs
}

func main() {

}
