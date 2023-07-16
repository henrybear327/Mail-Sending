package main

import (
	"context"
	"log"

	protonDriveAPI "github.com/henrybear327/Proton-API-Bridge"
	"github.com/henrybear327/Proton-API-Bridge/common"
	proton "github.com/henrybear327/go-proton-api"
)

var (
	NumberOfEmails        = 1
	EmailSubject          = "Test subject"
	RecipientEmailAddress = "a@pm.me"
	TemplateFile          = "./template/darkMode.html"
	EmailAttachments      = []string{"template/moon.png", "template/sun.png"}
	EmailContentIDs       = []string{"moon", "sun"}
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	config := common.NewConfigForIntegrationTests()
	protonDrive, _, err := protonDriveAPI.NewProtonDrive(ctx, config, func(auth proton.Auth) {}, func() {})
	if err != nil {
		log.Fatalln(err)
	}

	errChan := make(chan error)
	for i := 0; i < NumberOfEmails; i++ {
		go protonDrive.SendEmail(ctx, i, errChan, &protonDriveAPI.MailSendingParameters{
			TemplateFile:          TemplateFile,
			EmailSubject:          EmailSubject,
			RecipientEmailAddress: RecipientEmailAddress,
			EmailAttachments:      EmailAttachments,
			EmailContentIDs:       EmailContentIDs,
		})
	}

	for i := 0; i < NumberOfEmails; i++ {
		if err := <-errChan; err != nil {
			log.Fatalln(err)
		}
	}
}
