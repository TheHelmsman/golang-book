/*
	The `io` package consists of a few functions, but mostly interfaces used in other packages. 
	The two main interfaces are `Reader` and `Writer`. 
	`Readers` support reading via the `Read` method. 
	`Writers` support writing via the `Write` method. 
	Many functions in Go take `Readers` or `Writers` as arguments.


	For example the `io` package has a `Copy` function which copies data from a `Reader` to a `Writer`:

		func Copy(dst Writer, src Reader) (written int64, err error)

	To read or write to a `[]byte` or a `string` you can use the `Buffer` struct found in the `bytes` package:

		var buf bytes.Buffer
		buf.Write([]byte("test"))

	A `Buffer` doesn't have to be initialized and supports both the `Reader` and `Writer` interfaces. 
	You can convert it into a `[]byte` by calling `buf.Bytes()`

	If you only need to read from a string you can also use the `strings.NewReader` function 
	which is more efficient than using a buffer.
*/