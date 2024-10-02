package subcriber

import (
	"context"
	"log"

	"github.com/lehau17/food_delivery/common"
	appcontext "github.com/lehau17/food_delivery/components/app_context"
	"github.com/lehau17/food_delivery/components/job"
	"github.com/lehau17/food_delivery/components/pubsub"
)

type cosumerjob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appcontext.AppContect
}

func NewConsumerEngine(appCtx appcontext.AppContect) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (engine *consumerEngine) Start() error {
	engine.startSubTopic(common.TopicUserLikeRestaurant, true, IncreaseLikeCountAfterUserDisLikeRestaurant(engine.appCtx))
	engine.startSubTopic(common.TopicUserUnLikeRestaurant, true, DesLikeCountAfterUserDisLikeRestaurant(engine.appCtx))
	return nil
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, cosumerJob ...cosumerjob) error {
	c, _ := engine.appCtx.GetPubSub().Subcribe(context.Background(), topic)
	for _, item := range cosumerJob {
		log.Println("Set up ", item.Title)
	}
	getJobHanlder := func(job *cosumerjob, message *pubsub.Message) job.JobHandler {
		return func(context context.Context) error {
			log.Println("Running ", job.Title)
			return job.Hld(context, message)
		}
	}
	go func() {
		for {
			mess := <-c
			jobHdlArr := make([]job.Job, len(cosumerJob))
			for i := range cosumerJob {
				jobHdl := getJobHanlder(&cosumerJob[i], mess)
				jobHdlArr[i] = job.NewJob(jobHdl)
			}
			group := job.NewJobManager(isConcurrent, jobHdlArr...)
			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()
	return nil
}
