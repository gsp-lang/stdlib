# Gimpy

[Gisp](https://github.com/jcla1/gisp) is a compiler that translates Gisp into
the Go AST. Gimpy is a simple overlay to facilitate easy access to the Go
stdlib from Gisp. Primarily, this comes down to performing type assertions
because all values in Gisp are core.Any (or interface{}).

## Installing

```bash
go get github.com/eatonphil/gimpy
```

## Example

Here is a simple "Hello World" server application written in Gisp, using Gimpy:

```lisp
(ns main
    "github.com/eatonphil/gimpy/fmt"
    "github.com/eatonphil/gimpy/net/http")

(def hello (fn [w r]
    (fmt/Fprintf w "Hello World")
    ()))

(def main (fn []
    (http/HandleFunc "/" hello)
    (http/ListenAndServe ":9090" nil)))
```

## Running

The Gisp compiler is really just a translator. It translates Gisp to Go. This
Go is outputted

## API

While the API is tremendously bare, the standard has been set to allow anyone
to easily add functions to Gimpy as they need access to the Go stdlib
counterpart.

The API is designed to mimick the Go stdlib exactly.

## License

MIT
