# Daypaper

[![Go Report Card](https://goreportcard.com/badge/github.com/matteoggl/daypaper)](https://goreportcard.com/report/github.com/matteoggl/daypaper)

Daypaper sets your GNOME wallpaper based on the time of day from a random and relevant [Unsplash](https://unsplash.com) image.

The retrieved image is random and the default query will be the time of day as defined by daypaper. If the user specifies additional search terms, they will be appendend to the original query.

Additionally, the results are filtered by what Unsplash calls "Topics"; basically curated collections of photos. The selected topics right now are:

- Textures & Patterns
- Arts & Culture
- Wallpapers

## Installation

You can either:
- Build daypaper from source; clone this repo and run `make build`.
- Download and run the latest binary from the [releases page](https://github.com/MatteoGgl/daypaper/releases).

## Configuration

1. You will need an Access Token from Unsplash: create an app [here](https://unsplash.com/oauth/applications/new) after registration.
2. Create the necessary configuration files:

```bash
mkdir ~/.config/daypaper/
touch ~/.config/daypaper/.token
```

3. Paste your Access Token inside `.token`
4. You're ready to go!

## Usage

```bash
$: daypaper -h
Usage:
  daypaper [OPTIONS]

Application Options:
  -c, --credit                                      Displays the current wallpaper author and link
  -f, --force                                       Forces a wallpaper refresh even when in the same time span
  -t, --time=[morning|noon|afternoon|evening|night] Specify a particular time of day
  -s, --search=                                     Additional text query to be added while searching
  -q, --quality=                                    The downloaded image quality (default: 75)
  -e, --ext=                                        The downloaded image extension (default: jpg)
  -w, --width=                                      The downloaded image width (default: 1920)
  -a, --api=                                        The API endpoint (default: https://api.unsplash.com/photos/random)
  -d, --dbus=                                       Set this variable to the value of $DBUS_SESSION_BUS_ADDRESS; only needed when running through cronjob

Help Options:
  -h, --help                                        Show this help message
```

You can activate daypaper manually. If you don't like the random photo downloaded for this time span, you can always force a new download or specify a different time. You can even add search terms to further customize your results.

The simplest way to activate daypaper automatically is to run it every hour. You will need to specify the contents of your env variable DBUS_SESSION_BUS_ADDRESS: it's required by `gsettings` to correctly set the wallpaper. It's not available inside cronjobs because they work with a reduced set of env variables. To find out it's value type:
```bash
$: echo $DBUS_SESSION_BUS_ADDRESS
```

Then pass it as an option:

```bash
$: crontab -e

0 * * * * ~/go/bin/daypaper --dbus="WHATEVER_THE_VALUE_IS" >> ~/.daypaper.log 2>&1
```

Daypaper will contact the API only when needed (i.e. the current day period has changed).

The periods are defined like this:
```go
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
```

## Like your wallpaper?

Find out details on who shot it by using `daypaper -c`!

## Can Daypaper please do x feature/support my DE?

Just open an issue/PR and I'll gladly help!