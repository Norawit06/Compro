package w5

func main() {

}

func Sum(number []int) int {
	sum := 0

	for n := range number {
		sum += n
	}
	return sum
}
