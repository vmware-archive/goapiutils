package goapiutils_test

import (
	. "github.com/cfmobile/goapiutils"

	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("NotFoundHandler", func() {
	var subject http.Handler

	Context("when using constructor NewNotFoundHandler()", func() {
		BeforeEach(func() {
			subject = NewNotFoundHandler()
		})

		It("should create a valid Http Handler", func() {
			Expect(subject).NotTo(BeNil())
		})

		It("should have created an object of type NotFoundHandler", func() {
			Expect(subject).To(BeAssignableToTypeOf(&NotFoundHandler{}))
		})
	})

	Context("when ServeHTTP() is called", func() {
		var responseRecorder *httptest.ResponseRecorder
		var request *http.Request

		BeforeEach(func() {
			responseRecorder = httptest.NewRecorder()
			request, _ = http.NewRequest("GET", "testurl.com", nil)
			subject = NewNotFoundHandler()

			subject.ServeHTTP(responseRecorder, request)
		})

		It("should have returned an error code 404 (not found)", func() {
			Expect(responseRecorder.Code).To(Equal(http.StatusNotFound))
		})

		It(fmt.Sprintf("should have returned the error message %s", ErrInvalidRequest.Message()), func() {
			expectedJson := getExpectedJsonError(ErrInvalidRequest.Message())
			Expect(responseRecorder.Body.Bytes()).To(Equal(expectedJson))
		})
	})
})
