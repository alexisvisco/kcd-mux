package kcdmux

import (
	"github.com/alexisvisco/kcd"
	"github.com/gorilla/mux"
	"net/http"
)

func Setup() {
	kcd.Config.StringsExtractors = append(kcd.Config.StringsExtractors, MuxPathExtractor{})
}

type MuxPathExtractor struct{}

func (m MuxPathExtractor) Extract(req *http.Request, res http.ResponseWriter, valueOfTag string) ([]string, error) {
	if value, ok := mux.Vars(req)[valueOfTag]; ok {
		return []string{value}, nil
	}

	return nil, nil
}

func (m MuxPathExtractor) Tag() string {
	return "path"
}
