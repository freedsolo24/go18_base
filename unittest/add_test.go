package unittest_test

import (
	"os"
	"testing"
)

// run test: 直接跑,
// run debug: 一步一步的调试
func TestAdd(t *testing.T) { // 针对Add函数的单测。传参必须是*testing.T
	t.Log(os.Getenv("CONFIG_PATH"))

	// 只是想运行下, 可以调t.Log方法，打印结果
	// 自己手工判断, 打印log判断
	// 配置vscode打印单元测试的日志
	// 如果没有打印日志,配置vscode 打印单元测试日志 -v -count=1
	// t.Log(unittest.Add(1, 2))

	// 通过程序断言判断
	// if unittest.Add(1, 2) != 4 {
	// 	t.Fatalf("1+2!=4")
	// }

	// 专门的断言库
	// should := assert.New(t) // 构建一个断言对象
	// // 调equal方法,
	// should.Equal(3, unittest.Add(1, 2)) // 断言函数的Add的返回是否等于3

}
