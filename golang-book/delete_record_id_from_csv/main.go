package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
  fmt.Println("Starting the application...")
  response, err := http.Get("https://api.prod.cwsplatform.com/api/documentation#!/OEM_Specs/delete_specs_specId")
  if err != nil {
      fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
      data, _ := ioutil.ReadAll(response.Body)
      fmt.Println(string(data))
  }
  jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
  jsonValue, _ := json.Marshal(jsonData)
  response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
  if err != nil {
      fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
      data, _ := ioutil.ReadAll(response.Body)
      fmt.Println(string(data))
  }
  fmt.Println("Terminating the application...")
}
