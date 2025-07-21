# Ligaturizer

[![GitHub Releases](https://img.shields.io/github/v/release/nhatthm/ligaturizer)](https://github.com/nhatthm/ligaturizer/releases/latest)
[![Build Status](https://github.com/nhatthm/ligaturizer/actions/workflows/release-edge.yaml/badge.svg)](https://github.com/nhatthm/ligaturizer/actions/workflows/release-edge.yaml)
[![Go Report Card](https://goreportcard.com/badge/go.nhat.io/ligaturizer)](https://goreportcard.com/report/go.nhat.io/ligaturizer)
[![GoDevDoc](https://img.shields.io/badge/dev-doc-00ADD8?logo=go)](https://pkg.go.dev/go.nhat.io/ligaturizer)
[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](https://www.paypal.com/donate/?hosted_button_id=PJZSGJN57TDJY)
<!--[![codecov](https://codecov.io/gh/nhatthm/ligaturizer/branch/master/graph/badge.svg?token=eTdAgDE2vR)](https://codecov.io/gh/nhatthm/ligaturizer)-->

Copy the ligatures (glyphs and rendering information) from [Fira Code v3.1](https://github.com/tonsky/FiraCode) into any other `TrueType` or `OpenType` font.

> [!Note]
> The ligatures are scale-corrected, but otherwise copied as is from Fira Code; it doesn't create new ligature graphics based on the font you're modifying.

![image](https://github.com/nhatthm/ligaturizer/assets/1154587/c635112d-947f-4f4a-95b4-7abf559c4f96)

## Table of Contents

- [Prerequisites](#prerequisites)
- [Install](#install)
- [Usage](#usage)
    - [Using docker](#using-docker)
- [Development](#development)
    - [Prerequisites](#prerequisites-1)
- [Credit](#credit)
- [Contribution](#contribution)
- [Related Projects](#related-projects)

## Prerequisites

- `Python = 3.11` with `fontforge` python bindings
    - For Debian/Ubuntu they are available in `python3-fontforge` package.
    - For OpenSUSE and NixOS, they are included in the `fontforge` package.
    - For macOS, they are available via brew (`brew install fontforge`).

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

## Install

You can download the [latest stable version](https://github.com/nhatthm/ligaturizer/releases/latest) or
the [nightly build](https://github.com/nhatthm/ligaturizer/releases/tag/edge) (`edge` version).

Once downloaded, the binary can be run from anywhere. Ideally, though, you should move it into your `$PATH` for easy use. `/usr/local/bin` is a popular location for this.

If you want to build from source, run `make build`, the binary will be available at `./out/ligaturizer`.

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

## Usage

```shell
$ ligaturizer /path/to/font-file
  --ligature-font-file /path/to/firacode-v3-otf-font-file \
```

For example:

```shell
$ ligaturizer /opt/fonts/DejaVuSansMono.ttf \
  --ligature-font-file /opt/fonts/FiraCode3/otf/FiraCode-Regular.otf \
```

or

```shell
$ ligaturizer /path/to/font-dir
  --ligature-font-file /path/to/firacode-v3-otf-font-file \
```

For example:

```shell
$ ligaturizer /opt/fonts/DejaVuSansMono/ \
  --ligature-font-file /opt/fonts/FiraCode3/otf/FiraCode-Regular.otf \
```

> [!Note]
> If you don't provide a ligature font file, you must specify the `--ligature-font-dir` option, point it to the directory that contains the FiraCode OTF font files. For example:
>
> ```shell
> $ ls -1 /opt/fonts/FiraCode3/otf/
> FiraCode-Bold.otf
> FiraCode-Light.otf
> FiraCode-Medium.otf
> FiraCode-Regular.otf
> FiraCode-Retina.otf
> FiraCode-SemiBold.otf
> ```
> The tool will pick a font file corresponding to the weight of the input font file.

The ligaturized font will be generated in the current working directory. If you want to specify the output directory, use `--output-dir` option.

For all available options, run:

```shell
$ ligaturizer -h
```

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

### Using docker

If you have problem with installing `fontforge` or `python3.11`, you can use the docker version instead. You don't have to specify the `--ligature-font-file` or `--ligature-font-dir` option, the docker image already contains the FiraCode font files.

```shell
$ docker run --rm \
  -v /path/to/local/font-dir:/opt/fonts/output \
  -w /opt/fonts/output \
  ghcr.io/nhatthm/ligaturizer:latest DejaVuSansMono.ttf
```

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

## Development

### Prerequisites

- `Go >= 1.23` with `cgo` enabled
- `golangci-lint >= 1.55.2` (optional)
- `Python = 3.11` with `fontforge` python bindings

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

## Credit

- The ligaturize script was originally written by [IlyaSkriblovsky](https://github.com/IlyaSkriblovsky) for adding ligatures to DejaVuSans Mono ([dv-code-font](https://github.com/IlyaSkriblovsky/dv-code-font)).
- [Navid Rojiani](https://github.com/rojiani) made a few changes to generalize the script so that it works for any font.
- [ToxicFrog](https://github.com/ToxicFrog) made a large number of contributions with the [Ligaturizer](https://github.com/ToxicFrog/Ligaturizer) project.
- [Nhat](https://github.com/nhatthm) has ported the project to Golang with some improvements.

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

## Contribution

Contributions always welcome! Please submit a Pull Request, or create an Issue if you have an idea for a feature/enhancement (or bug).

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)

## Related Projects

For more awesome programming fonts with ligatures, check out:

1. [FiraCode](https://github.com/tonsky/FiraCode)
2. [Hasklig](https://github.com/i-tu/Hasklig)

[<sub><sup>[table of contents]</sup></sub>](#table-of-contents)
