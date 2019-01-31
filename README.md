# MyMovieList

## App

React native app which consumes the public MyMovieList API at mymovielist.benjen.me

## API

JSON interface offering the following endpoints:

### GET:
- /user[id=ID]
- /lists
- /list?id=ID
- /logout

### POST:
- /useroptions

Request body:
```
{
  "public":true
}
```
- /createlist

Request body:
```
{
  "name":"NAME"
}
```
- /deletelist

Request body:
```
{
  "id":ID
}
```
- /listoptions

Request body:
```
{
  "public":true
  "name":"NAME"
}
```
- /addtolist

Request body:
```
{
  "items":[
    {
      "id":ID,
      "title":"TITLE",
      "posterPath":"POSTER_PATH",
      "status":"STATUS",
      "runtime":RUNTIME,
      "overview":"OVERVIEW"
    }
  ]
}
```
- /removefromlist

Request body:
```
{
  "items":[
    {
      "id":ID
    }
  ]
}
```

### Google and Facebook login:
- /facebook/login
- /facebook/callback
- /google/login
- /google/callback
