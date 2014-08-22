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
  " _|": "9",

  " _ " +
  "| |" +
  "|_|": "0",

}

func ParseAccountNumber(accountNumberLines string) (resultNumber, status string) {
  resultNumber = ""

  parseEachNumber(accountNumberLines, func(position int, number, ocrNumber string, ok bool) {
    if !ok {
      status = "ILL"
      resultNumber += "?"
    } else {
      resultNumber += number
    }
  })

  if status == "" && !Checksum(resultNumber) {
    status = "ERR"
  }

  if status != "" {
    possibilities := findPossibleAccountNumbers(accountNumberLines, resultNumber)

    if len(possibilities) > 1 {
      status = "AMB [" + string(strings.Join(possibilities, ", ")) + "]"

    } else if len(possibilities) == 1 {
      resultNumber = possibilities[0]
      status = ""

    } else {
      status = "ERR"
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

      if status != "" {
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

func findPossibleAccountNumbers(accountNumberLines, resultNumber string) []string {
  possibilities := []string{}

  parseEachNumber(accountNumberLines, func(position int, number, ocrNumber string, ok bool) {
    eachNumberCombination(ocrNumber, func(numberCombination byte) {
      tmpResultNumber := []byte(resultNumber)
      tmpResultNumber[position] = numberCombination

      if !bytes.Contains(tmpResultNumber, []byte("?")) && Checksum(string(tmpResultNumber)) {
        possibilities = append(possibilities, string(tmpResultNumber))
      }
    })
  })

  return possibilities
}

func eachNumberCombination(ocrNumber string, callback func(byte)) {
  for position := 0; position < len(ocrNumber); position++ {
    bytesNumber := []byte(ocrNumber)

    if bytesNumber[position] == '_' || bytesNumber[position] == '|' {
      bytesNumber[position] = ' '

    } else if position == 1 || position == 4 || position == 7 {
      bytesNumber[position] = '_'

    } else if position == 3 || position == 5 || position == 6 || position == 8 {
      bytesNumber[position] = '|'
    }

    value, ok := OCRToNumber[string(bytesNumber)]

    if ok {
      callback([]byte(value)[0])
    }
  }
}

func parseEachNumber(accountNumberLines string, callback func(int, string, string, bool)) {
  rows := strings.Split(accountNumberLines, "\n")[0:3]

  for i := 0; i < 27; i += 3 {
    currentNumber := ""

    for _, row := range rows {
      currentNumber += row[i:i + 3]
    }

    value, ok := OCRToNumber[currentNumber]
    index := i / 3
    callback(index, value, currentNumber, ok)
  }
}


