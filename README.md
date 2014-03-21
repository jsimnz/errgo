Go(lang) Error package
======================

A simple and more verbose error package for Go (which is somewhat backwards compatible), that allows better use of typed error handling.

This can be helpful if you return the same error over and over again either in the same or different functions, but have specific variables inside the error string.

IE.
```
err := errors.New("Couldnt connect to computer with IP %v")
```

To mitigate this problem you could either use fmt.Sprintf, or get away with string additions, but that can lead to a whole basket of headaches.


### Creating Error types
Using this errors package you can just easily define a new error type with dynamic parameters.

```
errgo.Register(NO_CONNECT_ERR, "Couldnt connect to computer with IP %v")
err := errgo.New(NO_CONNECT_ERR, "192.156.22.10")
```
where `NO_CONNECT_ERR` is some int you define, as a constant or however else you choose.

### Handling errors
Error handling is managed the same way as the normal Go error package, but with the added ability to quickly identify the error type

```
	err := somefunction()
	if err != nil {
		panic(err)
	}

	// OR

	err := somefunction()
	switch err.Type {
	case NO_CONNECT_ERR:
		// handle
	case NO_ADDRESS_ERR:
		// handle
	default:
		// generic error
		// can also use the NO_TYPE
	}

	// ... etc
```

This allows more verbose error handling, so you can handle the appropriate error

### Backwards Compatibility

I tried to make this package 100% backwards compatible with the current errors package, however for ease of use that wasnt completly possible. The main Err type in this package implements the Go error interface, so it is compatible in that sense. Where the compatibility stops (sort of). The public function `New()` defined acts very similar to the `errors.New()`.

```
	errors.New("Some error") // standard go error package
	errgo.New("Some error") // equivalent 
```

However the function signatures differ to allow the dynamic variables and types this package is based behind

```
	func New(string) error
	func New(interface{}, ...interface{}) Err
```

This package can still act as a drop-in replacement in most situations