# senml

[![Travis-CI](https://travis-ci.org/objenious/senml.svg?branch=master)](https://travis-ci.org/objenious/senml)  [![GoDoc](https://godoc.org/github.com/objenious/senml?status.svg)](http://godoc.org/github.com/objenious/senml)
[![GoReportCard](https://goreportcard.com/badge/github.com/objenious/senml)](https://goreportcard.com/report/github.com/objenious/senml)
[![Coverage Status](https://coveralls.io/repos/github/objenious/senml/badge.svg?branch=master)](https://coveralls.io/github/objenious/senml?branch=master)

`go get github.com/objenious/senml`

## Status: beta (no breaking changes expected)

This package implements the SenML format (Sensor Measurement Lists, formerly known as Sensor Markup Language), as defined in https://tools.ietf.org/html/draft-ietf-core-senml-16

This package is used in production on the Objenious LoRaWAN platform, and is maintained.

## Encoding/decoding

Encoding to/from JSON and XML is managed by the standard library.

```
s := senml.Pack{{Name:"foo", Value: senml.Float(32)}}
err := json.NewEncoder(w).Encode(s)
```

```
s := senml.Pack{}
err := json.NewDecoder(req.Body).Decode(&s)
```

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
