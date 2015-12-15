package goapiutils_test

import (
	. "github.com/cfmobile/goapiutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("ApiResponse", func() {
	var response *ApiResponse
	var testData = []byte{0x00, 0x01, 0x02, 0x03}

	var commonTests = func() {
		It("should return a valid ApiResponse object", func() {
			Expect(response).NotTo(BeNil())
		})

		It("should contain the data that was passed in as the parameter", func() {
			Expect(response.Body).To(Equal(testData))
		})
	}

	Describe("NewApiResponse", func() {
		BeforeEach(func() {
			response = NewApiResponse(testData)
		})

		commonTests()

		It("should contain an http OK response code", func() {
			Expect(response.Code).To(Equal(http.StatusOK))
		})
	})

	Describe("NewApiResponseWithCode", func() {
		var testCode = 123

		BeforeEach(func() {
			response = NewApiResponseWithCode(testData, testCode)
		})

		commonTests()

		It("should contain the response code provided as argument", func() {
			Expect(response.Code).To(Equal(testCode))
		})
	})
})
