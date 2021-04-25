package mock

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestMock(t *testing.T) {
	//mock controller
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	//mock person
	mockPerson := NewMockPerson(mockController)

	//当参数是 "name" 时执行函数(注意后面函数的参数个数需要和SetName参数个数一样)
	mockPerson.EXPECT().SetName("name1").AnyTimes().Do(func(name string) {
		fmt.Println("name1")
	})

	//当参数是 "name1" 时,执行函数并返回,(注意后面函数的参数个数需要和SetName参数个数一样)
	mockPerson.EXPECT().SetName("name2").AnyTimes().DoAndReturn(func(name string) string {
		fmt.Println(name)
		return "name2"
	})

	//当参数是 "name1" 时,执行函数并返回,(注意后面函数的参数个数需要和SetName参数个数一样)
	//gomock.Any() 指定参数匹配条件
	//AnyTimes 指定可以执行任意次
	mockPerson.EXPECT().SetName(gomock.Any()).AnyTimes().Do(func(name string) {
		fmt.Println("any", name)
	})

	mockPerson.SetName("name3")
	mockPerson.SetName("name1")
	mockPerson.SetName("name1")
	mockPerson.SetName("name2")
	mockPerson.SetName("name3")
	//output
	//any name3
	//name1
	//name1
	//name2
	//any name3
}
