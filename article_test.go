package guardian

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestArticleCast(t *testing.T) {
	ga := GuardianArticle{
		Id_: "testId",
		WebTitle: "testTitle",
		Blocks: ArticleBlocks{
			Body: []*BlockType{
				{
					BodyTextSummary: "testBody",
				},
			},
		},
	}
	a := Article(ga)
	assert.Equal(t, a.Id(), "testId")
	assert.Equal(t, a.Title(), "testTitle")
	assert.Equal(t, a.Body(), "testBody")
}

