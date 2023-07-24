package config

#Deployment: {
	apiVersion: "apps/v1"
	kind:       "Deployment"
	metadata: name: string
  let appname = metadata.name
	spec: {
		replicas: <4 // ensure the replicas less than 4
		selector: {
			matchLabels: app: appname
		}
		template: {
      metadata: labels: app: appname
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
