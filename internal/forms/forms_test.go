package forms

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	isValid := form.Valid()
	if !isValid {
		t.Error("got invalid when it should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("got Valid when it should be In-valid")
	}

	postData := url.Values{}
	postData.Add("a", "a")
	postData.Add("b", "b")
	postData.Add("c", "c")

	r = httptest.NewRequest("POST", "/whatever", nil)
	r.PostForm = postData
	form = New(r.PostForm)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("got In-valid when it should be Valid")
	}
}

func TestForm_Has(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	has := form.Has("a")
	if has {
		t.Error("got Field when it is not exist")
	}

	postData = url.Values{}
	postData.Add("a", "a")
	form = New(postData)

	has = form.Has("a")
	if !has {
		t.Error("got false when the field exist")
	}
}

func TestForm_MinLength(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.MinLength("a", 10)
	if form.Valid() {
		t.Error("got Valid when it is In-Valid")
	}

	isError := form.Errors.Get("a")
	if isError == "" {
		t.Error("should have an error but not got one")
	}

	postData = url.Values{}
	postData.Add("a", "some_field")
	form = New(postData)

	form.MinLength("a", 100)

	if form.Valid() {
		t.Error("got Valid when it is In-Valid")
	}

	form = New(postData)
	form.MinLength("a", 10)
	if !form.Valid() {
		t.Error("got In-valid when it is Valid")
	}

	isError = form.Errors.Get("a")
	if isError != "" {
		t.Error("should not have an error but got one")
	}

}

func TestForm_IsEmail(t *testing.T) {
	postData := url.Values{}
	form := New(postData)

	form.IsEmail("email")
	if form.Valid() {
		t.Error("got Valid when it is non existing field")
	}

	postData = url.Values{}
	postData.Add("email", "not_email@notvalid")
	form = New(postData)

	form.IsEmail("email")

	if form.Valid() {
		t.Error("got Valid when it is In-Valid Email")
	}

	postData = url.Values{}
	postData.Add("email", "email@valid.com")
	form = New(postData)
	form.IsEmail("email")
	if !form.Valid() {
		t.Error("got In-valid when it is Valid Email")
	}
}

func TestNew(t *testing.T) {

}

func Test_errors_Add(t *testing.T) {

}

func Test_errors_Get(t *testing.T) {

}
