# cligen

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]
[![version-img]][version-url]

Generate CLI app skeleton

## Features

## Installation
```shell
go install github.com/neotoolkit/cligen/cmd/cligen@latest
```

## Usage
.cligen.yml
```yaml
name: "cli"
commands:
  - name: "command"
    alias: "c"
    description: "command description"
```
```shell
cligen
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/neotoolkit/cligen/workflows/build/badge.svg
[build-url]: https://github.com/neotoolkit/cligen/actions
[pkg-img]: https://pkg.go.dev/badge/neotoolkit/cligen
[pkg-url]: https://pkg.go.dev/github.com/neotoolkit/cligen
[reportcard-img]: https://goreportcard.com/badge/neotoolkit/cligen
[reportcard-url]: https://goreportcard.com/report/neotoolkit/cligen
[coverage-img]: https://codecov.io/gh/neotoolkit/cligen/branch/main/graph/badge.svg
[coverage-url]: https://codecov.io/gh/neotoolkit/cligen
[version-img]: https://img.shields.io/github/v/release/neotoolkit/cligen
[version-url]: https://github.com/neotoolkit/cligen/releases

## Sponsors
<p>
  <a href="https://evrone.com/?utm_source=github&utm_campaign=dotenv-linter">
    <img src="https://raw.githubusercontent.com/neotoolkit/.github/main/assets/sponsored_by_evrone.svg"
      alt="Sponsored by Evrone">
  </a>
</p>
