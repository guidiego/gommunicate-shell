# Gommunicate Shell
A simple shell wrapper to fmt.Scan to create a easy shell/user communicate

### Install
```
go get github.com/guidiego/gommunicate-shell
```

### API

#### .Ask(ask string, defaultValue string)
This method is print the **ask** and in case o blank response (enter), will put the **defaultValue**

*Example* 
```golang
package main

import (
  "fmt"
  "github.com/guidiego/gommunicate-shell"
)

func main() {
  answer1 := gshell.Ask("Say me a message!", "Hppy Xmas")
  fmt.Println(answer1)
}
```


#### .AskTilBlankEnter(ask string, repeatAlert string)
This method is print the **ask** + **repeatAlert** while enter not return a blank value.

*Example* 
```golang
package main

import (
  "fmt"
  "github.com/guidiego/gommunicate-shell"
)

func main() {
  answer1 := gshell.Ask("Say me a message!", "Hppy Xmas")
  fmt.Println(answer1)
  
  answer2 := gshell.AskTilBlankSpace("For what e-mails I need to send that?", "if want stop digit enter")
  fmt.Println(answer2)
  
  //Send a1 to all a2? how knows lol
}
```
