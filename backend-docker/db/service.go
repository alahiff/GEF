package db

import (
	"time"
)

// Bind describes the binding between an IOPort and a docker volume
type Bind struct {
	IOPort   IOPort
	VolumeID string
}

// Service describes metadata for a GEF service (used to serialize JSON)
type Service struct {
	ID          ServiceID
	ImageID     string
	Name        string
	RepoTag     string
	Description string
	Version     string
	Created     time.Time
	Size        int64
	Input       []IOPort
	Output      []IOPort
}

// ServiceID exported
type ServiceID string

// IOPort is an i/o specification for a service
// The service can only read data from volumes and write to a single volume
// Path specifies where the volumes are mounted
type IOPort struct {
	ID   string
	Name string
	Path string
}
