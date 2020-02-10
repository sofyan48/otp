package transmiter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
)

func (trs *Transmiter) wavecellActionShard(history string, payload *entity.HistoryItem) {
	reformatPayload := &entity.WavecellRequest{}
	reformatPayload.Destination = payload.Payload.Msisdn
	reformatPayload.Source = os.Getenv("WAVECELL_ACC_ID")
	reformatPayload.Text = payload.Payload.Text
	reformatPayload.DLRCallback = os.Getenv("WAVECELL_CALLBACK_URL")
	wavecelSendURL := "https://api.wavecell.com/sms/v1/" + os.Getenv("WAVECELL_SUB_ACC_ID_GENERAL") + "/single"
	if payload.Payload.OTP == true {
		wavecelSendURL = "https://api.wavecell.com/sms/v1/" + os.Getenv("WAVECELL_SUB_ACC_ID") + "/single"
	}
	wavecellReformatPayload, err := json.Marshal(reformatPayload)
	client, err := trs.Requester.CLIENT("POST", wavecelSendURL, wavecellReformatPayload)
	if err != nil {
		log.Println("Error: ", err)
	}
	requester := &http.Client{}
	client.Header.Set("Content-Type", "application/json")
	client.Header.Set("Authorization", "Bearer "+os.Getenv("WAVECELL_ACC_TOKEN"))
	response, err := requester.Do(client)
	if err != nil {
		log.Println("Wavecell Transmitter: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Wavecell Transmitter: ", err)
	}
	wavecellResponse := &entity.WavecellResponse{}
	json.Unmarshal(body, wavecellResponse)
	_, err = trs.updateDynamoTransmitt(payload.CallbackData,
		wavecellResponse.Status.Code,
		string(body))
	if err != nil {
		log.Println("Wavecell Transmitter Dynamo: ", err)
	}
}
