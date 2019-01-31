# MyMovieList

## App

React native app which consumes the public MyMovieList API at mymovielist.benjen.me

## API

JSON interface offering the following endpoints:

GET:
- /user[id=ID]
- /lists
- /list?id=ID
- /logout

POST:
- /useroptions

Request body:
```
{
  "public":true
}
```
- /createlist
- /deletelist
- /listoptions
- /addtolist
- /removefromlist

Google and Facebook login:
- /facebook/login
- /facebook/callback
- /google/login
- /google/callback
