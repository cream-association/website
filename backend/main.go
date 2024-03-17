package main

import (
	"fmt"
	"log"
	"net/mail"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/daos"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir
	app.OnBeforeServe().Add(
		func(e *core.ServeEvent) error {
			e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
			return nil
		})

	// register migration command
	// loosely check if it was executed with go run (dev mod)
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto migration from ui on dev mode
		Automigrate: isGoRun,
	})

	// hook to send mail after message saved to db
	app.OnRecordAfterCreateRequest("message").Add(
		func(e *core.RecordCreateEvent) error {
			app.Dao().RunInTransaction(
				func(txDao *daos.Dao) error {
					app.Logger().Info(fmt.Sprintf("New message recorded: '%s'. Sending mail ...", e.Record.GetId()))

					message := &mailer.Message{
						From: mail.Address{
							Address: e.Record.GetString("email"),
							Name:    e.Record.GetString("full_name"),
						},
						To:      []mail.Address{{Address: app.Settings().Meta.SenderAddress}},
						Subject: e.Record.GetString("subject"),
						HTML:    e.Record.GetString("body"),
					}

					if err := app.NewMailClient().Send(message); err != nil {
						return nil
					}

					e.Record.Set("mail_sent", true)

					if err := txDao.SaveRecord(e.Record); err != nil {
						return err
					}

					return nil
				})

			return nil
		})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
