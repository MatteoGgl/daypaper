package main

import (
	"os"
)

type App struct {
	config struct {
		ConfigPath string
		Token      string
	}
	opts struct {
		Force       bool   `short:"f" long:"force" description:"Forces a wallpaper refresh even when in the same time span"`
		Time        string `short:"t" long:"time" description:"Specify a particular time of day" choice:"morning" choice:"noon" choice:"afternoon" choice:"evening" choice:"night"`
		Search      string `short:"s" long:"search" description:"Additional text query to be added while searching"`
		Quality     int    `short:"q" long:"quality" description:"The downloaded image quality" default:"75"`
		Extension   string `short:"e" long:"ext" description:"The downloaded image extension" default:"jpg"`
		ScreenWidth string `short:"w" long:"width" description:"The downloaded image width" default:"1920"`
		Endpoint    string `short:"a" long:"api" description:"The API endpoint" default:"https://api.unsplash.com/photos/random"`
	}
}

func main() {
	app := &App{}

	err := app.newConfig()
	if err != nil {
		panic(err)
	}
	err = app.parseOpts()
	if err != nil {
		panic(err)
	}

	if app.opts.Force || app.shouldRefresh() {
		err = app.setWallpaper()
		if err != nil {
			panic(err)
		}
	}

	os.Exit(0)
}
