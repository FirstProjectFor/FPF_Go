package server_test

import (
	"fmt"
	"github.com/onsi/ginkgo/config"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGink(t *testing.T) {
	RegisterFailHandler(Fail)
	//测试结果生成xml文件
	//junitReporter := reporters.NewJUnitReporter("junit.xml")
	//RunSpecsWithDefaultAndCustomReporters(t, "Foo Suite", []Reporter{junitReporter})
	RunSpecs(t, "Gink Suite")
}

//在每个spec执行前执行一次
var _ = BeforeSuite(func() {
	node := config.GinkgoConfig.ParallelNode
	fmt.Println(node)
})

var _ = AfterSuite(func() {

})
