package polygon

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/big"
	"net/http"
	validators2 "validators2/interface"
)

// Создаем структуру для декодирования JSON
type Response struct {
	Status  string `json:"status"`
	Success bool   `json:"success"`
	Result  struct {
		Id                        int         `json:"id"`
		Name                      string      `json:"name"`
		ClaimedReward             json.Number `json:"claimedReward"`
		ValidatorUnclaimedRewards json.Number `json:"validatorUnclaimedRewards"`
		PerformanceIndex          int         `json:"performanceIndex"`
	} `json:"result"`
}

type ValidatorReward struct {
	Platform string
	Slashes  string
	Reward   string
}

func Polygon() *ValidatorReward {
	var matic = ValidatorReward{}
	return &matic
}

func (v ValidatorReward) GetValidatorReward() *validators2.ValidatorReward {
	const api = "https://staking-api.polygon.technology/api/v2/validators/31"

	client := http.DefaultClient

	// Создаем запрос GET к API
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Printf("Ошибка при создании запроса: %v\n", err)
	}

	// Отправляем запрос и получаем ответ
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Ошибка при выполнении запроса: %v\n", err)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Ошибка при чтении ответа: %v\n", err)
	}

	// Декодируем JSON-данные
	var formattedData bytes.Buffer
	err = json.Indent(&formattedData, body, "", "\t")
	if err != nil {
		log.Printf("Ошибка при форматировании JSON: %v\n", err)
	}

	// Декодируем JSON в структуру
	var data Response
	err = json.Unmarshal(formattedData.Bytes(), &data)
	if err != nil {
		log.Printf("Ошибка при декодировании JSON: %v\n", err)
	}

	totalRewards := new(big.Int)
	claimedReward := new(big.Int)
	validatorUnclaimedRewards := new(big.Int)
	// performanceIndex := data.Result.PerformanceIndex
	// fmt.Println("Performance Index:", performanceIndex, "%")

	claimedReward.SetString(data.Result.ClaimedReward.String(), 10)
	validatorUnclaimedRewards.SetString(data.Result.ValidatorUnclaimedRewards.String(), 10)

	totalRewards.Add(totalRewards, claimedReward)
	totalRewards.Add(totalRewards, validatorUnclaimedRewards)

	// Convert totalRewards to MATIC
	divisor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	maticValue := new(big.Float).Quo(new(big.Float).SetInt(totalRewards), divisor)
	maticStr := maticValue.Text('f', 4) + " MATIC"

	// Выводим отформатированные данные
	return &validators2.ValidatorReward{
		Platform: "Polygon",
		Slashes:  "",
		Reward:   maticStr,
	}
}
