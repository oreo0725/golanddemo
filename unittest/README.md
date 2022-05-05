# Unit test

## 甚麼是工作單元？
1. 被呼叫的公開方法回傳一個結果值
2. 在呼叫方法的前後，系統可見的狀態或行為發生變化
3. 呼叫一個不受測試所控制的第三方系統

## 單元測試應該具備以下特質：
1. Repeatable Automation
2. 容易被實現
3. 它到第二天應該還有存在意義（不是臨時性的）
4. Easy to run/trigger
5. Fast
6. 執行結果應該一致
7. 能完全掌控被測試的單元
8. 是完全被隔離的（獨立於其他測試）
9. 如果它的執行結果是失敗的，應該要很簡單清楚地呈現我們的期望為何，問題在哪

## To write a unit test
- `math.go`, the test file `math_test.go`
- use the [Testify](https://github.com/stretchr/testify) library 
    ```go
    got := Add(5, 5)
    want := 10
    assert.Equal(t, want, got)
    ```

## Run the test

```sh
go test -v ./unittest/math
```
  - `-v`: verbose

```sh
go test -coverprofile=coverage.out
```

## Write a Table-driven test

```go
testcases := []struct {
    name   string
    arg1   int
    arg2   int
    expect int
}{
    {
        name: "ok WHEN arg1, arg2 are both positive",
        arg1: 10, arg2: 12,
        expect: 24,
    },
    {
        name: "overflow when arg1 + arg2 too large",
        arg1: 1, arg2: 9223372036854775807,
        expect: -9223372036854775808,
    },
}
for _, tc := range testcases {
    t.Run(tc.name, func(t *testing.T) {
        got := Add(tc.arg1, tc.arg2)
        require.Equal(t, tc.expect, got)
    })
}
```

## Mock generation

```shell
# install mockgen
go install github.com/golang/mock/mockgen@v1.6.0

# to generate the mock object
go generate ./unittest/ex2product/credit/...
```
