// Copyright 2019-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"fmt"
	"github.com/onosproject/onos-test/pkg/kube"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

// newTestWorker returns a new test worker
func newTestWorker(test *TestConfig) (Worker, error) {
	kubeAPI, err := kube.GetAPI(test.JobID)
	if err != nil {
		return nil, err
	}
	return &TestWorker{
		client: kubeAPI.Client(),
		config: test,
	}, nil
}

// newBenchmarkWorker returns a new test worker
func newBenchmarkWorker(config *BenchmarkConfig) (Worker, error) {
	kubeAPI, err := kube.GetAPI(config.JobID)
	if err != nil {
		return nil, err
	}
	return &BenchmarkWorker{
		client: kubeAPI.Client(),
		config: config,
	}, nil
}

// Worker runs a single test suite
type Worker interface {
	// Run runs a test suite
	Run() error
}

// TestWorker runs a test job
type TestWorker struct {
	client client.Client
	config *TestConfig
}

// Run runs a test
func (w *TestWorker) Run() error {
	test, ok := Registry.tests[w.config.Suite]
	if !ok {
		return fmt.Errorf("unknown test suite %s", w.config.Suite)
	}

	tests := []testing.InternalTest{
		{
			Name: w.config.Suite,
			F: func(t *testing.T) {
				RunTests(t, test, w.config)
			},
		},
	}

	// Hack to enable verbose testing.
	os.Args = []string{
		os.Args[0],
		"-test.v",
	}

	testing.Main(func(_, _ string) (bool, error) { return true, nil }, tests, nil, nil)
	return nil
}

// BenchmarkWorker runs a benchmark job
type BenchmarkWorker struct {
	client client.Client
	config *BenchmarkConfig
}

// Run runs a benchmark
func (w *BenchmarkWorker) Run() error {
	benchmark, ok := Registry.benchmarks[w.config.Suite]
	if !ok {
		return fmt.Errorf("unknown benchmark suite %s", w.config.Suite)
	}
	RunBenchmarks(benchmark, w.config)
	return nil
}