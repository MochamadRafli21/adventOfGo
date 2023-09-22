package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "strconv"
)

func parseValue(text string)(int,int,int,int){
  pair := strings.Split(text, ",")
  first := strings.Split(pair[0], "-")
  second := strings.Split(pair[1], "-")
  lowestFirst, err := strconv.Atoi(first[0])
  highestFirst, err := strconv.Atoi(first[1])
  lowestSecond, err := strconv.Atoi(second[0])
  highestSecond, err := strconv.Atoi(second[1])
  if err!= nil{
    panic("cant convert value")
  }
  return int(lowestFirst), int(highestFirst),int(lowestSecond),int(highestSecond)
}

func main(){
  cals, err :=  os.Open("./input.txt")
  if err != nil {
      panic(err)
  }

  defer cals.Close()
  count :=0
  count2 :=0
  scaner := bufio.NewScanner(cals)
  scaner.Split(bufio.ScanLines)
  
  for scaner.Scan() {
    lowestFirst, highestFirst, lowestTwo, highestTwo := parseValue(scaner.Text())
    // first puzzle 
    if(lowestFirst <= lowestTwo && highestFirst >= highestTwo ||
    lowestTwo <= lowestFirst&& highestTwo>= highestFirst){
      count++
    }
    // second puzzle
    if(lowestTwo >= lowestFirst && highestFirst >= lowestTwo ||
    highestFirst >= lowestTwo && highestFirst <= highestTwo ||
    highestTwo >= lowestFirst && highestTwo <= highestFirst){
      count2++
    }
  }

  fmt.Println(count)
  fmt.Println(count2)
}
