package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/progrium/go-basher"
	"gopkg.in/yaml.v2"
        "encoding/json"
)

var Version string

func YamlKeys(args []string) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var m interface{}
	err = yaml.Unmarshal(bytes, &m)
	if err != nil {
		log.Fatal(err)
	}
	for _, arg := range args {
		if m == nil {
			break
		}
		m = m.(map[interface{}]interface{})[arg]
	}
	n, ok := m.(map[interface{}]interface{})
	if ok {
		for key := range n {
			fmt.Printf("%s\n", key)
		}
	}
}

func YamlGet(args []string) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var m interface{}
	err = yaml.Unmarshal(bytes, &m)
	if err != nil {
		log.Fatal(err)
	}
	for _, arg := range args {
		if m == nil {
			break
		}
		m = m.(map[interface{}]interface{})[arg]
	}
	switch val := m.(type) {
	case string:
		fmt.Println(val)
	case map[interface{}]interface{}:
		for key := range val {
			fmt.Printf("%s=%s\n", key, val[key])
		}
	case []interface{}:
		for _, v := range val {
			fmt.Printf("%s\n", v)
		}
	}
}

func AppJsonGet(args []string) {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var m interface{}
        err = json.Unmarshall(bytes, &m)
	if err != nil {
		log.Fatal(err)
	}
	for _, arg := range args {
		if m == nil {
			break
		}
		m = m.(map[interface{}]interface{})[arg]
	}
        switch val := m.(type) {
	case string:
		fmt.Println(val)
	case map[interface{}]interface{}:
		for key := range val {
			fmt.Printf("%s=%s\n", key, val[key])
		}
	case []interface{}:
		for _, v := range val {
			fmt.Printf("%s\n", v)
		}
	}
}

func AssetCat(args []string) {
	for _, asset := range args {
		data, err := Asset(asset)
		if err != nil {
			os.Exit(2)
		}
		os.Stdout.Write(data)
	}
}

func main() {
	os.Setenv("HEROKUISH_VERSION", Version)
	basher.Application(map[string]func([]string){
		"yaml-keys": YamlKeys,
		"yaml-get":  YamlGet,
                "app-json":  AppJsonGet,
		"asset-cat": AssetCat,
	}, []string{
		"include/herokuish.bash",
		"include/fn.bash",
		"include/cmd.bash",
		"include/buildpack.bash",
		"include/procfile.bash",
		"include/slug.bash",
	}, Asset, true)
}
