package utils

import (
	"net/http"
	"github.com/fifciu/what-can-i-do/server/models"
	"strconv"
	"strings"
)

const CACHE_TAGS_HEADER = "X-Cache-Tags"

func PopulateCacheTags(w http.ResponseWriter, data map[string]interface{}) {
	tags := map[string]bool{}
	_, dataProblem := data["problem"]
	if (dataProblem) {
		populateTagsProblem(tags, data["problem"])
	}
	_, dataProblems := data["problems"]
	if (dataProblems) {
		populateTagsProblems(tags, data["problems"])
	}

	tagsAsString := ""
	for tag, _ := range tags {
		tagsAsString += tag + " "
	}
	w.Header().Add(CACHE_TAGS_HEADER, strings.Trim(tagsAsString, " "))
}

func populateTagsProblem(tags map[string]bool, problem interface{}) {
	problemWithType := problem.(*models.Problem)
	if problemWithType.ID > 0 {
		tags["problem:"+strconv.Itoa(int(problemWithType.ID))] = true
	}
	if problemWithType.Ideas != nil {
		populateTagsIdeas(tags, problemWithType.Ideas)
	}
}

func populateTagsProblems(tags map[string]bool, problems interface{}) {
	problemsWithType := problems.([]*models.Problem)
	for _, problem := range problemsWithType {
		populateTagsProblem(tags, problem)
	}
}

func populateTagsIdea(tags map[string]bool, idea interface{}) {
	ideaWithType := idea.(*models.Idea)
	if ideaWithType.ID > 0 {
		tags["idea:"+strconv.Itoa(int(ideaWithType.ID))] = true
	}
}

func populateTagsIdeas(tags map[string]bool, ideas interface{}) {
	ideasWithType := ideas.([]*models.Idea)
	for _, idea := range ideasWithType {
		populateTagsIdea(tags, idea)
	}
}