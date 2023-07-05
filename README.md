# LOLESPORTDLE API

This API is made to interact with league of legends pro players simple data.
You can retrieve player's data such as :

````json
{
  "id": "109549750903142546",
  "summonerName": "Casting",
  "firstName": "Minje",
  "lastName": "Shin",
  "birthDate": "2003-10-30 00:00:00 +0000 UTC",
  "team": "Gen.G",
  "country": "KR",
  "teamUrl": "http://static.lolesports.com/teams/1655210113163_GenG_logo_200407-05.png",
  "role": "mid",
  "imageUrl": "http://static.lolesports.com/players/1686471944296_CL_GEN_Casting.png",
  "league": "LCK"
}
````

There is a simple documentation of consumable endpoints.

## Install

If you want to install this project, some prerequisites are mandatory.
First of all, you must create a file named `.env`, this file will contain the database's
connection string and the used port of the API (search for the .env template [right there](.env.template)).

Install all dependencies : 
```cmd
go mod download
````

Run the project
````cmd
go run .
````