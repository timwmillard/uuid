# UUID generator

Command line tool to generate a random UUID.

## Installation

```sh
go install github.com/timwmillard/uuid@latest
```

## Usage

Generate a random UUID:
```sh
uuid
```

Generate a nil UUID:
```sh
uuid -nil
```

## Why not `uuidgen`

- `uuidgen` generates UUID's in capital letters (personal preference)
- `uuid` is shorter to type (I know I could create an alias but this is funner).
