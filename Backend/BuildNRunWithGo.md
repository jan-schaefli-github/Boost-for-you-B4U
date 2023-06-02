# How to build the server.go to executable and run it..

## Requirements
- Have GO installed on you system. ---> https://go.dev
- Have the .env file configured correctly.

### .env
```
ACCESS_TOKEN='<ClashRoyal-API_TOKEN>'
CLAN_TAG='<ClanTag>'
PORT=<Port>
TRUSTED_PROXIES='*'
```
##### :information_source: Info:
- ClashRoyal-API_TOKEN: You can get at https://developer.clashroyale.com
- ClanTag: You can look in Game.
- Port: Is the Port you wish the application to be running on.
- TRUSTED_PROXIES: Trusted proxies is a security measurement if you want to restrict access to the api in your internal network. You can leave it at * for no restrictions.


## Build you exe
First you'll have to manually add a build directory if it does not exist already.
```
$ go build -o build/b4u_backend.exe server.go
```

 ## Start the go application for testing
```
$ go run .
```

### :information_source: Info:
You have to run the commands above in the Folder with the go.mod file or in this Case the Backend Folder or the won't work.