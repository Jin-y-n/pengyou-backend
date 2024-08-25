package common

import (
	"fmt"
	"pengyou/global/config"
	"pengyou/utils/log"
	"strconv"
	"sync"
	"time"
)

const (
	// Epoch time in milliseconds (January 1, 2023).
	epoch = 1672531200000 // Unix timestamp for 2023-01-01T00:00:00Z
)

var snowflakeIdGenerator *SnowflakeIDGenerator

// SnowflakeIDGenerator generates unique IDs.
type SnowflakeIDGenerator struct {
	mu       sync.Mutex
	workerID uint16
	lastTime int64
	sequence uint16
}

func NextSnowflakeID() uint64 {
	return snowflakeIdGenerator.NextID()
}

// NewSnowflakeIDGenerator creates a new SnowflakeIDGenerator.
func NewSnowflakeIDGenerator(workerID uint16) (*SnowflakeIDGenerator, error) {
	if workerID > 1023 || workerID < 0 {
		return nil, fmt.Errorf("workerID must be between 0 and 1023")
	}
	return &SnowflakeIDGenerator{
		workerID: workerID,
	}, nil
}

// NextSnowflakeID generates the next unique ID.
func (s *SnowflakeIDGenerator) NextID() uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentTime := currentTimeMillis()
	if currentTime < s.lastTime {
		panic(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", s.lastTime-currentTime))
	}

	if currentTime == s.lastTime {
		s.sequence++
		if s.sequence > 4095 {
			currentTime = waitForNextMillis(s.lastTime)
		}
	} else {
		s.sequence = 0
	}

	s.lastTime = currentTime

	// convert uint64 to int64 to ensure the consistency of the generated IDs
	workerIDInt64 := int64(s.workerID)
	sequenceInt64 := int64(s.sequence)

	return uint64(((currentTime - epoch) << 22) | (workerIDInt64 << 12) | sequenceInt64)
}

// currentTimeMillis returns the current time in milliseconds.
func currentTimeMillis() int64 {
	return time.Now().UnixMilli()
}

// waitForNextMillis waits until the next millisecond.
func waitForNextMillis(lastTime int64) int64 {
	for currentTime := range time.Tick(time.Millisecond) {
		if currentTime.UnixMilli() > lastTime {
			return currentTime.UnixMilli()
		}
	}
	return 0
}

func InitSnowflakeIdGenerator(snowflake config.Snowflake) {
	var err error
	snowflakeIdGenerator, err = NewSnowflakeIDGenerator(snowflake.WorkerID)
	if err != nil {
		panic(err)
	}

	log.Logger.Info("Snowflake ID generator initialized with worker ID: " + strconv.Itoa(int(snowflake.WorkerID)))
}
