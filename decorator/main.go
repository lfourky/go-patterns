package main

import (
	"fmt"

	"github.com/lfourky/go-patterns/decorator/log"
	"github.com/lfourky/go-patterns/decorator/mail"
	"github.com/lfourky/go-patterns/decorator/persistence"
)

func main() {
	logger := log.New()
	logger.Debugf("This is a %s %s", "general", "case")
	logger.Errorf("Something caused an error. The answer to life is %d", 42)

	fmt.Println()

	logger = log.WithMailAlerts(logger, mail.MockMailer{})
	logger.Debugf("This is a %s %s", "general", "case")
	logger.Errorf("Something caused an error. The answer to life is %d", 42)

	fmt.Println()

	//This one is composed of both regular, mailer and persistent logger (because the previous logger was already composed of regular+mailer)
	logger = log.WithPersistence(logger, persistence.MockPersister{})
	logger.Debugf("This is a %s %s", "general", "case")
	logger.Errorf("Something caused an error. The answer to life is %d", 42)

	fmt.Println()

	//This one is composed of just a regular and persistent logger
	logger = log.WithPersistence(log.New(), persistence.MockPersister{})
	logger.Debugf("This is a %s %s", "general", "case")
	logger.Errorf("Something caused an error. The answer to life is %d", 42)
}
