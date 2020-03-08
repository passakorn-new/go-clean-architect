package firebase

import (
	"fmt"

	"golang.org/x/net/context"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func ConnectFirebase() (*firebase.App, error) {
	opt := option.WithCredentialsFile("firebase_admin_sdk_config.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil
}
