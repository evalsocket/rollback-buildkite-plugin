package main

import (
	"encoding/json"
	"fmt"
	"github.com/buildkite/go-buildkite/v3/buildkite"
	"github.com/hasura/go-buildkite-dsl/pipeline"
	"github.com/hasura/go-buildkite-dsl/step"
	"log"
	"os"
)

const API_TOKE = "BUILDKITE_API_TOKE"
const ORG = "BUILDKITE_ORGANIZATION_SLUG"
const PIPELINE_NAME = "BUILDKITE_PIPELINE_SLUG"
const BRANCH_NAME = "BUILDKITE_PIPELINE_DEFAULT_BRANCH"

var label = "Rollback"

func main() {
	token := os.Getenv(API_TOKE)
	if len(token) == 0 {
		log.Fatalf("Please pass the buildkite api token on env %s", API_TOKE)
		os.Exit(1)
	}

	org := os.Getenv(ORG)
	if len(token) == 0 {
		log.Fatalf("%s env variable in not available", ORG)
		os.Exit(1)
	}

	pipelineName := os.Getenv(PIPELINE_NAME)
	if len(token) == 0 {
		log.Fatalf("%s env variable in not available", PIPELINE_NAME)
		os.Exit(1)
	}

	branch := os.Getenv(BRANCH_NAME)
	if len(token) == 0 {
		log.Fatalf("%s env variable in not available", BRANCH_NAME)
		os.Exit(1)
	}

	config, err := buildkite.NewTokenConfig(token, false)
	if err != nil {
		log.Fatalf("client config failed: %s", err)
		os.Exit(1)
	}
	client := buildkite.NewClient(config.Client())

	builds, _, err := client.Builds.ListByPipeline(org, pipelineName, &buildkite.BuildsListOptions{
		Branch: branch,
		State:  []string{"passed"},
		ListOptions: buildkite.ListOptions{
			Page:    1,
			PerPage: 1,
		},
	})
	if err != nil {
		log.Fatalf("List build failed: %s", BRANCH_NAME)
		os.Exit(1)
	}
	rollbackCommit := builds[0].Commit
	dynamicPipeline := pipeline.New(fmt.Sprintf("rollback-%s", pipelineName))

	triggerStep := step.Trigger{
		Label:        &label,
		PipelineSlug: pipeline.Slug(pipelineName),
		Branches:     &branch,
		Build: &step.Build{
			Message: &label,
			Commit:  rollbackCommit,
			Branch:  &branch,
		},
	}
	dynamicPipeline.Steps = append(dynamicPipeline.Steps, triggerStep)

	b, err := json.MarshalIndent(dynamicPipeline, "", "	")
	if err != nil {
		log.Fatalf("json marshel failed: %s", err)
	}
	fmt.Println(string(b))
}
