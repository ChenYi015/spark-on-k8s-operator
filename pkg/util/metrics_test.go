/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kubeflow/spark-operator/pkg/common"
)

func TestPositiveGauge_EmptyLabels(t *testing.T) {
	gauge := NewPositiveGauge("testGauge", "test-description", []string{})
	emptyMap := map[string]string{}
	gauge.Dec(emptyMap)
	assert.InEpsilon(t, float64(0), fetchGaugeValue(gauge.gaugeMetric, emptyMap), common.Epsilon)

	gauge.Inc(emptyMap)
	assert.InEpsilon(t, float64(1), fetchGaugeValue(gauge.gaugeMetric, emptyMap), common.Epsilon)
	gauge.Dec(map[string]string{})
	assert.InEpsilon(t, float64(0), fetchGaugeValue(gauge.gaugeMetric, emptyMap), common.Epsilon)
}

func TestPositiveGauge_WithLabels(t *testing.T) {
	gauge := NewPositiveGauge("testGauge1", "test-description-1", []string{"app_id"})
	app1 := map[string]string{"app_id": "test1"}
	app2 := map[string]string{"app_id": "test2"}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			gauge.Inc(app1)
		}
		for i := 0; i < 5; i++ {
			gauge.Dec(app1)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 5; i++ {
			gauge.Inc(app2)
		}
		for i := 0; i < 10; i++ {
			gauge.Dec(app2)
		}
		wg.Done()
	}()

	wg.Wait()
	assert.InEpsilon(t, float64(5), fetchGaugeValue(gauge.gaugeMetric, app1), common.Epsilon)
	// Always Positive Gauge.
	assert.InEpsilon(t, float64(0), fetchGaugeValue(gauge.gaugeMetric, app2), common.Epsilon)
}
