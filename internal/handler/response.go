package handler

import (
	"encoding/json"
	"loaders/internal/models"
	"net/http"
)

type Response struct {
	Result 		string 			`json:"result,omitempty"`
	HttpStatus 	int 			`json:"status_code,omitempty"`
	ID			int64			`json:"id,omitempty"`
	Username	string			`json:"username,omitempty"`
	Role		string			`json:"role,omitempty"`
	Balance		int				`json:"balance,omitempty"`
	Loaders		[]models.Loader	`json:"loaders,omitempty"`
	Tasks	  	[]models.Task	`json:"tasks,omitempty"`
	Token		string			`json:"token,omitempty"`
	Salary		int				`json:"salary,omitempry"`
	Weight		int				`json:"weight,omitempty"`
	Fatigue		int				`json:"fatigue,omitempty"`
	Drunk		bool			`json:"drunk,omitempty"`
}

func renderResponse(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
