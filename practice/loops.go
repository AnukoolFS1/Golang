package main

import "fmt"

func Sqrt(x float64) float64 {
    z := 1.0
    for {
        next := z - (z*z - x) / (2 * z)
		fmt.Println(abs(next-z),",", next,",", z,",")
        if abs(next - z) < 1e-10 {  // close enough
            z = next
            break
        }
        z = next
        fmt.Println(z)
    }
    return z
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    // fmt.Println(Sqrt(25))
    anotherFor()
}