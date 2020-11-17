package server

import (
	"database/sql"
	"log"
	"net/http"

	storage "github.com/promo/DB"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	DB     *sql.DB
	Router *httprouter.Router
}

func Launch(port string) {

	s := Server{
		DB:     storage.New(),
		Router: httprouter.New(),
	}

	s.Router.GET("/api/user/promotions", s.GetUserPromosHandler)
	s.Router.POST("/api/user/promotions", s.CreateUserPromosHandler)
	s.Router.PUT("/api/user/promotions/:id", s.UpdateUserPromoHandler)
	s.Router.GET("/api/user/promotions/:id", s.GetUserPromoByIDHandler)
	s.Router.DELETE("/api/user/promotions/:id", s.DeleteUserPromoHandler)

	s.Router.GET("/api/inc/promotions", s.GetIncPromosHandler)
	s.Router.POST("/api/inc/promotions", s.CreateIncPromosHandler)
	s.Router.PUT("/api/inc/promotions/:id", s.UpdateIncPromoHandler)
	s.Router.GET("/api/inc/promotions/:id", s.GetIncPromoByIDHandler)
	s.Router.DELETE("/api/inc/promotions/:id", s.DeleteIncPromoHandler)

	if err := http.ListenAndServe(":"+port, s.Router); err != nil {
		log.Fatalln(err)
	}

}
