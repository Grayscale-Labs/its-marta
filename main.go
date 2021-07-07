package main

import (
  "fmt"
  "os"
  "net/http" // TODO: switch to "github.com/hashicorp/go-retryablehttp"
  "io/ioutil"
  "log"
  "time"
  "encoding/json"
  "github.com/replit/database-go"
)

var out *os.File

var baseUrl = "http://developer.itsmarta.com/RealtimeTrain/RestServiceNextTrain/GetRealtimeArrivals?apikey=%s"
var reqUrl = fmt.Sprintf(baseUrl, os.Getenv("API_KEY"))
var httpClient = &http.Client{}

var timeStrLayout = "2006-01-02 15:04:05.999999999 -0700 MST"
var lastWriteKey = "lastWrite"

type Stop struct {
  TrainId string
  Waiting_Seconds string
}

func main() {
  fmt.Println("MARTA is smarta") // TODO: remove

  var err error
  out, err = os.OpenFile("out.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  checkError(err)
  defer out.Close()

  // lastWriteFinished will serve as our mutex; see https://gobyexample.com/channel-synchronization
  lastWriteFinished := make(chan bool)
  for {
    lastWrite, err := getLastWriteTime()
    if (err != nil) {
      lastWrite = time.Now()
      setLastWriteTime(lastWrite)
    }

    if (time.Now().Sub(lastWrite) >= (0 * time.Second)) { // hmm, this seems wrong
      go writeFilteredStops(lastWriteFinished)
      <-lastWriteFinished // "release" mutex
      time.Sleep(10 * time.Second)
    }
  }
}

func getLastWriteTime() (time.Time, error) {
  lastWriteStr, _ := database.Get(lastWriteKey)
  return time.Parse(timeStrLayout, lastWriteStr)
}

func setLastWriteTime(t time.Time) {
  database.Set(lastWriteKey, t.Format(timeStrLayout))
}

func writeFilteredStops(finished chan bool) {
  stops := filterStops(fetchStops())

  marshalled, err := json.Marshal(stops)
  checkError(err)

  out.Write(append(marshalled, '\n'))
  out.Sync()
  setLastWriteTime(time.Now())

  finished <- true
}

func fetchStops() []Stop {
  resp, err := httpClient.Get(reqUrl)
  checkError(err)

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  checkError(err)

  var stops []Stop
  err = json.Unmarshal(body, &stops)
  checkError(err)
  return stops
}

func filterStops(stops []Stop) []Stop {
  // TODO: filter to only stops w/ waiting_seconds < 120
  return stops
}

func checkError(err error) {
  if (err != nil) {
    log.Panic(err)
  }
}
