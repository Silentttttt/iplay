package smartcontract

import (
	"fmt"
	"iplay/go-iplay/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateQuizze(t *testing.T) {
	theme := "巴西 vs 德国"

	opts := []*models.ChoiceOpt{
		{1, "德国胜", 2.45, 0.5, 1, 2, nil},
		{2, "巴西胜", 3, 0.5, 1, 2, nil},
		{2, "法国胜", 3.5, 0.5, 1, 2, nil},
	}
	txHash, err := CreateQuizze(nil, 1, 1, time.Now().Unix()*1000+3600*1000, 1, theme, opts)
	go func() {
		CreateQuizze(nil, 1, 1, time.Now().Unix()*1000+3600*1000, 1, theme, opts)
	}()
	go func() {
		CreateQuizze(nil, 1, 1, time.Now().Unix()*1000+3600*1000, 1, theme, opts)
	}()
	time.Sleep(2 * time.Second)
	assert.Nil(t, err)
	fmt.Println(txHash)
}
