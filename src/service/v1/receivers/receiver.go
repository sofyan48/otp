package receivers

import (
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/callbackprovider"
	"github.com/sofyan48/cimol/src/util/helper/libaws"
)

// Receiver ...
type Receiver struct {
	AwsLib   libaws.AwsInterface
	Callback callbackprovider.ProviderCallbackInterface
}

// ReceiverHandler ...
func ReceiverHandler() *Receiver {
	return &Receiver{
		AwsLib:   libaws.AwsHAndler(),
		Callback: callbackprovider.ProviderCallbackHandler(),
	}
}

// ReceiverInterface ...
type ReceiverInterface interface {
	InfobipReceiver(ID string, data *entity.InfobipCallbackRequest)
	WavecellReceiver(ID string, data *entity.WavecellCallBackRequest) (string, error)
}

// TwilioReceiver ...
func (rcv *Receiver) TwilioReceiver(ID string, data *entity.WavecellCallBackRequest) (string, error) {

	return "", nil
}
