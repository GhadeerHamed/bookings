package handlers

import (
	"context"
	"github.com/ghadeerhamed/bookings/internal/models"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	expectedStatusCode int
}{
	{"Home page", "/", "GET", http.StatusOK},
	{"About page", "/about", "GET", http.StatusOK},
	{"General room", "/rooms/generals-quarters", "GET", http.StatusOK},
	{"Major suite", "/rooms/major-suit", "GET", http.StatusOK},
	{"Search availability", "/search-availability", "GET", http.StatusOK},
	{"Contact", "/contact", "GET", http.StatusOK},

	//{"Make reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	//
	////Post routes tests
	//{"POST search availability", "/search-availability", "POST", []postData{
	//	{key: "start", value: "2020-01-01"},
	//	{key: "end", value: "2020-01-03"},
	//}, http.StatusOK},
	//{"POST search availability JSON", "/search-availability-json", "POST", []postData{
	//	{key: "start", value: "2020-01-01"},
	//	{key: "end", value: "2020-01-03"},
	//}, http.StatusOK},
	//{"POST make reservation", "/make-reservation", "POST", []postData{
	//	{key: "first_name", value: "Joe"},
	//	{key: "last_name", value: "Smith"},
	//	{key: "email", value: "email@test.com"},
	//	{key: "phone", value: "555555555"},
	//}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()

	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range theTests {
		if test.method == "GET" {
			res, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if res.StatusCode != test.expectedStatusCode {
				t.Errorf("Error: For %s, expected %d, but got %d", test.name, test.expectedStatusCode, res.StatusCode)
			} else {
				t.Logf("Success: For %s: expected %d, and got %d", test.name, test.expectedStatusCode, res.StatusCode)
			}
		}
	}

}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/make-reservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	reqRecorder := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.Reservation)
	handler.ServeHTTP(reqRecorder, req)

	if reqRecorder.Code != http.StatusOK {
		t.Errorf("Reservation handler returns wrong response code: got %d, want: %d", reqRecorder.Code, http.StatusOK)
	}

	//testcase where reservation not in session (reset everything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	reqRecorder = httptest.NewRecorder()
	handler.ServeHTTP(reqRecorder, req)

	if reqRecorder.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returns wrong response code: got %d, want: %d", reqRecorder.Code, http.StatusTemporaryRedirect)
	}

	//testcase where no-existing room (reset everything)
	req, _ = http.NewRequest("GET", "/make-reservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	reqRecorder = httptest.NewRecorder()

	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)

	handler.ServeHTTP(reqRecorder, req)
	if reqRecorder.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returns wrong response code: got %d, want: %d", reqRecorder.Code, http.StatusTemporaryRedirect)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}

	return ctx
}
