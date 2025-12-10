package solutions

import (
	"encoding/json"

	S "github.com/espinosajuanma/slingr-go"
)

type Solutions struct {
	App  *S.App
	User *User
}

type User struct {
	Name     string
	Id       string
	Email    string
	Password string
	Country  string
}

type ManyResponse[T any] struct {
	Items []T `json:"items"`
}

type Relationship struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type CurrentUser struct {
	Localization struct {
		Timezone     string `json:"timezone"`
		TimezoneMode string `json:"timezoneMode"`
		Lang         string `json:"lang"`
	} `json:"localization"`
	LastName          string `json:"lastName"`
	IdentityProviders []struct {
		Id         string      `json:"id"`
		Name       string      `json:"name"`
		ExternalId interface{} `json:"externalId"`
		Label      string      `json:"label"`
	} `json:"identityProviders"`
	FullName string `json:"fullName"`
	Groups   []struct {
		Id       string `json:"id"`
		Primary  bool   `json:"primary"`
		External bool   `json:"external"`
		Name     string `json:"name"`
		Label    string `json:"label"`
	} `json:"groups"`
	Version                    int         `json:"version"`
	FirstName                  string      `json:"firstName"`
	PhoneNumber                interface{} `json:"phoneNumber"`
	TwoFactorAuthenticationKey interface{} `json:"twoFactorAuthenticationKey"`
	Permissions                struct {
		Developer       bool `json:"developer"`
		UsersManagement bool `json:"usersManagement"`
	} `json:"permissions"`
	TwoFactorAuthentication bool          `json:"twoFactorAuthentication"`
	Id                      string        `json:"id"`
	Integrations            []interface{} `json:"integrations"`
	Email                   string        `json:"email"`
	Status                  string        `json:"status"`
}

func NewSolutions() *Solutions {
	return &Solutions{
		App:  S.NewApp("solutions", S.EnvProd),
		User: &User{},
	}
}

func (s *Solutions) SetToken(token string) {
	s.App.Token = token
}

func (s *Solutions) GetCurrentUser() (CurrentUser, error) {
	var current CurrentUser
	r, err := s.App.Get("/users/current", nil)
	if err != nil {
		return current, err
	}
	err = json.Unmarshal(r, &current)
	if err != nil {
		return current, err
	}
	return current, nil
}

func (s *Solutions) Login() (*S.LoginResponse, error) {
	r, err := s.App.Login(s.User.Email, s.User.Password)
	if err != nil {
		return r, err
	}
	return r, nil
}
