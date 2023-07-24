package config

#Deployment: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
  _name: "web-app"
	metadata: name: _name // make sure the name is consistent
	spec: {
		replicas: <4 // ensure the replicas less than 4
		selector: {
			matchLabels: app: _name
		}
		template: {
      metadata: labels: app: _name
			spec: containers: [{
				name:  "web-app-container"
				image: =~ "^.*@sha256:[0-9a-f]{64}$" // ensure it sues image digest
				ports: [{
					containerPort: 80
				}]
			}]
		}
	}
}

#Default: {
	apiVersion: "example.com/v1"
	kind:       "GPGKey"
	metadata: name: string
	spec: {
		metadata: {
			user:    string
			purpose: string
		}
	}
}
