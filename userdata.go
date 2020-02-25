package ftapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
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

type CampusUser struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}

func GetAuthorizedUserData(accessToken string) UserData {
	bytes := DoFTRequest("/v2/me", accessToken)
	var userData UserData
	json.Unmarshal(bytes, &userData)
	return userData
}

func RequestUserData(login string, accessToken string) UserData {
	bytes := DoFTRequest("/v2/users/"+login, accessToken)
	var userData UserData
	json.Unmarshal(bytes, &userData)
	return userData
}

func RequestCampusUsers(campusID int, accessToken string) []CampusUser {
	campusUsers := make([]CampusUser, 0)
	i := 1
	for {
		bytes := DoFTRequest("/v2/campus/"+strconv.Itoa(campusID)+"/users?page="+strconv.Itoa(i), accessToken)
		if len(bytes) <= 2 {
			break
		}
		campusUserPage := make([]CampusUser, 0)
		err := json.Unmarshal(bytes, &campusUserPage)
		if err != nil {
			fmt.Println(err)
		}
		for i := range campusUserPage {
			campusUsers = append(campusUsers, campusUserPage[i])
		}
		i++
	}
	return campusUsers
}

func RequestAllCampusUsersData(campusID int, accessToken string) []UserData {
	campusUsers := RequestCampusUsers(campusID, accessToken)
	userData := make([]UserData, 0)
	for _, campusUser := range campusUsers {
		fmt.Print(".")
		u := RequestUserData(campusUser.Login, accessToken)
		userData = append(userData, u)
	}
	fmt.Print("\n")
	return userData
}

func LoadUserData(filename string) ([]UserData, error) {
	userData := make([]UserData, 0)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &userData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return userData, nil
}

func SaveUserData(filename string, userData []UserData) {
	jsonString, err := json.Marshal(userData)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(filename, jsonString, 0644)
}
