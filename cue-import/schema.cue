package config

gpg: {
	apiVersion: "example.com/v1"
	kind:       "GPGKey"
  test: string

	metadata: name: string
	spec: {
		metadata: {
			user:    string
			purpose: string
		}
	}
}
