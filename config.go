package main

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
)

const REQUIRED_ENV string = "DBUS_SESSION_BUS_ADDRESS"

var ErrNoRequiredDBUSEnv = errors.New("daypaper requires the -d option to be set when running through cronjob")

func (app *App) newConfig() error {
	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	appConfigPath := filepath.Join(userConfigPath, "daypaper")
	if f, err := os.Open(appConfigPath); err != nil {
		defer f.Close()

		err = os.Mkdir(appConfigPath, 0700)
		if err != nil {
			return err
		}
	}

	app.config.ConfigPath = appConfigPath

	err = app.readToken()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) parseOpts() error {
	_, err := flags.Parse(&app.opts)
	if err != nil {
		var e *flags.Error
		if errors.As(err, &e) {
			if err.(*flags.Error).Type == flags.ErrHelp {
				os.Exit(0)
			}
		}
		return err
	}

	if os.Getenv(REQUIRED_ENV) == "" && app.opts.DBUSEnv == "" {
		return ErrNoRequiredDBUSEnv
	}

	return nil
}

func (app *App) readToken() error {
	c, err := ioutil.ReadFile(filepath.Join(app.config.ConfigPath, ".token"))
	if err != nil {
		return ErrNoToken
	}

	app.config.Token = strings.TrimSuffix(string(c), "\n")
	return nil
}

func (app *App) getCurrentDaytime() string {
	h := time.Now().Hour()

	if h >= 6 && h <= 10 {
		return "morning"
	} else if h >= 11 && h <= 13 {
		return "noon"
	} else if h >= 14 && h <= 17 {
		return "afternoon"
	} else if h >= 18 && h <= 20 {
		return "evening"
	}

	return "night"
}

func (app *App) getDaytimeForDownload() string {
	if app.opts.Time != "" {
		return app.opts.Time
	}

	return app.getCurrentDaytime()
}
