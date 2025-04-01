// package main
//
// import (
//
//	"bufio"
//	"os"
//	"strconv"
//	"strings"
//
// )
//
// /*
//
//	这里说当程序需要使用不同方式处理不同种类请求， 而且请求类型和顺序预先未知时， 可以使用责任链模式。
//	暂且没有get到
//	https://dunwu.github.io/design/pages/b25735/#%E6%84%8F%E5%9B%BE
//
// */
//
//	type Handler interface {
//		Handle(request request)
//		SetNext(handler Handler)
//	}
//
//	type SupervisorHandler struct {
//		next Handler
//	}
//
//	func (s *SupervisorHandler) Handle(req request) {
//		if s != nil {
//			if req.reqDays <= 3 {
//				println(req.name + " Approved by Supervisor.")
//			} else {
//				if s.next != nil {
//					s.next.Handle(req)
//				}
//			}
//		}
//	}
//
//	func (s *SupervisorHandler) SetNext(handler Handler) {
//		if s != nil {
//			s.next = handler
//		}
//	}
//
//	type ManagerHandler struct {
//		next Handler
//	}
//
//	func (s *ManagerHandler) Handle(req request) {
//		if s != nil {
//			if req.reqDays <= 7 {
//				println(req.name + " Approved by Manager.")
//			} else {
//				if s.next != nil {
//					s.next.Handle(req)
//				}
//			}
//		}
//	}
//
//	func (s *ManagerHandler) SetNext(handler Handler) {
//		if s != nil {
//			s.next = handler
//		}
//	}
//
//	type DirectorHandler struct {
//		next Handler
//	}
//
//	func (s *DirectorHandler) Handle(req request) {
//		if s != nil {
//			if req.reqDays <= 10 {
//				println(req.name + " Approved by Director.")
//			} else {
//				println(req.name + " Denied by Director.")
//			}
//		}
//	}
//
//	func (s *DirectorHandler) SetNext(handler Handler) {
//		if s != nil {
//			s.next = handler
//		}
//	}
//
//	type request struct {
//		name    string
//		reqDays int
//	}
//
//	func main() {
//		scanner := bufio.NewScanner(os.Stdin)
//
//		dHandler := &DirectorHandler{}
//
//		mHandler := &ManagerHandler{}
//		mHandler.SetNext(dHandler)
//
//		hHanler := &SupervisorHandler{}
//		hHanler.SetNext(mHandler)
//
//		scanner.Scan()
//		times, _ := strconv.Atoi(strings.Fields(scanner.Text())[0])
//
//		reqs := make([]request, 0, 0)
//		for ; times >= 1; times-- {
//			scanner.Scan()
//			line := scanner.Text()
//			parts := strings.Fields(line)
//
//			reqs = append(reqs, request{
//				name: parts[0],
//				reqDays: func() int {
//					num, _ := strconv.Atoi(parts[1])
//					return num
//				}(),
//			})
//		}
//
//		for _, req := range reqs {
//			hHanler.Handle(req)
//		}
//	}

// DeepSeek优化后的代码
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Handler interface {
	Handle(request request)
	SetNext(handler Handler)
}

type SupervisorHandler struct{ next Handler }
type ManagerHandler struct{ next Handler }
type DirectorHandler struct{ next Handler }

func (s *SupervisorHandler) Handle(req request) {
	if req.reqDays <= 3 {
		fmt.Println(req.name + " Approved by Supervisor.")
	} else if s.next != nil {
		s.next.Handle(req)
	}
}

func (s *SupervisorHandler) SetNext(handler Handler) { s.next = handler }

func (s *ManagerHandler) Handle(req request) {
	if req.reqDays <= 7 {
		fmt.Println(req.name + " Approved by Manager.")
	} else if s.next != nil {
		s.next.Handle(req)
	}
}

func (s *ManagerHandler) SetNext(handler Handler) { s.next = handler }

func (s *DirectorHandler) Handle(req request) {
	if req.reqDays <= 10 {
		fmt.Println(req.name + " Approved by Director.")
	} else {
		fmt.Println(req.name + " Denied by Director.")
	}
}

func (s *DirectorHandler) SetNext(handler Handler) { s.next = handler }

type request struct {
	name    string
	reqDays int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 构建责任链
	director := &DirectorHandler{}
	manager := &ManagerHandler{}
	manager.SetNext(director)
	supervisor := &SupervisorHandler{}
	supervisor.SetNext(manager)

	// 读取输入
	line, _ := reader.ReadString('\n')
	times, _ := strconv.Atoi(strings.TrimSpace(line))

	for i := 0; i < times; i++ {
		line, _ = reader.ReadString('\n')
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		days, _ := strconv.Atoi(parts[1])
		supervisor.Handle(request{name: parts[0], reqDays: days})
	}
}
