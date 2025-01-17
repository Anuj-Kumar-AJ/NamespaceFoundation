import (
	"namespacelabs.dev/foundation/std/fn"
	"namespacelabs.dev/foundation/std/secrets"
	"namespacelabs.dev/foundation/std/core/info"
	"namespacelabs.dev/foundation/std/go/grpc/interceptors"
	"namespacelabs.dev/foundation/std/go/http/middleware"
)

extension: fn.#Extension & {
	hasInitializerIn: "GO"

	instantiate: {
		dsn: secrets.#Exports.Secret & {
			name: "sentry-dsn"
		}
		serverInfo:     info.#Exports.ServerInfo
		"interceptors": interceptors.#Exports.InterceptorRegistration
		"middleware":   middleware.#Exports.Middleware
	}
}
