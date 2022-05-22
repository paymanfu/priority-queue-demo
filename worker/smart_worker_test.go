package worker

import (
	"context"
	"github.com/fujiahui/talnet-challenge-payman/common"
	"github.com/fujiahui/talnet-challenge-payman/server"
	"sync"
	"testing"
	"time"
)

func getJobArray(created int64) *common.JobInfoArray {
	jobArray1 := &common.JobInfoArray{
		JobInfos: []*common.JobInfo{
			{
				ID:       1,
				Created:  1,
				Priority: common.LowPriority,
				Tasks:    []uint16{5, 6, 7},
			},
		},
	}

	jobArray2 := &common.JobInfoArray{
		JobInfos: []*common.JobInfo{
			{
				ID:       2,
				Created:  3,
				Priority: common.HighPriority,
				Tasks:    []uint16{3, 5},
			},
		},
	}

	jobInfoMap := make(map[int64]*common.JobInfoArray)
	jobInfoMap[1] = jobArray1
	jobInfoMap[3] = jobArray2

	if jobArray, ok := jobInfoMap[created]; ok {
		return jobArray
	}

	return nil
}

// Task 1.2
func TestNewBaseWorker(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "/Users/fujiahui/go-workspace/talent-challenge-payman/warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := int64(-1)
	w := NewBaseWorker(startTimestamp)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(2000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.1
func TestNewWorkerWithCapacity(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "/Users/fujiahui/go-workspace/talent-challenge-payman/warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := int64(-1)
	capacity := uint16(10)
	w := NewWorkerWithCapacity(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(5000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.2
func TestNewWorkerWithSimplePriority(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "/Users/fujiahui/go-workspace/talent-challenge-payman/warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := int64(-1)
	capacity := uint16(10)
	w := NewWorkerWithSimplePriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(5000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.3
func TestNewWorkerWithSmartPriority(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "/Users/fujiahui/go-workspace/talent-challenge-payman/warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := int64(-1)
	capacity := uint16(10)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(5000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 3.1
func TestNewWorkerWithNumPriority(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "/Users/fujiahui/go-workspace/talent-challenge-payman/warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := int64(-1)
	capacity := uint16(10)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(5000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 3.2
func TestNewWorkerWithTaskSpeed(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	//dirPath := "/Users/fujiahui/go-workspace/talent-challenge-payman/warehouse/data_num_priority/"
	//dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := int64(-1)
	// capacity := uint16(10)
	// w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	w := NewBaseWorker(startTimestamp)
	w.EnableTaskSpeed()
	// w.DisableTaskSpeed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		w.Start(ctx, getJobArray)
		// w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(2000) * time.Millisecond)
	cancel()
	wg.Wait()
}