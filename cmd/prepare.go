package prepare

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Chandan94f/mb-view/internal/controller"
	"github.com/Chandan94f/mb-view/internal/dao"
	"github.com/Chandan94f/mb-view/internal/service"

	// "github.com/Chandan94f/mb-view/internal/utility/logger"

	// response "github.com/Chandan94f/mb-view/internal/utility/response"
	"github.com/Chandan94f/mb-utility/utility/logger"

	"github.com/Chandan94f/mb-utility/utility/response"
	"github.com/gorilla/mux"
)

func Prepare() error {
	fmt.Println("connecting to data base here")

	// Create a new mux router
	router := mux.NewRouter()

	resp := response.NewResponse()

	daoService := dao.NewDaoService(dao.DoaParams{
		Logger: logger.NewBasicLogger(),
	})

	srv := service.NewService(service.ServiceParams{
		Dao: daoService,
	})

	// Register routes from the controller
	// controller := -/ to be added lated to enabel the wait gp
	controller.NewController(controller.ControllerParams{
		Resp:    resp,
		Dao:     daoService,
		Service: srv,
	}).SetupRoutes(router)

	// Start the HTTP server
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return err
	}

	return nil

	// controller.wait()  // to be added later
}
