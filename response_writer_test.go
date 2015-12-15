package goapiutils_test

import (
	"encoding/json"
	. "github.com/cfmobile/goapiutils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Response Writer", func() {
	var recorder *httptest.ResponseRecorder
	var header http.Header
	var bodyValue interface{}

	BeforeEach(func() {
		bodyValue = "hello"
		recorder = httptest.NewRecorder()
		header = http.Header{}
		header.Add("some header", "header value")
	})

	checkBodyDataAsExpected := func() {
		jsonData, _ := json.Marshal(bodyValue)
		Expect(recorder.Body.Bytes()).To(Equal(jsonData))
	}

	expectDefaultHeaders := func() {
		Expect(recorder.HeaderMap.Get("Content-Type")).To(Equal("application/json"))
	}

	Context("when running WriteJSON()", func() {
		BeforeEach(func() {
			WriteJSON(recorder, bodyValue, http.StatusAccepted, header)
		})

		It("should write the specified status code to the response", func() {
			Expect(recorder.Code).To(Equal(http.StatusAccepted))
		})

		It("should write the specified headers to the response", func() {
			expectDefaultHeaders()
		})

		It("should write the appropriate json response", func() {
			Expect(recorder.HeaderMap.Get("some header")).To(Equal("header value"))
		})

		It("should write the expected json data to the response body", func() {
			checkBodyDataAsExpected()
		})
	})

	Context("when running WriteJSONWithHeader()", func() {
		BeforeEach(func() {
			WriteJSONWithHeader(recorder, bodyValue, header)
		})

		It("should automatically return status code 200 (OK)", func() {
			Expect(recorder.Code).To(Equal(http.StatusOK))
		})

		It("should write the specified headers in the response", func() {
			expectDefaultHeaders()
		})

		It("should write the expected json data to the response body", func() {
			checkBodyDataAsExpected()
		})
	})

	Context("when running WriteJSONWithCode()", func() {
		BeforeEach(func() {
			WriteJSONWithCode(recorder, bodyValue, http.StatusAccepted)
		})

		It("should write the specified status code to the reponse", func() {
			Expect(recorder.Code).To(Equal(http.StatusAccepted))
		})

		It("should automatically include no headers", func() {
			expectDefaultHeaders()
		})

		It("should write the expected json data to the response body", func() {
			checkBodyDataAsExpected()
		})
	})

	Context("when running WriteJSONData() with a json response as the data", func() {
		BeforeEach(func() {
			WriteJSONData(recorder, bodyValue)
		})

		It("should automatically return status code 200 (OK)", func() {
			Expect(recorder.Code).To(Equal(http.StatusOK))
		})

		It("should automatically include no headers", func() {
			expectDefaultHeaders()
		})

		It("should write the expected json data to the response body", func() {
			checkBodyDataAsExpected()
		})
	})

	Context("when running WriteJSONData() with an ApiError as the data", func() {
		BeforeEach(func() {
			error := NewApiError("Message", http.StatusForbidden)
			bodyValue = error
			WriteJSONData(recorder, error)
		})

		It("should write the status code specified in the ApiError to the response", func() {
			Expect(recorder.Code).To(Equal(http.StatusForbidden))
		})

		It("should automatically include no headers", func() {
			expectDefaultHeaders()
		})

		It("should write the expected json data to the response body", func() {
			checkBodyDataAsExpected()
		})
	})
})
