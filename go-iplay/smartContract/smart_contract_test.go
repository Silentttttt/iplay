package smartcontract

import (
	"fmt"
	"testing"
	"time"

	"github.com/Silentttttt/iplay/go-iplay/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateQuizze(t *testing.T) {
	theme := "巴西 vs 德国"

	opts := []models.ChoiceOpt{
		{1, "德国胜", 2, 0.5, 1, 2, nil},
		{2, "巴西胜", 3, 0.5, 1, 2, nil},
	}
	txHash, err := createQuizze(1, 1, time.Now().Unix()*1000+3600*1000, 1, theme, opts)

	assert.Nil(t, err)
	fmt.Println(txHash)

	// tests := []struct {
	// 	filepath       string
	// 	expectedErr    error
	// 	expectedResult string
	// }{
	// 	{"test/test_multi_lib_version_require.js", nil, "\"\""},
	// 	{"test/test_uint.js", nil, "\"\""},
	// 	{"test/test_date_1.0.5.js", nil, "\"\""},
	// 	{"test/test_crypto.js", nil, "\"\""},
	// 	{"test/test_blockchain_1.0.5.js", nil, "\"\""},
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.filepath, func(t *testing.T) {
	// 		data, err := ioutil.ReadFile(tt.filepath)
	// 		assert.Nil(t, err, "filepath read error")
	// 		mem, _ := storage.NewMemoryStorage()
	// 		context, _ := state.NewWorldState(dpos.NewDpos(), mem)
	// 		addr, _ := core.AddressParse("n1FF1nz6tarkDVwWQkMnnwFPuPKUaQTdptE")
	// 		owner, err := context.GetOrCreateUserAccount(addr.Bytes())
	// 		assert.Nil(t, err)
	// 		owner.AddBalance(newUint128FromIntWrapper(1000000000000))
	// 		addr, _ = core.AddressParse("n1p8cwrrfrbFe71eda1PQ6y4WnX3gp8bYze")
	// 		contract, _ := context.CreateContractAccount(addr.Bytes(), nil, &corepb.ContractMeta{Version: "1.0.5"})
	// 		ctx, err := NewContext(mockBlockForLib(2000000), mockTransaction(), contract, context)

	// 		engine := NewV8Engine(ctx)
	// 		engine.SetExecutionLimits(10000000, 10000000)
	// 		result, err := engine.RunScriptSource(string(data), 0)
	// 		assert.Equal(t, tt.expectedErr, err)
	// 		assert.Equal(t, tt.expectedResult, result)
	// 		engine.Dispose()
	// 	})
	// }
}
