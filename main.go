package main

/*
Todo:
- Create bioguide struct
- Parse JSON
- Query json by state, should return 2 senators
- Return corresponding bioguide structs
*/

import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
)

type Senator struct {
  name string `json:"name"`
  state string `json:"state_name"`
  term string `json:"term_end"`
  website string `json:"website"`
  facebook string `json:"facebook_url"`
  twitter string `json:"twitter_url"`
  photo string `json:"photo_url"`
}

func main(){
  jsonFile, err := os.Open("bioguide_sample.json")

  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Opened bioguide_sample.json")
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  // var senators Senator
  var senators map[string]interface{}
  json.Unmarshal([]byte(byteValue), &senators)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Print(senators)

}
