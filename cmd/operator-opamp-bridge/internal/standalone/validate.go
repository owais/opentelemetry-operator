// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package standalone

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"sigs.k8s.io/yaml"

	"github.com/open-telemetry/opentelemetry-operator/apis/v1beta1"
	"github.com/open-telemetry/opentelemetry-operator/internal/webhook"
)

func validateCollectorConfigEntry(body string) error {
	if strings.TrimSpace(body) == "" {
		return errors.New("config value is empty")
	}

	var cfg v1beta1.Config
	if err := yaml.Unmarshal([]byte(body), &cfg); err != nil {
		return fmt.Errorf("failed to parse collector config: %w", err)
	}

	collector := &v1beta1.OpenTelemetryCollector{
		Spec: v1beta1.OpenTelemetryCollectorSpec{
			Mode:   v1beta1.ModeDeployment,
			Config: cfg,
		},
	}
	if _, err := (webhook.CollectorWebhook{}).Validate(context.Background(), collector); err != nil {
		return fmt.Errorf("collector config failed operator validation: %w", err)
	}

	return nil
}
