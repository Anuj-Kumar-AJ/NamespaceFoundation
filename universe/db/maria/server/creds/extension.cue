import (
	"namespacelabs.dev/foundation/std/fn"
)

extension: fn.#Extension & {
	import: [
		"namespacelabs.dev/foundation/universe/db/maria/incluster/creds",
	]
}

configure: fn.#Configure & {
	with: binary: "namespacelabs.dev/foundation/universe/db/maria/server/creds/tool"
}
