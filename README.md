# Stdlib

The Gsp Stdlib is a simple overlay to facilitate easy access to the Go
stdlib from Gsp. Primarily, this comes down to performing type assertions
because all values in Gsp are core.Any (or interface{}).

## API

While the API is tremendously bare, the standard has been set to allow anyone
to easily add functions to Gimpy as they need access to the Go stdlib
counterpart. The Go stdlib will continue to be exposed over time.

The API is designed to mimick the Go stdlib exactly.

## License

MIT
