package server_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cooldogee/cap-that-pic/data"
	"github.com/cooldogee/cap-that-pic/server"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performRequestBuf(r http.Handler, method, path string, buf []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewReader(buf))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var _ = Describe("Server", func() {
	var (
		router   *gin.Engine
		response *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		router = server.CreateRouter()
		data.Reload()
	})

	Describe("Version 1 API at /api/v1", func() {
		Describe("The / endpoint", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/api/v1/")
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("Returns the String 'Hello World'", func() {
				Expect(response.Body.String()).To(Equal("Hello World"))
			})
		})

		Describe("GET the /getcaption?fileName=ImgUrl endpoint", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/api/v1/getcaption?fileName=https://wisconsinskydivingcenter.com/wp-content/uploads/2018/08/skydiving-lessons-1200x641.jpg")
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("Returns with caption", func() {
				var actual, expect string
				json.Unmarshal(response.Body.Bytes(), &actual)
				expect = "I done paid taxes, paid dues, paid bills my whole life\nSo I'ma throw money in the air, like I'm tryin' to pay the sky\nTold myself that man if I can't live like this, I'd rather die"
				Expect(actual).Should(Equal(expect))
			})
		})
	})
})
