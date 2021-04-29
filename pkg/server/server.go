package server

//初始化suite文件
//go:generate ginkgo bootstrap
//初始化包测试文件
//go:generate ginkgo generate server

//go:generate ginkgo -v --progress
//显示进度信息
//go:generate ginkgo --progress
//显示详细信息
//go:generate ginkgo -v
//并发运行
//go:generate ginkgo -p
//打印堆栈追踪信息
//go:generate ginkgo -trace
// 跳过性能测试程序
//go:generate ginkgo --skipMeasurements

//go:generate ginkgo

//打乱测试用例执行顺序
//go:generate ginkgo watch --randomizeAllSpecs

//go:generate ginkgo watch -p

//go:generate ginkgo watch

type HelloServer interface {
	HandleMessage(message string)
}
