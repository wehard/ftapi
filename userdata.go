package ftapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
)

const (
	FortyTwo = 1
	Piscine  = 4
)

type Skill struct {
	ID    int     `json:"id"`
	Level float64 `json:"level"`
	Name  string  `json:"name"`
}

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
		Skills       []Skill     `json:"skills"`
		User         struct {
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
			fmt.Println("request campus user: ", err)
		}
		for i := range campusUserPage {
			campusUsers = append(campusUsers, campusUserPage[i])
		}
		fmt.Print(".")
		i++
	}
	fmt.Print("\n")
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

func LoadAccessToken(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	token := string(bytes)
	fmt.Println("token loaded", token)
	return token
}

func SaveAccessToken(filename, token string) {
	ioutil.WriteFile(filename, []byte(token), 0644)
}

func GetUserDataByLevel(level int, cursusID int, allUserData []UserData) []UserData {
	userData := make([]UserData, 0)
	for _, u := range allUserData {
		if len(u.CursusUsers) == 0 {
			continue
		}
		for i, _ := range u.CursusUsers {
			if u.CursusUsers[i].CursusID == cursusID && u.CursusUsers[i].Level < float64(level+1) && u.CursusUsers[i].Level >= float64(level) {
				userData = append(userData, u)
			}
		}
	}
	return userData
}

func GetUserDataByCursus(cursusID int, allUserData []UserData) []UserData {
	userData := make([]UserData, 0)
	for _, u := range allUserData {
		if len(u.CursusUsers) == 0 {
			continue
		}
		for i, _ := range u.CursusUsers {
			if u.CursusUsers[i].CursusID == cursusID && u.CursusUsers[i].Level > 0 {
				userData = append(userData, u)
			}
		}
	}
	return userData
}

func GetUserData(compareFunc func(int, float64) bool, allUserData []UserData) []UserData {
	userData := make([]UserData, 0)
	for _, u := range allUserData {
		if len(u.CursusUsers) == 0 {
			continue
		}
		for i, _ := range u.CursusUsers {
			if compareFunc(u.CursusUsers[i].CursusID, u.CursusUsers[i].Level) {
				userData = append(userData, u)
			}
		}
	}
	return userData
}

func GetUserDataByLogin(login string, allUserData []UserData) UserData {
	for _, u := range allUserData {
		if u.Login == login {
			return u
		}
	}
	fmt.Println("user not found")
	return UserData{}
}

func GetUserLevel(user UserData, cursusID int) float64 {
	for i, _ := range user.CursusUsers {
		if user.CursusUsers[i].CursusID == cursusID {
			return user.CursusUsers[i].Level
		}
	}
	return -1.0
}

func GetUserSkills(user UserData, cursusID int) []Skill {
	skills := make([]Skill, 0)
	for i, _ := range user.CursusUsers {
		if user.CursusUsers[i].CursusID == cursusID {
			for s, _ := range user.CursusUsers[i].Skills {
				skills = append(skills, user.CursusUsers[i].Skills[s])
			}
		}
	}
	return skills
}

func GetRandomUserSkill(user UserData, cursusID int) Skill {
	skills := GetUserSkills(user, cursusID)
	return skills[rand.Intn(len(skills)-1)]
}
