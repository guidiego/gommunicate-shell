package gshell

import (
  "fmt"
)

func Ask(ask string, defaultResponse string) string {
  response, err := baseShellRead(ask + " (default:" + defaultResponse + ")");

  if err != nil {
    return defaultResponse
  }

  return response
}

func AskTilBlankEnter(ask string, repeatAlert string) []string {
  var responseList []string;
  var err error;

  for err == nil {
    var response string

    response, err = baseShellRead(ask + " " + repeatAlert)
    responseList = append(responseList, response)
  }

  return responseList
}

func baseShellRead(ask string) (string, error) {
  var response string;

  fmt.Println(ask)
  _, err := fmt.Scanln(&response)

  return response, err
}
