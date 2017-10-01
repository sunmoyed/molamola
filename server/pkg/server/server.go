package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sunmoyed/molamola/server/pkg/log"
)

var logger = log.DefaultLogger

type State struct {
	addr     string // Address to listen on
	datadir  string // Directory to store data
	assetdir string // Directory of assets

	assets map[string]struct{} // All assets
}

const assetPath string = "/assets"
const assetPathSlash string = assetPath + "/"

func NewServer(addr, datadir, assetdir string) (*State, error) {
	assets := make(map[string]struct{})
	assetInfo, assetInfoErr := ioutil.ReadDir(assetdir)
	if assetInfoErr != nil {
		return nil, assetInfoErr
	}
	for _, ai := range assetInfo {
		if ai.IsDir() {
			continue
		}
		assets[ai.Name()] = struct{}{}
		logger.Println("new server asset", ai.Name())
	}

	return &State{
		addr:     addr,
		datadir:  datadir,
		assetdir: assetdir,
		assets:   assets,
	}, nil
}

func (s *State) Run() error {
	// public: /public/<username>
	// edit:   /edit/<username>    < we shouldn't have username here cause we should have a user session already
	// api:    /api/<...>
	// assets: /assets/<...>

	http.HandleFunc("/", s.handleDefault)
	http.HandleFunc("/edit/", s.handleEdit)
	http.HandleFunc(assetPathSlash, s.handleAssets)

	logger.Printf("listening on %s", s.addr)
	return http.ListenAndServe(s.addr, nil)
}

func (s *State) handleDefault(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		s.handleError(w, r, http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, fmt.Sprintf("%s/%s", s.assetdir, "index.html"))
}

func (s *State) handleError(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	switch status {
	case http.StatusNotFound:
		fmt.Fprint(w, "404 molamola not found")
	}
}

func (s *State) handleAssets(w http.ResponseWriter, r *http.Request) {
	asset, assetErr := getAsset(r, assetPathSlash)
	if assetErr != nil {
		logger.Printf("handle assets: %s", assetErr)
		return
	}

	if _, ok := s.assets[asset]; !ok {
		s.handleError(w, r, http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, fmt.Sprintf("%s/%s", s.assetdir, asset))
}

func getAsset(r *http.Request, prefix string) (string, error) {
	logger.Println(r.URL.Path)
	if !strings.HasPrefix(r.URL.Path, prefix) {
		return "", fmt.Errorf("%s missing prefix %s", r.URL.Path, prefix)
	}
	cleanPrefix := strings.TrimPrefix(r.URL.Path, prefix)
	cleanPath := strings.TrimSuffix(cleanPrefix, "/")
	splitPath := strings.Split(cleanPath, "/")
	if len(splitPath) != 1 {
		return "", fmt.Errorf("could not get asset: %s", r.URL.Path)
	}
	return splitPath[0], nil
}
