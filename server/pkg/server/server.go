package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sunmoyed/molamola/server/pkg/log"
	"github.com/sunmoyed/molamola/server/pkg/user"
	"github.com/sunmoyed/molamola/server/pkg/util"
)

var logger = log.DefaultLogger

type State struct {
	addr    string          // Address to listen on
	datadir string          // Directory to store data
	webdir  string          // Directory of webfiles
	us      *user.UserState // User state
}

const (
	webPath   string = "/assets"
	loginPath string = "/login"

	webPathSlash string = webPath + "/"
)

func NewServer(addr, datadir, webdir string) (*State, error) {
	logger.Println("addr", addr)
	logger.Println("data", datadir)
	logger.Println("webfiles", webdir)

	us, usErr := user.NewUserState(datadir)
	if usErr != nil {
		return nil, usErr
	}

	return &State{
		addr:    addr,
		datadir: datadir,
		webdir:  webdir,
		us:      us,
	}, nil
}

func (s *State) Run() error {
	// public: /public/<username>/<listname>
	// edit:   /edit/<listname>    < we shouldn't have username here cause we should have a user session already
	// api:    /api/<...>
	// assets: /assets/<...>

	http.HandleFunc("/", s.handleDefault)
	http.HandleFunc("/edit/", s.handleEdit)
	http.HandleFunc(loginPath, s.handleLogin)
	http.HandleFunc(webPathSlash, s.handleWeb)

	logger.Printf("listening on %s", s.addr)
	return http.ListenAndServe(s.addr, nil)
}

func (s *State) webfilepath(path string) string {
	return s.webdir + "/" + path
}

func (s *State) handleDefault(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.handleError(w, r, http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, s.webfilepath("index.html"))
}

func (s *State) handleError(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	switch status {
	case http.StatusUnauthorized:
		fmt.Fprint(w, "401 molamola unauthorized")
	case http.StatusNotFound:
		fmt.Fprint(w, "404 molamola not found")
	case http.StatusInternalServerError:
		fmt.Fprint(w, "500 molamola internal server error")
	}
}

func (s *State) handleWeb(w http.ResponseWriter, r *http.Request) {
	webpath := strings.TrimPrefix(r.URL.Path, "/")

	if err := validateWebPath(r.URL.Path); err != nil {
		logger.Println(err)
		s.handleError(w, r, http.StatusNotFound)
		return
	}

	path := s.webfilepath(webpath)
	if ok, err := util.FileExists(path); err != nil {
		logger.Printf("handleWeb error: %s", err)
		s.handleError(w, r, http.StatusInternalServerError)
		return
	} else if !ok {
		s.handleError(w, r, http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, path)
}

func validateWebPath(path string) error {
	// XXX Check that this is actually working
	split := strings.Split(path, "/")
	for _, s := range split {
		switch s {
		case ".":
			fallthrough
		case "..":
			return fmt.Errorf("invalid web path: %s", s)
		}
	}
	return nil
}
