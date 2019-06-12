package router

import (
	"go-todo/models"
	"github.com/gin-gonic/gin"

func InitRouter(cfg *models.Config) (*gin.Engine, error) {
	// change this to gin.ReleaseMode on production
	gin.SetMode(gin.DebugMode)
	router := gin.New()

	router.Use(cors.CORSMiddleware())
	router.Use(gin.Logger())

	// start swagger
	InitSwaggerRouter(router)

	// Setup No Route Message
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "Route Not Found"})
	})

	// Initiate Routers

	client, err := connection.GetMongoClient(cfg)
	if err != nil {
		msg := fmt.Sprintf("InitRouter :: %s", err.Error())
		fmt.Println(msg)
	} else {

		var dbService interfaces.DbService
		dbService, err = services.GetMongoService(cfg, client, "user")
		if err != nil {
			// if user collection doesn't exist,
			// it will be automatically created on insert
			fmt.Println(err)
		}

		InitAuthRouter(router, cfg, &dbService)

	}

	if cfg.Camera == "true" {

		// get mongo client
		client, err := connection.GetMongoClient(cfg)
		if err != nil {
			msg := fmt.Sprintf("InitRouter :: %s", err.Error())
			fmt.Println(msg)
		} else {

			var dbService interfaces.DbService
			dbService, err = services.GetMongoService(cfg, client, "camera")
			if err != nil {
				// if camera collection doesn't exist,
				// it will be automatically created on insert
				fmt.Println(err)
			}

			InitCameraRouter(router, cfg, &dbService)

		}
	}

	if cfg.Camera == "true" {

		// get mongo client
		client, err := connection.GetRedisClient(cfg)
		if err != nil {
			msg := fmt.Sprintf("InitRouter :: %s", err.Error())
			fmt.Println(msg)
		} else {

			InitStreamRouter(router, cfg, client)

		}
	}

	if cfg.Rules == "true" {

		// get mongo client
		client, err := connection.GetMongoClient(cfg)
		if err != nil {
			msg := fmt.Sprintf("InitRouter :: %s", err.Error())
			fmt.Println(msg)
		} else {

			var dbService interfaces.DbService
			dbService, err = services.GetMongoService(cfg, client, "rules")
			if err != nil {
				// if rules collection doesn't exist,
				// it will be automatically created on insert
				fmt.Println(err)
			}

			InitRulesRouter(router, cfg, &dbService)

		}
	}

	if cfg.Alert == "true" {

		// get mongo client
		client, err := connection.GetMongoClient(cfg)
		if err != nil {
			msg := fmt.Sprintf("InitRouter :: %s", err.Error())
			fmt.Println(msg)
		} else {

			var dbService interfaces.DbService
			dbService, err = services.GetMongoService(cfg, client, "alerts")
			if err != nil {
				// if rules collection doesn't exist,
				// it will be automatically created on insert
				fmt.Println(err)
			}

			InitAlertRouter(router, cfg, &dbService)

		}
	}

	if cfg.LPRService == "true" {

		// get mongo client
		client, err := connection.GetMongoClient(cfg)
		if err != nil {
			msg := fmt.Sprintf("InitRouter :: %s", err.Error())
			fmt.Println(msg)
		} else {

			var dbService interfaces.DbService
			dbService, err = services.GetMongoService(cfg, client, "refined_tracks")
			if err != nil {
				// if refined_tracks collection doesn't exist,
				// someone needs to create it first before using the API
				fmt.Println(err)
			}

			InitLprRouter(router, cfg, &dbService)

		}
	}

	if cfg.FaceService == "true" {

		// get mongo client
		client, err := connection.GetMongoClient(cfg)
		if err != nil {
			msg := fmt.Sprintf("InitRouter :: %s", err.Error())
			fmt.Println(msg)
		} else {

			var rDbService interfaces.DbService
			rDbService, err = services.GetMongoService(cfg, client, "results_fr")
			if err != nil {
				// if results_fr collection doesn't exist,
				// someone needs to create it first before using the API
				fmt.Println(err)
			}

			var iDbService interfaces.DbService
			iDbService, err = services.GetMongoService(cfg, client, "identities")
			if err != nil {
				// if identities collection doesn't exist,
				// someone needs to create it first before using the API
				fmt.Println(err)
			}

			InitFaceRouter(router, cfg, &rDbService, &iDbService)

		}
	}

	if cfg.MapService == "true" {

		// get mongo client
		client, err := connection.GetMongoClient(cfg)
		if err != nil {
			msg := fmt.Sprintf("InitRouter :: %s", err.Error())
			fmt.Println(msg)
		} else {

			var dbService interfaces.DbService
			dbService, err = services.GetMongoService(cfg, client, "map")
			if err != nil {
				// if map collection doesn't exist,
				// it will be automatically created on insert
				fmt.Println(err)
			}

			InitMapRouter(router, cfg, &dbService)

		}
	}

	return router, nil
}