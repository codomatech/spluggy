# `spluggy`: compile-time plugins for go

Organizing your code as plugins has many benefits, Go provides [excellent support for them](https://golang.org/pkg/plugin/).
The caveat (beside dependency hell) is that you lose the cool, single-binary advantage of Go.

`spluggy` gets you the best of both worlds: static plugins.
You define your plugins as sub-packages within the same package. As long as they all expose a function with the same name,
`spluggy` will discover them automatically and make the function accessible for external packages.


## Usage

A functioning app is in [example](example) which demonstrates how to use `spluggy` using `go generate`.
The commend line interface is:

```
Usage: spluggy <flags> <package directory>

With flags:

  -func string
    	The exposed function name
  -out string
    	Output file name (default "plugins.go")
  -pkg string
    	The base package
  -v	Enable verbose output

```

---
`spluggy` is a work of :heart: by [Codoma.tech](https://www.codoma.tech/).
