/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scheduling

import (
	opmetrics "github.com/awslabs/operatorpkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	crmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"

	"sigs.k8s.io/karpenter/pkg/metrics"
)

const (
	ControllerLabel    = "controller"
	schedulingIDLabel  = "scheduling_id"
	schedulerSubsystem = "scheduler"

	logyballNodeNameLabel       = "node_name"
	logyballTopologyDomainLabel = "domain"
)

var (
	DurationSeconds = opmetrics.NewPrometheusHistogram(
		crmetrics.Registry,
		prometheus.HistogramOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "scheduling_duration_seconds",
			Help:      "Duration of scheduling simulations used for deprovisioning and provisioning in seconds.",
			Buckets:   metrics.DurationBuckets(),
		},
		[]string{
			ControllerLabel,
		},
	)
	QueueDepth = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "queue_depth",
			Help:      "The number of pods currently waiting to be scheduled.",
		},
		[]string{
			ControllerLabel,
			schedulingIDLabel,
		},
	)
	UnfinishedWorkSeconds = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "unfinished_work_seconds",
			Help:      "How many seconds of work has been done that is in progress and hasn't been observed by scheduling_duration_seconds.",
		},
		[]string{
			ControllerLabel,
			schedulingIDLabel,
		},
	)
	IgnoredPodCount = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Name:      "ignored_pod_count",
			Help:      "Number of pods ignored during scheduling by Karpenter",
		},
		[]string{},
	)
	UnschedulablePodsCount = opmetrics.NewPrometheusGauge(
		crmetrics.Registry,
		prometheus.GaugeOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "unschedulable_pods_count",
			Help:      "The number of unschedulable Pods.",
		},
		[]string{
			ControllerLabel,
		},
	)

	LogyballNewSchedulersCreated = opmetrics.NewPrometheusCounter(
		crmetrics.Registry,
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "logyball_new_scheduler_created_total",
			Help:      "The number of schedulers created (calls to NewScheduler).",
		},
		[]string{},
	)

	LogyballNewSchedulerStateNodesConsidered = opmetrics.NewPrometheusCounter(
		crmetrics.Registry,
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "logyball_new_scheduler_existing_nodes_total",
			Help:      "The number of nodes existing in the cluster as considered by all runs of NewScheduler.",
		},
		[]string{},
	)

	LogyballNewSchedulerTopologiesConsidered = opmetrics.NewPrometheusCounter(
		crmetrics.Registry,
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "logyball_new_scheduler_topologies_considered_total",
			Help:      "The number of topologies per node in the cluster as considered by all runs of NewScheduler.",
		},
		[]string{
			logyballNodeNameLabel,
		},
	)

	LogyballNewSchedulerInverseTopologiesConsidered = opmetrics.NewPrometheusCounter(
		crmetrics.Registry,
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "logyball_new_scheduler_inverse_topologies_considered_total",
			Help:      "The number of inverse topologies per node in the cluster as considered by all runs of NewScheduler.",
		},
		[]string{
			logyballNodeNameLabel,
		},
	)

	LogyballTopologyGroupDomainsScanned = opmetrics.NewPrometheusCounter(
		crmetrics.Registry,
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "logyball_topology_groups_domains_scanned_total",
			Help:      "The number domains scanned when registering topologies.",
		},
		[]string{
			logyballTopologyDomainLabel,
		},
	)

	LogyballTopologyGroupDomainsInserted = opmetrics.NewPrometheusCounter(
		crmetrics.Registry,
		prometheus.CounterOpts{
			Namespace: metrics.Namespace,
			Subsystem: schedulerSubsystem,
			Name:      "logyball_topology_groups_domains_inserted_total",
			Help:      "The number domains inserted into a new set when registering topologies.",
		},
		[]string{
			logyballTopologyDomainLabel,
		},
	)
)
