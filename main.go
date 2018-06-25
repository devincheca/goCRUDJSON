package main

import
(
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

type input struct {
  Index int
  Data string
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
  if fromReq.Index == 0 {
    fmt.Print(w, "Request received without index.")
    fmt.Fprint(w, "An index is required to create a doc.")
    panic("An index is required to create a doc.")
  }
  filename := fmt.Sprintf("data", string(fromReq.Index), ".json")
  file, err := os.Create(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()
  contents := fmt.Sprintf("{\"Index:\":\"", string(fromReq.Index), "\",\"Data\":\"", fromReq.Data, "\"}")
  file.WriteString(contents)
  fmt.Println("file write complete")
}

func main() {
  http.HandleFunc("/create", create)
  fmt.Println("listening on port 3000...")
  if err := http.ListenAndServe(":3000", nil); err != nil {
    log.Fatal(err)
  }
}
