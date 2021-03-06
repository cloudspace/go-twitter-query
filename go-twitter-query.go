package main

import(
  "fmt"
  "github.com/ChimeraCoder/anaconda"
  "net/url"
  "encoding/json"
  "os"
)

type Args struct {
  Consumer_key         string `json:"consumer_key"`
  Consumer_secret      string `json:"consumer_secret"`
  Access_token         string `json:"access_token"`
  Access_token_secret  string `json:"Access_token_secret"`
  Query                string `json:"query"`
  Limit                string `json:"limit"`
  Since                string `json:"since"`
}

func search_request(a *Args) (anaconda.SearchResponse) {
  anaconda.SetConsumerKey(a.Consumer_key)
  anaconda.SetConsumerSecret(a.Consumer_secret)
  api:= anaconda.NewTwitterApi(a.Access_token, a.Access_token_secret)

  // make the request to search end point with optional params or defaults
  v := url.Values{}
  v.Set("count", string(a.Limit))
  v.Set("since_id", string(a.Since))
  result, err := api.GetSearch(string(a.Query), v)
  if err != nil {
    fmt.Println(err)
    //return anaconda.SearchResponse(err)
  }
  return result
}

func main() {
  a := &Args{Limit: "1", Since: "0"}

  a.Consumer_key = os.Args[1]
  a.Consumer_secret = os.Args[2]
  a.Access_token = os.Args[3]
  a.Access_token_secret = os.Args[4]
  a.Query = os.Args[5]
  a.Limit = os.Args[6]
  a.Since = os.Args[7]

  //json_str := []byte(os.Args[1])
  //err := json.Unmarshal(json_str, a)

  
  //if err != nil {
    //fmt.Println(err)
  //}

  result := search_request(a)

  result_json,err := json.Marshal(result)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(result_json))
}
