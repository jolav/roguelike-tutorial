/* */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func prettyPrintStruct(s interface{}) {
	result, _ := json.MarshalIndent(s, "", "    ") //"\t")
	fmt.Print(string(result), "\n")
}

func loadConfigJSON(c *Configuration) {
	err := json.Unmarshal(getGlobalConfigJSON(), &c)
	if err != nil {
		log.Fatal("Error parsing JSON config => \n", err)
	}
}
func checkFlags() {
	versionFlag := flag.Bool("v", false, "Show current version and exit")
	flag.Parse()
	switch {
	case *versionFlag:
		fmt.Printf("Version:\t: %s\n", version)
		fmt.Printf("Date   :\t: %s\n", releaseDate)
		os.Exit(0)
	}
}
