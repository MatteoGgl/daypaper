# Daypaper

Daypaper sets your GNOME wallpaper based on the time of day from a random and relevant [Unsplash](https://unsplash.com) image.

The retrieved image is random and the default query will be the time of day as defined by daypaper. If the user specifies additional search terms, they will be appendend to the original query.

Additionally, the results are filtered by what Unsplash calls "Topics"; basically curated collections of photos. The selected topics right now are:

- Textures & Patterns
- Arts & Culture
- Wallpapers

## Installation

1. You will need an Access Token from Unsplash: create an app [here](https://unsplash.com/oauth/applications/new) after registration.
2. Create the necessary configuration files:

```bash
mkdir ~/.config/daypaper/
touch ~/.config/daypaper/.token
```

3. Paste your Access Token inside `.token`
4. Clone this repository
5. Run `go install`
4. You're ready to go!

## Usage

```bash
$: daypaper -h
Usage:
  daypaper [OPTIONS]

Application Options:
  -f, --force                                       Forces a wallpaper refresh even when in the same time span
  -t, --time=[morning|noon|afternoon|evening|night] Specify a particular time of day
  -s, --search=                                     Additional text query to be added while searching
  -q, --quality=                                    The downloaded image quality (default: 75)
  -e, --ext=                                        The downloaded image extension (default: jpg)
  -w, --width=                                      The downloaded image width (default: 1920)
  -a, --api=                                        The API endpoint (default: https://api.unsplash.com/photos/random)

Help Options:
  -h, --help                                        Show this help message
```

The simplest way to activate daypaper is to run it every hour:

```bash
$: crontab -e

0 * * * * ~/go/bin/daypaper >> ~/.daypaper.log 2>&1
```

Daypaper will contact the API only when needed (i.e. the current day period has changed).

The periods are defined like this:
```go
h := time.Now().Hour()

if h >= 7 && h <= 11 {
    return "morning"
} else if h >= 12 && h <= 14 {
    return "noon"
} else if h >= 15 && h <= 17 {
    return "afternoon"
} else if h >= 18 && h <= 21 {
    return "evening"
}

return "night"
```

## Like your wallpaper?

Find out who shot it by looking in `~/.config/daypaper/credits.txt`!

## Can Daypaper please do x feature/support my DE?

Just open an issue/PR and I'll gladly help!