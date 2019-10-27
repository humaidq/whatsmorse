# WhatsMorse

## 1. Description
![Screenshot of WhatsMorse page](https://humaidq.ae/projects/screenshots/WhatsMorse.gif)

WhatsMorse is a web messaging application which translates all your messages to morse code written in a two hour ["Stupid" Hackathon by Transcend](https://www.meetup.com/transcenddubai/events/245505285/) in January of 2018.
The goal of the hackathon was to create something useless (can be anything, not limited to computer software).  

The web app can be accessed from the URL of the project above.

## 2. Requirements

The following packages must be installed on your system.

- Go
- Git

## 3. Copying and contributing

This program is written by Humaid AlQassimi, and is distributed
under the [MIT](https://humaidq.ae/license/mit) license.  

## 4. Download and install

```sh
$ go get -u git.sr.ht/~humaid/whatsmorse
$ go install git.sr.ht/~humaid/whatsmorse
```

## 5. Usage
To run the web app, `$PORT` must be set in the enviornment.
```sh
$ export PORT=8080
$ whatsmorse
```
The web app will be accessible at `http://localhost:8080`.
