/*
Sumo Logic API

Testing ArchiveManagementAPIService

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

func Test_sumologic_ArchiveManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ArchiveManagementAPIService CreateArchiveJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.ArchiveManagementAPI.CreateArchiveJob(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArchiveManagementAPIService DeleteArchiveJob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string
		var id string

		httpRes, err := apiClient.ArchiveManagementAPI.DeleteArchiveJob(context.Background(), sourceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArchiveManagementAPIService ListArchiveJobsBySourceId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var sourceId string

		resp, httpRes, err := apiClient.ArchiveManagementAPI.ListArchiveJobsBySourceId(context.Background(), sourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArchiveManagementAPIService ListArchiveJobsCountPerSource", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ArchiveManagementAPI.ListArchiveJobsCountPerSource(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
