package repositories

import (
	"github.com/marcoagpegoraro/marco_blog/enum"
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

var PostRepository = PostRepositoryStruct{}

type PostRepositoryStruct struct {
}

func (service PostRepositoryStruct) GetPosts(isAuth bool, currentPage int, pageSize int, language string, tag string, showDrafts bool) []models.Post {

	var posts []models.Post
	dbQuery := initializers.DB.Preload("Tags")

	if isAuth {
		dbQuery.Where("is_draft = ?", showDrafts)
		if showDrafts {
			dbQuery.Order("created_at desc")
		} else {
			dbQuery.Order("publicated_at desc")
		}
	} else {
		dbQuery.Where("is_draft = ?", "false")
		dbQuery.Order("publicated_at desc")
	}

	if language != "All" {
		var languageKey uint8
		for k, v := range enum.LanguageEnumValues() {
			if v == language {
				languageKey = k
				break
			}
		}
		if languageKey != 0 {
			dbQuery.Where("language = ?", languageKey)
		}
	}

	dbQuery.Limit(pageSize).Offset(pageSize * (currentPage - 1))
	dbQuery.Find(&posts)

	if tag != "" { //i know... i know... this is way to ugly, but it is good enough for now
		var filteredPostsByTag []models.Post
		for _, post := range posts {
			for _, postTag := range post.Tags {
				if postTag.Name == tag {
					filteredPostsByTag = append(filteredPostsByTag, post)
					break
				}
			}
		}
		return filteredPostsByTag
	}

	return posts
}

func (service PostRepositoryStruct) GetTotalPostsCount(isAuth bool, showDrafts bool) int64 {
	var count int64
	dbQuery := initializers.DB.Model(&models.Post{})

	if !isAuth {
		dbQuery.Where("is_draft = ?", "false")
	} else {
		dbQuery.Where("is_draft = ?", showDrafts)
	}

	dbQuery.Count(&count)

	return count
}
