package main

import (
  "fmt"
  "time"
  "sync"
  "bufio"
  "os"
  "strings"
)

var reader = bufio.NewReader(os.Stdin)
var printStatement, _ = fmt.Println("Provide your Api Key:")
var  takeInput, _ = reader.ReadString('\n')
var apiKey = strings.TrimSuffix(takeInput, "\n")

var baseSoccer = fmt.Sprintf(
  "https://api.the-odds-api.com/v3/odds/?apiKey=%s&sport=soccer&region=eu&mkt=h2h",
  apiKey)

var baseBasketBall = fmt.Sprintf(
  "https://api.the-odds-api.com/v3/odds/?apiKey=%s&sport=basketball&region=eu&mkt=h2h",
  apiKey)

var baseBaseBall = fmt.Sprintf(
  "https://api.the-odds-api.com/v3/odds/?apiKey=%s&sport=baseball&region=eu&mkt=h2h",
  apiKey)

var baseRugby = fmt.Sprintf(
  "https://api.the-odds-api.com/v3/odds/?apiKey=%s&sport=rugbyleague&region=eu&mkt=h2h",
  apiKey)

var sport Sports
var sites Sites
var data Data

func main() {

  sportSoccer := new(Sports)
  sportBaseBall := new(Sports)
  sportBasketBall := new(Sports)
  sportRugby := new(Sports)


  db := DbController{}
  db.connectDatabase()

  db.dropAndAutoMigrate(&sites)
  db.dropAndAutoMigrate(&data)
  db.dropAndAutoMigrate(&sport)

  var wg sync.WaitGroup
  wg.Add(4)

  go func() {
    getAllSports(baseSoccer, &sportSoccer)
    fmt.Println(sportSoccer)
    db.createModel(sportSoccer)
    wg.Done()
  }()

  go func() {
    getAllSports(baseBaseBall, &sportBaseBall)
    fmt.Println(sportBaseBall)
    db.createModel(sportBaseBall)
    wg.Done()
  }()

  go func() {
    getAllSports(baseBasketBall, &sportBasketBall)
    fmt.Println(sportBasketBall)
    db.createModel(sportBaseBall)
    wg.Done()
  }()

  go func() {
    getAllSports(baseRugby, &sportRugby)
    fmt.Println(sportRugby)
    db.createModel(sportRugby)
    wg.Done()
  }()

  wg.Wait()

  db.retrieveDataAndSites()
}
