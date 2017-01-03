# Termput
A simple shell wrapper to fmt.Scan to create a easy shell/user inputs

### Install
```
go get github.com/purrgil/termput
```

### API

#### .Input(placeholder string, defaultValue string)
This method is print the **placeholder** and in case o blank response (enter), will put the **defaultValue**

*Example* 
```golang
package main

import (
  "fmt"
  "github.com/purrgil/termput"
)

func main() {
  answer1 := termput.Input("Say me a message!", "Hppy Xmas")
  fmt.Println(answer1)
}
```


#### .LoopInput(placeholder string, repeatAlert string)
This method is print the **placeholder** + **repeatAlert** while enter not return a blank value.

*Example* 
```golang
package main

import (
  "fmt"
  "github.com/purrgil/termput"
)

func main() {
  answer1 := termput.Input("Say me a message!", "Hppy Xmas")
  fmt.Println(answer1)
  
  answer2 := termput.LoopInput("For what e-mails I need to send that?", "if want stop digit enter")
  fmt.Println(answer2)
  
  //Send a1 to all a2? how knows lol
}
```
