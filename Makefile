GO = go build
GOARGS = -ldflags "-s -w"

all:
	# Building BibleBot for local machine; architecture local...
	$(GO) -o biblebot $(GOARGS)
	# Done.

darwin:
	# Building BibleBot for Darwin; architecture 386...
	GOOS="darwin" GOARCH="386" $(GO) -o biblebot_darwin_386 $(GOARGS)
	# echo Building BibleBot for Darwin; architecture AMD...
	GOOS="darwin" GOARCH="adm" $(GO) -o biblebot_darwin_amd $(GOARGS)
	# Done.

linux:
	# Building BibleBot for GNU/Linux; architecture 386...
	GOOS="linux" GOARCH="386" $(GO) -o biblebot_linux_386 $(GOARGS)
	# Building BibleBot for GNU/Linux; architecture AMD...
	GOOS="linux" GOARCH="amd" $(GO) -o biblebot_linux_amd $(GOARGS)
	# Building BibleBot for GNU/Linux; architecture ARM...
	GOOS="linux" GOARCH="arm" $(GO) -o biblebot_linux_arm $(GOARGS)
	# Done.

windows:
	# Building BibleBot for Windows; architecture 386...
	GOOS="windows" GOARCH="386" $(GO) -o biblebot_win_386.exe $(GOARGS)
	# Building BibleBot for Windows; architecture AMD...
	GOOS="windows" GOARCH="amd" $(GO) -o biblebot_win_amd.exe $(GOARGS)
	# Done.

clean:
	rm biblebot* # remove any compiled executables
	# Clean.

license:
	make credits

credits:
	# Copyright (c) 2021, Nathanael Bracy

	# This program is free software; you can redistribute it and/or modify
	# it under the terms of the GNU General Public License as published by
	# the Free Software Foundation; either version 3 of the License, or
	# (at your option) any later version.
	#
	# This program is distributed in the hope that it will be useful,
	# but **WITHOUT ANY WARRANTY;** without even the implied warranty of
	# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	# GNU General Public License for more details.
	#
	# You should have received a copy of the GNU General Public License
	# along with this program; if not, write to the Free Software
	# Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston,
	# MA 02110-1301, USA.

version:
	# BibleBot version 1.0.0
