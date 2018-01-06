package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-macaron/session"
	macaron "gopkg.in/macaron.v1"
)

type UserData struct {
	transmissions map[int]string
}

var connectedUsers map[string]*UserData

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	rand.Seed(time.Now().UTC().UnixNano())
	connectedUsers = make(map[string]*UserData)
	macaron.Env = macaron.PROD
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())

	m.Get("/", func(ctx *macaron.Context, sess session.Store) {
		if sess.Get("Code") == nil {
			number := fmt.Sprint(rand.Intn(8999) + 1000)
			connectedUsers[number] = &UserData{make(map[int]string)}
			sess.Set("Code", number)
			ctx.Data["Code"] = number
		} else {
			ctx.Data["Code"] = sess.Get("Code")
		}

		ctx.HTML(200, "index")
	})

	m.Get("/transmit", func(ctx *macaron.Context, sess session.Store) {
		_, ok := connectedUsers[ctx.Query("to")]
		if sess.Get("Code") != nil && ctx.Query("message") != "" && ctx.Query("to") != "" && ok {
			connectedUsers[ctx.Query("to")].transmissions[len(connectedUsers[ctx.Query("to")].transmissions)] = fmt.Sprint(sess.Get("Code")) + ": " + toMorse(ctx.Query("message"))
			fmt.Println("All okay!")
		}
	})

	m.Get("/clear", func(ctx *macaron.Context, sess session.Store) {
		connectedUsers[fmt.Sprint(sess.Get("Code"))] = &UserData{make(map[int]string)}
	})

	m.Get("/incoming", func(ctx *macaron.Context, sess session.Store) {
		if sess.Get("Code") != nil {
			ctx.Data["Trans"] = connectedUsers[fmt.Sprint(sess.Get("Code"))].transmissions
		}
		ctx.HTML(200, "incoming")
	})
	log.Println(http.ListenAndServe("0.0.0.0:"+port, m))

}

// Function modified from https://github.com/pravj/morser
func toMorse(text string) string {
	text = strings.ToUpper(text)
	reverseMap := make(map[string]string)

	// Alphabates
	reverseMap["A"] = ".-"
	reverseMap["B"] = "-..."
	reverseMap["C"] = "-.-."
	reverseMap["D"] = "-.."
	reverseMap["E"] = "."
	reverseMap["F"] = "..-."
	reverseMap["G"] = "--."
	reverseMap["H"] = "...."
	reverseMap["I"] = ".."
	reverseMap["J"] = ".---"
	reverseMap["K"] = "-.-"
	reverseMap["L"] = ".-.."
	reverseMap["M"] = "--"
	reverseMap["N"] = "-."
	reverseMap["O"] = "---"
	reverseMap["P"] = ".--."
	reverseMap["Q"] = "--.-"
	reverseMap["R"] = ".-."
	reverseMap["S"] = "..."
	reverseMap["T"] = "-"
	reverseMap["U"] = "..-"
	reverseMap["V"] = "...-"
	reverseMap["W"] = ".--"
	reverseMap["X"] = "-..-"
	reverseMap["Y"] = "-.--"
	reverseMap["Z"] = "--.."

	// Decimals
	reverseMap["1"] = ".----"
	reverseMap["2"] = "..---"
	reverseMap["3"] = "...--"
	reverseMap["4"] = "....-"
	reverseMap["5"] = "....."
	reverseMap["6"] = "-...."
	reverseMap["7"] = "--..."
	reverseMap["8"] = "---.."
	reverseMap["9"] = "----."
	reverseMap["0"] = "-----"

	// Punctuation marks and miscellaneous signs
	reverseMap["."] = ".-.-.-"
	reverseMap[","] = "--..--"
	reverseMap[":"] = "---..."
	reverseMap["?"] = "..--.."
	reverseMap["'"] = ".----."
	reverseMap["-"] = "-....-"
	reverseMap["/"] = "-..-."
	reverseMap["("] = "-.--."
	reverseMap[")"] = "-.--.-"
	reverseMap["+"] = ".-.-."
	reverseMap["×"] = "-..-"
	reverseMap["@"] = ".--.-."
	result := ""

	for j := 0; j < len(text); j++ {
		if text[j] == ' ' {
			result = result + " / "
		} else {
			result = result + " "
		}

		// validates the input text, character by character
		_, isPresent := reverseMap[fmt.Sprintf("%c", text[j])]

		// invalid input provided
		if isPresent {
			result = result + reverseMap[fmt.Sprintf("%c", text[j])]
		}

	}
	return result

}
