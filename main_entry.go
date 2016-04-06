package cpn 

import (
    "flag"
    "fmt"
    "os"
)

const APP_VERSION = "0.1"

/*
 *  A list of all parsible flags
 */
// -v <...> : the version label. The default value is 'v1'
var versionFlag string

// -e <...> : the name of the endpoint to use
var endpointFlag string

// -m <...> : the model (for PUT/POST requests)
var modelFlag string

// -ops <...> : the rest operands as a one string like 'op1/op2/op3/...' 
var operandsFlag string

// -help <...> : the help 
var helpFlag string

// -c <...> : the command to use
var cmdFlag string

/*
 *  The main entry point of the application
 */ 
func main() {
	
	fmt.Printf("Arguments 1: %+v\n", os.Args)
	
    flag.StringVar(&versionFlag, "v", "v1", "The version of the API endpoint")
    flag.StringVar(&endpointFlag, "e", "", "The name of the API endpoint")
    flag.StringVar(&modelFlag, "m", "", "The model of the request")
    flag.StringVar(&operandsFlag, "op", "", "The operands of the request")
    flag.StringVar(&helpFlag, "help", "", "The help")
    flag.StringVar(&cmdFlag, "c", "", "The command to use")
    
    flag.Parse() // Scan the arguments list

	// If found the help flag - writing the help and exit
	if (helpFlag != "") {
		fmt.Printf("Copernicus Gold CLI: version %s\n", APP_VERSION)
		fmt.Printf("Usage: cpn [command] -e <endpoint name> [flags]\n\n", APP_VERSION)
		fmt.Printf("Command:  auth\n")
		fmt.Printf("          create\n")
		fmt.Printf("          rm\n")
		fmt.Printf("Flags:    -v   : the label of the endpoint's version\n")
		fmt.Printf("          -ops : the rest parts of the endpoint as a single string\n")
		fmt.Printf("          -m   : the model of the request, i.e. JSON\n")
		os.Exit(0);
	}

	if (cmdFlag != "") {
		fmt.Printf("Command: %s\n", cmdFlag)	
	}
	if (versionFlag != "") {
		fmt.Printf("Version: %s\n", versionFlag)	
	}
	if (endpointFlag != "") {
		fmt.Printf("Endpoint: %s\n", endpointFlag)	
	}
	if (modelFlag != "") {
		fmt.Printf("Model: %s\n", modelFlag)	
	}
	if (operandsFlag != "") {
		fmt.Printf("Operands: %s\n", operandsFlag)	
	}
	fmt.Printf("other args: %+v\n", flag.Args())
}
