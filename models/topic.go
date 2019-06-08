package models

import (
	_ "github.com/jinzhu/gorm"
)

func init() {
}

type (
	Topics struct {
		Data []Topic `json:"data"`

		Meta struct {
			Pagination struct {
				Count       int `json:"count"`
				CurrentPage int `json:"current_page"`
				Links       struct {
					Next     string      `json:"next"`
					Previous interface{} `json:"previous"`
				} `json:"links"`
				PerPage    int `json:"per_page"`
				Total      int `json:"total"`
				TotalPages int `json:"total_pages"`
			} `json:"pagination"`
		} `json:"meta"`
	}

	Replies struct {
		Data []ReplyUser `json:"data"`
	}

	TopicReplies struct {
		Replies
		Meta `json:"meta"`
	}

	ReplyUser struct {
		Reply
		User `json:"user"`
	}
)
