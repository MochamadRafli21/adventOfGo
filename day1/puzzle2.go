package main

import (
  "fmt"
  "os"
  "sort"
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
    
    elves := make(map[int]int)
    counter := 0
    var biggestVal []int
    for scaner.Scan() {
      val := scaner.Text()
      if val == "" {
        biggestVal = append(biggestVal, elves[counter])
        counter++
        continue
      }

      intVal, err := strconv.Atoi(val)
      if err != nil {
        panic(err)
      }

      elves[counter] = elves[counter] + intVal
    }

    sort.Ints(biggestVal)
    fmt.Println("Total Calories of 3 elves: ",biggestVal[len(biggestVal)-1]+ biggestVal[len(biggestVal)-2]+biggestVal[len(biggestVal)-3])
}
