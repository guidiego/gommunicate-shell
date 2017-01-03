package termput

import (
  "fmt"
  "strings"
)

func Input(placeholder string, defaultResponse string) string {
  defaultPlaceholder := ""

  if strings.TrimSpace(defaultResponse) != "" {
    defaultPlaceholder = " (default:" + defaultResponse + ")"
  }

  response, err := baseShellRead(placeholder + defaultPlaceholder);

  if err != nil {
    return defaultResponse
  }

  return response
}

func LoopInput(placeholder string, repeatAlert string) []string {
  var responseList []string;
  var err error;

  for err == nil {
    var response string

    response, err = baseShellRead(placeholder + " " + repeatAlert)

    if strings.TrimSpace(response) != "" {
      responseList = append(responseList, response)
    }
  }

  return responseList
}

func baseShellRead(placeholder string) (string, error) {
  var response string;

  fmt.Println(placeholder)
  _, err := fmt.Scanln(&response)

  return response, err
}
