package main

func main() {
	tests := []struct {
		list     []int
		start    int
		end      int
		reversed []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 0, 3, []int{4, 3, 2, 1, 5, 6}},
	}
}
