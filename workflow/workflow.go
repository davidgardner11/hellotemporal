package workflow

import (
	"context"
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context, name string) (string error) {
	return "", nil
}

func Activity(ctx context.Context, name string) (string error) {
        return "", nil
}
