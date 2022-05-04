// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"
	"time"

	"github.com/authzed/controller-idioms"
)

type FakeControlRequeueAPIErr struct {
	DoneStub        func()
	doneMutex       sync.RWMutex
	doneArgsForCall []struct {
	}
	RequeueStub        func()
	requeueMutex       sync.RWMutex
	requeueArgsForCall []struct {
	}
	RequeueAPIErrStub        func(error)
	requeueAPIErrMutex       sync.RWMutex
	requeueAPIErrArgsForCall []struct {
		arg1 error
	}
	RequeueAfterStub        func(time.Duration)
	requeueAfterMutex       sync.RWMutex
	requeueAfterArgsForCall []struct {
		arg1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeControlRequeueAPIErr) Done() {
	fake.doneMutex.Lock()
	fake.doneArgsForCall = append(fake.doneArgsForCall, struct {
	}{})
	stub := fake.DoneStub
	fake.recordInvocation("Done", []interface{}{})
	fake.doneMutex.Unlock()
	if stub != nil {
		fake.DoneStub()
	}
}

func (fake *FakeControlRequeueAPIErr) DoneCallCount() int {
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	return len(fake.doneArgsForCall)
}

func (fake *FakeControlRequeueAPIErr) DoneCalls(stub func()) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = stub
}

func (fake *FakeControlRequeueAPIErr) Requeue() {
	fake.requeueMutex.Lock()
	fake.requeueArgsForCall = append(fake.requeueArgsForCall, struct {
	}{})
	stub := fake.RequeueStub
	fake.recordInvocation("Requeue", []interface{}{})
	fake.requeueMutex.Unlock()
	if stub != nil {
		fake.RequeueStub()
	}
}

func (fake *FakeControlRequeueAPIErr) RequeueCallCount() int {
	fake.requeueMutex.RLock()
	defer fake.requeueMutex.RUnlock()
	return len(fake.requeueArgsForCall)
}

func (fake *FakeControlRequeueAPIErr) RequeueCalls(stub func()) {
	fake.requeueMutex.Lock()
	defer fake.requeueMutex.Unlock()
	fake.RequeueStub = stub
}

func (fake *FakeControlRequeueAPIErr) RequeueAPIErr(arg1 error) {
	fake.requeueAPIErrMutex.Lock()
	fake.requeueAPIErrArgsForCall = append(fake.requeueAPIErrArgsForCall, struct {
		arg1 error
	}{arg1})
	stub := fake.RequeueAPIErrStub
	fake.recordInvocation("RequeueAPIErr", []interface{}{arg1})
	fake.requeueAPIErrMutex.Unlock()
	if stub != nil {
		fake.RequeueAPIErrStub(arg1)
	}
}

func (fake *FakeControlRequeueAPIErr) RequeueAPIErrCallCount() int {
	fake.requeueAPIErrMutex.RLock()
	defer fake.requeueAPIErrMutex.RUnlock()
	return len(fake.requeueAPIErrArgsForCall)
}

func (fake *FakeControlRequeueAPIErr) RequeueAPIErrCalls(stub func(error)) {
	fake.requeueAPIErrMutex.Lock()
	defer fake.requeueAPIErrMutex.Unlock()
	fake.RequeueAPIErrStub = stub
}

func (fake *FakeControlRequeueAPIErr) RequeueAPIErrArgsForCall(i int) error {
	fake.requeueAPIErrMutex.RLock()
	defer fake.requeueAPIErrMutex.RUnlock()
	argsForCall := fake.requeueAPIErrArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeControlRequeueAPIErr) RequeueAfter(arg1 time.Duration) {
	fake.requeueAfterMutex.Lock()
	fake.requeueAfterArgsForCall = append(fake.requeueAfterArgsForCall, struct {
		arg1 time.Duration
	}{arg1})
	stub := fake.RequeueAfterStub
	fake.recordInvocation("RequeueAfter", []interface{}{arg1})
	fake.requeueAfterMutex.Unlock()
	if stub != nil {
		fake.RequeueAfterStub(arg1)
	}
}

func (fake *FakeControlRequeueAPIErr) RequeueAfterCallCount() int {
	fake.requeueAfterMutex.RLock()
	defer fake.requeueAfterMutex.RUnlock()
	return len(fake.requeueAfterArgsForCall)
}

func (fake *FakeControlRequeueAPIErr) RequeueAfterCalls(stub func(time.Duration)) {
	fake.requeueAfterMutex.Lock()
	defer fake.requeueAfterMutex.Unlock()
	fake.RequeueAfterStub = stub
}

func (fake *FakeControlRequeueAPIErr) RequeueAfterArgsForCall(i int) time.Duration {
	fake.requeueAfterMutex.RLock()
	defer fake.requeueAfterMutex.RUnlock()
	argsForCall := fake.requeueAfterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeControlRequeueAPIErr) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	fake.requeueMutex.RLock()
	defer fake.requeueMutex.RUnlock()
	fake.requeueAPIErrMutex.RLock()
	defer fake.requeueAPIErrMutex.RUnlock()
	fake.requeueAfterMutex.RLock()
	defer fake.requeueAfterMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeControlRequeueAPIErr) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ libctrl.ControlRequeueAPIErr = new(FakeControlRequeueAPIErr)