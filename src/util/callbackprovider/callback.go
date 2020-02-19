package callbackprovider

import (
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/helper/libaws"
	"github.com/sofyan48/cimol/src/util/helper/provider"
)

// ProviderCallback ...
type ProviderCallback struct {
	AwsLib   libaws.AwsInterface
	Provider provider.ProvidersInterface
}

// ProviderCallbackHandler ...
func ProviderCallbackHandler() *ProviderCallback {
	return &ProviderCallback{
		AwsLib:   libaws.AwsHAndler(),
		Provider: provider.ProvidersHandler(),
	}
}

// ProviderCallbackInterface ...
type ProviderCallbackInterface interface {
	InfobipCallback(dynamo *entity.DynamoItemResponse, data *entity.InfobipCallbackRequest, history *entity.HistoryItem)
	WavecellCallback(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
	wavecellMessagesNotSuccess(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
	infobipMessagesNotSuccess(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
}
