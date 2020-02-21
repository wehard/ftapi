package ftapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UserData struct {
	Achievements []struct {
		Description  string      `json:"description"`
		ID           int         `json:"id"`
		Image        string      `json:"image"`
		Kind         string      `json:"kind"`
		Name         string      `json:"name"`
		NbrOfSuccess interface{} `json:"nbr_of_success"`
		Tier         string      `json:"tier"`
		UsersURL     string      `json:"users_url"`
		Visible      bool        `json:"visible"`
	} `json:"achievements"`
	Campus []struct {
		Address  string `json:"address"`
		City     string `json:"city"`
		Country  string `json:"country"`
		Facebook string `json:"facebook"`
		ID       int    `json:"id"`
		Language struct {
			CreatedAt  string `json:"created_at"`
			ID         int    `json:"id"`
			Identifier string `json:"identifier"`
			Name       string `json:"name"`
			UpdatedAt  string `json:"updated_at"`
		} `json:"language"`
		Name        string `json:"name"`
		TimeZone    string `json:"time_zone"`
		Twitter     string `json:"twitter"`
		UsersCount  int    `json:"users_count"`
		VogsphereID int    `json:"vogsphere_id"`
		Website     string `json:"website"`
		Zip         string `json:"zip"`
	} `json:"campus"`
	CampusUsers []struct {
		CampusID  int  `json:"campus_id"`
		ID        int  `json:"id"`
		IsPrimary bool `json:"is_primary"`
		UserID    int  `json:"user_id"`
	} `json:"campus_users"`
	CorrectionPoint int `json:"correction_point"`
	CursusUsers     []struct {
		BeginAt      string      `json:"begin_at"`
		BlackholedAt interface{} `json:"blackholed_at"`
		Cursus       struct {
			CreatedAt string `json:"created_at"`
			ID        int    `json:"id"`
			Name      string `json:"name"`
			Slug      string `json:"slug"`
		} `json:"cursus"`
		CursusID     int         `json:"cursus_id"`
		EndAt        interface{} `json:"end_at"`
		Grade        string      `json:"grade"`
		HasCoalition bool        `json:"has_coalition"`
		ID           int         `json:"id"`
		Level        float64     `json:"level"`
		Skills       []struct {
			ID    int     `json:"id"`
			Level float64 `json:"level"`
			Name  string  `json:"name"`
		} `json:"skills"`
		User struct {
			ID    int    `json:"id"`
			Login string `json:"login"`
			URL   string `json:"url"`
		} `json:"user"`
	} `json:"cursus_users"`
	Displayname     string        `json:"displayname"`
	Email           string        `json:"email"`
	ExpertisesUsers []interface{} `json:"expertises_users"`
	FirstName       string        `json:"first_name"`
	Groups          []interface{} `json:"groups"`
	ID              int           `json:"id"`
	ImageURL        string        `json:"image_url"`
	LanguagesUsers  []struct {
		CreatedAt  string `json:"created_at"`
		ID         int    `json:"id"`
		LanguageID int    `json:"language_id"`
		Position   int    `json:"position"`
		UserID     int    `json:"user_id"`
	} `json:"languages_users"`
	LastName      string        `json:"last_name"`
	Location      string        `json:"location"`
	Login         string        `json:"login"`
	Partnerships  []interface{} `json:"partnerships"`
	Patroned      []interface{} `json:"patroned"`
	Patroning     []interface{} `json:"patroning"`
	Phone         string        `json:"phone"`
	PoolMonth     string        `json:"pool_month"`
	PoolYear      string        `json:"pool_year"`
	ProjectsUsers []struct {
		CurrentTeamID int    `json:"current_team_id"`
		CursusIds     []int  `json:"cursus_ids"`
		FinalMark     int    `json:"final_mark"`
		ID            int    `json:"id"`
		Marked        bool   `json:"marked"`
		MarkedAt      string `json:"marked_at"`
		Occurrence    int    `json:"occurrence"`
		Project       struct {
			ID       int         `json:"id"`
			Name     string      `json:"name"`
			ParentID interface{} `json:"parent_id"`
			Slug     string      `json:"slug"`
		} `json:"project"`
		RetriableAt string `json:"retriable_at"`
		Status      string `json:"status"`
		Validated_  bool   `json:"validated?"`
	} `json:"projects_users"`
	Staff_      bool          `json:"staff?"`
	Titles      []interface{} `json:"titles"`
	TitlesUsers []interface{} `json:"titles_users"`
	URL         string        `json:"url"`
	Wallet      int           `json:"wallet"`
}

func DoFTRequest(endpoint string, accessToken string) []byte {
	var data struct {
		Token string `json:"access_token"`
	}
	data.Token = accessToken
	b, err := json.Marshal(data)
	req, err := http.NewRequest("GET", "https://api.intra.42.fr"+endpoint, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	return (body)
}