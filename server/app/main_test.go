package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// reference: https://flaviocopes.com/go-random/
func randomString(pool string, l int) string {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

func TestShowRequestPathSimpleTest(t *testing.T) {
	// characters '#' and '?' excluded for our simple URL
	pool := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz/0123456789-._~[]@!$&'()*+,;=:"
	expected := randomString(pool, 32)
	req, err := http.NewRequest("GET", "/"+expected, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRoot())
	handler.ServeHTTP(rr, req)

	if rr.Body.String() != expected {
		t.Error("test failed with path: " + expected)
	}
}

func TestHandlePrime(t *testing.T) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandlePrime())

	req, err := http.NewRequest("GET", "/prime/18", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)
	expected := "61"
	if result := rr.Body.String(); result != expected {
		t.Errorf("test failed: expected '%s' but got '%s'", expected, result)
	}

	rr = httptest.NewRecorder()
	req, err = http.NewRequest("GET", "/prime/abc", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("test failed: expected status as '%d' but got '%d'", http.StatusBadRequest, status)
	}
}

func TestNthPrime(t *testing.T) {
	for input, expected := range map[uint]uint{10: 29, 100: 541, 1000: 7919} {
		if result, _ := NthPrime(input); result != expected {
			t.Errorf("test failed: expected '%d' but got '%d'", expected, result)
		}
	}

	if _, err := NthPrime(0); err == nil {
		t.Errorf("test failed: expected non-nil error but got '%v' error", err)
	}
}

func BenchmarkNthPrime(b *testing.B) {
	values := []uint{
		10,
		100,
		1000,
		10000,
	}

	for _, val := range values {
		b.Run(
			fmt.Sprintf("nth=%d", val),
			func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					NthPrime(val)
				}
			},
		)
	}
}
