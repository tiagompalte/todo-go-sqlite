package server

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/tiagompalte/todo-go-sqlite/contract"
	"github.com/tiagompalte/todo-go-sqlite/entity"
)

type Server struct {
	service contract.Service
}

func NewServer(service contract.Service) Server {
	return Server{service}
}

type templateHtml struct {
	Todos entity.Todos
}

func (s Server) listAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	todos, err := s.service.List(req.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := template.Must(template.New("index.html").ParseFiles("server/index.html"))

	err = t.Execute(w, templateHtml{
		Todos: todos,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s Server) save(w http.ResponseWriter, req *http.Request) {
	text := req.FormValue("text")

	if strings.Trim(text, "") == "" {
		http.Redirect(w, req, "", http.StatusSeeOther)
		return
	}

	err := s.service.Create(req.Context(), text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "", http.StatusSeeOther)
}

func (s Server) markDone(w http.ResponseWriter, req *http.Request) {
	formId := req.FormValue("id")

	id, err := strconv.ParseUint(formId, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.service.MarkDone(req.Context(), uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "", http.StatusSeeOther)
}

func (s Server) delete(w http.ResponseWriter, req *http.Request) {
	formId := req.FormValue("id")

	id, err := strconv.ParseUint(formId, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.service.Delete(req.Context(), uint32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "", http.StatusSeeOther)
}

func (s Server) Handle() {

	http.HandleFunc("/", s.listAll)
	http.HandleFunc("/save", s.save)
	http.HandleFunc("/mark-done", s.markDone)
	http.HandleFunc("/delete", s.delete)

	http.ListenAndServe(":8080", nil)
}
