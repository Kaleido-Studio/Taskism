package main_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"taskism/controllers"
	"taskism/handlers"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var router *gin.Engine

var _ = BeforeSuite(func() {
	router = controllers.GinEngine()
})

var _ = Describe("Backend", func() {
	Describe("Get user", func() {
		It("Should 400 if id is not uuid", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/user/info/1", nil)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
		It("Should 200 if id is uuid", func() {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/user/info/97e910d1-a572-4624-bd32-84a7a0c76bde", nil)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
		It("Should 404 if method invalid", func() {
			cc, _ := json.Marshal(gin.H{"dsf": "sdf"})
			dd := bytes.NewBuffer(cc)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/user/info/97e910d1-a572-4624-bd32-84a7a0c76bde", dd)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusNotFound))
		})
	})

	Describe("Register user", Ordered, func() {
		It("Should able to register", func() {
			body, _ := json.Marshal(handlers.UserLogRegReqBody{Name: fmt.Sprint(GinkgoRandomSeed()), Password: "test"})
			buffer := bytes.NewBuffer(body)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/user/register", buffer)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
			resBody, _ := io.ReadAll(w.Body)
			var resBodyJson handlers.RegisterResBody
			err := json.Unmarshal(resBody, &resBodyJson)
			Expect(err).To(BeNil())
			Expect(resBodyJson.Username).To(Equal(fmt.Sprint(GinkgoRandomSeed())))
			Expect(resBodyJson.Token).ToNot(Equal(nil))
		})
		It("Shouldn't able to register if conflict[409]", func() {
			body, _ := json.Marshal(gin.H{"name": fmt.Sprint(GinkgoRandomSeed()), "password": "test"})
			buffer := bytes.NewBuffer(body)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/user/register", buffer)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusConflict))
		})
		It("Shouldn't able to register if request is invalid[400]", func() {
			body, _ := json.Marshal(gin.H{"name": GinkgoRandomSeed(), "password": "test"})
			buffer := bytes.NewBuffer(body)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/user/register", buffer)
			router.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusBadRequest))
		})
	})

	Describe("Login user", func() {
		It("Should able to login", func() {
			reqBody, _ := json.Marshal(gin.H{"name": fmt.Sprint(GinkgoRandomSeed()), "password": "test"})
			buffer := bytes.NewBuffer(reqBody)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/api/user/login", buffer)
			router.ServeHTTP(w, req)
			resBody, _ := io.ReadAll(w.Body)
			var resBodyJson handlers.LoginResBodyJson
			err := json.Unmarshal(resBody, &resBodyJson)
			Expect(err).To(BeNil())
			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(resBodyJson.Token).ToNot(BeNil())
		})
	})
})

func TestBackend(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Backend Suite")
}
