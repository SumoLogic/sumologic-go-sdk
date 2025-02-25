/*
Sumo Logic API

Testing ExtractionRuleManagementAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package sumologic

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/SumoLogic/sumologic-go-sdk"
)

func Test_sumologic_ExtractionRuleManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExtractionRuleManagementAPIService CreateExtractionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExtractionRuleManagementAPI.CreateExtractionRule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtractionRuleManagementAPIService DeleteExtractionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ExtractionRuleManagementAPI.DeleteExtractionRule(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtractionRuleManagementAPIService GetExtractionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExtractionRuleManagementAPI.GetExtractionRule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtractionRuleManagementAPIService ListExtractionRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ExtractionRuleManagementAPI.ListExtractionRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtractionRuleManagementAPIService UpdateExtractionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ExtractionRuleManagementAPI.UpdateExtractionRule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
