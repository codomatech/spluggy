# Example App with `spluggy`

This is an example app which uses `spluggy` code generation to organize code as
plugins.

## Build and Run

Install `spluggy` and make sure `$GOPATH/bin` is in your `$PATH`

```
go get github.com/codomatech/spluggy
go install github.com/codomatech/spluggy
```

Build:

```
go generate
go build
```

Run:
```
$ ./app
2020/10/09 15:33:46 processors.plugins: map[nested/proc2:0x4a62c0 proc1:0x4a63e0]
2020/10/09 15:33:46 Processed log:
`
processed by proc2.Process()
processed by proc1.Process()

`
```

## What Happened?

The directory `processors` contains plugins to process some record, a plugin is
a package which exposes a function with the following signature:
```
func Process(string, *common.Record)
```

In `main.go` we added a line to automatically discover plugins and write code to
expose them for later consumption:
```
//go:generate spluggy -pkg github.com/example/app/processors ./processors
```

After executing this line (i.e. `go generate`), the package `processors` has a
function `Plugins` which is a map from subpackage name to a function pointer.
