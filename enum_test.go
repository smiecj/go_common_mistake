// package go_common_mistake 参考博客 The Top 10 Most Common Mistakes I’ve Seen in Go Projects 编写的测试 go 常见错误代码
package go_common_mistake

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type StateEnum int

// 正确的枚举定义
const (
	StateUnknown StateEnum = iota
	StateOpen
	StateClose
)

// 不合理的枚举定义
// const (
// 	StateOpen StateEnum = iota
// 	StateClose
// 	StateUnknown
// )

type ReadFileResponse struct {
	Name  string
	Rsp   string
	State StateEnum
}

func TestEnum(t *testing.T) {
	readSuccessStr := fmt.Sprintf(`{"Name": "test.txt", "Rsp": "Read file finish", "State": %d}`, StateOpen)
	var rsp ReadFileResponse
	err := json.Unmarshal([]byte(readSuccessStr), &rsp)
	require.Equal(t, nil, err)
	require.Equal(t, StateOpen, rsp.State)

	rsp = ReadFileResponse{}
	readUnknownStr := `{"Name": "test.txt", "Rsp": "Read file finish"}`
	_ = json.Unmarshal([]byte(readUnknownStr), &rsp)
	require.Equal(t, StateUnknown, rsp.State)
}
