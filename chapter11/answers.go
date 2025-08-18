/*

	1. Why do we use packages?

		1. It reduces the chance of having overlapping names. 
		This keeps our function names short and succinct

		2. It organizes code so that its easier to find code you want to reuse.

		3. It speeds up the compiler by only requiring recompilation of smaller 
		chunks of a program. Although we use the package fmt, we don't have to 
		recompile it every time we change our program.

	2. What is the difference between an identifier that starts with a capital letter and 
	one which doesn’t? (Average vs average)

		In Go if something starts with a capital letter that means other packages (and programs) 
		are able to see it.

	3. What is a package alias? How do you make one?

		In Go, a package alias provides an alternative, usually shorter, name for an imported 
		package within a specific source file. This is primarily used to resolve naming conflicts 
		when importing multiple packages that happen to have the same default name, or to provide 
		a more convenient or descriptive name for a package.
		
		How to make a package alias:
			You create a package alias during the import statement by prepending the desired 
			alias name followed by a space before the package import path.

		package main

		import (
			f "fmt" // 'f' is an alias for the 'fmt' package
			"io/ioutil"
		)

		func main() {
			f.Println("Hello from aliased fmt!")
			// You can still use other imported packages normally
			_, err := ioutil.ReadFile("somefile.txt")
			if err != nil {
				// handle error
			}
		}

	4. We copied the average function from chapter 7 to our new package. 
	Create Min and Max func- tions which find the minimum and maximum values in a slice of float64s.

		see `myutils/myutils.go`

	5. How would you document the functions you created in #3?

		In Go, when documenting functions, the documentation is typically placed directly above 
		the function declaration. This applies whether the function is a regular function or a 
		function type that is aliased.
		
		Documenting a Function Type Alias:
			When you create a type alias for a function signature, you document the alias itself, 
			explaining what kind of function it represents.

		// MyFunctionType represents a function that takes an integer and returns a string.
		type MyFunctionType func(int) string

		// Example usage of MyFunctionType
		func processData(fn MyFunctionType, value int) string {
			return fn(value)
		}
*/