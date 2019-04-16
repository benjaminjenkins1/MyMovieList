# MyMovieList

## App

React native app which consumes the public MyMovieList API at mymovielist.benjen.me

## API

JSON interface with the following endpoints:

### GET:
#### /user?id=[ID]
#### /lists
#### /list?id=[ID]
#### /logout

### POST:
#### /useroptions

Request body:
```
{
  "public":[bool]
}
```
#### /createlist

Request body:
```
{
  "name":[string]
}
```
#### /deletelist

Request body:
```
{
  "id":[int]
}
```
#### /listoptions

Request body:
```
{
  "public":[bool]
  "name":[string]
}
```
#### /addtolist

Request body:
```
{
  "items":[
    {
      "id":[int],
      "title":[string],
      "posterPath":[string],
      "status":[string],
      "runtime":[int],
      "overview":[string]
    }
  ]
}
```
#### /removefromlist

Request body:
```
{
  "items":[
    {
      "id":[int]
    }
  ]
}
```

### Google and Facebook authentication (OAuth2):
#### /facebook/login
#### /facebook/callback
#### /google/login
#### /google/callback
