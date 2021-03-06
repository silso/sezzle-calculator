package gin_test

import (
	//. "github.com/sezzle-calculator/gin"

	"bytes"
	"encoding/json"
	"flag"
	"github.com/golang/glog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func init() {
	if testing.Verbose() {
		flag.Set("alsologtostderr", "true")
		flag.Set("v", "5")
	}
}

var _ = Describe("SignupHandler", func() {
	RegisterFailHandler(Fail)

	var (
		jsonErr error
		// newCustomerToken string
		// phoneNumber      string = "1234567898"
	)

	BeforeEach(func() {
		// setup database
		// migrate db
		// setup gin
		endpointHeaders = make(http.Header)
		endpointHeaders.Add("X-Real-IP", "74.37.200.161") //Setting a fake IP address for our login security tests.
		form = gin.H{}
	})

	JustBeforeEach(func() {
		glog.Error("Fire request")
		response = httptest.NewRecorder()

		if (endpointMethod == "POST") || (endpointMethod == "PUT") || endpointMethod == "PATCH" {
			jsonString, _ := json.Marshal(form)
			var err error
			request, err = http.NewRequest(endpointMethod, endpointURL, bytes.NewReader(jsonString))
			if err != nil {
				glog.Error(err)
			}
			request.Header = endpointHeaders
			request.Header.Add("Content-Type", "application/json")

		} else {
			request, _ = http.NewRequest(endpointMethod, endpointURL, nil)
			request.Header = endpointHeaders
		}

		s.ServeHTTP(response, request)
	})

	//g.GET("/debug/test",)
	Describe("GET /debug/test", func() {
		var responseKeyValue keyValueResp
		BeforeEach(func() {
			endpointMethod = "GET"
			endpointURL = "http://localhost:8000/debug/test"
		})

		Context("On submitting a bad number", func() {
			BeforeEach(func() {
				form = gin.H{}
			})

			JustBeforeEach(func() {
				jsonErr = DecodeTestJson(response, &responseKeyValue)
			})

			It("should return an error", func() {
				Ω(response.Code).Should(Equal(http.StatusOK))
				Ω(response.HeaderMap["Content-Type"][0]).Should(Equal("application/json; charset=utf-8"))
				Ω(jsonErr).ShouldNot(HaveOccurred())
				Ω(responseKeyValue).ShouldNot(BeEmpty())
			})
		})

		Context("On next context", func() {
			BeforeEach(func() {
				form = gin.H{}
				endpointMethod = "POST"
				endpointURL = "http://localhost:8000/v1/user"
			})

			JustBeforeEach(func() {
				jsonErr = DecodeTestJson(response, &responseKeyValue)
			})

			It("should ask for an OTP", func() {
				Ω(response.Code).Should(Equal(http.StatusOK))
				Ω(response.HeaderMap["Content-Type"][0]).Should(Equal("application/json; charset=utf-8"))
				Ω(jsonErr).ShouldNot(HaveOccurred())
				Ω(responseKeyValue).ShouldNot(BeEmpty())
				Ω(responseKeyValue).Should(HaveKeyWithValue("name", "Test Name"))
			})
		})
	})

})
