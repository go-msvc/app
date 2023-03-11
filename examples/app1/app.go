package main

import (
	"context"

	"github.com/go-msvc/app"
	"github.com/go-msvc/errors"
)

func New() app.App {
	//start immediately with a menu for the user to select from
	mainMenu := app.Section(app.StaticText("Main Menu"),
		app.Text(app.StaticText("Some extra text...")),
		app.Button(app.StaticText("Register"), showRegister),
		app.Button(app.StaticText("Login"), showLogin),
		app.Button(app.StaticText("My Profile"), showMyProfile),
		app.Button(app.StaticText("Users"), showUsers),
		app.Text(app.StaticText("Some footer text...")),
	)
	return mainMenu
}

func showRegister(ctx context.Context) app.App {
	return app.Section(app.StaticText("Register"),
		app.Text(app.StaticText("Enter the email you want to use for login.")),
		app.Input(app.StaticText("email"), "email"),
		app.Button(app.StaticText("Register"), doRegister),
	)
}

func doRegister(ctx context.Context) error {
	return errors.Errorf("NYI")
} //doRegister()

func showLogin(ctx context.Context) app.App {
	return app.Section(app.StaticText("Login"),
		app.Input(app.StaticText("email"), "email"),
		app.Input(app.StaticText("password"), "password"),
		app.Button(app.StaticText("Login"), doLogin),
	)
}

func doLogin(ctx context.Context) error {
	return errors.Errorf("NYI")
} //doLogin()

func showUsers(ctx context.Context) app.App {
	return app.Section(app.StaticText("Users"),
		app.Button(app.StaticText("user1"), nil),
		app.Button(app.StaticText("user2"), nil),
		app.Button(app.StaticText("user3"), nil),
	)
}

func showMyProfile(ctx context.Context) app.App {
	return app.Group(
		app.Button(app.StaticText("name"), nil),
		app.Button(app.StaticText("dob"), nil),
	)
}
