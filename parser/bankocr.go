package parser

import (
  "bytes"
  "io/ioutil"
  "strings"
  "strconv"
)

var OCRToNumber = map[string]string{
  "   " +
  "  |" +
  "  |": "1",

  " _ " +
  " _|" +
  "|_ ": "2",

  " _ " +
  " _|" +
  " _|": "3",

  "   " +
  "|_|" +
  "  |": "4",

  " _ " +
  "|_ " +
  " _|": "5",

  " _ " +
  "|_ " +
  "|_|": "6",

  " _ " +
  "  |" +
  "  |": "7",

  " _ " +
  "|_|" +
  "|_|": "8",

  " _ " +
  "|_|" +
  "  |": "9",

  " _ " +
  "| |" +
  "|_|": "0",

}

func ParseAccountNumber(accountNumberLines string) (resultNumber, status string) {
  rows := strings.Split(accountNumberLines, "\n")[0:3]

  resultNumber = ""

  for i := 0; i < 27; i += 3 {
    currentNumber := ""

    for _, row := range rows {
      currentNumber += row[i:i + 3]
    }

    value, ok := OCRToNumber[currentNumber]

    if !ok {
      status = "ILL"
      resultNumber += "?"
    } else {
      resultNumber += value
    }

  }

  return resultNumber, status
}

func ParseAccountNumbersFile(filename string) (accountNumbers []string) {
  accountNumbersFile, readError := ioutil.ReadFile(filename)

  if readError != nil {
    panic(readError)
  }

  lines := bytes.Split(accountNumbersFile, []byte{'\n'})

  for i := 0; i < len(lines); i += 4 {
    if i + 4 < len(lines) {
      accountNumberOcr := string(bytes.Join(lines[i:i + 4], []byte{'\n'}))
      accountNumber, status := ParseAccountNumber(accountNumberOcr)

      if status == "" && !Checksum(accountNumber) {
        accountNumber += " ERR"
      } else if status != "" {
        accountNumber += " " + status
      }

      accountNumbers = append(accountNumbers, accountNumber)
    }
  }

  return accountNumbers
}

func Checksum(accountNumber string) bool {
  sum := 0

  for index := len(accountNumber); index > 0; index-- {
    val, convError := strconv.Atoi(string(accountNumber[index - 1]))

    if convError != nil {
      panic(convError)
    }

    sum += val * (len(accountNumber) - index + 1)
  }

  return sum % 11 == 0
}
