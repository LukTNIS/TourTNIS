package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPGetUsers(t *testing.T) {
	t.Run("it should return httpCode 200", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/dloan-payment/v1/calculate-installment-amount", nil)
		if err != nil {
			t.Error(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(resp.UserHandler)
		handler.ServeHTTP(resp, req)
		if status := resp.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})
}
