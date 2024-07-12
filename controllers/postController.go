package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"tutorial1.go.emp10.com/initializers"
	"tutorial1.go.emp10.com/models"
	"tutorial1.go.emp10.com/utils"
)

func PostsIndex(c *gin.Context) {
	containedTitle := c.Query("title")
	containedBody := c.Query("body")

	var posts []struct {
		ID    uint
		Title string
		Body  string
	}

	query := initializers.DB.Model(&models.Post{})

	if containedTitle != "" {
		query = query.Where("title LIKE ?", "%"+containedTitle+"%")
	}

	if containedBody != "" {
		query = query.Where("body LIKE ?", "%"+containedBody+"%")
	}

	query.Find(&posts)

	c.JSON(
		http.StatusOK,
		utils.ResponseSuccess(
			"Posts successfully retrieved",
			map[string]any{
				"posts": posts,
			},
		),
	)
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post struct {
		ID    uint
		Title string
		Body  string
	}

	result := initializers.DB.Model(&models.Post{}).First(&post, id)

	if result.Error != nil {
		c.JSON(
			http.StatusNotFound,
			utils.ResponseError("Post failed to be retrieved. Post ID was not found"),
		)

		return
	}

	c.JSON(
		http.StatusOK,
		utils.ResponseSuccess(
			"Post successfully retrieved",
			map[string]any{
				"post": post,
			},
		),
	)
}

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	if body.Title == "" {
		c.JSON(
			http.StatusBadRequest,
			utils.ResponseError("Post failed to be created. Title field cannot be empty"),
		)

		return
	}

	if body.Body == "" {
		c.JSON(
			http.StatusBadRequest,
			utils.ResponseError("Post failed to be created. Body field cannot be empty"),
		)

		return
	}

	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(
			http.StatusInternalServerError,
			utils.ResponseError("Post failed to be created. Internal Server Error"),
		)

		return
	}

	c.JSON(
		http.StatusOK,
		utils.ResponseSuccess(
			"Post successfully created",
			map[string]any{
				"postId": post.ID,
			},
		),
	)
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	if body.Title == "" && body.Body == "" {
		c.JSON(
			http.StatusBadRequest,
			utils.ResponseError("Post failed to be updated. Title or body field both cannot be empty"),
		)

		return
	}

	var post models.Post

	result := initializers.DB.Model(&models.Post{}).First(&post, id)

	if result.Error != nil {
		c.JSON(
			http.StatusNotFound,
			utils.ResponseError("Post failed to be updated. Post ID was not found"),
		)

		return
	}

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(
		http.StatusOK,
		utils.ResponseSuccess(
			"Post successfully updated",
			nil,
		),
	)
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	result := initializers.DB.Model(&models.Post{}).First(&post, id)

	if result.Error != nil {
		c.JSON(
			http.StatusNotFound,
			utils.ResponseError("Post failed to be deleted. Post ID was not found"),
		)

		return
	}

	initializers.DB.Delete(&post)

	c.JSON(
		http.StatusOK,
		utils.ResponseSuccess(
			"Post successfully deleted",
			nil,
		),
	)
}
