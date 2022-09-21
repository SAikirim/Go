package hdinsightapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/hdinsight/2018-11-01-preview/hdinsight"
	"io"
)

// JobClientAPI contains the set of methods on the JobClient type.
type JobClientAPI interface {
	DeleteSparkBatch(ctx context.Context, batchID int32, requestedBy string) (result hdinsight.SparkJobDeletedResult, err error)
	DeleteSparkSessionJob(ctx context.Context, sessionID int32, requestedBy string) (result hdinsight.SparkJobDeletedResult, err error)
	DeleteSparkStatementJob(ctx context.Context, sessionID int32, statementID int32, requestedBy string) (result hdinsight.SparkStatementCancellationResult, err error)
	Get(ctx context.Context, jobID string) (result hdinsight.JobDetailRootJSONObject, err error)
	GetAppState(ctx context.Context, appID string) (result hdinsight.AppState, err error)
	GetSparkBatchJob(ctx context.Context, batchID int32) (result hdinsight.SparkBatchJob, err error)
	GetSparkBatchLog(ctx context.Context, batchID int32, from *int32, size *int32) (result hdinsight.SparkJobLog, err error)
	GetSparkBatchState(ctx context.Context, batchID int32) (result hdinsight.SparkJobState, err error)
	GetSparkSessionJob(ctx context.Context, sessionID int32) (result hdinsight.SparkSessionJob, err error)
	GetSparkSessionLog(ctx context.Context, sessionID int32, from *int32, size *int32) (result hdinsight.SparkJobLog, err error)
	GetSparkSessionState(ctx context.Context, sessionID int32) (result hdinsight.SparkJobState, err error)
	GetSparkStatementJob(ctx context.Context, sessionID int32, statementID int32) (result hdinsight.SparkStatement, err error)
	Kill(ctx context.Context, jobID string) (result hdinsight.JobDetailRootJSONObject, err error)
	List(ctx context.Context) (result hdinsight.ListJobListJSONObject, err error)
	ListAfterJobID(ctx context.Context, jobid string, numrecords *int32) (result hdinsight.ListJobListJSONObject, err error)
	ListSparkBatchJob(ctx context.Context, from *int32, size *int32) (result hdinsight.SparkBatchJobCollection, err error)
	ListSparkSessionJob(ctx context.Context, from *int32, size *int32) (result hdinsight.SparkSessionCollection, err error)
	ListSparkStatementJob(ctx context.Context, sessionID int32) (result hdinsight.SparkStatementCollection, err error)
	SubmitHiveJob(ctx context.Context, content io.ReadCloser) (result hdinsight.JobSubmissionJSONResponse, err error)
	SubmitMapReduceJob(ctx context.Context, content io.ReadCloser) (result hdinsight.JobSubmissionJSONResponse, err error)
	SubmitMapReduceStreamingJob(ctx context.Context, content io.ReadCloser) (result hdinsight.JobSubmissionJSONResponse, err error)
	SubmitPigJob(ctx context.Context, content io.ReadCloser) (result hdinsight.JobSubmissionJSONResponse, err error)
	SubmitSparkBatchJob(ctx context.Context, sparkBatchJobRequest hdinsight.SparkBatchJobRequest, requestedBy string) (result hdinsight.SparkBatchJob, err error)
	SubmitSparkSessionJob(ctx context.Context, sparkSessionJobRequest hdinsight.SparkSessionJobRequest, requestedBy string) (result hdinsight.SparkSessionJob, err error)
	SubmitSparkStatementJob(ctx context.Context, sessionID int32, sparkStatementRequest hdinsight.SparkStatementRequest, requestedBy string) (result hdinsight.SparkStatement, err error)
	SubmitSqoopJob(ctx context.Context, content io.ReadCloser) (result hdinsight.JobSubmissionJSONResponse, err error)
}

var _ JobClientAPI = (*hdinsight.JobClient)(nil)
