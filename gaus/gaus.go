package main
import "fmt"

func gaus(n int)int{
    return n * (n+1)/2
}
func main(){
    fmt.Println(gaus(100))
}
