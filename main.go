package main

import "fmt"

func main() {
    var w string
    fmt.Scanln(&w)
    results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
    slice := append(results, w)
   // var s []string = results[0:]
    res := 0
    for _, a := range slice {
      if a == "w" {
    res += 3   
      } else if a == "d"{
    res += 1
      }
    }
    fmt.Println(res)

}
