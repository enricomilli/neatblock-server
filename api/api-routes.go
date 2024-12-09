package api

import (
	"github.com/enricomilli/neat-server/api/v1/pools"
	"github.com/enricomilli/neat-server/api/v1/wallets"
	"github.com/enricomilli/neat-server/middleware"
	"github.com/go-chi/chi/v5"
)

func CreateRoutes(router *chi.Mux) {

	router.Route("/api/", func(apiRouter chi.Router) {
		apiRouter.Route("/v1/", func(v1Router chi.Router) {

			// available only to users
			v1Router.Group(func(privRoute chi.Router) {
				privRoute.Use(middleware.Make(middleware.ValidJWToken))

				privRoute.Post("/pools/add", pools.HandleAddPool)
				privRoute.Post("/wallets/transactions/all", wallets.HandleAllTransactions)
			})

			// internal APIs
			// v1Router.Group(func(internal chi.Router) {
			// })
		})
	})

}
