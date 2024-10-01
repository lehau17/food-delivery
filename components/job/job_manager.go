package job

import (
	"context"
	"log"
	"sync"
)

type JobManager struct {
	jobs         []Job
	isConcurrent bool
	wg           *sync.WaitGroup
}

func NewJobManager(isConcurrent bool, jobs ...Job) *JobManager {
	return &JobManager{jobs: jobs,
		isConcurrent: isConcurrent,
		wg:           new(sync.WaitGroup)}
}

func (jm *JobManager) RunJob(ctx context.Context, j Job) error {
	if err := j.Excute(ctx); err != nil {
		for {
			log.Println(err)
			if j.State() == StateRetryFailed {
				return err
			}
			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}
	return nil
}

func (jm *JobManager) Run(ctx context.Context) error {
	jm.wg.Add(len(jm.jobs))
	errChan := make(chan error, len(jm.jobs))
	for i := range jm.jobs {
		if jm.isConcurrent {
			go func(j Job) {
				errChan <- jm.RunJob(ctx, j)
				jm.wg.Done()
			}(jm.jobs[i])
			continue
		}
		job := jm.jobs[i]
		errChan <- jm.RunJob(ctx, job)
		jm.wg.Done()
	}
	var err error
	for i := 1; i <= len(jm.jobs); i++ {
		if v := <-errChan; v != nil {
			err = v
		}
	}
	jm.wg.Wait()
	return err
}
