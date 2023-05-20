package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/manan-rastogi/currencyConvertor/currencyProto"
	"github.com/manan-rastogi/currencyConvertor/helpers"
)

type Service interface {
	CurrencyConveterService(ctx context.Context, input *currencyProto.CurrencyRequest) (output *currencyProto.CurrencyResponse, err error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) CurrencyConveterService(ctx context.Context, input *currencyProto.CurrencyRequest) (output *currencyProto.CurrencyResponse, err error) {
	err = godotenv.Load("../.env")
	if err != nil {
		helpers.ErrorLogF("Error loading .env file-- " + err.Error())
	}

	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/convert?to=%v&from=%v&amount=%v", input.To, input.From, input.Amount)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		helpers.ErrorLog("Got Error while generating request: " + err.Error())
		return
	}
	// defer req.Body.Close()

	req.Header.Add("Accept", "*/*")
	req.Header.Add("apikey", os.Getenv("API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			helpers.ErrorLog("Got timeout while making request: " + err.Error())
			return
		}

		helpers.ErrorLog("Got Error while making request: " + err.Error())
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		helpers.ErrorLog("Error Reading API Response: " + err.Error())
		return
	}

	resp := make(map[string]interface{})
	err = json.Unmarshal(body, &resp)
	if err != nil {
		helpers.ErrorLog("Error Unmarshalling Response: " + err.Error())
		return
	}

	helpers.ValueLog("API Response", fmt.Sprint(resp))

	var response currencyProto.CurrencyResponse
	if resp["success"].(bool) {
		response.Success = true
		response.InputAmount = input.Amount
		response.ExchangeAmount = float32(resp["result"].(float64))
		response.Rate = float32(resp["info"].(map[string]interface{})["rate"].(float64))
	} else {
		response.Success = false
		response.InputAmount = input.Amount
	}

	output = &response
	return
}
