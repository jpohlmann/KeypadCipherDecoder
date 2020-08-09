### Description

Keypad cipher decoder for Hunt a Killer Cadence Theater episode.  Tries all combinations of keypads, and scores them based on the number of dictionary words in result.

Written in go

### Requirements
Requires go1.14.4.  Recommended to have a program like Postman, but you can use a browser too

### Usage

From the command line, in the root of the project, run 

    go build ./main.go 
    ./main

using Postman (or the browser), make a GET request to 

    http://localhost:3000/find?code=????
    
You should retrieve a result with the best match.  You can also see all matches tested in the logs.