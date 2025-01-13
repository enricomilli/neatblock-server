package api

import (
	"github.com/enricomilli/neat-server/api/v1/pools"
	v2pools "github.com/enricomilli/neat-server/api/v2/pools"
	"github.com/enricomilli/neat-server/middleware"
	"github.com/go-chi/chi/v5"
)

func CreateRoutes(router *chi.Mux) {

	router.Route("/api/", func(apiRouter chi.Router) {

		apiRouter.Get("/v2/foreman", v2pools.ForemanTests)

		apiRouter.Route("/v1/", func(v1Router chi.Router) {
			// v1Router.Get("/wallets/transaction", wallets.TestWalletInfo)

			// available only to users
			v1Router.Group(func(privRoute chi.Router) {
				privRoute.Use(middleware.Make(middleware.ValidJWToken))

				privRoute.Post("/pools/add", pools.HandleAddPool)
				privRoute.Delete("/pools", pools.HandlePoolDelete)
				// privRoute.Post("/wallets/transactions/all", wallets.HandleAllTransactions)
			})

			// internal APIs
			v1Router.Group(func(internal chi.Router) {
				internal.Use(middleware.Make(middleware.ValidateAPIToken))

				internal.Get("/pools/update-all", pools.HandleUpdateAll)
			})
		})
	})

}
