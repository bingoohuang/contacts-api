package contactapi

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func Init(config Configuration) {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	logrus.Info("Init\n %+d", config)
	client := ConnectMongoDb(config.Database.Url)

	repo := NewContactRepository(&config, client)
	handler := NewContactHandler(client, repo)

	router.GET("/", handler.GetAllContacts)
	router.GET("/contacts/:email", handler.GetContactByEmail)
	router.POST("/contact/delete/:id", handler.DeleteContact)

	router.GET("/health", handler.HealthCheck)

	log.Info("port is :8080\n", config.Database.Url)

	// PORT environment variable was defined.
	router.Run(":" + config.Server.Port + "")
}
