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
)

type bioInfo struct {
  name string `json:"name"`
  state string `json:"state_name"`
  term string `json:"term_end"`
  website string `json:"website"`
  facebook string `json:"facebook_url"`
  twitter sting `json:"twitter_url"`
  photo string `json:"photo_url"`
}

func main(){
  fmt.Println("Hello World")
}
