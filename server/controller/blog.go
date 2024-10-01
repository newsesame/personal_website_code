package controller

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
	"github.com/newsesame/jblog/database"
	"github.com/newsesame/jblog/model"
)

/*
InsertEmployee ->
FindEmployeeID ->
FindAllEmployee ->
UpdateEmployeeByID ->
DeleteEmployeeByID ->
DeleteAllEmployee ->
*/
type BlogRepo struct {
	MongoCollection *mongo.Collection
}

func (r *BlogRepo) InsertBlog(blog *model.Blog) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), blog)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BlogRepo) FindBlogID(blogID string) (*model.Blog, error) {
	var blog model.Blog

	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "blog_id", Value: blogID}}).Decode(&blog)

	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (r *BlogRepo) FindAllBlog() ([]model.Blog, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var blogs []model.Blog

	// Decode
	err = results.All(context.Background(), &blogs)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}
	return blogs, nil
}

func (r *BlogRepo) UpdateBlogByID(blogID string, updateBlog *model.Blog) (int64, error) {
	results, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "blog_id", Value: blogID}},
		bson.D{{Key: "$set", Value: updateBlog}})

	if err != nil {
		return 0, err
	}

	return results.ModifiedCount, nil

}

func (r *BlogRepo) DeleteBlogByID(blogID string) (int64, error) {
	results, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "blog_id", Value: blogID}})

	if err != nil {
		return 0, err
	}
	return results.DeletedCount, nil

}

func (r *BlogRepo) DeleteAllBlog() (int64, error) {
	results, err := r.MongoCollection.DeleteMany(context.Background(),
		bson.D{})

	if err != nil {
		return 0, err
	}
	return results.DeletedCount, nil

}

// Blog list
func EmpList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	// Sleep to add some delay in API response
	// time.Sleep(time.Millisecond * 1500)

	// db := database.DBConn
	coll := database.Blog_Coll_Conn
	blogRepo := BlogRepo{MongoCollection: coll}

	// Declare a variable to hold the list of blog records.
	// var records []model.Employee

	// Query the database to find all blog records.
	results, err := blogRepo.FindAllBlog()

	if err != nil {
		log.Print("get operation failed", err)
		c.Status(400)
		return c.JSON(context)
	}

	log.Println("Blogs: ", results)

	// Add the retrieved blog records to the context map.
	context["blog_records"] = results

	c.Status(200)
	return c.JSON(context)
}

// Testing
func Testing(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Server Running",
	}

	c.Status(200)
	return c.JSON(context)
}

// Blog detail page
func BlogDetail(c *fiber.Ctx) error {
	c.Status(400)
	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	// Extract the param from URL
	id := c.Params("id")

	// var record model.Blog

	coll := database.Blog_Coll_Conn
	blogRepo := BlogRepo{MongoCollection: coll}

	record, err := blogRepo.FindBlogID(id)

	if err != nil {
		log.Println("Find Operation Failed", err)
		context["msg"] = "Record not Found."

		c.Status(404)
		return c.JSON(context)
	}

	context["record"] = record
	context["statusText"] = "Ok"
	context["msg"] = "Blog Detail"
	c.Status(200)
	return c.JSON(context)
}

// // Add a Blog into Database
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}

	record := new(model.Blog)

	// When the body content is not as expected
	if err := c.BodyParser(record); err != nil {
		log.Println("Error in parsing request.")
		log.Println(record)
		context["statusText"] = ""
		context["msg"] = "Something went wrong."

		// Bad Request 400
		c.Status(400)
		return c.JSON(context)
	}

	// Create the record
	coll := database.Blog_Coll_Conn
	blogRepo := BlogRepo{MongoCollection: coll}

	_, err := blogRepo.InsertBlog(record)
	// result := database.DBConn.Create(record)

	if err != nil {
		log.Println("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Something went wrong."
	}

	context["msg"] = "Record is saved successully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

// Update  a Blog
func BlogUpdate(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Blog",
	}

	// //http://localhost:8000/2

	id := c.Params("id")

	record := new(model.Blog)
	coll := database.Blog_Coll_Conn
	blogRepo := BlogRepo{MongoCollection: coll}

	record, err := blogRepo.FindBlogID(id)

	// database.DBConn.First(&record, id)

	// Record not Found.
	if err != nil {
		log.Println("Record not Found.")

		// For front end guy to watch
		context["statusText"] = ""

		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}

	// When the body content is not as expected
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")

		context["msg"] = "Something went wrong."
		c.Status(400)
		return c.JSON(context)
	}

	// Save the changers
	count, err := blogRepo.UpdateBlogByID(id, record)

	if err != nil {
		log.Println("Error in saving data.")

		context["msg"] = "Error in saving data."
		c.Status(500)
		return c.JSON(context)
	}

	context["msg"] = "Record updated successfully. No. of updated blog: " + strconv.FormatInt(count, 10)
	context["data"] = record

	c.Status(200)
	return c.JSON(context)
}

// Delete a Blog
func BlogDelete(c *fiber.Ctx) error {

	c.Status(400)
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Delete Blog for the given ID",
	}

	id := c.Params("id")

	// var record model.Blog

	coll := database.Blog_Coll_Conn
	blogRepo := BlogRepo{MongoCollection: coll}

	_, err := blogRepo.FindBlogID(id)

	if err != nil {
		log.Println("Record not Found.")
		context["msg"] = "Record not Found."

		return c.JSON(context)
	}

	// // Remove image
	// filename := record.Image

	// err := os.Remove(filename)
	// if err != nil {
	// 	log.Println("Error in deleting file.", err)
	// }

	_, err = blogRepo.DeleteBlogByID(id)
	// result := database.DBConn.Delete(record)

	if err != nil {
		context["msg"] = "Delete operation went wrong."
		return c.JSON(context)
	}

	context["statusText"] = "Ok"
	context["msg"] = "Record deleted successfully."
	c.Status(200)
	return c.JSON(context)
}
