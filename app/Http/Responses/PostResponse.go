package Responses

import (
	"mvcGolang/app/Helpers"
	"mvcGolang/app/Models/Posts"
	"strconv"
)

type PostResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func FormatPost(post Posts.Post) PostResponse {
	formatter := PostResponse{
		ID:    Helpers.Encrypt(strconv.Itoa(post.ID)),
		Title: post.Title,
	}

	return formatter
}

func FormatPosts(posts []Posts.Post) []PostResponse {
	postsFormatter := []PostResponse{}

	for _, post := range posts {
		postFormatter := FormatPost(post)
		postsFormatter = append(postsFormatter, postFormatter)
	}

	return postsFormatter
}
