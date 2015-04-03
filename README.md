# go-twitter-query

## Input
Passed in values in order are:

  - consumer_key
  - consumer_secret
  - access_token
  - access_token_secret
  - query
  - limit
  - since

`./go-twitter-query "CK_VAL" "CS_VAL" "AT_VAL" "ATS_VAL" "Q_VAL" "L_VAL" "S_VAL"`

## Output
Returns JSON

Keys: `statuses`, `search_metadata`

`statuses` is an `[]` of JSON representing the relevant tweets

`search_metadata`:

`{"completed_in":0.012,"max_id":583390603335086080,"max_id_str":"583390603335086080","query":"%40cloudspace","refresh_url":"?since_id=583390603335086080\u0026q=%40cloudspace\u0026include_entities=1","count":1,"since_id":0,"since_id_str":"0"}}`

## Build

`docker build -t cloudspace/go-twitter-query ./`

## Run

`docker run -ti cloudspace/go-twitter-query:0.2.0 ./go-twitter-query "CK_VAL" "CS_VAL" "AT_VAL" "ATS_VAL" "Q_VAL" "L_VAL" "S_VAL"`
