package main

import
(
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
)

type input struct {
  index int
  data string
}

func create(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err)
  }
  var fromReq input
  err = json.Unmarshal(body, &fromReq)
  if err != nil {
    panic(err)
  }
  if fromReq.index == 0 {
    fmt.Print(w, "Request received without index.")
    fmt.Fprint(w, "An index is required to create a doc.")
    panic("An index is required to create a doc.")
  }
  fmt.Println("continue with file write")
}

func main() {
  http.HandleFunc("/create", create)
  fmt.Println("listening on port 3000...")
  if err := http.ListenAndServe(":3000", nil); err != nil {
    log.Fatal(err)
  }
}
