package job

import (
	"context"
	"time"
)

type Job interface {
	Excute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDuration(time []time.Duration)
}

const (
	defaultMaxTimeout = 3
)

var defaultRetryTime = []time.Duration{time.Second, time.Second * 3, time.Second * 4}

type JobHandler func(ctx context.Context) error

type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	return &job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeout,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		state:      StateInit,
		retryIndex: -1,
		stopChan:   make(chan bool),
	}
}

func (j *job) Excute(ctx context.Context) error {
	j.state = StateRunning

	err := j.handler(ctx)
	if err != nil {
		j.state = StateFailed
		return err
	}
	j.state = StateCompleted
	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1
	time.Sleep(j.config.Retries[j.retryIndex])

	err := j.handler(ctx)
	//handle success
	if err == nil {
		j.state = StateCompleted
		return nil
	}
	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return err
	}
	j.state = StateFailed
	return err

}

func (j *job) State() JobState { return j.state }

func (j *job) SetRetryDuration(time []time.Duration) { j.config.Retries = time }
