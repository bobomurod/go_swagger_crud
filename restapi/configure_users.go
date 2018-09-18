// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"users-api/restapi/operations"
	"users-api/restapi/operations/users"
)

//go:generate swagger generate server --target .. --name users-api --spec ../../../../Downloads/LearningTech_learning-users-api_0.0.1_swagger.yaml

func configureFlags(api *operations.UsersAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UsersAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.UsersAddUserHandler = users.AddUserHandlerFunc(func(params users.AddUserParams) middleware.Responder {
		return middleware.NotImplemented("operation users.AddUser has not yet been implemented")
	})
	api.DelUserHandler = operations.DelUserHandlerFunc(func(params operations.DelUserParams) middleware.Responder {
		return middleware.NotImplemented("operation .DelUser has not yet been implemented")
	})
	api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation users.GetUserByID has not yet been implemented")
	})
	//api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {
	//	return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
	//})

	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {

		return users.NewGetUsersOK()       // вызывает функцию NewGetUserOK() - файл get_users_response.go
											// возвращает массив из



		//return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
		})




	api.UsersUpdateUserByIDHandler = users.UpdateUserByIDHandlerFunc(func(params users.UpdateUserByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation users.UpdateUserByID has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
