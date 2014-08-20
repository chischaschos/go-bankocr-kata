package bankocr_test

import (
  parser "github.com/chischaschos/bankocr/parser"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
  Context("When parsing individual account numbers", func() {
    It("should parse valid numbers", func() {
      number := "    _  _  _  _  _  _     _ \n" +
                "|_||_|| ||_||_   |  |  ||_ \n" +
                "  |  ||_||_||_|  |  |  | _|\n"
      number, status := parser.ParseAccountNumber(number)
      Expect(number).To(Equal("490867715"))
      Expect(status).To(BeEmpty())
    })

    It("should return errors on invalid numbers", func() {
      number := "    _  _     _  _  _  _  _ \n" +
                " _| _| _||_||_ |_   ||_||_|\n" +
                "  ||_  _|  | _||_|  ||_|  |\n"
      number, status := parser.ParseAccountNumber(number)
      Expect(number).To(Equal("?23456789"))
      Expect(status).To(Equal("ILL"))
    })
  })

  Context("When parsing an account numbers file", func() {
    It("should return a list of found account numbers", func() {
      Expect(parser.ParseAccountNumbersFile("account_numbers.txt")).To(Equal([]string{
        "123456789",
        "664371495 ERR",
        "?23456789 ILL",
      }))
    })
  })

  Context("Checksum calculator", func() {
    It("should check if an account number has a valid checksum", func() {
      Expect(parser.Checksum("490867715")).To(Equal(true))
      Expect(parser.Checksum("345882865")).To(Equal(true))
      Expect(parser.Checksum("123456789")).To(Equal(true))
    })

    It("should check if an account number has an invalid checksum", func() {
      Expect(parser.Checksum("664371495")).To(Equal(false))
    })
  })
})
