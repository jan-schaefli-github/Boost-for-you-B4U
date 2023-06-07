# How to build the server.go to executable and run it..

## Requirements

- Have GO installed on you system. ---> https://go.dev
- Have the .env file configured correctly.
- My Sql Server with our database script running ---> see Database Readme

## Steps
1. Setup the .env file
2. Build or test run the application.

### .env

```
# Request enviorment variables
ACCESS_TOKEN='<ClashRoyal-API_TOKEN>'
CLAN_TAG='<ClanTag>'
PORT=<Port>
TRUSTED_PROXIES='*'

# Database environment variables
DB_HOST=<Host of Database for local 'localhost'>
DB_PORT=<PORT of the Database>
DB_USER=<Datenbank User>
DB_PASS=<Datenbank Passwort>
DB_NAME=<Datenbank Name>
```

##### :information_source: Info - Request environment variables:
- ClashRoyal-API_TOKEN: You can get at https://developer.clashroyale.com
- ClanTag: You can look in Game.
- Port: Is the Port you wish the application to be running on.
- TRUSTED_PROXIES: Trusted proxies is a security measurement if you want to restrict access to the api in your internal
  network. You can leave it at * for no restrictions.

##### :information_source: Info - Database environment variables:
- DB_HOST: Where is your Database hosted? For local hosting just put ins 'localhost'
- DB_PORT: On which port is your Database running on? To find out type "SELECT @@port;" command in mysql console or set with option -P <The Port you want> on login.
- DB_USER= Which Database User should be used?
- DB_PASS: The password for that User?
- DB_NAME: The name of the database? If you run our database.sql script it should be "b4u". Go to the database md to find out.

## Build you exe

#### If you're on Windows just run: build.bat

#### If you're running on a linux mmachine
First you'll have to manually add a build directory if it does not exist already.

```
Standard (Windows)
$ go build -o build/b4u_backend.exe server.go

REM Build for Linux:
GOOS=linux
GOARCH=amd64
go build -o builds/b4u_backend_linux
```

## Start the go application for testing

```
$ go run .
```

### :information_source: Info:

You have to run the commands above in the Folder with the go.mod file or in this Case the Backend Folder or the won't
work.