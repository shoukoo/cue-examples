package config

apiVersion: "apps/v1"
kind:       "Deployment"
metadata: name: "web-app"
spec: {
	replicas: 6 // Number of replicas to run (you can adjust this based on your requirements)
	selector: {
		matchLabels: app: "web-app2"
	}
	template: {
		metadata: labels: app: "web-app"
		spec: containers: [{
			name:  "web-app-container"
			image: "your-docker-registry/your-web-app-image:latest"
			ports: [{
				containerPort: 80
			}]
		}]
	}
}
