package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/newsesame/jblog/controller"
)

// Setup routing information
func SetupRoutes(app *fiber.App) {

	/*
	   InsertEmployee ->
	   FindEmployeeID ->
	   FindAllEmployee ->
	   UpdateEmployeeByID ->
	   DeleteEmployeeByID ->
	   DeleteAllEmployee ->
	*/

	app.Get("/", controller.EmpList)
	app.Get("/cover/:id", controller.CoverImageHandler)
	app.Get("/allsongs", controller.SongList)
	app.Get("/songs", controller.SongListSorted)
	app.Post("/songcreate", controller.SongCreate)

	app.Get("/test", controller.Testing)
	app.Get("/:id", controller.BlogDetail)
	app.Post("/", controller.BlogCreate)
	app.Put("/:id", controller.BlogUpdate)
	app.Delete("/:id", controller.BlogDelete)

}

/*
InsertEmployee ->
FindEmployeeID ->
FindAllEmployee ->
UpdateEmployeeByID ->
DeleteEmployeeByID ->
DeleteAllEmployee ->
*/

// // Setup routing information
// func SetupRoutes(app *fiber.App) {

// 	// list => get
// 	// read blog => get (id)
// 	// add => post
// 	// update => put
// 	// delete => delete

// 	app.Get("/", controller.BlogList)
// 	app.Get("/test", controller.Testing)
// 	app.Get("/:id", controller.BlogDetail)
// 	app.Post("/", controller.BlogCreate)
// 	app.Put("/:id", controller.BlogUpdate)
// 	app.Delete("/:id", controller.BlogDelete)

// }
