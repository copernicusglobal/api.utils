package cpn 

import (
    "testing"
    "os"
)

func TestParsing(t *testing.T) {

	os.Args = make([]string, 1)
	os.Args = append(os.Args,"-c", "create", "-v","v2", "-e","customers" ,"-m","'{\"id\" : \"22\"}'", "xxx") 

	main()
	
	if versionFlag != "v2" {
	    t.Error("Unable to parse the version")
	}
	if cmdFlag != "create" {
	    t.Error("Unable to parse the command")
	}
	if endpointFlag != "customers" {
	    t.Error("Unable to parse the API endpoint")
	}
	if modelFlag != "'{\"id\" : \"22\"}'" {
	    t.Error("Unable to parse the model")
	}
	if operandsFlag != "" {
	    t.Error("Unable to parse the operands")
	}
}