package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

func main()  {
  response, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?id=524901&appid=bb7302fe136e4097821c1f74e3a51765")

  if err != nil {
    fmt.Print(err.Error())
    os.Exit(1)
  }

  responseData, err := ioutil.ReadAll(response.Body)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(responseData))

}
