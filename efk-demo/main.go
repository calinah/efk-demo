package main

import (
  "net/http"
  "fmt"
  "log"
  "encoding/json"
  "time"
  "github.com/icrowley/fake"
  "math/rand"
)

var (
  spamToggle bool
  behaveStatus bool
  validCodes []int16
)

type logMessage struct {
  Level interface{} `json:"level"`
  Code interface{} `json:"code"`
  Message interface{} `json:"message"`
  Date interface{} `json:"date"`
}

func main() {
  spamToggle = false
  behaveStatus = true
  validCodes = []int16{200, 404, 503}
  fmt.Print("Hello world")

  http.HandleFunc("/spam", func(w http.ResponseWriter, r *http.Request) {
    spamToggle = !spamToggle
    fmt.Fprintf(w, "Turned spam %v", spamToggle)
  })

  http.HandleFunc("/behave", func(w http.ResponseWriter, r *http.Request) {
    behaveStatus = !behaveStatus
    fmt.Fprintf(w, "Am I behaving? %v", behaveStatus)
  })

  http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "I'm up and running")
  })

  go func() {
    for {
      if spamToggle {
        fmt.Println(generateLog())
      }
    }
  }()

  log.Fatal(http.ListenAndServe(":4003", nil))

}

func generateLog () string {
  var (
    code interface{}
    msg interface{}
  )
  if behaveStatus || rand.Intn(2) == 0 {
    code = fake.FirstName()
    msg = fake.Paragraph()
  } else {
    code = validCodes[rand.Intn(len(validCodes))]
    msg = rand.Int()
  }
  date := time.Now().String()

  log := &logMessage{
    Level: "Info",
    Code: code,
    Message: msg,
    Date: date,
  }
  jsonLog, _ := json.Marshal(log)
  return string(jsonLog)

}



// truthy
