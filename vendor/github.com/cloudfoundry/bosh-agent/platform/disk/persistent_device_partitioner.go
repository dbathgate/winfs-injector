package disk

import (
	"errors"

	"github.com/cloudfoundry/bosh-utils/logger"
)

const MaxFdiskPartitionSize = uint64(2 * 1024 * 1024 * 1024 * 1024)

var ErrGPTPartitionEncountered = errors.New("sfdisk detected a GPT partition")

type PersistentDevicePartitioner struct {
	sfDiskPartitioner Partitioner
	partedPartitioner Partitioner
	deviceUtil        Util
	logger            logger.Logger
}

func NewPersistentDevicePartitioner(
	sfDiskPartitioner Partitioner,
	partedPartitioner Partitioner,
	deviceUtil Util,
	logger logger.Logger,
) *PersistentDevicePartitioner {
	return &PersistentDevicePartitioner{
		sfDiskPartitioner: sfDiskPartitioner,
		partedPartitioner: partedPartitioner,
		deviceUtil:        deviceUtil,
		logger:            logger,
	}
}

func (p *PersistentDevicePartitioner) Partition(devicePath string, partitions []Partition) error {
	size, err := p.deviceUtil.GetBlockDeviceSize(devicePath)
	if err == nil && size > MaxFdiskPartitionSize {
		p.logger.Debug("persistent-disk-partitioner", "Using parted partitioner because disk size is too large: %d", size)
		return p.partedPartitioner.Partition(devicePath, partitions)
	}

	p.logger.Debug("persistent-disk-partitioner", "Attempting to partition with sfdisk partitioner")
	err = p.sfDiskPartitioner.Partition(devicePath, partitions)
	if IsGPTError(err) {
		p.logger.Debug("persistent-disk-partitioner", "GPT partition detected, falling back to parted")
		return p.partedPartitioner.Partition(devicePath, partitions)
	}

	return err
}

func (p *PersistentDevicePartitioner) GetDeviceSizeInBytes(devicePath string) (uint64, error) {
	return p.sfDiskPartitioner.GetDeviceSizeInBytes(devicePath)
}

func IsGPTError(err error) bool {
	return err == ErrGPTPartitionEncountered
}
