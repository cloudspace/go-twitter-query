package main

import(
  "flag"
  "fmt"
  "github.com/ChimeraCoder/anaconda"
  "reflect"
  "errors"
  "net/url"
  "encoding/json"
)

type Credentials struct {
  Consumer_key        string
  Consumer_secret     string
  Access_token        string
  Access_token_secret string
}

func inArray(a string, list []interface{}) (bool, error) {
    for _, b := range list {
        if b == a {
            return true, errors.New("flag missing")
        }
    }
    return false, nil
}

func check(c Credentials) (bool, error){
  // loop through struct fields and assign to values array
  vof := reflect.ValueOf(c)

  values := make([]interface{}, vof.NumField())

  for i := 0; i < vof.NumField(); i++ {
    values[i] = vof.Field(i).Interface()
  }

  // if values contains "blank", they missed a flag and we need to error
  return inArray("blank", values)
}

func format_results(r anaconda.SearchResponse) ([]interface{}){
  values := make([]interface{}, len(r.Statuses))

  for i := 0; i < len(r.Statuses); i++ {
    values[i] = r.Statuses[i].Text
  }

  return values
}

func main() {

  // declare struct
  c := new(Credentials)

  // declare cli flags (name, default, description)
  consumer_key := flag.String("consumer_key", "blank", "consumer key")
  consumer_secret := flag.String("consumer_secret", "blank", "consumer secret")
  access_token := flag.String("access_token", "blank", "access token")
  access_token_secret := flag.String("access_secret", "blank", "access token secret")
  query := flag.String("query", "blank", "search term")
  limit := flag.String("limit", "1", "number of results to return")
  since := flag.String("since", "0", "return results since this")
  
  flag.Parse()

  // assign parsed flags to struct
  c.Consumer_key = *consumer_key
  c.Consumer_secret = *consumer_secret
  c.Access_token = *access_token
  c.Access_token_secret = *access_token_secret

  // make sure none of them are still default values
  // throw simple 'flag missing' error if any of them are
  _, e := check(*c)
  if e != nil {
    fmt.Println(e)
  }

  // set up anaconda package creds
  anaconda.SetConsumerKey(c.Consumer_key)
  anaconda.SetConsumerSecret(c.Consumer_secret)
  api:= anaconda.NewTwitterApi(c.Access_token, c.Access_token_secret)

  // make the request to search end point
  v := url.Values{}
  v.Set("count", *limit)
  v.Set("since_id", *since)
  result, _ := api.GetSearch(*query, v)

  // return results json
  r,err := json.Marshal(result)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(r))

}
