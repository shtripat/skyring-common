/*Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the  specific language governing permissions and
limitations under the License.
*/
package models

import (
	"fmt"
	"github.com/skyrings/skyring-common/tools/uuid"
	"time"
)

type AddStorageNodeRequest struct {
	Hostname       string `json:"hostname"`
	SshFingerprint string `json:"sshfingerprint"`
	User           string `json:"user"`
	Password       string `json:"password"`
	SshPort        int    `json:"sshport"`
}

type AddClusterRequest struct {
	Name              string            `json:"name"`
	CompatVersion     string            `json:"compat_version"`
	Type              string            `json:"type"`
	WorkLoad          string            `json:"workload"`
	Tags              []string          `json:"tags"`
	Options           map[string]string `json:"options"`
	OpenStackServices []string          `json:"openstackservices"`
	Nodes             []ClusterNode     `json:"nodes"`
	Networks          ClusterNetworks   `json:"networks"`
}

type ClusterNode struct {
	NodeId   string              `json:"nodeid"`
	NodeType []string            `json:"nodetype"`
	Devices  []ClusterNodeDevice `json:"disks"`
	Options  map[string]string   `json:"options"`
}

type ClusterNodeDevice struct {
	Name    string            `json:"name"`
	FSType  string            `json:"fstype"`
	Options map[string]string `json:"options"`
}

type AddStorageRequest struct {
	Name             string            `json:"name"`
	Type             string            `json:"type"`
	Tags             []string          `json:"tags"`
	Size             string            `json:"size"`
	Replicas         int               `json:"replicas"`
	Profile          string            `json:"profile"`
	SnapshotsEnabled bool              `json:"snapshots_enabled"`
	SnapshotSchedule SnapshotSchedule  `json:"snapshot_schedule"`
	QuotaEnabled     bool              `json:"quota_enabled"`
	QuotaParams      map[string]string `json:"quota_params"`
	Options          map[string]string `json:"options"`
}

type SnapshotScheduleRequest struct {
	Recurrence    string   `json:"recurrence"`
	Interval      int      `json:"interval"`
	ExecutionTime string   `json:"execution_time"`
	Days          []string `json:"days"`
	StartFrom     string   `json:"start_from"`
	EndBy         string   `json:"endby"`
}

type Nodes []Node

type NodeEvent struct {
	Timestamp time.Time         `json:"timestamp"`
	Node      string            `json:"node"`
	Tag       string            `json:"tag"`
	Tags      map[string]string `json:"tags"`
	Message   string            `json:"message"`
	Severity  string            `json:"severity"`
}

type Event struct {
	EventId   uuid.UUID         `json:"event_id"`
	ClusterId uuid.UUID         `json:"cluster_id"`
	NodeId    uuid.UUID         `json:"node_id"`
	Timestamp time.Time         `json:"timestamp"`
	Tag       string            `json:"tag"`
	Tags      map[string]string `json:"tags"`
	Message   string            `json:"message"`
	Severity  string            `json:"severity"`
}

const (
	DEFAULT_SSH_PORT                = 22
	DEFAULT_FS_TYPE                 = "xfs"
	REQUEST_SIZE_LIMIT              = 1048576
	COLL_NAME_STORAGE               = "storage"
	COLL_NAME_NODE_EVENTS           = "node_events"
	COLL_NAME_STORAGE_NODES         = "storage_nodes"
	COLL_NAME_STORAGE_CLUSTERS      = "storage_clusters"
	COLL_NAME_STORAGE_LOGICAL_UNITS = "storage_logical_units"
	COLL_NAME_TASKS                 = "tasks"
	COLL_NAME_SESSION_STORE         = "skyring_session_store"
)

type Clusters []Cluster
type Storages []Storage

type UnmanagedNode struct {
	Name            string `json:"name"`
	SaltFingerprint string `json:"saltfingerprint"`
}

type UnmanagedNodes []UnmanagedNode

type ClusterStatus int

// Status values for the cluster
const (
	CLUSTER_STATUS_INACTIVE = 1 + iota
	CLUSTER_STATUS_NOT_AVAILABLE
	CLUSTER_STATUS_ACTIVE_AND_AVAILABLE
	CLUSTER_STATUS_CREATING
	CLUSTER_STATUS_FAILED
)

var ClusterStatuses = [...]string{
	"inactive",
	"not available",
	"active and available",
	"creating",
	"failed",
}

// Storage logical unit types
const (
	CEPH_OSD = 1 + iota
)

var StorageLogicalUnitTypes = [...]string{
	"osd",
}

const (
	STATUS_UP   = "up"
	STATUS_DOWN = "down"
)

func (c ClusterStatus) String() string { return ClusterStatuses[c-1] }

type AsyncResponse struct {
	TaskId uuid.UUID `json:"taskid"`
}

func (s Status) String() string {
	return fmt.Sprintf("%s %s", s.Timestamp, s.Message)
}

type TaskStatus int

const (
	TASK_STATUS_NONE = iota
	TASK_STATUS_SUCCESS
	TASK_STATUS_FAILURE
)

var TaskStatuses = [...]string{
	"none",
	"success",
	"failed",
}

func (t TaskStatus) String() string { return TaskStatuses[t] }
