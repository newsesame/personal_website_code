package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	// "image"
	// "io"
	"log"
	// "net/http"
	"os/exec"
	"reflect"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
type SongRepo struct {
	MongoCollection *mongo.Collection
}

func (r *SongRepo) InsertSong(song *model.Song) (interface{}, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), song)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *SongRepo) FindSongID(songID string) (*model.Song, error) {
	var song model.Song

	err := r.MongoCollection.FindOne(context.Background(), bson.D{{Key: "song_id", Value: songID}}).Decode(&song)

	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (r *SongRepo) FindAllSong() ([]model.Song, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var songs []model.Song

	// Decode
	err = results.All(context.Background(), &songs)

	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}

	return songs, nil
}

func sortedKeys(m interface{}, desc bool) ([]string, error) {
	var keys []string

	switch myMap := m.(type) {
	case map[string]map[string][]model.Song:
		for k := range myMap {
			keys = append(keys, k)
		}
	case map[string][]model.Song:
		for k := range myMap {
			keys = append(keys, k)
		}
	default:
		return nil, fmt.Errorf("m is not of the expected type")
	}
	// fmt.Println(reflect.TypeOf(m))
	// myMap, ok := m.(map[string]any)
	// fmt.Println(myMap)
	// fmt.Println(reflect.TypeOf(myMap))
	// if ok {
	// 	for k := range myMap {
	// 		keys = append(keys, k)
	// 	}
	// } else {
	// 	log.Fatal("m is not a map[string]interface{}")

	// }

	// 排序
	sort.Strings(keys)
	if desc {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	}
	fmt.Println("Sorted keys: ", keys)

	return keys, nil
}

type SortedSongs struct {
	Year  string       `json:"year"`
	Month string       `json:"month"`
	Songs []model.Song `json:"songs"`
}

func (r *SongRepo) GroupSongByMonth() ([]SortedSongs, error) {
	songs, err := r.FindAllSong()

	if err != nil {
		return nil, err
	}
	fmt.Println("Songs: ", songs)
	// Group the songs by years and months
	groupedSongs := make(map[string]map[string][]model.Song)

	for _, song := range songs {

		date, err := time.Parse("2006-01-03", song.AddedDate)

		if err != nil {
			log.Fatal("Failed to compile the date of record in the database.", err)
		}

		year := date.Format("2006")

		fmt.Println(reflect.TypeOf(year))
		month := date.Format("01")

		if groupedSongs[year] == nil {
			groupedSongs[year] = make(map[string][]model.Song)
		}

		groupedSongs[year][month] = append(groupedSongs[year][month], song)
	}

	fmt.Println("grouped songs: ", groupedSongs)
	// Sort the groups
	sortedYears, err := sortedKeys(groupedSongs, true)

	if err != nil {
		log.Fatal("sorted keys problem", err)
	}

	var sortedSongs []SortedSongs

	for _, year := range sortedYears {
		fmt.Println(year)
		months := groupedSongs[year]
		fmt.Println(months)
		sortedMonths, _ := sortedKeys(months, true)
		fmt.Println("Sorted Months ", sortedMonths)

		for _, month := range sortedMonths {
			sortedSongs = append(sortedSongs, SortedSongs{
				Year:  year,
				Month: month,
				Songs: months[month],
			})
		}
	}
	fmt.Println("Sorted Songs as array: ", sortedSongs)
	return sortedSongs, nil

}

// Blog list
func SongList(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	// db := database.DBConn
	coll := database.Song_Coll_Conn
	songRepo := SongRepo{MongoCollection: coll}

	// Query the database to find all blog records.
	results, err := songRepo.FindAllSong()

	if err != nil {
		log.Print("Get Operation Failed", err)
		c.Status(400)
		context["msg"] = "Get Operation Failed"
		return c.JSON(context)
	}

	log.Println("Songs: ", results)

	// Add the retrieved blog records to the context map.
	context["song_records"] = results

	c.Status(200)
	return c.JSON(context)
}
func SongListSorted(c *fiber.Ctx) error {

	context := fiber.Map{
		"statusText": "",
		"msg":        "",
	}

	// db := database.DBConn
	coll := database.Song_Coll_Conn
	songRepo := SongRepo{MongoCollection: coll}

	// Query the database to find all blog records.
	results, err := songRepo.GroupSongByMonth()

	if err != nil {
		log.Print("Get Operation Failed", err)
		c.Status(400)
		context["msg"] = "Get Operation Failed"
		return c.JSON(context)
	}

	log.Println("Songs: ", results)

	// Add the retrieved blog records to the context map.
	context["song_records"] = results

	c.Status(200)
	return c.JSON(context)
}

func SongCreate(c *fiber.Ctx) error {

	// HTTP Response
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Song",
	}

	record := new(model.Song)

	// Get the Parameter from the request
	type SongCreateRequest struct {
		Keyword   string `json:"keyword"`
		AddedDate string `json:"added_date"`
	}

	request := new(SongCreateRequest)

	if err := c.BodyParser(request); err != nil {
		log.Println("Error in parsing request.")
		log.Println(request)
		context["statusText"] = ""
		context["msg"] = "Error in parsing request."

		// Bad Request 400
		c.Status(400)
		return c.JSON(context)
	}

	// Create the record
	coll := database.Song_Coll_Conn
	songRepo := SongRepo{MongoCollection: coll}

	// Obtain the image URL
	cmd := exec.Command("python3", "python/cover_image.py", request.Keyword)

	stderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Python Song Scraper Error Occurred.")
		log.Println(string(stderr))
		log.Fatal(err)
	}

	log.Println(string(stderr))
	err = json.Unmarshal(stderr, &record)
	if err != nil {
		log.Println("Converting Scraper Result to JSON file Error Occurer.")
		log.Fatal(err)
	}

	record.AddedDate = request.AddedDate
	fmt.Println(*record)

	// Download the image
	resp, err := http.Get((*record).CoverURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Upload image to the grid.fs bucket
	uploadStream, err := database.SongCover_Coll_Conn.OpenUploadStream(record.Title + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(imgData)
	if err != nil {
		log.Fatal(err)
	}

	// Get the distinct key of the image
	fileID := uploadStream.FileID.(primitive.ObjectID)
	record.CoverImage = fileID

	_, err = songRepo.InsertSong(record)

	if err != nil {
		log.Println("Error in saving data.")
		context["statusText"] = ""
		context["msg"] = "Error in saving data."
		return c.JSON(context)
	}

	context["msg"] = "Record is saved successully."
	context["data"] = record

	c.Status(201)
	return c.JSON(context)
}

// Image Handler
func CoverImageHandler(c *fiber.Ctx) error {
	imageID := c.Params("id")

	// change ImageID to ObjectID
	fileID, err := primitive.ObjectIDFromHex(imageID)
	if err != nil {
		return c.Status(400).SendString("Invalid image ID")
	}

	var buf bytes.Buffer
	_, err = database.SongCover_Coll_Conn.DownloadToStream(fileID, &buf)
	if err != nil {
		log.Println("Failed to download image:", err)
		return c.Status(404).SendString("Image not found")
	}

	c.Set("Content-Type", "image/jpeg")
	return c.Send(buf.Bytes())
}
