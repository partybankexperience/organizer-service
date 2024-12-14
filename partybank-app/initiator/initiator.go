package initiator

import (
	"github.com/djfemz/organizer-service/partybank-app/config"
	"github.com/djfemz/organizer-service/partybank-app/repositories"
	"github.com/djfemz/organizer-service/partybank-app/security/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"log"
	"os"
	"time"
)

func init() {
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading configuration: ", err)
	}
	log.Println("connecting to db")
	db = repositories.Connect()
	log.Println("Connected to db: ", db)
}

func Initiate() {
	config.GoogleConfig()
	go startCron()
	router := gin.Default()
	configureAppComponents()
	middlewares.AddComponentsToRouters(router,
		organizerController,
		eventController, seriesController,
		ticketController, authService,
		attendeeController, authController, attendeeRepository)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err = router.Run(":" + port)
	if err != nil {
		log.Println("Error starting server: ", err)
	}
}

func startCron() {
	loc, err := time.LoadLocation("Africa/Lagos")
	if err != nil {
		log.Fatal("timezone not found")
	}
	job := cron.NewWithLocation(loc)
	err = job.AddFunc("5 0 * * * *", func() {
		err = eventRepository.RemovePastEvents()
		if err != nil {
			log.Println("failed to Remove past events")
		}
	})
	if err != nil {
		log.Println("failed to schedule remove past events job")
	}
	log.Println("Starting scheduler...")
	job.Start()
}
