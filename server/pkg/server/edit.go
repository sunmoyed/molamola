package server

import (
	"fmt"
	"net/http"
)

func (s *State) handleEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "edit unimplemented")
}
