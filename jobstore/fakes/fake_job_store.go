// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/Bo0mer/os-agent/jobstore"
	"github.com/Bo0mer/os-agent/model"
)

type FakeJobStore struct {
	SetStub        func(model.Job)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		arg1 model.Job
	}
	GetStub        func(string) (model.Job, bool)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 model.Job
		result2 bool
	}
}

func (fake *FakeJobStore) Set(arg1 model.Job) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		arg1 model.Job
	}{arg1})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(arg1)
	}
}

func (fake *FakeJobStore) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *FakeJobStore) SetArgsForCall(i int) model.Job {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].arg1
}

func (fake *FakeJobStore) Get(arg1 string) (model.Job, bool) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeJobStore) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeJobStore) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1
}

func (fake *FakeJobStore) GetReturns(result1 model.Job, result2 bool) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 model.Job
		result2 bool
	}{result1, result2}
}

var _ jobstore.JobStore = new(FakeJobStore)
