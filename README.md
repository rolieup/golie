# Golie - A Golang ROLIE Implementation

![Golie Build CI](https://github.com/rolieup/golie/workflows/Golie%20Build%20CI/badge.svg)

A client/server implementation of the Resource-Oriented Lightweight Information Exchange (ROLIE) specification [RFC8322](https://tools.ietf.org/html/rfc8322)

## Usage

### Building rolie resources
```
golie new -h
Rolie files will be placed to the same directory enabling easy serving by http server software of your choice.

Usage:
  golie new SCAP-CONTENT-PATH [flags]

Flags:
  -h, --help              help for new
      --id string         ID for the rolie feed
      --root-uri string   URI to the feed itself. Example 'https://acme.org/my_rolie_content/
      --title string      Title for the rolie feed

Global Flags:
      --debug             Run in debug mode
      --loglevel string   Set log verbosity. Options are "debug", "info", "warn" or "error". (default "error")
```

### Fetching remote rolie resources
```
golie clone -h
Reads the rolie resource from the given URI and traverses any referenced documents. Everything is fetched locally with the same directory structure.

Usage:
  golie clone URI [flags]

Flags:
      --dir string   Directory to clone the feed into. (default "./")
  -h, --help         help for clone

Global Flags:
      --debug             Run in debug mode
      --loglevel string   Set log verbosity. Options are "debug", "info", "warn" or "error". (default "error")

```

## Installation

```
go get -u -v github.com/rolieup/golie/cmd/golie
```

## Building

Change directory to the `build` directory:
```
cd build
```

To build the Golie client, run the following from the `build` directory:

```
go build ../cmd/golie
```

To build the Golie server run the following from the `build` directory:

```
go build ../cmd/golied
```

You can then run the Golie client `golie` and Golie server `golied` from the
`build` directory.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
