package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

//Conf ... generic structure
//taken from goofbot
type Conf struct {
	Owner  string
	Chan   string
	Server string
	Port   int
	Nick   string
	Pass   string
	User   string
	Name   string
	SSL    bool
}

func main() {
	//check for config file specified by command line flag -c
	jsonlocation := flag.String("c", "config.json", "Path to config file in JSON format")
	//spit out config file structure if requested
	jsonformat := flag.Bool("j", false, "Describes JSON config file fields")

	flag.Parse()

	//command line switch = true
	//display expected format of json file
	if *jsonformat == true {
		fmt.Println(`Here is the format for the JSON config file:
            {
                "owner": "YourNickHere",
                "chan": "#bots",
                "server": "irc.tilde.chat",
                "port": 6697,
                "nick": "goofbot",
                "pass": "",
                "user": "goofbot",
                "name": "Goofus McBotus",
                "ssl": true
            }`)
		os.Exit(0)
	}
	//read the config file into a byte array
	jsonconf, err := ioutil.ReadFile(*jsonlocation)
	if err != nil {
		panic(err)
	}

	// unmarshal the json byte array into struct conf
	var conf Conf
	err = json.Unmarshal(jsonconf, &conf)
	if err != nil {
		panic(err)
	}
}
