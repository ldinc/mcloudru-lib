package mcloudru_ru

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Login struct {
	login string
	pswd  string
}

func NewLogin(login, pswd string) *Login {
	l := new(Login)
	l.login = login
	l.pswd = pswd

	return l
}

func (l *Login) Connect() error {
	if l.login == "" {
		return errors.New("Empty login field")
	}
	if l.pswd == "" {
		return errors.New("Empty password field")
	}
	//reqstr := "Login=" + l.login + "&Domain=" + Domain + "&Password=" + l.pswd
	//fmt.Println("req str: " + reqstr)
	cjar, _ := cookiejar.New(nil)
	client := http.Client{Jar: cjar}

	form := url.Values{}
	form.Add("Login", l.login)
	form.Add("Domain", Domain)
	form.Add("Password", l.pswd)

	fmt.Print("form: ")
	fmt.Println(form)
	fmt.Println(form.Encode())

	//resp, err := http.PostForm(AuthCGI, form)

	req, err := http.NewRequest("POST", AuthCGI, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-ww-form-urlencoded")

	fmt.Println(req)
	resp, err := client.Do(req)

	fmt.Println(err)
	fmt.Println(resp)
	fmt.Println(resp.Location())
	fmt.Println(resp.Cookies())

	return nil
}
