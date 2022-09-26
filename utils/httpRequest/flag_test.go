package httprequest

import (
	"net/http"
	"testing"

	"github.com/flagship-io/flagship/models"
	"github.com/flagship-io/flagship/utils"
	"github.com/jarcoal/httpmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestHTTPGetFlag(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testFlag := models.Flag{
		ID:          "testFlagID",
		Name:        "testFlagName",
		Type:        "string",
		Description: "testFlagDescription",
		Source:      "manual",
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+testFlag.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testFlag)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPGetFlag("testFlagID")

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testFlagID", respBody.ID)
	assert.Equal(t, "testFlagName", respBody.Name)
}

func TestHTTPListFlag(t *testing.T) {

	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testFlagList := []models.Flag{
		{
			ID:          "testFlagID",
			Name:        "testFlagName",
			Type:        "string",
			Description: "testFlagDescription",
			Source:      "manual",
		},
		{
			ID:          "testFlagID1",
			Name:        "testFlagName1",
			Type:        "string",
			Description: "testFlagDescription1",
			Source:      "manual",
		},
	}

	resp := HTTPListResponse[models.Flag]{
		Items:             testFlagList,
		CurrentItemsCount: 2,
		CurrentPage:       1,
		TotalCount:        2,
		ItemsPerPage:      10,
		LastPage:          1,
	}

	httpmock.RegisterResponder("GET", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, resp)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPListFlag()

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, "testFlagID", respBody[0].ID)
	assert.Equal(t, "testFlagName", respBody[0].Name)

	assert.Equal(t, "testFlagID1", respBody[1].ID)
	assert.Equal(t, "testFlagName1", respBody[1].Name)
}

func TestHTTPCreateFlag(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testFlag := models.Flag{
		ID:          "testFlagID",
		Name:        "testFlagName",
		Type:        "string",
		Description: "testFlagDescription",
		Source:      "manual",
	}

	data := "{\"name\":\"testFlagName\", \"type\":\"string\", \"description\":\"testFlagDescription\", \"source\":\"manual\"}"

	httpmock.RegisterResponder("POST", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags",
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testFlag)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPCreateFlag(data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testFlagID\",\"name\":\"testFlagName\",\"type\":\"string\",\"description\":\"testFlagDescription\",\"source\":\"manual\"}"), respBody)
}

func TestHTTPEditFlag(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testFlag := models.Flag{
		ID:          "testFlagID",
		Name:        "testFlagName1",
		Type:        "string",
		Description: "testFlagDescription1",
		Source:      "manual",
	}

	data := "{\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"manual\"}"

	httpmock.RegisterResponder("PATCH", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+testFlag.ID,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, testFlag)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)

	respBody, err := HTTPEditFlag(testFlag.ID, data)

	assert.NotNil(t, respBody)
	assert.Nil(t, err)

	assert.Equal(t, []byte("{\"id\":\"testFlagID\",\"name\":\"testFlagName1\",\"type\":\"string\",\"description\":\"testFlagDescription1\",\"source\":\"manual\"}"), respBody)
}

func TestHTTPDeleteFlag(t *testing.T) {
	ViperNotSet(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	testFlag := models.Flag{
		ID:          "testFlagID",
		Name:        "testFlagName",
		Type:        "string",
		Description: "testFlagDescription",
		Source:      "manual",
	}

	httpmock.RegisterResponder("DELETE", utils.Host+"/v1/accounts/"+viper.GetString("account_id")+"/flags/"+testFlag.ID,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(204, ""), nil

		},
	)

	err := HTTPDeleteFlag("testFlagID")

	assert.Nil(t, err)
}