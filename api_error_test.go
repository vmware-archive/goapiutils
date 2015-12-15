package goapiutils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"encoding/json"
	. "github.com/cfmobile/goapiutils"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("ApiError", func() {
	kErrorCode := http.StatusNotFound
	kErrorMessage := "error message"

	getExpectedJsonError := func(message string) []byte {
		expectedJson := map[string]interface{}{
			"error": map[string]interface{}{
				"message": message,
			},
		}

		expectedBytes, _ := json.Marshal(expectedJson)
		return expectedBytes
	}

	var apiError *ApiError

	BeforeEach(func() {
		apiError = nil
	})

	Describe("NewApiError", func() {
		BeforeEach(func() {
			apiError = NewApiError(kErrorMessage, kErrorCode)
		})

		It("should create a new ApiError object", func() {
			Expect(apiError).NotTo(BeNil())
		})

		It("should set the error message that was passed in", func() {
			Expect(apiError.Message()).To(Equal(kErrorMessage))
		})

		It("should set the error code that was passed in", func() {
			Expect(apiError.Code()).To(Equal(kErrorCode))
		})
	})

	Describe("WriteError", func() {
		var fakeResponseWriter *httptest.ResponseRecorder

		BeforeEach(func() {
			fakeResponseWriter = httptest.NewRecorder()
			apiError = NewApiError(kErrorMessage, kErrorCode)

			WriteError(fakeResponseWriter, apiError)
		})

		It("should have written the expected error code in the header", func() {
			Expect(fakeResponseWriter.Code).To(Equal(kErrorCode))
		})

		It("should have written the expected error code in the response body", func() {
			expectedBytes := getExpectedJsonError(kErrorMessage)
			Expect(fakeResponseWriter.Body.Bytes()).To(Equal(expectedBytes))
		})
	})

	Describe("ApiError Methods", func() {
		BeforeEach(func() {
			apiError = NewApiError(kErrorMessage, kErrorCode)
		})

		Describe("MarshalJSON", func() {
			var byteArrayResponse []byte
			var errorResponse error

			BeforeEach(func() {
				byteArrayResponse, errorResponse = apiError.MarshalJSON()
			})

			It("should not have returned an error", func() {
				Expect(errorResponse).To(BeNil())
			})

			It("should have serialized the message in the expected format", func() {
				expectedBytes := getExpectedJsonError(kErrorMessage)
				Expect(byteArrayResponse).To(Equal(expectedBytes))
			})
		})
	})
})
