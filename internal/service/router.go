package service

import (
	"github.com/go-chi/chi"
	"github.com/rarimo/ipfs-svc/internal/service/handlers"
	"github.com/rarimo/ipfs-svc/internal/service/helpers"
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
				r.Post("/upload-file", handlers.UploadFile)

			})
		})
	})

	return r
}
