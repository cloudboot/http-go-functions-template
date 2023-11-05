package cloudboot

import (
  "encoding/json"
  "fmt"
  "html"
  "net/http"

  "github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
   functions.HTTP("Main", main)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func main(w http.ResponseWriter, r *http.Request) {
  var d struct {
    Name string `json:"name"`
  }
  if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
    fmt.Fprint(w, "It works!")
    return
  }
  if d.Name == "" {
    fmt.Fprint(w, "It works!")
    return
  }
  fmt.Fprintf(w, "It works, %s!", html.EscapeString(d.Name))
}
