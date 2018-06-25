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
  Key string
  Value string
}

type query struct {
  Key string
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
  if fromReq.Key == "" {
    fmt.Print(w, "Request received without key.")
    fmt.Fprint(w, "A key is required to create a doc.")
    panic("A key is required to create a doc.")
  }
  filename := fmt.Sprintf("data/data" + fromReq.Key + ".json")
  file, err := os.Create(filename)
  if err != nil {
    panic(err)
  }
  defer file.Close()
  contents := fmt.Sprintf("{\"Index:\":\"" + fromReq.Key + "\",\"Data\":\"" + fromReq.Value + "\"}")
  file.WriteString(contents)
  fmt.Println(filename + " write complete")
  fmt.Fprint(w, filename + " write complete")
}

func read(w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err)
  }
  var fromReq query
  err = json.Unmarshal(body, &fromReq)
  if err != nil {
    panic(err)
  }
  filename := fmt.Sprintf("data/data" + fromReq.Key + ".json")
  stream, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }
  fmt.Println(string(stream))
  fmt.Fprint(w, string(stream))
}

func main() {
  http.HandleFunc("/create", create)
  http.HandleFunc("/read", read)
  fmt.Println("listening on port 3000...")
  if err := http.ListenAndServe(":3000", nil); err != nil {
    log.Fatal(err)
  }
}
