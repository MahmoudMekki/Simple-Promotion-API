package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/promo/view"

	"github.com/promo/internal"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) GetUserPromosHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	userA := internal.UserInt{}
	u := userA.Plug()
	promos, err := u.GetAll(s.DB)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(promos)
}

func (s *Server) CreateUserPromosHandler(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	promo := view.Promo{}
	json.NewDecoder(req.Body).Decode(&promo)
	userA := internal.UserInt{}
	u := userA.Plug()
	err := u.CreatePromo(s.DB, &promo)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{"Status": "Created Successfully"})
}

func (s *Server) UpdateUserPromoHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	q := p.ByName("id")
	id, err := strconv.Atoi(q)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(res).Encode(map[string]string{"error": "Wrong ID "})
		return
	}
	userA := internal.UserInt{}
	u := userA.Plug()
	promo := view.Promo{}
	json.NewDecoder(req.Body).Decode(&promo)
	err = u.UpdatePromo(id, s.DB, &promo)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{"Status": "Updated Successfully"})
}
func (s *Server) GetUserPromoByIDHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	q := p.ByName("id")
	id, err := strconv.Atoi(q)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(res).Encode(map[string]string{"error": "Wrong ID "})
		return
	}
	userA := internal.UserInt{}
	u := userA.Plug()
	err = u.GetPromo(id, s.DB)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(u)
}

func (s *Server) DeleteUserPromoHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	q := p.ByName("id")
	id, err := strconv.Atoi(q)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotAcceptable)
		json.NewEncoder(res).Encode(map[string]string{"error": "Wrong ID "})
		return
	}
	userA := internal.UserInt{}
	u := userA.Plug()
	err = u.DeletePromo(id, s.DB)
	if err != nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{"error": err.Error()})
		return
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{"Status": "Deleted Successfully"})
}
