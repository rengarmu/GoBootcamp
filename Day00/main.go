package main

import( 
    "fmt"
    "strconv"
    "sort"
    "math"
    )


func main() {
var numsList[]int

fmt.Println ("Введите числа(для завершения ввода введите 'exit'):")
var input string
for {
    _, err := fmt.Scan(&input)
    if input == "exit" {
        break
    }
    num, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println ("Ошибка ввода")
    } else if num >= -100000 && num <= 100000 {
        numsList = append(numsList, num)
    }else {
        fmt.Println("Выход за пределы допустивых значений")
    }
    
}
sort.Ints(numsList)
length := len(numsList)
boolsArray := [4]bool{}
fmt.Println("Что вы хотите вывести: mean, median, mode, sd, all?")
var output string
for {
    fmt.Scan(&output)
    if output == "exit" {
        break
    } else if output == "all" {
        for i := range boolsArray {
            boolsArray[i] = true 
        }
        break
    } else if output == "mean" {
     boolsArray[0] = true
    } else if output == "median" {
     boolsArray[1] = true
    } else if output == "mode" {
     boolsArray[2] = true
    } else if output == "sd" {
     boolsArray[3] = true
    }
}

//output
if boolsArray[0] {
fmt.Printf("Mean: %.2f\n", Mean(numsList, length))
}
if boolsArray[1] {
fmt.Printf("Median: %.2f\n", Median(numsList, length))
}
if boolsArray[2] {
fmt.Println("Mode: ", Mode(numsList))
}
if boolsArray[3] {
fmt.Printf("SD: %.2f\n", SD(numsList, length, Mean(numsList, length)))
}
}


// MEAN
func Mean (numsList[]int, length int) float64 {
var sum float64 = 0
for _, num := range numsList {
    sum += float64(num)
}
return sum / float64(length)
}
 

// MEDIAN
func Median (numsList[]int, length int) float64 {
var median float64 = 0
if length % 2 == 0 {
    n1 := numsList[(length / 2)-1]
    n2 := numsList[(length / 2)]
    median = float64(n1 + n2) / 2.0
} else {
    median = float64(numsList[length / 2])
}
return median
}


// MODE
func Mode (numsList[]int) int {
var mode int = 0
var maxCount int = 0
for _, num := range numsList {
     count := 0
     for _, n := range numsList {
         if num == n {
            count++
       }
   }
 	if count > maxCount{
      maxCount = count
      mode = num
  } }
return mode
}

// SD
func SD (numsList[]int, length int, mean float64) float64{
var sumDev float64 = 0
for _, num := range numsList {
    sumDev += (float64(num) - mean) * (float64(num) - mean)
}
sd :=  math.Sqrt(sumDev / float64(length))
return sd
}