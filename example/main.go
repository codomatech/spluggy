package main

import "log"

import processors "github.com/example/app/processors"
import common "github.com/example/app/common"

func main() {
	var r common.Record
	plugins := processors.Plugins()
	log.Printf("processors.plugins: %+v\n", plugins)
	for _, fn := range plugins {
		fn("", &r)
	}
	log.Printf("Processed log:\n`\n%s\n`\n", r.Log)
}

// code generation

//go:generate spluggy -pkg github.com/example/app/processors ./processors
