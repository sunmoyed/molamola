package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LoginReqBody struct {
	Auth LoginReqBodyAuth `json:"auth"`
}

type LoginReqBodyAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespBody struct {
	Auth LoginRespBodyAuth `json:"auth"`
}

type LoginRespBodyAuth struct {
	Token string `json:"token"`
}

func (s *State) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != loginPath {
		logger.Printf("handleLogin error: badpath %s", r.URL.Path)
		s.handleError(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		http.ServeFile(w, r, s.webfilepath("index.html"))
	case http.MethodPost:
		s.handleLoginPost(w, r)
	}
}

func (s *State) handleLoginPost(w http.ResponseWriter, r *http.Request) {
	reqB, reqBErr := ioutil.ReadAll(r.Body)
	if reqBErr != nil {
		logger.Printf("handleLogin error: %s", reqBErr)
		s.handleError(w, r, http.StatusInternalServerError)
		return
	}

	logger.Printf("handle login post body: %s", string(reqB))
	loginReq := LoginReqBody{}
	if err := json.Unmarshal(reqB, &loginReq); err != nil {
		logger.Printf("handleLogin error: %s", err)
		s.handleError(w, r, http.StatusInternalServerError)
		return
	}

	tokStr, tokStrErr := s.us.LoginUser(loginReq.Auth.Username, loginReq.Auth.Password)
	if tokStrErr != nil {
		logger.Printf("handleLogin error: %s", tokStrErr)
		s.handleError(w, r, http.StatusUnauthorized)
		return
	}

	loginResp := LoginRespBody{
		Auth: LoginRespBodyAuth{
			Token: tokStr,
		},
	}
	respB, respBErr := json.Marshal(loginResp)
	if respBErr != nil {
		logger.Printf("handleLogin error: %s", respBErr)
		s.handleError(w, r, http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(respB); err != nil {
		logger.Printf("handleLogin error: %s", err)
	}
}
