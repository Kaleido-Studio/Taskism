package main_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"taskism/controllers"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Backend", Ordered, func() {
	var router *gin.Engine
	BeforeAll(func() { router = controllers.GinEngine() })

	It("Smoke Test", func() {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)
		Expect(w.Code).To(Equal(http.StatusOK))
	})

	Describe("/api/{id}", func() {
		It("Should 400 if id is not uuid", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/1", nil)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
		It("Should 200 if id is uuid", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/97e910d1-a572-4624-bd32-84a7a0c76bde", nil)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
		It("Should 404 if method invalid", func() {
			cc, _ := json.Marshal(gin.H{"dsf": "sdf"})
			dd := bytes.NewBuffer(cc)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/97e910d1-a572-4624-bd32-84a7a0c76bde", dd)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})
})

func TestBackend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Backend Suite")
}
