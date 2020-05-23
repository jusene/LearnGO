## Go语言 测试

```go
import "fmt"

func main() {
	fmt.Println(Age(-7))
}

func Age(n int) int {
	if n > 0 {
		return n
	}
	n = 0
	return n
}
```

- 单元测试
```go
package main

import "testing"

func TestAge(t *testing.T) {
	var (
		input = -100
		expected = 0
	)

	actual := Age(input)
	if actual != expected {
		t.Errorf("Age(%d) = %d 预期 %d", input, actual, expected)
	}
}
```

- 表组测试
```go
package main

import "fmt"

func main() {
	fmt.Println(isPrime(25))
}

func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5;i * i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}
```

```go
package  main

import "testing"

func TestPrime(t *testing.T) {
	var primeTests = []struct{
		input int
		expected bool
	} {
		{1, false},
		{2, true},
	}

	for _, tt := range primeTests {
		actual := isPrime(tt.input)
		if actual != tt.expected {
			t.Errorf("%d %v %v", tt.input, actual, tt.expected)
		}
	}
}
```

- Fail: 记录失败信息，然后继续执行后续用例
- FailNow: 记录失败信息，所以测试中断
- SkipNow: 不会记录失败的用例信息，然后中断测试
- Skip: 记录失败信息，中断后续测试
- Skipf: 相比前者多了一个格式化输出
- Log: 输出错误信息，在单元测试中，默认不输出成功的用例的信息，不会中断后续测试
- Logf: 输出格式化的信息，不中断后续测试
- Error: 相当于Log + Fail，不会中断后续测试
- Errorf: 相当于Logf + Fail，同上
- Fatal: 相当于Log + FailNow，会中断后续测试
- Fatalf: 相当于Logf + FailNow，同上



