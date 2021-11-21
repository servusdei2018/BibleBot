# biblebot (version 1.0.0)
A cross-platform, fast and lightweight discord bot to quote the Bible.

# Features
 - Fully cross-platform (Mac, Windows, Linux)
 - Fully configurable via the command line
 - Performant, lightweight and optimized design 
 - Limit the amount of verses that may be requested at one time
 - Configure how long the bot pauses after posting a message

# Commands
 - `#biblebot BOOK CHAPTER VERSE` - quote the specified verse
 - `#biblebot BOOK CHAPTER VERSE:VERSE` - quote the specified verses
 - `#biblebot status` - report whether the bot is online
 - `#biblebot version` - display biblebot's version
 - `#biblebot help` - display information on biblebot's commands
 
# Installation

 - First, download one of the following prebuilt binaries below. Just match up your OS and chipset.
 - Second, make sure you have a valid discord bot token. Refer to [discord's official documentation](discordapp.com/developers/applications) if you don't know how to get one.
 - Run the executable. On Mac/Linux, this shall be something like `./biblebot -token="PUT_YOUR_TOKEN_HERE"`. On Windows, you'll type something like `biblebot.exe -token="PUT_YOUR_TOKEN_HERE".`
 - If you like, you can configure the Max_Verses and Timeout configuration. For this, see the section entitled "Configuration."

## Downloads

### Mac

 - [biblebot_darwin_amd](https://github.com/servusdei2018/BibleBot/releases/download/1.0.0/BibleBot_1.0.0_darwin_amd64.tar.gz) - Mac (AMD processors)

### Windows

 - [biblebot_win_386.exe](https://github.com/servusdei2018/BibleBot/releases/download/1.0.0/BibleBot_1.0.0_windows_386.tar.gz) - Windows XP/Vista/7/8/10 (i386 processors)
 - [biblebot_win_amd.exe](https://github.com/servusdei2018/BibleBot/releases/download/1.0.0/BibleBot_1.0.0_windows_amd64.tar.gz) - Windows XP/Vista/7/8/10 (AMD processors)

### Linux

 - [biblebot_linux_386](https://github.com/servusdei2018/BibleBot/releases/download/1.0.0/BibleBot_1.0.0_linux_386.tar.gz) - Universal Linux (i386 processors)
 - [biblebot_linux_amd](https://github.com/servusdei2018/BibleBot/releases/download/1.0.0/BibleBot_1.0.0_linux_amd64.tar.gz) - Universal Linux (AMD processors)
 - [biblebot_linux_arm](https://github.com/servusdei2018/BibleBot/releases/download/1.0.0/BibleBot_1.0.0_linux_arm.tar.gz) - Universal Linux (ARM processors / Raspberry Pi)

## Configuration

BibleBot is highly configurable. You may configure the maximum amount of verses it can quote at a time (default is 5) as well as how long it pauses between posts, to make sure it doesn't spam (the default is 3 seconds). BibleBot is configured via commandline flags.

To configure the **maximum amount of verses**, use the "-verses=YOUR_AMOUNT_HERE" flag. For example, if I want my bot to only allow up to 10 verses in one quote, I'd use "-verses=10".

To configure the **timeout** in seconds to pause after posting a message, use the "-timeout=YOUR_TIMEOUT_HERE" flag. For example, if I want my bot to pause 30 seconds after posting a message, I'd use "-timeout=30".

Here are examples of running BibleBot on Linux or Mac where I put those flags into use. If you're on windows, simply think "biblebot.exe" instead of "./biblebot".

**Example One: Maximum verses of 4**
```
./biblebot -token="MY_TOKEN" -verses=4
```

**Example Two: Timeout of 5 seconds**
```
./biblebot -token="MY_TOKEN" -timeout=5
```

**Example Three: Maximum verses of 3 and timeout of 10 seconds**
```
./biblebot -token="MY_TOKEN" -verses=3 -timeout=10
```

# Building

If you are a developer and want to compile your own version of discordbot:

 - Make sure you have a working Go environment with Go installed.
 - Download this repository, and unzip it.
 - From a terminal, run either `go build -ldflags "-s -w"`, or if you have GNU Make, just type `make`.

# Contributing

Contributions are welcome. If you have a project that uses BibleBot, please tell us about it and we'll link it below.

# Copyright

Copyright (c) 2021, Nathanael Bracy

# License

**This program is free software;** you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but **WITHOUT ANY WARRANTY;** without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software
Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
MA 02110-1301, USA.
