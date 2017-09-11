// nolint
package dto

import (
	"fmt"
	"github.com/moira-alert/moira"
	"github.com/moira-alert/moira/api/middleware"
	"github.com/moira-alert/moira/checker"
	"github.com/moira-alert/moira/expression"
	"github.com/moira-alert/moira/target"
	"net/http"
	"time"
)

type TriggersList struct {
	Page  *int64               `json:"page,omitempty"`
	Size  *int64               `json:"size,omitempty"`
	Total *int64               `json:"total,omitempty"`
	List  []moira.TriggerCheck `json:"list"`
}

func (*TriggersList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type Trigger struct {
	moira.Trigger
	Throttling int64 `json:"throttling"`
}

func (trigger *Trigger) Bind(request *http.Request) error {
	if len(trigger.Targets) == 0 {
		return fmt.Errorf("targets is required")
	}
	if trigger.WarnValue == nil && trigger.Expression == nil {
		return fmt.Errorf("warn_value is required")
	}
	if trigger.ErrorValue == nil && trigger.Expression == nil {
		return fmt.Errorf("error_value is required")
	}

	triggerExpression := expression.TriggerExpression{
		AdditionalTargetsValues: make(map[string]float64),
		WarnValue:               trigger.WarnValue,
		ErrorValue:              trigger.ErrorValue,
		PreviousState:           checker.NODATA,
		Expression:              trigger.Expression,
	}

	logger := middleware.GetLoggerEntry(request)

	if err := resolvePatterns(request, trigger, &triggerExpression); err != nil {
		logger.Infof("Invalid graphite targets %s: %s\n", trigger.Targets, err.Error())
		return fmt.Errorf("Invalid graphite targets: %s", err.Error())
	}
	if _, err := triggerExpression.Evaluate(); err != nil {
		logger.Infof("Invalid expression %s: %s\n", moira.UseString(trigger.Expression), err.Error())
		return err
	}
	return nil
}

func resolvePatterns(request *http.Request, trigger *Trigger, expressionValues *expression.TriggerExpression) error {
	now := time.Now().Unix()
	targetNum := 1
	trigger.Patterns = make([]string, 0)
	timeSeriesNames := make(map[string]bool)

	for _, tar := range trigger.Targets {
		database := middleware.GetDatabase(request)
		result, err := target.EvaluateTarget(database, tar, now-600, now, true)
		if err != nil {
			return err
		}
		trigger.Patterns = append(trigger.Patterns, result.Patterns...)
		for _, timeSeries := range result.TimeSeries {
			timeSeriesNames[timeSeries.Name] = true
		}
		if targetNum == 1 {
			expressionValues.MainTargetValue = 42
		} else {
			targetName := fmt.Sprintf("t%v", targetNum)
			expressionValues.AdditionalTargetsValues[targetName] = 42
		}
		targetNum++
	}
	middleware.SetTimeSeriesNames(request, timeSeriesNames)
	return nil
}

func (*Trigger) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type TriggerCheck struct {
	*moira.CheckData
	TriggerID string `json:"trigger_id"`
}

func (*TriggerCheck) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type MetricsMaintenance map[string]int64

func (*MetricsMaintenance) Bind(r *http.Request) error {
	return nil
}

type ThrottlingResponse struct {
	Throttling int64 `json:"throttling"`
}

func (*ThrottlingResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type SaveTriggerResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func (*SaveTriggerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type TriggerMetrics map[string][]moira.MetricValue

func (*TriggerMetrics) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
