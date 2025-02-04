package app

import (
	"fmt"
	"os"
	"strconv"
)

func getArgs() (rootPath string, topLen int) {
	rootPath = "."
	topLen = 10

	var err error
	args := os.Args

	for i, arg := range args {
		switch arg {
		case "-p", "--path":
			if i+1 >= len(args) {
				fmt.Println("Missing directory path for path argument!")
				printHelp()
				os.Exit(1)
			}
			rootPath = args[i+1]
			if _, err = os.ReadDir(rootPath); err != nil {
				fmt.Printf("Wrong\\unreadable directotry path '%s'\n", rootPath)
				printHelp()
				os.Exit(1)
			}
		case "-tl", "--toplen":
			if i+1 > len(args) {
				fmt.Println("Missing integer input for top length argument!")
				printHelp()
				os.Exit(1)
			}
			if topLen, err = strconv.Atoi(args[i+1]); err != nil || topLen < 0 {
				fmt.Printf("Cannot use '%s' an a top length!\n", args[i+1])
				printHelp()
				os.Exit(1)
			}
		case "-h", "--help":
			printHelp()
			os.Exit(0)
		}
	}
	return
}

func printHelp() {
	fmt.Println("Usage: fatdir -p [PATH_TO_DIR] -tl [INTEGER]",
		"\nBase call arguments: -p='.', -tl=10",
		"\n\nAvailable arguments:",
		"\n-h   --help      Print help information and arguments",
		"\n-p   --path      Provides path to the root directory",
		"\n-tl  --toplen    How many entries should the top contain")
}
