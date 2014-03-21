Go(lang) Error package
======================

A simple and more verbose error package for Go (which is somewhat backwards compatible), that allows better use of typed error handling.

This can be helpful if you return the same error over and over again either in the same or different functions, but have specific variables inside the error string.

IE.
```
err := errgo.New("Couldnt connect to computer with IP %v")
```

To mitigate this problem you could either use fmt.Sprintf, or get away with string additions, but that can lead to a whole basket of headaches.


### Creating Error types
Using this errors package you can just easily define a new error type with dynamic parameters.

```
errgo.Register(NO_CONNECT_ERR, "Couldnt connect to computer with IP %v")
err := errgo.New(NO_CONNECT_ERR, "192.156.22.10")
```
where `NO_CONNECT_ERR` is some int constant you define

### Handling errors
Error handling is managed the same way as the normal Go error package, but with the added ability to quickly identify the error type

```
	err := somefunction()
	if err != nil {
		panic(err)
	}

	// OR

	err := somefunction()
	switch {
	case err.IsType(NO_CONNECT_ERR):
		// handle
	case err.IsType(NO_ADDRESS_ERR):
		// handle

	// ... etc
```

This allows more verbose error handling, so you can handle the appropriate error

