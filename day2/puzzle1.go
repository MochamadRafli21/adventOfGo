package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func decrypt(key string)string{
  switch key {
 
  case "A", "X":
    return "Rock"
 
  case "B","Y":
    return "Paper"
  
  case "C","Z":
    return "Scissors"
 
  }
  
  panic("Not Found")
}

func getPoints(hands string)int{
  points := 0
  listHands := strings.Split(hands, " ") 
  myHand := decrypt(listHands[1])
  yourHand := decrypt(listHands[0])
  
  switch myHand {
  case "Rock":
   points += 1
  case "Paper":
    points += 2
  case "Scissors":
    points += 3
  }

  if myHand == yourHand{
    points += 3
  }else if myHand == "Rock" && yourHand == "Scissors" || myHand == "Paper" && yourHand == "Rock" || myHand == "Scissors" && yourHand == "Paper"{
    points += 6 
  }else{
    points +=0
  }

  return points
}

func main(){
  cals, err :=  os.Open("./input.txt")
  if err != nil {
    panic(err)
  }

  defer cals.Close()

  scaner := bufio.NewScanner(cals)
  scaner.Split(bufio.ScanLines)
  points := 0
  for scaner.Scan() {
    lines := scaner.Text()
    points += getPoints(lines)
  }
  fmt.Println(points)
}
