package main

import (
	"flag"
	"fmt"
	"plugin"
)

func loadPluginAndExecute(name string) {
	p, err := plugin.Open(name)
	if err != nil {
		panic(err)
	}

	varName := "V"

	if name == "logger_plugin.so" {
		varName = "W"
	}

	v, err := p.Lookup(varName)
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7

	f.(func())()
}

func main() {
	pluginName := flag.String("plugin", "plugin.so", "Plugin file to be loaded")

	flag.Parse()

	fmt.Println("Loading plugin: ", *pluginName)

	loadPluginAndExecute(*pluginName)
}
