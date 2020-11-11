package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func Test_POST_Key(t *testing.T) {
	key := "testK"
	router := httprouter.New()
	router.POST("/metric/"+key, HandlePostMetric)
	req, _ := http.NewRequest(http.MethodPost, "/metric/"+key, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}

func Test_POST_MANY(t *testing.T) {
	key := "testK"
	router := httprouter.New()
	router.POST("/metric/"+key, HandlePostMetric)
	router.GET("/metric/"+key+"/sum", HandleGetMetric)
	req, _ := http.NewRequest(http.MethodPost, "/metric/"+key, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	req, _ = http.NewRequest(http.MethodGet, "/metric/"+key+"/sum", nil)
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	resp := rr.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, rr.Code)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"Value":2}`, string(body))
}
func Test_POST_NotExist(t *testing.T) {
	key := "testKNotExist"
	router := httprouter.New()
	router.GET("/metric/"+key+"/sum", HandleGetMetric)
	req, _ := http.NewRequest(http.MethodGet, "/metric/"+key, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
