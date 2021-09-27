package protoc

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strings"
)


type config struct {
	Interfaces map[string]string `yaml:"interfaces"`
}

var cfg config

func init() {
	yfile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yfile, &cfg)
	if err != nil {
		panic(err)
	}
	//cfg.Interfaces = make(map[string]string)
	//cfg.Interfaces["cosmos.gov.Content"] = "github.com/cosmos/cosmos-sdk/x/gov/types"
	//cfg.Interfaces["cosmos.staking.Validator"] = "github.com/cosmos/cosmos-sdk/x/staking/types"
}

// cosmos proto option definitions
var (
	// field options
	acceptsInterfaceFieldOption = "93001"
	scalarFieldOption           = "93002"

	// message options
	implementsInterfaceMessageOption = "93001"
)

func unwrapOption(o string) (string, string) {
	split := strings.Split(o, ":")
	if len(split) != 2 {
		panic("invalid option format")
	}
	return cleanString(split[0]), cleanString(split[1])
}

// for some reasons option strings have quotes and spaces diluting them
// so we use this func to clean it up.
func cleanString(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\"")
	return s
}

func getIdentifier(n string) string {
	// should take something like "cosmos.gov.Content"
	// and simply return the identifier associated with it.
	// in the example above, this function would return "Content"
	identifier := strings.Split(n, ".")
	return identifier[len(identifier)-1]
}

// ideally this should look up whatever `o` is in a yaml config
// i.e. cosmos.gov.Content: github.com/cosmos/cosmos-sdk/x/gov/types
func getInterfaceImport(o string) string {
	return cfg.Interfaces[o]
}