package appInitialize

import "go-backend-template/internal/app"

var (
	apps = make([]app.Module, 0)
)

func GetApps() []app.Module {
	return apps
}
