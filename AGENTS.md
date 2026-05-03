# AGENTS.md — mab-go/nmea

This file is the authoritative briefing for any AI agent working on this project. Read it in full before writing any code or making any changes.

---

## Project Summary

`mab-go/nmea` is a Go library for parsing NMEA 0183 sentences from GPS and navigation data streams. The module path is `github.com/mab-go/nmea` and the project lives at `/home/matt/Projects/mab-go/nmea`.

Each NMEA sentence type is implemented as its own sub-package under `sentence/`. Adding a new sentence type is the primary extension workflow — follow the patterns in `sentence/gpgga/` exactly.

---

## Project Structure

```
nmea.go                    — package doc only; no logic
sentence/
  sentence.go              — NMEASentence interface
  checksum.go              — VerifyChecksum (called by SegmentParser.Parse)
  checksum_test.go
  errors.go                — ParsingError type
  errors_test.go
  parser.go                — SegmentParser: Parse, As*, Require* methods
  parser_test.go
  testhelp/
    testhelp.go            — ReadTestData, EnsureFloat32/64, OptInt8/16, etc.
                             (used only by sentence package tests, not sub-packages)
  _testdata/               — YAML fixtures for sentence package tests only
    good-data.yaml
    bad-invalid-checksums.yaml
    bad-malformed.yaml
  gpgga/                   — GPS fix data (GPGGA)
    gpgga.go               — GPGGA struct, GetSentenceType(), Parse()
    parser.go              — SegmentParser embedding sentence.SegmentParser
    enum.go                — NorthSouth, EastWest, FixQuality types + //go:generate
    enum_gen.go            — generated; do not edit by hand
    gpgga_test.go          — inline Go test data (goodTestData, badTestData maps)
  gpgll/                   — Geographic position (GPGLL)
    gpgll.go / parser.go / enum.go / enum_gen.go / gpgll_test.go
  gpgsa/                   — GPS DOP and active satellites (GPGSA)
    gpgsa.go / parser.go / enum.go / enum_gen.go / gpgsa_test.go
tools/
  test-cover.sh            — coverage script wrapper
.agents/
  skills/                  — agent skills (symlinked from .claude/skills)
.claude/
  skills/                  — symlink → ../.agents/skills/
Makefile                   — build, test, lint, fmt, generate targets
.golangci.yaml             — golangci-lint v2 config
CLAUDE.md                  — symlink → AGENTS.md
```

---

## Adding a New Sentence Type

Create `sentence/<TYPE>/` (lowercase, e.g. `gprmc`) with these five files:

### `<TYPE>.go`

```go
package <TYPE>

import "github.com/mab-go/nmea/sentence"

type <TYPE> struct {
    // Each field comment must note the sentence element index, e.g.:
    // FixTime is ... It is element [1] of a <TYPE> sentence.
    FieldName FieldType
    // ...
}

func (s <TYPE>) GetSentenceType() string { return "<TYPE>" }

// Ensure <TYPE> implements NMEASentence
var _ sentence.NMEASentence = <TYPE>{}

func Parse(s string) (*<TYPE>, error) {
    segments := &SegmentParser{}
    if err := segments.Parse(s); err != nil {
        return nil, err
    }

    _ = segments.RequireString(0, "<TYPE>")
    result := &<TYPE>{
        FieldName: segments.AsFieldType(1),
        // ...
    }

    if err := segments.Err(); err != nil {
        return nil, err
    }

    return result, nil
}
```

### `enum.go`

```go
package <TYPE>

type MyEnum int

const (
    ValueA MyEnum = iota + 1 // A
    ValueB                    // B
)

//go:generate go run github.com/dmarkham/enumer@v1.6.3 -type=MyEnum -text -linecomment -transform=first-upper -output=enum_gen.go
```

Run `make generate` after creating or modifying `enum.go` to regenerate `enum_gen.go`. Commit `enum_gen.go`.

### `parser.go`

Embed `sentence.SegmentParser` and add one method per enum field:

```go
package <TYPE>

import (
    "fmt"
    "github.com/mab-go/nmea/sentence"
)

type SegmentParser struct {
    sentence.SegmentParser
    err error
}

func (p *SegmentParser) Err() error {
    if err := p.SegmentParser.Err(); err != nil {
        return err
    }
    return p.err
}

func (p *SegmentParser) AsMyEnum(i int8) MyEnum {
    val, err := MyEnumString(p.AsString(i))
    if err != nil {
        p.err = &sentence.ParsingError{
            Segment: i,
            Message: fmt.Sprintf("must be parsable as a MyEnum but was %q", p.AsString(i)),
        }
        return MyEnum(0)
    }
    return val
}
```

### `<TYPE>_test.go`

Use inline Go test data (maps, not YAML) — see `gpgga/gpgga_test.go` for the exact pattern:

- `goodTestData` — map of title → `testVec{input, expected, ""}` covering real device samples
- `badTestData` — map of title → `testVec{input, zero-value, errMsg}` with one bad field per case
- `TestParse_goodData`, `TestParse_badSegments`, `TestParse_invalidChecksum`, `TestGetSentenceType`
- `ExampleParse` with `// Output:` comment

The `testhelp` package is **not** used by sentence sub-package tests — it is only used by the `sentence` package's own tests.

---

## SegmentParser Error Pattern

`sentence.SegmentParser` accumulates the first error; all subsequent `As*`/`Require*` calls no-op. Each sentence sub-package's `SegmentParser` embeds it and adds its own `err` field for enum parsing errors. The combined `Err()` method checks both. Always call `segments.Err()` at the end of `Parse()` before returning the struct.

---

## Enum Pattern

- All enums use `iota + 1` — zero is never a valid enum value.
- Line comments on each constant (e.g., `// N`) are the string representations used by `enumer`.
- Generated file is always named `enum_gen.go` and is committed, not gitignored.
- `enumer` flags: `-text -linecomment -transform=first-upper`; list all types in a single `-type=A,B,C` flag.

---

## Verification

Treat any change as incomplete until **`make test lint`** passes. Run `make generate` whenever `enum.go` is created or modified.

| Command | Purpose |
|---|---|
| `make test` | `go test -race ./...` |
| `make lint` | golangci-lint (run `make setup` first if not installed) |
| `make generate` | regenerate `enum_gen.go` files |
| `make fmt` | goimports formatting |
| `make setup` | install golangci-lint, goimports, gocyclo into `./bin` |

---

## Documentation Maintenance

After any structural change, update the corresponding documentation:

| Change made | What to review/update |
|---|---|
| New sentence sub-package added | Project Structure tree in this file |
| New field added to a sentence struct | Field comments (include element index) |
| New enum type added or renamed | enum.go line comments, re-run `make generate` |
| New `As*` method added to `sentence.SegmentParser` | parser.go doc comment |
| Interface `NMEASentence` changed | All sentence structs (each has a compile-time check) |
| New convention established | Conventions sections in this file |

Only update docs when something has genuinely changed — no cosmetic edits.
