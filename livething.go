package livething

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yaaaaashiki/livething/db"
	"github.com/yaaaaashiki/livething/domain/repository"
	"github.com/yaaaaashiki/livething/interfaceadapter/controller"
	"github.com/yaaaaashiki/livething/usecase"
)

// This holds database connection and router settings based on gin.
type Server struct {
	db  *gorm.DB
	gin *gin.Engine
}

// New returns server object.
func New() *Server {
	return &Server{}
}

// Close makes the database connection to close.
func (s *Server) Close() error {
	return s.db.Close()
}

// Init initialize server state. Connecting to database, compiling templates,
// and settings router.
func (s *Server) Init(dbconf, env string, debug bool) {
	cs, err := db.NewConfigsFromFile(dbconf)
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	db, err := cs.Open(env)
	if err != nil {
		log.Fatalf("db initialization failed: %s", err)
	}

	s.db = db
	s.gin = gin.Default()
	s.Route()
}

// Run starts running http server.
func (s *Server) Run(addr string) {
	log.Printf("start listening on %s", addr)

	s.gin.Run(addr)
}

func (s *Server) Route() {

	r := s.gin

	r.LoadHTMLGlob("view/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	api := r.Group("/api")

	objectRepository := repository.NewObjectRepository(s.db)
	setCurrentObjectStatusUseCase := usecase.NewSetCurrentObjectStatusUseCase(objectRepository)
	setCurrentObjectStatusController := controller.NewSetCurrentObjectStatusController(setCurrentObjectStatusUseCase)

	/*
		slackService := slack.NewSlackAPIService(env.SlackAPIUrl)
		notificationUsecase := usecase.NewNotificationUsecase(slackService)

		r.POST("/slack", func(c *gin.Context) {
			//
			c.JSON(http.StatusOK, gin.H{
				"status": "OK!",
			})
		})
		r.Run(":3000")
	*/

	api.POST("/objects", setCurrentObjectStatusController.Execute)

	//wifi.PingStaticIP()
}
