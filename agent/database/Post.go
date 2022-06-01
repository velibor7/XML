package database

import "github.com/XML/agent/model"

func GetPosts(company_id string) []model.Post {
	db := DBConn
	var posts []model.Post
	db.Find(&posts, "company_id = ?", company_id)
	return posts
}

func CreatePost(post *model.Post) error {
	db := DBConn
	result := db.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func LoadPosts(company_id string, comp *model.Company) {
	posts := GetPosts(company_id)
	for _, post := range posts {
		comp.Posts = append(comp.Posts, post)
	}
}
