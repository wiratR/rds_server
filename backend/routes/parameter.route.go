package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wiratR/rds_server/controllers"
	"github.com/wiratR/rds_server/middleware"
)

func ParameterRoute(route fiber.Router) {
	// routes "/api/parameters"
	route.Get("/card_types/:id", middleware.DeserializeUser, controllers.GetCardType)       // Get a single CardType by ID
	route.Get("/card_types", middleware.DeserializeUser, controllers.GetCardTypeAll)        // Get all CardTypes
	route.Put("/card_types/:id", middleware.DeserializeUser, controllers.UpdateCardType)    // Update a CardType by ID
	route.Delete("/card_types/:id", middleware.DeserializeUser, controllers.DeleteCardType) // Delete a CardType by ID

	route.Get("/media_types/:id", middleware.DeserializeUser, controllers.GetMediaType)       // Get a single MediaType by ID
	route.Get("/media_types", middleware.DeserializeUser, controllers.GetMediaTypeAll)        // Get all MediaTypes
	route.Put("/media_types/:id", middleware.DeserializeUser, controllers.UpdateMediaType)    // Update a MediaType by ID
	route.Delete("/media_types/:id", middleware.DeserializeUser, controllers.DeleteMediaType) // Delete a MediaType by ID

	route.Get("/stations/:id", middleware.DeserializeUser, controllers.GetStation)       // Get a single Station by ID
	route.Get("/stations", middleware.DeserializeUser, controllers.GetStationAll)        // Get all Stations
	route.Put("/stations/:id", middleware.DeserializeUser, controllers.UpdateStation)    // Update a Station by ID
	route.Delete("/stations/:id", middleware.DeserializeUser, controllers.DeleteStation) // Delete a Station by ID

	route.Get("/lines/:id", middleware.DeserializeUser, controllers.GetLine)       // Get a single Line by ID
	route.Get("/lines", middleware.DeserializeUser, controllers.GetLineAll)        // Get all Lines
	route.Put("/lines/:id", middleware.DeserializeUser, controllers.UpdateLine)    // Update a Line by ID
	route.Delete("/lines/:id", middleware.DeserializeUser, controllers.DeleteLine) // Delete a Line by ID

	route.Get("/service_providers/:id", middleware.DeserializeUser, controllers.GetServiceProvider)       // Get a single ServiceProvider by ID
	route.Get("/service_providers", middleware.DeserializeUser, controllers.GetServiceProviderAll)        // Get all ServiceProviders
	route.Put("/service_providers/:id", middleware.DeserializeUser, controllers.UpdateServiceProvider)    // Update a ServiceProvider by ID
	route.Delete("/service_providers/:id", middleware.DeserializeUser, controllers.DeleteServiceProvider) // Delete a ServiceProvider by ID

	route.Get("/fares/:id", middleware.DeserializeUser, controllers.GetFare)       // Get a single Fare by ID
	route.Get("/fares", middleware.DeserializeUser, controllers.GetFareAll)        // Get all Fares
	route.Put("/fares/:id", middleware.DeserializeUser, controllers.UpdateFare)    // Update a Fare by ID
	route.Delete("/fares/:id", middleware.DeserializeUser, controllers.DeleteFare) // Delete a Fare by ID

	route.Get("/txn_types/:id", middleware.DeserializeUser, controllers.GetTxnType)       // Get a single TxnType by ID
	route.Get("/txn_types", middleware.DeserializeUser, controllers.GetTxnTypeAll)        // Get all TxnTypes
	route.Put("/txn_types/:id", middleware.DeserializeUser, controllers.UpdateTxnType)    // Update a TxnType by ID
	route.Delete("/txn_types/:id", middleware.DeserializeUser, controllers.DeleteTxnType) // Delete a TxnType by ID
}
