package smartcontract

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Silentttttt/iplay/go-iplay/models"
)

func createQuizze(
	payType uint32,
	gameType uint32,
	deadLine time.Time,
	amount uint64,
	theme string,
	opts []models.ChoiceOpt) (uint64, error) {
	smartContractOpt := make([]*option, 0)

	opt1 := option{1, "fs"}
	a, _ := json.Marshal(opt1)
	fmt.Println(string(a))
	for _, opt := range opts {
		smartContractOpt = append(smartContractOpt, &option{opt.Odds, opt.Name})
	}
	params := make([]interface{}, 0)
	params = append(params, payType)
	params = append(params, gameType)
	params = append(params, deadLine.Unix()*1000)
	params = append(params, amount)
	params = append(params, theme)
	params = append(params, smartContractOpt)

	b, _ := json.Marshal(params)
	fmt.Println(string(b))
	// req, err := http.NewRequest("POST", "url", bytes.NewBuffer(b))
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	return 1, nil
}
