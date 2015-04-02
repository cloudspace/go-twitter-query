# go-twitter-query

## Input
Takes a JSON string. Does not play well with `\'`, `\"` preferred.

Required keys are:
  - consumer_key
  - consumer_secret
  - access_token
  - access_token_secret
  - query

`limit` and `since` are optional

`./go-twitter-query "{\"consumer_key\" : \"VAL\", \"consumer_secret\" : \"VAL\", \"access_token\" : \"VAL\", \"access_token_secret\" : \"VAL\", \"query\" : \"VAL\", \"since\" : \"VAL\", \"limit\" : \"VAL\"}"`

## Output
Returns JSON

Keys: `statuses`, `search_metadata`

`statuses` is an `[]` of JSON representing the relevant tweets

`search_metadata`:

`{"completed_in":0.012,"max_id":583390603335086080,"max_id_str":"583390603335086080","query":"%40cloudspace","refresh_url":"?since_id=583390603335086080\u0026q=%40cloudspace\u0026include_entities=1","count":1,"since_id":0,"since_id_str":"0"}}`

## Build

`docker build -t cgmoore120/go-twitter-query ./`

## Run

`docker run -ti cgmoore120/go-twitter-query:0.1.1 ./go-twitter-query "{\"consumer_key\" : \"VAL\", \"consumer_secret\" : \"VAL\", \"access_token\" : \"VAL\", \"access_token_secret\" : \"VAL\", \"query\" : \"VAL\", \"limit\" : \"VAL\", \"since\" : \"VAL\"}"`
