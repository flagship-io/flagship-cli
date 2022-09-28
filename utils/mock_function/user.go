package mockfunction

import (
	"net/http"
	"net/url"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/flagship-io/flagship/utils/config"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
)

func APIUser() {
	config.SetViper()

	email := "example@abtasty.com"

	testUserList := []models.User{
		{
			Email: "example@abtasty.com",
			Role:  "ADMIN",
		},
		{
			Email: "example1@abtasty.com",
			Role:  "VIEWER",
		},
	}

	resp := utils.HTTPListResponse[models.User]{
		Items:             testUserList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("PUT", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(204, "")
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/account_environments/"+viper.GetString("account_environment_id")+"/users?emails[]="+url.QueryEscape(email),
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(204, "")
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}