package termput

import (
    "bufio"
    "fmt"
    "os"
)

type InputConfig struct {
	DefaultValue string
	Label 	         string
	CustomLabel func()
}

func Input(opts InputConfig) string {
	scanner := bufio.NewScanner(os.Stdin)

	if opts.CustomLabel != nil {
		opts.CustomLabel()
	} else if opts.Label != "" {
		fmt.Println(opts.Label)
	} else {
		fmt.Print(">> ")
	}

	scanner.Scan()
	s := scanner.Text()

	if s != "" {
		return s
	}

	return opts.DefaultValue
}


func InputLoop(opts InputConfig) []string {
	stringCache := []string{}
	continueLoop := true

	for continueLoop {
		response := Input(opts)

		if response != opts.DefaultValue {
			stringCache = append(stringCache, response)
		} else {
			continueLoop = false

			if opts.DefaultValue != "" {
				stringCache = append(stringCache, response)
			}
		}
	}

	return stringCache
}

