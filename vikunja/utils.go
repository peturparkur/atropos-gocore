// Package vikunja contains the types that are used in the Vikunja API
package vikunja

import (
	"fmt"
	"time"
)

// GetLatestComment returns the latest comment from a list of comments
func GetLatestComment(comments []Comment) (Comment, error) {
	if len(comments) == 0 {
		return Comment{}, fmt.Errorf("no comments available")
	}

	latestComment := comments[0]
	latestTime, err := time.Parse(time.RFC3339, latestComment.Updated)
	if err != nil {
		return Comment{}, err
	}

	for _, comment := range comments[1:] {
		currentTime, err := time.Parse(time.RFC3339, comment.Updated)
		if err != nil {
			return Comment{}, err
		}
		if currentTime.After(latestTime) {
			latestTime = currentTime
			latestComment = comment
		}
	}

	return latestComment, nil
}
