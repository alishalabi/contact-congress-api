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

// type Senator struct {
//   name string `json:"name"`
//   state string `json:"state_name"`
//   term string `json:"term_end"`
//   website string `json:"website"`
//   facebook string `json:"facebook_url"`
//   twitter string `json:"twitter_url"`
//   photo string `json:"photo_url"`
// }

type Senator struct {
  state_name string `json:"state_name"`
  state_name_slug string `json:"state_name_slug"`
  state_code string `json:"-"`
  state_code_slug string`json:"state_code_slug"`
  class string `json:"-"`
  bioguide string `json:"bioguide"`
  thomas string `json:"-"`
  opensecrets string`json:"-"`
  votesmart string `json:"-"`
  fec string `json:"-"`
  maplight string`json:"-"`
  wikidata string `json:"-"`
  google_entity_id string`json:"-"`
  title string`json:"-"`
  party string `json:"party"`
  name string`json:"name"`
  name_slug string `json:"-"`
  first_name string `json:"-"`
  middle_name string `json:"-"`
  last_name string `json:"-"`
  name_suffix string `json:"-"`
  goes_by string `json:"-"`
  pronunciation string `json:"-"`
  gender string `json:"gender"`
  ethnicity string `json:"-"`
  religion string `json:"-"`
  openly_lgbtq string `json:"-"`
  date_of_birth string `json:"-"`
  entered_office string `json:"entered_office"`
  term_end string `json:"term_end"`
  biography string `json:"biography"`
  phone string `json:"phone"`
  fax string `json:"fax"`
  latitude string `json:"-"`
  longitude string `json:"-"`
  address_complete string `json:"-"`
  address_number string `json:"-"`
  address_prefix string `json:"-"`
  address_street string `json:"-"`
  address_sec_unit_type string `json:"-"`
  address_sec_unit_num string`json:"-"`
  address_city string `json:"-"`
  address_state string `json:"-"`
  address_zipcode string `json:"-"`
  address_type string `json:"-"`
  website string `json:"website"`
  contact_page string `json:"-"`
  facebook_url string `json:"facebook_url"`
  twitter_handle string`json:"twitter_handle"`
  twitter_url string `json:"twitter_url"`
  photo_url string `json:"photo_url"`
}

func main(){
  jsonFile, err := os.Open("bioguide_sample.json")

  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("Opened bioguide_sample.json")
  defer jsonFile.Close()

  fmt.Println(&jsonFile)

  byteValue, _ := ioutil.ReadAll(jsonFile)

  // var senators Senator
  var senators map[string]interface{}
  _ = json.Unmarshal([]byte(byteValue), &senators)
  if err != nil {
    fmt.Println(err)
  }

  fmt.Print(senators)

}
