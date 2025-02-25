/*
Sumo Logic API

Testing DashboardManagementAPIService

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

func Test_sumologic_DashboardManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DashboardManagementAPIService CreateDashboard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DashboardManagementAPI.CreateDashboard(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService CreateScheduleReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DashboardManagementAPI.CreateScheduleReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService DeleteDashboard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.DashboardManagementAPI.DeleteDashboard(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService DeleteReportSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scheduleId string

		httpRes, err := apiClient.DashboardManagementAPI.DeleteReportSchedule(context.Background(), scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService GenerateDashboardReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DashboardManagementAPI.GenerateDashboardReport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService GetAsyncReportGenerationResult", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.DashboardManagementAPI.GetAsyncReportGenerationResult(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService GetAsyncReportGenerationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.DashboardManagementAPI.GetAsyncReportGenerationStatus(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService GetDashboard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DashboardManagementAPI.GetDashboard(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService GetDashboardMigrationResult", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.DashboardManagementAPI.GetDashboardMigrationResult(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService GetDashboardMigrationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var jobId string

		resp, httpRes, err := apiClient.DashboardManagementAPI.GetDashboardMigrationStatus(context.Background(), jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService GetReportSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scheduleId string

		resp, httpRes, err := apiClient.DashboardManagementAPI.GetReportSchedule(context.Background(), scheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService ListDashboards", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DashboardManagementAPI.ListDashboards(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService ListReportSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DashboardManagementAPI.ListReportSchedules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService MigrateReportToDashboard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DashboardManagementAPI.MigrateReportToDashboard(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService PreviewMigrateReportToDashboard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DashboardManagementAPI.PreviewMigrateReportToDashboard(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService UpdateDashboard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.DashboardManagementAPI.UpdateDashboard(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DashboardManagementAPIService UpdateReportSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scheduleId string

		resp, httpRes, err := apiClient.DashboardManagementAPI.UpdateReportSchedule(context.Background(), scheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
