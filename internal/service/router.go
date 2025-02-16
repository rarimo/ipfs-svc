package service

import (
	"github.com/Dmytro-Hladkykh/ipfs-svc/internal/service/handlers"
	"github.com/Dmytro-Hladkykh/ipfs-svc/internal/service/helpers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxConfig(s.cfg),
		),
	)
	r.Route("/integrations/ipfs-svc", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/public", func(r chi.Router) {
				r.Post("/upload", handlers.UploadJSON)

			})
		})
	})

	return r
}
