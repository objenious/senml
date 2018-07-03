# senml

[![Travis-CI](https://travis-ci.org/objenious/senml.svg)](https://travis-ci.org/objenious/senml)  [![GoDoc](https://godoc.org/github.com/objenious/senml?status.svg)](http://godoc.org/github.com/objenious/senml)
[![GoReportCard](https://goreportcard.com/badge/github.com/objenious/senml)](https://goreportcard.com/report/github.com/objenious/senml)
[![Coverage Status](https://coveralls.io/repos/github/objenious/senml/badge.svg)](https://coveralls.io/github/objenious/senml)

`go get github.com/objenious/senml`

## Status: alpha - breaking changes might happen

This package implements the SenML format (Sensor Measurement Lists, formerly known as Sensor Markup Language), as defined in https://tools.ietf.org/html/draft-ietf-core-senml-16

This package is used in production on the Objenious LoRaWAN platform, and is maintained.

## TODO

* CBOR Representation
* EXI Representation
* Fragment Identification

## Contribution guidelines

Contributions are welcome, as long as :
* unit tests & comments are included,
* no external package is added to the top-level package (but allowed in sub-packages).

## Licence

MIT - See LICENSE
