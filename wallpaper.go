package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type photo struct {
	ID   string
	URLs struct {
		Raw string
	}
	User struct {
		Name  string
		Links struct {
			HTML string
		}
	}
}

func (app *App) setWallpaper() error {
	p, err := app.getRandomPhoto()
	if err != nil {
		return err
	}

	downloadedPhotoPath, err := app.downloadPhoto(p)
	if err != nil {
		return err
	}
	app.saveCredits(p)

	app.cleanUnusedFiles(filepath.Base(downloadedPhotoPath))

	app.saveTime()

	err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", strconv.Quote("file://"+downloadedPhotoPath)).Run()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) getRandomPhoto() (*photo, error) {
	c := http.DefaultClient

	qs := url.Values{
		"client_id":      []string{app.config.Token},
		"topics":         []string{strings.Join([]string{"textures-patterns", "arts-culture", "wallpapers"}, ",")},
		"orientation":    []string{"landscape"},
		"content_filter": []string{"high"},
		"query":          []string{strings.Join([]string{app.getDaytimeForDownload(), app.opts.Search}, " ")},
	}

	apiURL, err := joinURLWithQuery(app.opts.Endpoint, qs)
	if err != nil {
		return nil, err
	}

	res, err := c.Get(apiURL)
	if err != nil {
		return nil, err
	}

	p := photo{}
	err = decodeJSONBody(res, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
