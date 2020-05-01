package controllers

//import (
//	"net/http"
//
//	u "github.com/fifciu/what-can-i-do/server/utils"
//	"github.com/fifciu/what-can-i-do/server/models"
//	"github.com/gorilla/context"
//)
//
//func GetMineIdeas(w http.ResponseWriter, r *http.Request) {
//	claims := context.Get(r, "CurrentUser").(*models.Claims)
//	ideas := models.GetUserIdeas(claims.ID, []models.IdeasMapper{models.MapperAddProblemsName})
//	response := u.Status(true)
//	response["ideas"] = ideas
//	u.RespondWithCode(w, response, http.StatusOK)
//}
