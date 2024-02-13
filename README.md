# ToBahtText

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org) [![Go Reference](https://pkg.go.dev/badge/github.com/buildingwatsize/toBahtText@v0.1.0.svg)](https://pkg.go.dev/github.com/buildingwatsize/toBahtText@v0.1.0) [![GitHub issues](https://img.shields.io/github/issues/buildingwatsize/toBahtText)](https://github.com/buildingwatsize/toBahtText/issues) [![GitHub forks](https://img.shields.io/github/forks/buildingwatsize/toBahtText)](https://github.com/buildingwatsize/toBahtText/network) [![GitHub stars](https://img.shields.io/github/stars/buildingwatsize/toBahtText)](https://github.com/buildingwatsize/toBahtText/stargazers)

ToBahtText เป็น Golang library สำหรับแปลงตัวเลขเป็นคำอ่านภาษาไทย | ToBahtText is a Golang library for converting a number to the words.

<img src="./images/toBahtText.png" height="240" alt="logo" />

## Table of Contents

- [ToBahtText](#tobahttext)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Examples](#examples)
  - [Dependencies](#dependencies)
  - [Want some more example?](#want-some-more-example)

## Installation

```bash
  go get -u github.com/buildingwatsize/toBahtText
```

## Examples

```go
  words := toBahtText.From("2021")
  fmt.Println(words) // "สองพันยี่สิบเอ็ด"
```

## Dependencies

- Only [https://github.com/stretchr/testify](https://github.com/stretchr/testify)

## Want some more example?

Go to [toBahtText_test.go](toBahtText_test.go)

> Made with ❤️ Powered by Watsize