package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
	})

	Describe("Caption Generate Algorithm", func() {
		Describe("The getLyricsLines function", func() {
			var lines []string
			BeforeEach(func() {
				lyrics := []string{"line1 for lyric1\nline2 for lyric1", "line2 for lyric2\nline2 for lyric2"}
				lines = GetLyricsLines(lyrics)
			})

			It("Returns with Lines", func() {
				Expect(lines).Should(ConsistOf("line1 for lyric1", "line2 for lyric1", "line2 for lyric2", "line2 for lyric2"))
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
	})
})
