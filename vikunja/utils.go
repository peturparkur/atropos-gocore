// Package vikunja contains the types that are used in the Vikunja API
package vikunja

import (
	"fmt"
	"time"

	"github.com/atropos112/gocore/utils"
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

// LabelsWithGivenTitles returns a list of labels with the given titles
// If the title of a label is not found, an error is returned, it is expected that you only provide valid titles
func LabelsWithGivenTitles(labels []Label, titles []string) ([]Label, error) {
	labelMap := map[string]Label{}
	for _, label := range labels {
		labelMap[label.Title] = label
	}

	result := []Label{}
	for _, title := range titles {
		if label, ok := labelMap[title]; ok {
			result = append(result, label)
		} else {
			return nil, &utils.DeveloperError{
				Message: "Label with title " + title + " not found",
			}
		}
	}

	return result, nil
}
