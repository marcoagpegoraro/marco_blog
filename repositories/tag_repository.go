package repositories

import (
	"github.com/marcoagpegoraro/marco_blog/initializers"
	"github.com/marcoagpegoraro/marco_blog/models"
)

var TagRepository = TagRepositoryStruct{}

type TagRepositoryStruct struct {
}

func (service TagRepositoryStruct) GetAllTags(isAuth bool) []models.Tag {

	sqlQuery := `
    SELECT distinct t.name
    FROM tags t
    LEFT OUTER JOIN posts_tags pt
    ON t.name = pt.tag_name
    LEFT OUTER JOIN posts p
    ON pt.post_id  = p.id  
    `

	var tags []models.Tag

	if !isAuth {
		sqlQuery += ` WHERE  p.is_draft = false`
	}

	sqlQuery += ` ORDER BY t.name`

	initializers.DB.Raw(sqlQuery).Scan(&tags)

	return tags
}

func (service TagRepositoryStruct) DeleteTagsFromPost(post *models.Post, tagsToBeDeleted []models.Tag) {
	initializers.DB.Model(post).Association("Tags").Delete(tagsToBeDeleted)
}
