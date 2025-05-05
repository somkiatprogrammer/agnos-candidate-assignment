package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatientSearchV1RoutePositive(t *testing.T) {

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patient/search/1", nil)
	router.ServeHTTP(w, req)

	jsonResult := `{"date_of_birth":"1980-12-11T00:00:00Z","email":"dummy1@dummy.com","first_name_en":"Dummy1 firstname","first_name_th":"ดัมมี่1 ชื่อ","gender":"M","last_name_en":"Dummy1 lastname","last_name_th":"ดัมมี่1 นามสกุล","middle_name_en":"Dummy1 middlename","middle_name_th":"ดัมมี่1 ชื่อกลาง","national_id":1,"passport_id":"DUMMY1","patient_hn":"dummy hospital","phone_number":"000000000"}`

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, jsonResult, w.Body.String())

}

func TestPatientSearchV1RouteNegative(t *testing.T) {

	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/patient/search/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestPatientSearchV2RoutePositive(t *testing.T) {

	router := SetupRouter()

	body := []byte(`{
		"first_name":"Dummy4 firstname"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patient/search", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDY0OTc2ODksImlkIjozfQ.aoXR564DdqJ10MYFGSDYgy4KlL0yeWBs-6Z32DDqM-E")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestPatientSearchV2RouteNegative(t *testing.T) {

	router := SetupRouter()

	body := []byte(`{
		"first_name":"Dummy4 firstname"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/patient/search", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

}

func TestStaffCreatNegative(t *testing.T) {

	router := SetupRouter()

	body := []byte(`{
		"username":"staff1",
		"password":"1234",
		"hospital":"dummy hospital"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/staff/create", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}

func TestStaffLoginPositive(t *testing.T) {

	router := SetupRouter()

	body := []byte(`{
		"username":"staff1",
		"password":"1234",
		"hospital":"dummy hospital"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/staff/login", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestStaffLoginNegative(t *testing.T) {

	router := SetupRouter()

	body := []byte(`{
		"username":"staff1",
		"password":"123456",
		"hospital":"dummy hospital"
	}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/staff/login", bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	jsonResult := `{"message":"Cannot Login!","status":"error"}`

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, jsonResult, w.Body.String())

}
