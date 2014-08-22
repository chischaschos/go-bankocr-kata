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

    It("should return errors on invalid numbers", func() {
      number := " _  _     _  _        _  _ \n" +
                "|_ |_ |_| _|  |  ||_||_||_ \n" +
                "|_||_|  | _|  |  |  |  | _|\n"
      number, status := parser.ParseAccountNumber(number)
      Expect(number).To(Equal("664371495"))
      Expect(status).To(Equal("ERR"))
    })

  })

  Context("When parsing an account numbers file", func() {
    It("should return a list of found account numbers", func() {
      Expect(parser.ParseAccountNumbersFile("account_numbers.txt")).To(Equal([]string{
        "711111111",
        "777777177",
        "200800000",
        "333393333",
        "888888888 AMB ['888886888', '888888880', '888888988']",
        "555555555 AMB ['555655555', '559555555']",
        "666666666 AMB ['666566666', '686666666']",
        "999999999 AMB ['899999999', '993999999', '999959999']",
        "490067715 AMB ['490067115', '490067719', '490867715']",
        "123456789",
        "000000051",
        "490867715",
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
