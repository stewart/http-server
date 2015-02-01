# http-server

A Very Simple™, pretty dang fast static-file HTTP Server in Go.

Basically just the built-in `http.FileServer` handler, wrapped in a tiny bit of config from the `flag` stdlib package.

## Installation

With Go installed:

    $ go get github.com/stewart/http-server

## Precompiled Binaries

Thanks to the excellent [gox][] tool by [@mitchellh][], precompiled binaries
are available for Windows, Linux, OS X, FreeBSD, and NetBSD.

To find them, check out the [Releases][] page.

[gox]: https://github.com/mitchellh/gox
[@mitchellh]: https://twitter.com/mitchellh
[Releases]: https://github.com/stewart/http-server/releases

## Usage

Serving files from `./` on port `3000`:

    $ http-server

Serving files from `./public` on port `8080`:

    $ http-server --port 8080 --path ./public

Serving files from `192.168.1.1` on port `8080`:

    $ http-server --host 192.168.1.1 --port 8080

Help:

    $ http-server --help
    Usage of http-server:
      -host="0.0.0.0": host to serve on
      -path="./": path to serve from
      -port=3000: port to serve on

## Options

- **`--port`** - which HTTP port to serve on. Defaults to `3000`.
- **`--host`** - which HTTP host to serve on. Defaults to `0.0.0.0`.
- **`--path`** - which path to serve to serve content from. Defaults to `./`.

## TODO

- Better logging (ideally, per-request w/ timestamps + requested file path)
- SSL support?
- Cache control?

## License

MIT. Plain and simple. Read `LICENSE.txt` for more details.
