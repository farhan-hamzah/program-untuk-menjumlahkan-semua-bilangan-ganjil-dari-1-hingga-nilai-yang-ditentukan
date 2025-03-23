package main
import "fmt"

func ganjil(n int) int {
    if n < 2 {
        return 1
    }
    if n%2 != 0 {
        return n + ganjil(n-2)
    } else {
        return ganjil(n - 1)
    }
}

func main(){
	var num int
	fmt.Scan(&num)
	fmt.Print(ganjil(num))
}
