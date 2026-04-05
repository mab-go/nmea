# mab-go/nmea

<p align="center">
  <a href="https://github.com/mab-go/nmea/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/mab-go/nmea/ci.yml?style=flat&logo=github&logoColor=white&labelColor=555555&label=build" alt="Build Status" /></a>
  <a href="https://codecov.io/gh/mab-go/nmea"><img src="https://img.shields.io/codecov/c/github/mab-go/nmea?style=flat&logo=codecov&logoColor=white&labelColor=555555" alt="Code Coverage" /></a>
  <a href="https://goreportcard.com/report/github.com/mab-go/nmea"><img src="https://goreportcard.com/badge/github.com/mab-go/nmea?style=flat" alt="Go Report Card" /></a>
  <a href="https://pkg.go.dev/github.com/mab-go/nmea"><img src="https://img.shields.io/badge/-reference-00ADD8?style=flat&logo=go&logoColor=white&labelColor=555555" alt="Go Reference" /></a>
  <a href="https://deepwiki.com/mab-go/nmea"><img src="https://img.shields.io/badge/DeepWiki-nmea-blue?style=flat&logoColor=white&labelColor=555555" alt="Ask DeepWiki" /></a>
  <a href="LICENSE"><img src="https://img.shields.io/github/license/mab-go/nmea?style=flat&labelColor=555555" alt="License: MIT" /></a>
</p>

A Go library for parsing NMEA 0183 sentences. Verify checksums, extract typed
fields, and decode sentence types like GPGGA and GPGLL into structured Go
values.

---

## Quick Start

```go
import (
    "fmt"
    "log"

    "github.com/mab-go/nmea/sentence/gpgga"
)

fix, err := gpgga.Parse("$GPGGA,174800.864,4002.741,N,07618.550,W,1,12,1.0,0.0,M,0.0,M,,*70")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Lat: %f %s, Lon: %f %s\n", fix.Latitude, fix.NorthSouth, fix.Longitude, fix.EastWest)
```

---

## Installation

```bash
go get github.com/mab-go/nmea
```

Requires Go 1.26 or later.

---

## Supported Sentences

| Package          | Sentence | Description                                                                       |
|------------------|----------|-----------------------------------------------------------------------------------|
| `sentence/gpgga` | GPGGA    | GPS fix data: time, lat/lon, fix quality, satellite count, HDOP, altitude         |
| `sentence/gpgll` | GPGLL    | Geographic position: lat/lon, fix time, data status, mode                         |
| `sentence/gpgsa` | GPGSA    | GPS DOP and active satellites: selection mode, fix mode, PRN list, PDOP/HDOP/VDOP |

The `sentence` package provides lower-level building blocks usable with any
NMEA 0183 sentence:

- **`VerifyChecksum`** — validates the `*XX` checksum on a raw sentence string
- **`SegmentParser`** — splits a sentence into typed fields (`AsFloat64`,
  `AsInt16`, `AsString`, and more)

---

## Development

First-time setup (installs golangci-lint, goimports, gocyclo into `./bin`):

```bash
make setup
```

Then:

```bash
make test        # Run all tests (with -race by default)
make test:cover  # Tests + HTML coverage report
make lint        # golangci-lint
make fmt         # goimports
```

Run `make help` for the full list of targets.

---

## License

MIT. See [LICENSE](LICENSE).
