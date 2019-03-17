package main
import (
    "fmt"
    //"io/ioutil"
    "os"
    "bufio"
    "strconv"
    "math"
)
var sc = bufio.NewScanner(os.Stdin)
func nextLine() string{
    sc.Scan()
    return sc.Text()
}


func stringToInt(arr []string, arrInt []float64) []float64{
    for _, i:= range arr{
            j, err :=strconv.ParseFloat(i, 64)
            if err != nil{
                panic(err)
            }
            arrInt = append(arrInt,j)
    }
    return arrInt 
}

func main() {
    // scanner := bufio.NewScanner(os.Stdin)
    // for scanner.Scan(){
    //     str := scanner.Text()
    //     fmt.Print(str, "\n")
    // }
    //_, n1, n2 := nextLine(),nextLine(),nextLine()
    var c []string 
    var c2 =[]float64{}
    for i:=0; i<3; i++{
        c = append(c, nextLine())
    }
    c2 = stringToInt(c, c2)
    //var sum float64
    var square float64
    //var squareSum float64
    var flo float64 = 1
    for i:=0; i<len(c2); i++{
        sum := 0.0
        squareSum := 0.0
        if(i != 0){
            for j:=flo; j<=c2[i]; j++{
                sum += j
                squareSum += math.Pow(j,2)    
            }
            square = math.Pow(sum,2)
            fmt.Println(square - squareSum)
        }
    }
}
