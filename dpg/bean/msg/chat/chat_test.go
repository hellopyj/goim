package chat
import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"testing"
)

type (
	Result struct {
		Code    int             `json:"code"`
		Message string          `json:"message"`
		Data    json.RawMessage `json:"data"`
	}

	TickersResult struct {
		Tickers []struct {
			ID              string  `json:"_id"`
			DisplayPairName string  `json:"display_pair_name"`
			Price           float64 `json:"price"`
			Volume          float64 `json:"volume"`
		} `json:"tickers"`
	}

	Student struct {
		Name   string  `json:"name"`
		Class  string  `json:"class"`
		Weight float64 `json:"weight"`
		Age    int     `json:"age"`
	}
	IBean interface{ Tag() }
)

func (*TickersResult) Tag() {}
func (*Student) Tag()       {}

func (r *Result) ParseData(iBean IBean) IBean {
	if err := json.Unmarshal([]byte(r.Data), &iBean); err != nil {
		return nil
	}
	return iBean
}

func TestJson(t *testing.T) {
	tickersJsonBytes := []byte(`{"code": 0, "message": "success", "data": { "tickers": [ { "_id": "5a33741cce79d2cf9bf5942c", "display_pair_name": "EOS/KRW", "price": 12.5284441, "volume": 202401841.79944953 }]}}`)
	var result Result
	if err := json.Unmarshal(tickersJsonBytes, &result); err != nil {
		t.Fatal(err)
	}
	var tickersResult TickersResult
	t.Log(result.ParseData(&tickersResult))

	studentJsonBytes := []byte(`{"code":0,"message":"success","data":{"name":"5a33741cce79d2cf9bf5942c","class":"A109","weight":12.5284441,"age":20}}`)
	var studentResult Result
	if err := json.Unmarshal(studentJsonBytes, &studentResult); err != nil {
		t.Fatal(err)
	}
	var student Student
	t.Log(studentResult.ParseData(&student))
}
func TestChatMsg_Init(t *testing.T) {
	msg:=ChatMsg{}
	msg.Init()
	fmt.Println(jsoniter.MarshalToString(msg))
}
