package config

apiVersion: "example.com/v1"
kind:       "GPGKey"
metadata: name: string
spec: {
  metadata: {
    user:    string
    purpose: string
  }
}
