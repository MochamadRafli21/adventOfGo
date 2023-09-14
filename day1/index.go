package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
    cals, err :=  os.Open("./input.txt")
    if err != nil {
      panic(err)
    }

    defer cals.Close()

    scaner := bufio.NewScanner(cals)
    scaner.Split(bufio.ScanLines)
    
    elves := make(map[int]int64)
    counter := 0
    biggest := 0
    biggestVal := int64(0)
    for scaner.Scan() {
      val := scaner.Text()
      if val == "" {
        counter++
        continue
      }

      intVal, err := strconv.ParseInt(val, 10, 0)
      if err != nil {
        panic(err)
      }

      elves[counter] = elves[counter] + intVal
      if elves[counter] > biggestVal{
        biggest = counter
        biggestVal = elves[counter]
      }
    }
    fmt.Println("Elves Number: ",biggest)
    fmt.Println("Total Calories: ",biggestVal)


}
