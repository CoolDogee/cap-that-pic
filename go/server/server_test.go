package server_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cooldogee/cap-that-pic/data"
	"github.com/cooldogee/cap-that-pic/server"
	. "github.com/cooldogee/cap-that-pic/server"
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
		router = CreateRouter()
		data.Reload()
	})

	Describe("Caption Generate Algorithm", func() {
		Describe("The GetLyricsLines function", func() {
			var lines []string
			BeforeEach(func() {
				lines = GetLyricsLines(&data.Song(-1, 0).List)
			})

			It("Returns with Lines", func() {
				Expect(lines).Should(ContainElement("What what, what, what"))
			})
		})

		Describe("The GenerateCaption function", func() {
			var caption string
			BeforeEach(func() {
				songs := data.Song(-1, 0).List
				tags := data.Tag(-1, 0).List
				log.Println("Len = ", len(tags))
				log.Println(tags[0].Name)

				caption = GenerateCaption(&songs, &tags)
			})

			It("Returns with caption", func() {
				Expect(caption).Should(Equal("I'm a hot air balloon that could go to space"))
			})
		})

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

		Describe("POST the /getcaption endpoint", func() {
			BeforeEach(func() {
				var img = server.Image{URL: "aaaa"}
				request, _ := json.Marshal(img)
				response = performRequestBuf(router, "POST", "/api/v1/getcaption", request)
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})

			It("Returns with caption", func() {
				var actual, expect server.Caption
				json.Unmarshal(response.Body.Bytes(), &actual)
				expect.Content = "I'm a hot air balloon that could go to space"
				Expect(actual).Should(Equal(expect))
			})
		})
	})
})
