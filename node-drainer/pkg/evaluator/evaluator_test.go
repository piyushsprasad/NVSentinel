package evaluator

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nvidia/nvsentinel/data-models/pkg/model"
	"github.com/nvidia/nvsentinel/data-models/pkg/protos"
)

func TestEvaluateEventWithDatabase_DrainOverridesSkipMarksAlreadyDrained(t *testing.T) {
	evaluator := &NodeDrainEvaluator{}

	result, err := evaluator.EvaluateEventWithDatabase(context.Background(), model.HealthEventWithStatus{
		HealthEvent: &protos.HealthEvent{
			NodeName: "test-node",
			DrainOverrides: &protos.BehaviourOverrides{
				Skip: true,
			},
		},
		HealthEventStatus: &protos.HealthEventStatus{
			NodeQuarantined: string(model.Quarantined),
			UserPodsEvictionStatus: &protos.OperationStatus{
				Status: string(model.StatusInProgress),
			},
		},
	}, nil, nil)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, ActionMarkAlreadyDrained, result.Action)
	require.Equal(t, model.AlreadyDrained, result.Status)
}
