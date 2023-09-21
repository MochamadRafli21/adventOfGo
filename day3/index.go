package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func splitRack(racks string) (string, string){
  half := len(racks)/2 
  return racks[:half], racks[half:]
}

func findCommonRack(rack1 string, rack2 string)rune{
  for _, key := range rack1{
    for _,val := range rack2{
      if key == val{
        return key
      }
    }
  }
  panic("cant find commond value")
}

func findCommonBadge(elves [3]string)rune{
  chars := elves[0]+elves[1]+elves[2] 
  for _, key := range chars{
    char := string(key)
    if strings.Contains(elves[0], char) && strings.Contains(elves[1], char) && strings.Contains(elves[2], char){
      return key
    } 
  }
  panic("cant find commond value")

}

func findIndex(char rune)int{
  if char < 92{
    return int(char-38)
  }

  return int(char-96)
}


func getPoint(racks string)int {
  rack1, rack2 := splitRack(racks)
  common := findCommonRack(rack1, rack2)
  return findIndex(common)
}

func getGroup(elves [3]string) int{
  common := findCommonBadge(elves)
  return findIndex(common)
}


func main(){
  cals, err :=  os.Open("./input.txt")
  if err != nil {
      panic(err)
  }

  defer cals.Close()
  sum :=0

  scaner := bufio.NewScanner(cals)
  //scaner.Split(bufio.ScanLines)
  // first puzzle 
  //for scaner.Scan() {
  //  point := getPoint(scaner.Text())
  //  sum += point
  //}

  texts := []string{}
  for scaner.Scan() {
    texts = append(texts, scaner.Text())
  }

  for i:=0 ; i< len(texts); i+=3{
    groups := [3]string{
      texts[i],
      texts[i+1],
      texts[i+2],
    }
    sum += getGroup(groups)
  }
  fmt.Println(sum)



}
