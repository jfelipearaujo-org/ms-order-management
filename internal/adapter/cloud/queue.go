package cloud

import (
	"context"
	"encoding/json"
	"log/slog"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/jfelipearaujo-org/ms-order-management/internal/service"
	"github.com/jfelipearaujo-org/ms-order-management/internal/service/order/process"
)

type QueueService interface {
	GetQueueName() string
	UpdateQueueUrl(ctx context.Context) error
	ConsumeMessages(ctx context.Context)
}

type AwsSqsService struct {
	QueueName string
	QueueUrl  string
	Client    *sqs.Client

	MessageProcessor service.ProcessMessageService[process.ProcessMessageDto]

	ChanMessage chan types.Message

	Mutex     sync.Mutex
	WaitGroup sync.WaitGroup
}

func NewQueueService(queueName string, config aws.Config, messageProcessor service.ProcessMessageService[process.ProcessMessageDto]) QueueService {
	client := sqs.NewFromConfig(config)

	return &AwsSqsService{
		QueueName: queueName,
		Client:    client,

		MessageProcessor: messageProcessor,

		ChanMessage: make(chan types.Message, 10),

		Mutex:     sync.Mutex{},
		WaitGroup: sync.WaitGroup{},
	}
}

func (s *AwsSqsService) GetQueueName() string {
	return s.QueueName
}

func (s *AwsSqsService) UpdateQueueUrl(ctx context.Context) error {
	output, err := s.Client.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: &s.QueueName,
	})
	if err != nil {
		return err
	}

	s.QueueUrl = *output.QueueUrl

	return nil
}

func (s *AwsSqsService) ConsumeMessages(ctx context.Context) {
	output, err := s.Client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &s.QueueUrl,
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     5,
	})
	if err != nil {
		slog.ErrorContext(ctx, "error receiving message from queue", "queue_url", s.QueueUrl, "error", err)
		return
	}

	s.WaitGroup.Add(len(output.Messages))

	for _, message := range output.Messages {
		go s.processMessage(ctx, message)
	}

	s.WaitGroup.Wait()
}

func (s *AwsSqsService) processMessage(ctx context.Context, message types.Message) {
	defer s.WaitGroup.Done()
	s.Mutex.Lock()

	slog.InfoContext(ctx, "message received", "message_id", *message.MessageId)

	var notification TopicNotification

	if err := json.Unmarshal([]byte(*message.Body), &notification); err != nil {
		slog.ErrorContext(ctx, "error unmarshalling message", "error", err)
	} else {
		if notification.Type != "Notification" {
			slog.ErrorContext(ctx, "message is not a notification", "message_id", *message.MessageId)
		} else {
			var request process.ProcessMessageDto

			if err := json.Unmarshal([]byte(notification.Message), &request); err != nil {
				slog.ErrorContext(ctx, "error unmarshalling message", "message_id", *message.MessageId, "error", err)
			} else {
				slog.InfoContext(ctx, "message unmarshalled", "request", request)
				if err := s.MessageProcessor.Handle(ctx, request); err != nil {
					slog.ErrorContext(ctx, "error processing message", "message_id", *message.MessageId, "error", err)
				}
			}
		}
	}

	if err := s.deleteMessage(ctx, message); err != nil {
		slog.ErrorContext(ctx, "error deleting message", "message_id", *message.MessageId, "error", err)
	}

	s.Mutex.Unlock()
}

func (s *AwsSqsService) deleteMessage(ctx context.Context, message types.Message) error {
	_, err := s.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &s.QueueUrl,
		ReceiptHandle: message.ReceiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
