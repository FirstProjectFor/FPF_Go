package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"time"
)

var _ = Describe("Server", func() {

	var (
		message1 string
		message2 string
	)

	//每个测试之后执行
	BeforeEach(func() {
		message1 = "hello1"
		message2 = "hello1"
	})

	//每个测试之前执行
	AfterEach(func() {
	})

	//测试message1
	Describe("消息内容", func() {
		When("消息名字是 message1", func() {
			It("内容应该是 hello1", func() {
				Expect(message1).To(Equal("hello1"))
			})
		})
	})

	//测试message2
	Describe("消息内容", func() {
		BeforeEach(func() {
			message2 = "hello2"
		})
		When("消息名字是 message2", func() {
			It("内容应该是 hello2", func() {
				Ω(func() (string, error) {
					return message2, nil
				}()).Should(Equal("hello2"))
			})
		})
	})

	//不执行测试,这个测试不会执行
	PIt("pending ", func() {
		Fail("pending")
	})

	//跳过测试
	It("skip", func() {
		Skip("skip")
		Fail("skip")
	})

	////只执行该测试,其余测试会跳过
	//FIt("focus", func() {
	//	Fail("focus")
	//})

	//获取测试环境信息
	It("test information", func() {
		description := CurrentGinkgoTestDescription()
		Expect(description).NotTo(Equal(nil))
	})

	//获取测试环境信息
	It("test by", func() {
		By("step 1")
		Expect("step 1").To(Equal("step 1"))
		By("step 1")
		Expect("step 2").To(Equal("step 2"))
	})

	//超时,执行时间超过指定超时时间时会报错，单位秒
	PIt("over time", func(done Done) {
		time.Sleep(time.Millisecond * 30)
		close(done)
	}, 0.2)

	//基准测试,运行200次
	Measure("measure", func(b Benchmarker) {
		b.Time("runtime", func() {
			Expect(17).To(Equal(17))
			//记录数据
			b.RecordValue("disk usage (in MB)", rand.Float64())
		})
	}, 2000)

	//使用assert
	It("assert", func() {
		assert.Equal(GinkgoT(), "assert", "assert")
	})
})
