package protoc

import "strings"

func extractOption(o string) (string, string) {
	split := strings.Split(o, ":")
	if len(split) != 2 {
		panic("invalid option format")
	}
	return split[0], strings.TrimSpace(split[1])
}

func getIdentifier(n string) string {
	// should take something like "cosmos.gov.Content"
	// and simply return the identifier associated with it.
	// in this case we want "Content"
	return "Content"
}

func lookUpOption(n string) string {
	// we should have some sort of yaml file or something that
	// will map us from the field option number (i.e. 90021 or whatever) to
	// "implements_interface" or "accepts_interface"
	return "implements_interface"
}

// ideally this should look up whatever `o` is in a yaml config
// i.e. cosmos.gov.Content: github.com/cosmos/cosmos-sdk/x/gov/types
func getInterfaceImport(o string) string {
	return "github.com/cosmos/cosmos-sdk/x/gov/types"
}