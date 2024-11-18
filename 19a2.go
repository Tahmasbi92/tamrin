import "fmt"

func main() {

	number := 1

	for number <= 10 {
		fmt.Printf("5 * %d = %d", number, 5*number)
		number += 1
	}
}
