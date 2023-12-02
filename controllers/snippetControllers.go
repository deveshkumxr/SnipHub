package controllers

import (
	"context"
	"strconv"

	"github.com/deveshkumxr/SnipHub/db"
	"github.com/deveshkumxr/SnipHub/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSnippets(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		return c.SendString(err.Error())
	}

	perPage, err := strconv.Atoi(c.Query("perPage", "10"))
	if err != nil || perPage < 1 {
		return c.SendString(err.Error())
	}

	skip := (page - 1) * perPage

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(perPage))

	cursor, err := db.Collection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return c.SendString(err.Error())
	}
	defer cursor.Close(context.Background())

	var snippets []models.Snippet
	if err := cursor.All(context.Background(), &snippets); err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"pageNum": page,
		"perPage": perPage,
		"data":    snippets,
	})
}

func CreateSnippet(c *fiber.Ctx) error {
	var snippet models.Snippet
	if err := c.BodyParser(&snippet); err != nil {
		return c.SendString(err.Error())
	}

	snippet.ID = primitive.NewObjectID()

	_, err := db.Collection.InsertOne(context.Background(), snippet)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"newUser": snippet,
	})
}

func UpdateSnippet(c *fiber.Ctx) error {
	id := c.Params("snippet_id")
	snippetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.SendString(err.Error())
	}

	var newSnippet map[string]interface{}
	if err := c.BodyParser(&newSnippet); err != nil {
		return c.SendString(err.Error())
	}

	filter := bson.M{"_id": snippetID}
	update := bson.M{"$set": newSnippet}

	if err := db.Collection.FindOneAndUpdate(context.Background(), filter, update).Err(); err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"updatedSnippetID" : id,
		"message" : "Snippet updated successfully!!",
	})
}

func DeleteSnippet(c *fiber.Ctx) error {
	id := c.Params("snippet_id")
	snippetID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.SendString(err.Error())
	}

	filter := bson.M{"_id": snippetID}

	if err := db.Collection.FindOneAndDelete(context.Background(), filter).Err(); err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(fiber.Map{
		"deletedSnippetID" : id,
		"message" : "Snippet deleted successfully!!",
	})
}
