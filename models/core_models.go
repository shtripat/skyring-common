/*Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package models

import (
	"github.com/skyrings/skyring-common/tools/uuid"
	"time"
)

type Node struct {
	NodeId        uuid.UUID          `json:"nodeid"`
	Hostname      string             `json:"hostname"`
	Tags          []string           `json:"tags"`
	ManagementIP4 string             `json:"management_ip4"`
	ClusterIP4    string             `json:"cluster_ip4"`
	PublicIP4     string             `json:"public_ip4"`
	ClusterId     uuid.UUID          `json:"clusterid"`
	Location      string             `json:"location"`
	Status        string             `json:"status"`
	Options       map[string]string  `json:"options"`
	CPUs          []CPU              `json:"cpus"`
	NetworkInfo   StorageNodeNetwork `json:"network_info"`
	StorageDisks  []Disk             `json:"storage_disks"`
	Memory        []Memory           `json:"memory"`
	OS            OperatingSystem    `json:"os"`
	Enabled       bool               `json:"enabled"`
}

type Disk struct {
	DevName    string
	FSType     string
	FSUUID     uuid.UUID
	Model      string
	MountPoint []string
	Name       string
	Parent     string
	Size       uint64
	Type       string
	Used       bool
	Vendor     string
}

type CPU struct {
	CPUId string `bson:"cpuid"`
}

type StorageNodeNetwork struct {
	Subnet []string `bson:"subnet"`
	Ipv4   []string `bson:"ipv4"`
	Ipv6   []string `bson:"ipv6"`
}

type Memory struct {
	Name       string `bson:"name"`
	Type       string `bson:"type"`
	TotalSize  int    `bson:"totalsize"`
	FreeSize   int    `bson:"freesize"`
	Attributes string `bson:"attribute"`
}

type OperatingSystem struct {
	Name                string `bson:"name"`
	OSVersion           string `bson:"osversion"`
	KernelVersion       string `bson:"kernel_version"`
	StorageSWVersion    string `bson:"storage_sw_version"`
	KdumpStatus         string `bson:"kdump_status"`
	MemPageShareStatus  string `bson:"mem_page_share_status"`
	AutomaticLargePages bool   `bson:"automatic_large_pages"`
	SELinuxMode         string `bson:"selinuxmode"`
}

type User struct {
	Username            string   `json:"username"`
	Email               string   `json:"email"`
	Hash                []byte   `json:"hash"`
	Role                string   `json:"role"`
	Groups              []string `json:"groups"`
	Type                int      `json:"type"`
	Status              bool     `json:"status"`
	FirstName           string   `json:"firstname"`
	LastName            string   `json:"lastname"`
	NotificationEnabled bool     `json:"notificationenabled"`
}

type Cluster struct {
	ClusterId         uuid.UUID         `json:"clusterid"`
	Name              string            `json:"name"`
	CompatVersion     string            `json:"compat_version"`
	Type              string            `json:"type"`
	WorkLoad          string            `json:"workload"`
	Status            int               `json:"status"`
	Tags              []string          `json:"tags"`
	Options           map[string]string `json:"options"`
	OpenStackServices []string          `json:"openstack_services"`
	Networks          ClusterNetworks   `json:"networks"`
	Enabled           bool              `json:"enabled"`
}

type ClusterNetworks struct {
	Cluster string `json:"cluster"`
	Public  string `json:"public"`
}

type StorageLogicalUnit struct {
	SluId             uuid.UUID         `json:"sluid"`
	Name              string            `json:"name"`
	Type              int               `json:"type"`
	ClusterId         uuid.UUID         `json:"clusterid"`
	NodeId            uuid.UUID         `json:"nodeid"`
	StorageId         uuid.UUID         `json:"storageid"`
	StorageDeviceId   uuid.UUID         `json:"storagedeviceid"`
	StorageDeviceSize uint64            `json:"storagedevicesize"`
	Status            string            `json:"status"`
	Options           map[string]string `json:"options"`
}

type Storage struct {
	StorageId           uuid.UUID         `json:"storageid"`
	Name                string            `json:"name"`
	Type                string            `json:"type"`
	Tags                []string          `json:"tags"`
	ClusterId           uuid.UUID         `json:"clusterid"`
	Size                string            `json:"size"`
	Status              string            `json:"status"`
	Replicas            int               `json:"replicas"`
	Profile             string            `json:"profile"`
	SnapshotsEnabled    bool              `json:"snapshots_enabled"`
	SnapshotScheduleIds []uuid.UUID       `json:"snapshot_schedule_ids"`
	QuotaEnabled        bool              `json:"quota_enabled"`
	QuotaParams         map[string]string `json:"quota_params"`
	Options             map[string]string `json:"options"`
}

type SnapshotSchedule struct {
	Id            uuid.UUID `json:"id"`
	Recurrence    string    `json:"recurrence"`
	Interval      int       `json:"interval"`
	ExecutionTime string    `json:"execution_time"`
	Days          []string  `json:"days"`
	StartFrom     string    `json:"start_from"`
	EndBy         string    `json:"endby"`
}

type Status struct {
	Timestamp time.Time
	Message   string
}

type AppTask struct {
	Id         uuid.UUID         `json:"id"`
	Name       string            `json:"name"`
	ParentId   uuid.UUID         `json:"parentid"`
	Started    bool              `json:"started"`
	Completed  bool              `json:"completed"`
	StatusList []Status          `json:"statuslist"`
	Tag        map[string]string `json:"tag"`
	SubTasks   []uuid.UUID       `json:"subtasks"`
	Status     TaskStatus        `json:"status"`
}
