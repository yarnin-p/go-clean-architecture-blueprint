package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-clean-architecture-blueprint/utils"
	"gorm.io/driver/sqlserver"

	_departmentModel "go-clean-architecture-blueprint/api/v1/models"
	_departmentInterface "go-clean-architecture-blueprint/api/v1/modules/department/interfaces"
	_departmentRepository "go-clean-architecture-blueprint/api/v1/modules/department/repositories"
	_departmentRoute "go-clean-architecture-blueprint/api/v1/modules/department/routes"
	_departmentUseCase "go-clean-architecture-blueprint/api/v1/modules/department/usecases"

	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	departmentUC _departmentInterface.DepartmentUseCase
}

func NewApp() *App {
	initLog()
	db := initDB()
	migrateDB(db)

	departmentRepo := _departmentRepository.NewDepartmentRepository(db)

	return &App{
		departmentUC: _departmentUseCase.NewDepartmentUseCase(departmentRepo),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	//router.Use(
	//	gin.Recovery(),
	//	gin.Logger(),
	//)

	basePath := router.Group("/api/v1")
	_departmentRoute.RegisterHTTPEndpoints(basePath, a.departmentUC)

	// HTTP Server with graceful shutdown
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB() *gorm.DB {
	dsn := "sqlserver://root:password@localhost:1433?database=" + viper.GetString("DATABASE")
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func migrateDB(db *gorm.DB) {
	if err := db.AutoMigrate(
		&_departmentModel.Department{},
	); err != nil {
		log.Fatal(err.Error(), "Failed to migrate database")
	}
}

func initLog() {
	// open a file
	logFileName := utils.GetCurrentTime("2006-01-02")
	f, err := os.OpenFile("./storages/logs/"+logFileName+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// don't forget to close it
	//defer f.Close()

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     false,
	})

	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)
}
