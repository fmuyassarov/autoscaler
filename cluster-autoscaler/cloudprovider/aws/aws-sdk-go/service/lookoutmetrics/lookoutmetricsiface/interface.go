// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lookoutmetricsiface provides an interface to enable mocking the Amazon Lookout for Metrics service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package lookoutmetricsiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/lookoutmetrics"
)

// LookoutMetricsAPI provides an interface to enable mocking the
// lookoutmetrics.LookoutMetrics service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Lookout for Metrics.
//	func myFunc(svc lookoutmetricsiface.LookoutMetricsAPI) bool {
//	    // Make svc.ActivateAnomalyDetector request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := lookoutmetrics.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockLookoutMetricsClient struct {
//	    lookoutmetricsiface.LookoutMetricsAPI
//	}
//	func (m *mockLookoutMetricsClient) ActivateAnomalyDetector(input *lookoutmetrics.ActivateAnomalyDetectorInput) (*lookoutmetrics.ActivateAnomalyDetectorOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockLookoutMetricsClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type LookoutMetricsAPI interface {
	ActivateAnomalyDetector(*lookoutmetrics.ActivateAnomalyDetectorInput) (*lookoutmetrics.ActivateAnomalyDetectorOutput, error)
	ActivateAnomalyDetectorWithContext(aws.Context, *lookoutmetrics.ActivateAnomalyDetectorInput, ...request.Option) (*lookoutmetrics.ActivateAnomalyDetectorOutput, error)
	ActivateAnomalyDetectorRequest(*lookoutmetrics.ActivateAnomalyDetectorInput) (*request.Request, *lookoutmetrics.ActivateAnomalyDetectorOutput)

	BackTestAnomalyDetector(*lookoutmetrics.BackTestAnomalyDetectorInput) (*lookoutmetrics.BackTestAnomalyDetectorOutput, error)
	BackTestAnomalyDetectorWithContext(aws.Context, *lookoutmetrics.BackTestAnomalyDetectorInput, ...request.Option) (*lookoutmetrics.BackTestAnomalyDetectorOutput, error)
	BackTestAnomalyDetectorRequest(*lookoutmetrics.BackTestAnomalyDetectorInput) (*request.Request, *lookoutmetrics.BackTestAnomalyDetectorOutput)

	CreateAlert(*lookoutmetrics.CreateAlertInput) (*lookoutmetrics.CreateAlertOutput, error)
	CreateAlertWithContext(aws.Context, *lookoutmetrics.CreateAlertInput, ...request.Option) (*lookoutmetrics.CreateAlertOutput, error)
	CreateAlertRequest(*lookoutmetrics.CreateAlertInput) (*request.Request, *lookoutmetrics.CreateAlertOutput)

	CreateAnomalyDetector(*lookoutmetrics.CreateAnomalyDetectorInput) (*lookoutmetrics.CreateAnomalyDetectorOutput, error)
	CreateAnomalyDetectorWithContext(aws.Context, *lookoutmetrics.CreateAnomalyDetectorInput, ...request.Option) (*lookoutmetrics.CreateAnomalyDetectorOutput, error)
	CreateAnomalyDetectorRequest(*lookoutmetrics.CreateAnomalyDetectorInput) (*request.Request, *lookoutmetrics.CreateAnomalyDetectorOutput)

	CreateMetricSet(*lookoutmetrics.CreateMetricSetInput) (*lookoutmetrics.CreateMetricSetOutput, error)
	CreateMetricSetWithContext(aws.Context, *lookoutmetrics.CreateMetricSetInput, ...request.Option) (*lookoutmetrics.CreateMetricSetOutput, error)
	CreateMetricSetRequest(*lookoutmetrics.CreateMetricSetInput) (*request.Request, *lookoutmetrics.CreateMetricSetOutput)

	DeactivateAnomalyDetector(*lookoutmetrics.DeactivateAnomalyDetectorInput) (*lookoutmetrics.DeactivateAnomalyDetectorOutput, error)
	DeactivateAnomalyDetectorWithContext(aws.Context, *lookoutmetrics.DeactivateAnomalyDetectorInput, ...request.Option) (*lookoutmetrics.DeactivateAnomalyDetectorOutput, error)
	DeactivateAnomalyDetectorRequest(*lookoutmetrics.DeactivateAnomalyDetectorInput) (*request.Request, *lookoutmetrics.DeactivateAnomalyDetectorOutput)

	DeleteAlert(*lookoutmetrics.DeleteAlertInput) (*lookoutmetrics.DeleteAlertOutput, error)
	DeleteAlertWithContext(aws.Context, *lookoutmetrics.DeleteAlertInput, ...request.Option) (*lookoutmetrics.DeleteAlertOutput, error)
	DeleteAlertRequest(*lookoutmetrics.DeleteAlertInput) (*request.Request, *lookoutmetrics.DeleteAlertOutput)

	DeleteAnomalyDetector(*lookoutmetrics.DeleteAnomalyDetectorInput) (*lookoutmetrics.DeleteAnomalyDetectorOutput, error)
	DeleteAnomalyDetectorWithContext(aws.Context, *lookoutmetrics.DeleteAnomalyDetectorInput, ...request.Option) (*lookoutmetrics.DeleteAnomalyDetectorOutput, error)
	DeleteAnomalyDetectorRequest(*lookoutmetrics.DeleteAnomalyDetectorInput) (*request.Request, *lookoutmetrics.DeleteAnomalyDetectorOutput)

	DescribeAlert(*lookoutmetrics.DescribeAlertInput) (*lookoutmetrics.DescribeAlertOutput, error)
	DescribeAlertWithContext(aws.Context, *lookoutmetrics.DescribeAlertInput, ...request.Option) (*lookoutmetrics.DescribeAlertOutput, error)
	DescribeAlertRequest(*lookoutmetrics.DescribeAlertInput) (*request.Request, *lookoutmetrics.DescribeAlertOutput)

	DescribeAnomalyDetectionExecutions(*lookoutmetrics.DescribeAnomalyDetectionExecutionsInput) (*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, error)
	DescribeAnomalyDetectionExecutionsWithContext(aws.Context, *lookoutmetrics.DescribeAnomalyDetectionExecutionsInput, ...request.Option) (*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, error)
	DescribeAnomalyDetectionExecutionsRequest(*lookoutmetrics.DescribeAnomalyDetectionExecutionsInput) (*request.Request, *lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput)

	DescribeAnomalyDetectionExecutionsPages(*lookoutmetrics.DescribeAnomalyDetectionExecutionsInput, func(*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, bool) bool) error
	DescribeAnomalyDetectionExecutionsPagesWithContext(aws.Context, *lookoutmetrics.DescribeAnomalyDetectionExecutionsInput, func(*lookoutmetrics.DescribeAnomalyDetectionExecutionsOutput, bool) bool, ...request.Option) error

	DescribeAnomalyDetector(*lookoutmetrics.DescribeAnomalyDetectorInput) (*lookoutmetrics.DescribeAnomalyDetectorOutput, error)
	DescribeAnomalyDetectorWithContext(aws.Context, *lookoutmetrics.DescribeAnomalyDetectorInput, ...request.Option) (*lookoutmetrics.DescribeAnomalyDetectorOutput, error)
	DescribeAnomalyDetectorRequest(*lookoutmetrics.DescribeAnomalyDetectorInput) (*request.Request, *lookoutmetrics.DescribeAnomalyDetectorOutput)

	DescribeMetricSet(*lookoutmetrics.DescribeMetricSetInput) (*lookoutmetrics.DescribeMetricSetOutput, error)
	DescribeMetricSetWithContext(aws.Context, *lookoutmetrics.DescribeMetricSetInput, ...request.Option) (*lookoutmetrics.DescribeMetricSetOutput, error)
	DescribeMetricSetRequest(*lookoutmetrics.DescribeMetricSetInput) (*request.Request, *lookoutmetrics.DescribeMetricSetOutput)

	DetectMetricSetConfig(*lookoutmetrics.DetectMetricSetConfigInput) (*lookoutmetrics.DetectMetricSetConfigOutput, error)
	DetectMetricSetConfigWithContext(aws.Context, *lookoutmetrics.DetectMetricSetConfigInput, ...request.Option) (*lookoutmetrics.DetectMetricSetConfigOutput, error)
	DetectMetricSetConfigRequest(*lookoutmetrics.DetectMetricSetConfigInput) (*request.Request, *lookoutmetrics.DetectMetricSetConfigOutput)

	GetAnomalyGroup(*lookoutmetrics.GetAnomalyGroupInput) (*lookoutmetrics.GetAnomalyGroupOutput, error)
	GetAnomalyGroupWithContext(aws.Context, *lookoutmetrics.GetAnomalyGroupInput, ...request.Option) (*lookoutmetrics.GetAnomalyGroupOutput, error)
	GetAnomalyGroupRequest(*lookoutmetrics.GetAnomalyGroupInput) (*request.Request, *lookoutmetrics.GetAnomalyGroupOutput)

	GetDataQualityMetrics(*lookoutmetrics.GetDataQualityMetricsInput) (*lookoutmetrics.GetDataQualityMetricsOutput, error)
	GetDataQualityMetricsWithContext(aws.Context, *lookoutmetrics.GetDataQualityMetricsInput, ...request.Option) (*lookoutmetrics.GetDataQualityMetricsOutput, error)
	GetDataQualityMetricsRequest(*lookoutmetrics.GetDataQualityMetricsInput) (*request.Request, *lookoutmetrics.GetDataQualityMetricsOutput)

	GetFeedback(*lookoutmetrics.GetFeedbackInput) (*lookoutmetrics.GetFeedbackOutput, error)
	GetFeedbackWithContext(aws.Context, *lookoutmetrics.GetFeedbackInput, ...request.Option) (*lookoutmetrics.GetFeedbackOutput, error)
	GetFeedbackRequest(*lookoutmetrics.GetFeedbackInput) (*request.Request, *lookoutmetrics.GetFeedbackOutput)

	GetFeedbackPages(*lookoutmetrics.GetFeedbackInput, func(*lookoutmetrics.GetFeedbackOutput, bool) bool) error
	GetFeedbackPagesWithContext(aws.Context, *lookoutmetrics.GetFeedbackInput, func(*lookoutmetrics.GetFeedbackOutput, bool) bool, ...request.Option) error

	GetSampleData(*lookoutmetrics.GetSampleDataInput) (*lookoutmetrics.GetSampleDataOutput, error)
	GetSampleDataWithContext(aws.Context, *lookoutmetrics.GetSampleDataInput, ...request.Option) (*lookoutmetrics.GetSampleDataOutput, error)
	GetSampleDataRequest(*lookoutmetrics.GetSampleDataInput) (*request.Request, *lookoutmetrics.GetSampleDataOutput)

	ListAlerts(*lookoutmetrics.ListAlertsInput) (*lookoutmetrics.ListAlertsOutput, error)
	ListAlertsWithContext(aws.Context, *lookoutmetrics.ListAlertsInput, ...request.Option) (*lookoutmetrics.ListAlertsOutput, error)
	ListAlertsRequest(*lookoutmetrics.ListAlertsInput) (*request.Request, *lookoutmetrics.ListAlertsOutput)

	ListAlertsPages(*lookoutmetrics.ListAlertsInput, func(*lookoutmetrics.ListAlertsOutput, bool) bool) error
	ListAlertsPagesWithContext(aws.Context, *lookoutmetrics.ListAlertsInput, func(*lookoutmetrics.ListAlertsOutput, bool) bool, ...request.Option) error

	ListAnomalyDetectors(*lookoutmetrics.ListAnomalyDetectorsInput) (*lookoutmetrics.ListAnomalyDetectorsOutput, error)
	ListAnomalyDetectorsWithContext(aws.Context, *lookoutmetrics.ListAnomalyDetectorsInput, ...request.Option) (*lookoutmetrics.ListAnomalyDetectorsOutput, error)
	ListAnomalyDetectorsRequest(*lookoutmetrics.ListAnomalyDetectorsInput) (*request.Request, *lookoutmetrics.ListAnomalyDetectorsOutput)

	ListAnomalyDetectorsPages(*lookoutmetrics.ListAnomalyDetectorsInput, func(*lookoutmetrics.ListAnomalyDetectorsOutput, bool) bool) error
	ListAnomalyDetectorsPagesWithContext(aws.Context, *lookoutmetrics.ListAnomalyDetectorsInput, func(*lookoutmetrics.ListAnomalyDetectorsOutput, bool) bool, ...request.Option) error

	ListAnomalyGroupRelatedMetrics(*lookoutmetrics.ListAnomalyGroupRelatedMetricsInput) (*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, error)
	ListAnomalyGroupRelatedMetricsWithContext(aws.Context, *lookoutmetrics.ListAnomalyGroupRelatedMetricsInput, ...request.Option) (*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, error)
	ListAnomalyGroupRelatedMetricsRequest(*lookoutmetrics.ListAnomalyGroupRelatedMetricsInput) (*request.Request, *lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput)

	ListAnomalyGroupRelatedMetricsPages(*lookoutmetrics.ListAnomalyGroupRelatedMetricsInput, func(*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, bool) bool) error
	ListAnomalyGroupRelatedMetricsPagesWithContext(aws.Context, *lookoutmetrics.ListAnomalyGroupRelatedMetricsInput, func(*lookoutmetrics.ListAnomalyGroupRelatedMetricsOutput, bool) bool, ...request.Option) error

	ListAnomalyGroupSummaries(*lookoutmetrics.ListAnomalyGroupSummariesInput) (*lookoutmetrics.ListAnomalyGroupSummariesOutput, error)
	ListAnomalyGroupSummariesWithContext(aws.Context, *lookoutmetrics.ListAnomalyGroupSummariesInput, ...request.Option) (*lookoutmetrics.ListAnomalyGroupSummariesOutput, error)
	ListAnomalyGroupSummariesRequest(*lookoutmetrics.ListAnomalyGroupSummariesInput) (*request.Request, *lookoutmetrics.ListAnomalyGroupSummariesOutput)

	ListAnomalyGroupSummariesPages(*lookoutmetrics.ListAnomalyGroupSummariesInput, func(*lookoutmetrics.ListAnomalyGroupSummariesOutput, bool) bool) error
	ListAnomalyGroupSummariesPagesWithContext(aws.Context, *lookoutmetrics.ListAnomalyGroupSummariesInput, func(*lookoutmetrics.ListAnomalyGroupSummariesOutput, bool) bool, ...request.Option) error

	ListAnomalyGroupTimeSeries(*lookoutmetrics.ListAnomalyGroupTimeSeriesInput) (*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, error)
	ListAnomalyGroupTimeSeriesWithContext(aws.Context, *lookoutmetrics.ListAnomalyGroupTimeSeriesInput, ...request.Option) (*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, error)
	ListAnomalyGroupTimeSeriesRequest(*lookoutmetrics.ListAnomalyGroupTimeSeriesInput) (*request.Request, *lookoutmetrics.ListAnomalyGroupTimeSeriesOutput)

	ListAnomalyGroupTimeSeriesPages(*lookoutmetrics.ListAnomalyGroupTimeSeriesInput, func(*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, bool) bool) error
	ListAnomalyGroupTimeSeriesPagesWithContext(aws.Context, *lookoutmetrics.ListAnomalyGroupTimeSeriesInput, func(*lookoutmetrics.ListAnomalyGroupTimeSeriesOutput, bool) bool, ...request.Option) error

	ListMetricSets(*lookoutmetrics.ListMetricSetsInput) (*lookoutmetrics.ListMetricSetsOutput, error)
	ListMetricSetsWithContext(aws.Context, *lookoutmetrics.ListMetricSetsInput, ...request.Option) (*lookoutmetrics.ListMetricSetsOutput, error)
	ListMetricSetsRequest(*lookoutmetrics.ListMetricSetsInput) (*request.Request, *lookoutmetrics.ListMetricSetsOutput)

	ListMetricSetsPages(*lookoutmetrics.ListMetricSetsInput, func(*lookoutmetrics.ListMetricSetsOutput, bool) bool) error
	ListMetricSetsPagesWithContext(aws.Context, *lookoutmetrics.ListMetricSetsInput, func(*lookoutmetrics.ListMetricSetsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*lookoutmetrics.ListTagsForResourceInput) (*lookoutmetrics.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *lookoutmetrics.ListTagsForResourceInput, ...request.Option) (*lookoutmetrics.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*lookoutmetrics.ListTagsForResourceInput) (*request.Request, *lookoutmetrics.ListTagsForResourceOutput)

	PutFeedback(*lookoutmetrics.PutFeedbackInput) (*lookoutmetrics.PutFeedbackOutput, error)
	PutFeedbackWithContext(aws.Context, *lookoutmetrics.PutFeedbackInput, ...request.Option) (*lookoutmetrics.PutFeedbackOutput, error)
	PutFeedbackRequest(*lookoutmetrics.PutFeedbackInput) (*request.Request, *lookoutmetrics.PutFeedbackOutput)

	TagResource(*lookoutmetrics.TagResourceInput) (*lookoutmetrics.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *lookoutmetrics.TagResourceInput, ...request.Option) (*lookoutmetrics.TagResourceOutput, error)
	TagResourceRequest(*lookoutmetrics.TagResourceInput) (*request.Request, *lookoutmetrics.TagResourceOutput)

	UntagResource(*lookoutmetrics.UntagResourceInput) (*lookoutmetrics.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *lookoutmetrics.UntagResourceInput, ...request.Option) (*lookoutmetrics.UntagResourceOutput, error)
	UntagResourceRequest(*lookoutmetrics.UntagResourceInput) (*request.Request, *lookoutmetrics.UntagResourceOutput)

	UpdateAlert(*lookoutmetrics.UpdateAlertInput) (*lookoutmetrics.UpdateAlertOutput, error)
	UpdateAlertWithContext(aws.Context, *lookoutmetrics.UpdateAlertInput, ...request.Option) (*lookoutmetrics.UpdateAlertOutput, error)
	UpdateAlertRequest(*lookoutmetrics.UpdateAlertInput) (*request.Request, *lookoutmetrics.UpdateAlertOutput)

	UpdateAnomalyDetector(*lookoutmetrics.UpdateAnomalyDetectorInput) (*lookoutmetrics.UpdateAnomalyDetectorOutput, error)
	UpdateAnomalyDetectorWithContext(aws.Context, *lookoutmetrics.UpdateAnomalyDetectorInput, ...request.Option) (*lookoutmetrics.UpdateAnomalyDetectorOutput, error)
	UpdateAnomalyDetectorRequest(*lookoutmetrics.UpdateAnomalyDetectorInput) (*request.Request, *lookoutmetrics.UpdateAnomalyDetectorOutput)

	UpdateMetricSet(*lookoutmetrics.UpdateMetricSetInput) (*lookoutmetrics.UpdateMetricSetOutput, error)
	UpdateMetricSetWithContext(aws.Context, *lookoutmetrics.UpdateMetricSetInput, ...request.Option) (*lookoutmetrics.UpdateMetricSetOutput, error)
	UpdateMetricSetRequest(*lookoutmetrics.UpdateMetricSetInput) (*request.Request, *lookoutmetrics.UpdateMetricSetOutput)
}

var _ LookoutMetricsAPI = (*lookoutmetrics.LookoutMetrics)(nil)
