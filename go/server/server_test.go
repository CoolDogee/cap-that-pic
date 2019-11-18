package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
		router = gin.Default()
		// data.Reload()
		server.SetupRoutes(router)
		router.Run()

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

	Describe("Version 1 API at /api/v1", func() {
		Describe("The /validateImageURL endpoint when link is an image", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/api/v1/validateImageURL?fileName=https://cms.hostelworld.com/hwblog/wp-content/uploads/sites/2/2017/08/girlgoneabroad.jpg")
			})

			It("Returns with Status 200", func() {
				Expect(response.Code).To(Equal(200))
			})
		})
		Describe("The /validateImageURL endpoint when link is valid but not an image", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/api/v1/validateImageURL?fileName=https://cms.hostelworld.com")
			})

			It("Returns with Status 400", func() {
				Expect(response.Code).To(Equal(http.StatusBadRequest))
			})
		})
		Describe("The /validateImageURL endpoint when link is invalid", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/api/v1/validateImageURL?fileName=htt")
			})

			It("Returns with Status 400", func() {
				Expect(response.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Describe("The /validateImageURL endpoint is valid but not an image", func() {
			BeforeEach(func() {
				response = performRequest(router, "GET", "/api/v1/validateImageURL?fileName=www.google.com")
			})

			It("Returns with Status 400", func() {
				Expect(response.Code).To(Equal(http.StatusBadRequest))
			})
		})

	})

	// Describe("GET the /getcaption?fileName=animals.jpg endpoint", func() {
	// 	BeforeEach(func() {
	// 		response = performRequest(router, "GET", "/api/v1/getcaption?fileName=https://cms.hostelworld.com/hwblog/wp-content/uploads/sites/2/2017/08/girlgoneabroad.jpg")
	// 	})

	// 	It("Returns with Status 200", func() {
	// 		Expect(response.Code).To(Equal(200))
	// 	})

	// 	It("Returns with caption", func() {
	// 		var actual, expect string
	// 		json.Unmarshal(response.Body.Bytes(), &actual)
	// 		expect = "Sunshine she's here, you can take a break\nI'm a hot air balloon that could go to space\nWith the air, like I don't care, baby, by the way"
	// 		Expect(actual).Should(Equal(expect))
	// 	})
	// })

	// 	Describe("GET the /getTagsfromImage endpoint", func() {
	// 		BeforeEach(func() {
	// 			response = performRequest(router, "GET", "/api/v1/getTagsfromImage")
	// 		})

	// 		It("Returns with Status 200", func() {
	// 			Expect(response.Code).To(Equal(200))
	// 		})

	// 		It("Returns with caption", func() {
	// 			var actual []models.Tag
	// 			json.Unmarshal(response.Body.Bytes(), &actual)
	// 			var expect = []models.Tag{
	// 				models.Tag{
	// 					Name:       "animal",
	// 					Confidence: 99.99402165412903,
	// 				},
	// 				models.Tag{
	// 					Name:       "aquatic mammal",
	// 					Confidence: 99.9920129776001,
	// 				},
	// 				models.Tag{
	// 					Name:       "mammal",
	// 					Confidence: 99.99026656150818,
	// 				},
	// 				models.Tag{
	// 					Name:       "seal",
	// 					Confidence: 99.93897676467896,
	// 				},
	// 				models.Tag{
	// 					Name:       "marine mammal",
	// 					Confidence: 98.87551069259644,
	// 				},
	// 				models.Tag{
	// 					Name:       "harbor seal",
	// 					Confidence: 98.05331230163574,
	// 				},
	// 				models.Tag{
	// 					Name:       "sea lion",
	// 					Confidence: 97.40312099456787,
	// 				},
	// 				models.Tag{
	// 					Name:       "ground",
	// 					Confidence: 97.18411564826965,
	// 				},
	// 				models.Tag{
	// 					Name:       "outdoor",
	// 					Confidence: 88.97097110748291,
	// 				},
	// 				models.Tag{
	// 					Name:       "earless seal",
	// 					Confidence: 85.63637733459473,
	// 				},
	// 				models.Tag{
	// 					Name:       "fur seal",
	// 					Confidence: 82.55710005760193,
	// 				},
	// 				models.Tag{
	// 					Name:       "california sea lion",
	// 					Confidence: 77.21861004829407,
	// 				},
	// 				models.Tag{
	// 					Name:       "baltic gray seal",
	// 					Confidence: 76.9413948059082,
	// 				},
	// 				models.Tag{
	// 					Name:       "standing",
	// 					Confidence: 75.22529363632202,
	// 				},
	// 				models.Tag{
	// 					Name:       "steller sea lion",
	// 					Confidence: 72.91018962860107,
	// 				},
	// 				models.Tag{
	// 					Name:       "bearded seal",
	// 					Confidence: 67.76590943336487,
	// 				},
	// 				models.Tag{
	// 					Name:       "otter",
	// 					Confidence: 55.91362714767456,
	// 				},
	// 				models.Tag{
	// 					Name:       "dirt",
	// 					Confidence: 20.025861263275146,
	// 				},
	// 			}
	// 			Expect(actual).Should(Equal(expect))
	// 		})
	// 	})

	// })
})
