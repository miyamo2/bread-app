package core

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

	fs "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Client = *fs.Client

type Context = context.Context

var client Client

var ctx Context

func init() {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		home = os.Getenv("APPDATA")
	}

	// Windows: %APPDATA%/.config/breadapp/service-account-file.json
	// Linux/Unix/Mac: $HOME/.config/breadapp/service-account-file.json
	credentialFilePath := filepath.Join(home, ".config", "breadapp", "firebase_credential.json")
	opt := option.WithCredentialsFile(credentialFilePath)
	ctx = context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalln(err)
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}

// Singleton
func GetContext() Context {
	return ctx
}

// Singleton
func GetClient() Client {
	return client
}
