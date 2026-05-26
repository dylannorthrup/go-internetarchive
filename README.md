# InternetArchive

[![license](https://img.shields.io/github/license/dylannorthrup/internetarchive.svg)](https://github.com/dylannorthrup/internetarchive/blob/master/LICENSE)
[![paypal](https://img.shields.io/badge/donate-paypal-009cdf?logo=paypal)](https://paypal.me/nektro)
[![goreportcard](https://goreportcard.com/badge/github.com/dylannorthrup/internetarchive)](https://goreportcard.com/report/github.com/dylannorthrup/internetarchive)
[![downloads](https://img.shields.io/github/downloads/dylannorthrup/internetarchive/total.svg)](https://github.com/dylannorthrup/internetarchive/releases)

`ia` is a command-line interface for interacting with https://archive.org/ written in Golang. It is an alternative to the offical Python client available here https://github.com/jjjake/internetarchive.

## Usage

- `ia`
```
Usage:
  ia [command]

Available Commands:
  download    download an item or collection
  help        Help about any command
  metadata    retrieve metadata for items and collections

Flags:
  -h, --help                help for ia
      --mbpp-bar-gradient   Enabling this will make the bar gradient from red/yellow/green.
```

- `ia download`
```
Usage:
  ia download {item_name} [flags]

Flags:
  -c, --concurrency int   number of concurrent download jobs to run at once (default 10)
      --dense             when enabled, stores items based on their creation date
  -h, --help              help for download
      --only-meta         when enabled, only saves _meta.xml files
  -o, --save-dir string    (default "./data")
      --resume            When enabled, performs a deeper check for item completion
```

## Built With
- https://github.com/spf13/cobra
- https://github.com/nektro/go-util
- https://github.com/PuerkitoBio/goquery

## License
MIT
