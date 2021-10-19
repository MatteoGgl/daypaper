package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var (
	ErrNoToken = errors.New("cannot find .token")
)

func (app *App) downloadPhoto(p *photo) (string, error) {
	downloadPath := filepath.Join(app.config.ConfigPath, p.ID+"."+app.opts.Extension)

	qs := url.Values{
		"fm": []string{app.opts.Extension},
		"q":  []string{fmt.Sprintf("%d", app.opts.Quality)},
		"w":  []string{app.opts.ScreenWidth},
	}

	downloadURL, err := joinURLWithQuery(p.URLs.Raw, qs)
	if err != nil {
		return "", err
	}

	r, err := http.Get(downloadURL)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	f, err := os.Create(downloadPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = io.Copy(f, r.Body)
	if err != nil {
		return "", err
	}

	return downloadPath, nil
}

func (app *App) cleanUnusedFiles(toPreserve string) error {
	files, err := os.ReadDir(app.config.ConfigPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.Name() != toPreserve && f.Name() != ".token" && f.Name() != "credits.txt" {
			os.Remove(filepath.Join(app.config.ConfigPath, f.Name()))
		}
	}

	return nil
}

func (app *App) saveTime() error {
	out, err := os.Create(filepath.Join(app.config.ConfigPath, app.getCurrentDaytime()))
	if err != nil {
		return err
	}
	defer out.Close()

	return nil
}

func (app *App) saveCredits(p *photo) error {
	f, err := os.Create(filepath.Join(app.config.ConfigPath, "credits.txt"))
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	ls := []string{
		"Photo by " + p.User.Name,
		p.User.Links.HTML,
	}
	for _, l := range ls {
		_, err := w.WriteString(l + "\n")
		if err != nil {
			return err
		}
	}

	w.Flush()
	return nil
}

func (app *App) showCredits() error {
	c, err := ioutil.ReadFile(filepath.Join(app.config.ConfigPath, "credits.txt"))
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, string(c))

	return nil
}

func (app *App) shouldRefresh() bool {
	f, err := os.Open(filepath.Join(app.config.ConfigPath, app.getCurrentDaytime()))
	if err != nil {
		return true
	}
	defer f.Close()

	return false
}
