package main

import (
	"fmt"
	"os"
)

var (
	version   string
	buildTime string
)

type App struct {
	config struct {
		ConfigPath string
		Token      string
	}
	opts struct {
		Version     bool   `short:"v" long:"version" description:"Displays version and build info"`
		Credit      bool   `short:"c" long:"credit" description:"Displays the current wallpaper author and link"`
		Force       bool   `short:"f" long:"force" description:"Forces a wallpaper refresh even when in the same time span"`
		Time        string `short:"t" long:"time" description:"Specify a particular time of day" choice:"morning" choice:"noon" choice:"afternoon" choice:"evening" choice:"night"`
		Search      string `short:"s" long:"search" description:"Additional text query to be added while searching"`
		Quality     int    `short:"q" long:"quality" description:"The downloaded image quality" default:"75"`
		Extension   string `short:"e" long:"ext" description:"The downloaded image extension" default:"jpg"`
		ScreenWidth string `short:"w" long:"width" description:"The downloaded image width" default:"1920"`
		Endpoint    string `short:"a" long:"api" description:"The API endpoint" default:"https://api.unsplash.com/photos/random"`
		DBUSEnv     string `short:"d" long:"dbus" description:"Set this variable to the value of $DBUS_SESSION_BUS_ADDRESS; only needed when running through cronjob" `
	}
}

func main() {
	app := &App{}

	err := app.parseOpts()
	if err != nil {
		os.Exit(1)
	}

	if app.opts.Version {
		fmt.Printf("Version:\t%s\n", version)
		fmt.Printf("Build time:\t%s\n", buildTime)
		os.Exit(0)
	}

	err = app.newConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if app.opts.Credit {
		err = app.showCredits()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if app.opts.Force || app.shouldRefresh() {
		err = app.setWallpaper()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
