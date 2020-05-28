/*
 * biblebot - A cross-platform, fast and lightweight discord bot to
 * quote the Bible.
 *
 * Version 1.0.0
 * Copyright (c) 2020, Free Software Foundation Inc., NJB/ServusDei2018 (servusdei@programmer.net)
 *
 * # Features
 *   - Fully cross-platform (Mac, Windows, Linux)
 * 	 - Fully configurable via the command line
 *   - Performant, lightweight and optimized design
 *   - Limit the amount of verses that may be requested at one time
 *   - Configure how long the bot pauses after posting a message
 *
 * # Commands (to be entered on a channel the bot has access to)
 *   - "#biblebot BOOK CHAPTER VERSE" - quote the specified verse
 *   - "#biblebot BOOK CHAPTER VERSE:VERSE" - quote the specified verses
 *   - "#biblebot status" - report whether the bot is online
 *   - "#biblebot version" - display biblebot's version
 *   - "#biblebot help" - display information on biblebot's commands
 *
 * # License
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
 * MA 02110-1301, USA.
 */

package main

// Import required packages
import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Declare global constants
const (
	Credits = `
	--- biblebot credits ---
	Copyright (c) 2020, Free Software Foundation Inc., NJB/ServusDei2018

	Distributed under the terms of the GNU General Public License as published by the Free Software Foundation.
	`
	Usage = `
	--- biblebot usage ---
#biblebot BOOK VERSE
#biblebot BOOK VERSE:VERSE
#biblebot credits
#biblebot status
#biblebot help
#biblebot version
	----------------------
	`
	Status = "BibleBot is online"
	Version = "BibleBot 1.0.0"
	API_Endpoint = `http://labs.bible.org/api/?passage=`
)

// Declare global variables
var (
	// Variables used for command line parameters
	Token string	// Bot token
	MaxVerses int64	// Maximum verses to quote
	Timeout int64	// Timeout in seconds after posting a message
	// Variables for internal operation
	Mutex sync.Mutex
)

// Initialization behavior
func init() {
	flag.StringVar(&Token, "token", "", "bot token")
	flag.Int64Var(&MaxVerses, "verses", 5, "maximum amount of verses to serve")
	flag.Int64Var(&Timeout, "timeout", 3, "rate limit in seconds")

	flag.Parse() // Parse command-line flags
}

// Main application entrypoint
func main() {
	// Check to make sure a token was provided
	if Token == "" {
		fmt.Println("A valid bot token must be provided.")
		fmt.Println(`See "./biblebot --help" for more information.`)

		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating discord session:", err)
		return
	}

	// Register messageCreate() as a callback for MessageCreate events
	dg.AddHandler(messageCreate)

	err = dg.Open() // Open a websocket connection to discord and listen
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	// Wait until CTRL-C or other term signal is received
	fmt.Println("BibleBot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close() // Close the discord session.
}

// messageCreate() is called every time a message is created on a
// channel that this bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ensure that we reply to one message at a time
	Mutex.Lock()
	defer Mutex.Unlock()

	// Ignore our own messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Put message in lowercase and split it by spaces
	args := strings.Split(strings.ToLower(m.Content), " ")

	// This message is in the proper format to be a verse(s) request
	// #biblebot John 3 2
	// #biblebot John 3 4-6
	if args[0] == "#biblebot" && len(args) == 4 {
		var bk, ch string
		var startVerse, endVerse int64; var err error

		bk = args[1]; ch = args[2]

		if strings.Contains(args[3], "-") {
			// Handle multiple verses
			verses := strings.Split(args[3], "-")

			startVerse, err = strconv.ParseInt(verses[0], 10, 8)
			if err != nil { return }
			endVerse, err = strconv.ParseInt(verses[1], 10, 8)
			if err != nil { return }
		} else {
			// They only want one verse
			startVerse, err = strconv.ParseInt(args[3], 10, 8)
			if err != nil { return }
			endVerse = startVerse
		}

		if startVerse > endVerse {
			// The starting verse occurs after the ending verse
			s.ChannelMessageSend(m.ChannelID, "error: starting verse is smaller than ending verse")
		} else if (endVerse - startVerse) > MaxVerses {
			// Too many verses are to be quoted
			s.ChannelMessageSend(m.ChannelID, "error: you are not allowed to quote that many verses")
		} else {
			// Get the requested verse(s) and send them
			msg, err := Get_Verses(bk, ch, startVerse, endVerse)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("api error: %v", err))
			} else {
				s.ChannelMessageSend(m.ChannelID, msg)
			}
			time.Sleep(time.Duration(Timeout)*time.Second)
		}
	} else if args[0] == "#biblebot" { // If it's not a verse request it could be a command
		if len(args) > 1 {
			switch strings.ToLower(args[1]) {
			case "credits":
				// Show them the credits
				s.ChannelMessageSend(m.ChannelID, Credits)
			case "status":
				// They want to know if we're there
				s.ChannelMessageSend(m.ChannelID, Status)
			case "version":
				// They want to know the version
				s.ChannelMessageSend(m.ChannelID, Version)
			default:
				// They didn't send a proper command
				s.ChannelMessageSend(m.ChannelID, Usage)
			}
			time.Sleep(time.Duration(Timeout)*time.Second)
		}
	}
}

// Get the specified verses from the Bible
func Get_Verses(book, chapter string, startVerse, endVerse int64) (string, error) {
	var query string

	if endVerse == startVerse {
		query = fmt.Sprintf("%s%s%s:%v&formatting=plain", API_Endpoint, book, chapter, startVerse)
	} else {
		query = fmt.Sprintf("%s%s%s:%v-%v&formatting=plain", API_Endpoint, book, chapter, startVerse, endVerse)
	}

	fmt.Println(query)

	resp, err := http.Get(query)
	defer resp.Body.Close()
	if err != nil { return "", err }

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil { return "", err }

	return string(content), nil
}
