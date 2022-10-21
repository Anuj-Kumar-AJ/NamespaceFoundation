import "namespacelabs.dev/foundation/std/fn"

test: fn.#Test & {
	name: "test-transcoding"

	binary: {
		from: go_package: "."
	}

	fixture: {
		serversUnderTest: [
			"namespacelabs.dev/foundation/internal/testdata/server/gogrpc",
		]
	}
}
