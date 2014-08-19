package bankocr_test

import (
  parser "github.com/chischaschos/bankocr/parser"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
  Context("When parsing individual account numbers", func() {
    It("should parse valid numbers", func() {
      number := "    _  _     _  _  _  _  _ \n" +
                "  | _| _||_||_ |_   ||_||_|\n" +
                "  ||_  _|  | _||_|  ||_|  |\n"
      Expect(parser.ParseAccountNumber(number)).To(Equal("123456789"))
    })

    It("should return errors on invalid numbers", func() {
      number := "    _  _     _  _  _  _  _ \n" +
                " _| _| _||_||_ |_   ||_||_|\n" +
                "  ||_  _|  | _||_|  ||_|  |\n"
      number, parseError := parser.ParseAccountNumber(number)
      Expect(number).To(Equal(""))
      Expect(parseError).To(HaveOccurred())
    })
  })

  Context("When parsing an account numbers file", func() {
    It("should return a list of found account numbers", func() {
      Expect(parser.ParseAccountNumbersFile("account_numbers.txt")).To(Equal([]string{
        "123456789",
        "123456789",
        "123456789",
      }))
    })
  })
})
