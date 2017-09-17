package util

import (
	"encoding/base64"
	"fmt"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/edgedagency/mulungu/logger"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

//PubSubTopicID returns topic id with namespace attached
func PubSubTopicID(namespace, topic string) string {
	return strings.Join([]string{namespace, topic}, "-")
}

//PubSubTopicSplitID returns namespace and topic
func PubSubTopicSplitID(topicID string) []string {
	return strings.Split(topicID, "-")
}

//PubSubTopicPublish sends data to topic
func PubSubTopicPublish(ctx context.Context, topicID string, payload map[string]interface{}, attributes map[string]string) (id string, err error) {
	topic, topicError := PubSubTopic(ctx, topicID)
	if topicError != nil {
		logger.Errorf(ctx, "pubsub util", "Failed to publish to topic: %s", topicID)
		return "", topicError
	}
	// NOTE:Concurrency settup
	// topic.PublishSettings = pubsub.PublishSettings{
	// 	NumGoroutines: 2,
	// }
	result := topic.Publish(ctx, &pubsub.Message{Data: []byte(base64.StdEncoding.EncodeToString([]byte(MapInterfaceToJSONString(payload)))), Attributes: attributes})
	resultID, resultErr := result.Get(ctx)
	if err != nil {
		logger.Errorf(ctx, "pubsub util", "Failed to publish to topic: %s", resultErr.Error())
		return "", err
	}

	return resultID, nil
}

//PubSubData obtain data item by key
func PubSubData(data map[string]interface{}, messageKey, key string) interface{} {
	if dataElement, ok := data[messageKey]; ok {
		if val, ok := dataElement.(map[string]interface{})[key]; ok {
			return val
		}
	}
	return nil
}

//PubSubClient retunrs a client
func PubSubClient(ctx context.Context) (*pubsub.Client, error) {
	pubsubClient, pubsubClientErr := pubsub.NewClient(ctx, appengine.AppID(ctx))
	if pubsubClientErr != nil {
		logger.Errorf(ctx, "pubsub util", "Failed to create client: %v", pubsubClientErr)
		return nil, pubsubClientErr
	}
	return pubsubClient, nil
}

//PubSubTopic returns a topic based on topic name
func PubSubTopic(ctx context.Context, topicID string) (*pubsub.Topic, error) {
	pubsubClient, pubsubClientErr := PubSubClient(ctx)
	if pubsubClientErr != nil {
		logger.Errorf(ctx, "pubsub util", "Failed to create client: %v", pubsubClientErr)
		return nil, pubsubClientErr
	}
	defer pubsubClient.Close()

	topic := pubsubClient.Topic(topicID)

	if topic == nil {
		logger.Errorf(ctx, "pubsub util", "Failed to obtain topic: %s", topicID)
		return nil, fmt.Errorf("failed to obtain topic %s", topicID)
	}

	return topic, nil
}

//PubSubCreateTopicSubscription creates a topic subscription
func PubSubCreateTopicSubscription(ctx context.Context, topicID, endpoint string) (*pubsub.Subscription, error) {
	topic, topicErr := PubSubTopic(ctx, topicID)
	if topicErr != nil {
		logger.Errorf(ctx, "pubsub util", "unable to obtain topic %s", topicErr.Error())
		return nil, topicErr
	}

	pubsubClient, pubsubClientErr := PubSubClient(ctx)
	if pubsubClientErr != nil {
		logger.Errorf(ctx, "pubsub util", "Failed to create client: %v", pubsubClientErr)
		return nil, pubsubClientErr
	}
	defer pubsubClient.Close()

	subscription, subscriptionErr := pubsubClient.CreateSubscription(ctx, topicID,
		pubsub.SubscriptionConfig{Topic: topic, PushConfig: pubsub.PushConfig{Endpoint: endpoint}})
	if subscriptionErr != nil {
		logger.Errorf(ctx, "pubsub subscription service", "subscription failed: %s", subscriptionErr.Error())
		return nil, subscriptionErr
	}

	return subscription, nil
}

//PubSubDeleteTopicSubscription delets a topic subscription
func PubSubDeleteTopicSubscription(ctx context.Context, subscriptionID string) (*pubsub.Subscription, error) {
	pubsubClient, pubsubClientErr := PubSubClient(ctx)
	if pubsubClientErr != nil {
		logger.Errorf(ctx, "pubsub util", "Failed to create client: %v", pubsubClientErr)
		return nil, pubsubClientErr
	}
	defer pubsubClient.Close()
	subscription := pubsubClient.Subscription(subscriptionID)
	if subscription == nil {
		return nil, fmt.Errorf("Uable to obtain subscription to delete %s", subscriptionID)
	}

	subscriptionErr := subscription.Delete(ctx)

	if subscriptionErr != nil {
		return nil, subscriptionErr
	}

	return subscription, nil
}
