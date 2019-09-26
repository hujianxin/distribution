// Package fds implements the Xiaomi Galaxy FDS(File Data Storage) driver backend
package fds

import (
	"fmt"

	fdsservice "github.com/XiaoMi/galaxy-fds-sdk-golang"
	// fdsmodel "github.com/XiaoMi/galaxy-fds-sdk-golang/Model"
	storagedriver "github.com/docker/distribution/registry/storage/driver"
	"github.com/docker/distribution/registry/storage/driver/base"
	"github.com/docker/distribution/registry/storage/driver/factory"
)

const driverName = "fds"

func init() {
	factory.Register(driverName, &fdsDriverFactory{})
}

type fdsDriverFactory struct{}

//DriverParameters A struct that encapsulates all of the driver parameters after all values have been set
type DriverParameters struct {
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	Endpoint        string
	Internal        bool
	ChunkSize       int64
	RootDirectory   string
}

type driver struct {
	Client          *fdsservice.FDSClient
	Bucket          string
	ChunkSize       int64
	Encrypt         bool
	RootDirectory   string
	EncryptionKeyID string
}

type baseEmbed struct{ base.Base }

// Driver is a storagedriver.StorageDriver implementation backed by
// Xiaomi Galaxy FDS Service.
type Driver struct{ baseEmbed }

func (factory *fdsDriverFactory) Create(parameters map[string]interface{}) (storagedriver.StorageDriver, error) {
	return FromParameters(parameters)
}

// FromParameters constructs a new Driver with a given parameters map.
func FromParameters(parameters map[string]interface{}) (*Driver, error) {
	accessKey, ok := parameters["accessKey"]
	if !ok || fmt.Sprint(accessKey) == "" {
		return nil, fmt.Errorf("no accessKey parameter provided")
	}

	return New(DriverParameters{})
}

// New constructs a new Driver with the given FDS credentials, endpoint and bucket name
func New(params DriverParameters) (*Driver, error) {
	return nil, nil
}
