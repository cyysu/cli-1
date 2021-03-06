// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/api/uaa/constant"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeAuthActor struct {
	AuthenticateStub        func(ID string, secret string, grantType constant.GrantType) error
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		ID        string
		secret    string
		grantType constant.GrantType
	}
	authenticateReturns struct {
		result1 error
	}
	authenticateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthActor) Authenticate(ID string, secret string, grantType constant.GrantType) error {
	fake.authenticateMutex.Lock()
	ret, specificReturn := fake.authenticateReturnsOnCall[len(fake.authenticateArgsForCall)]
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		ID        string
		secret    string
		grantType constant.GrantType
	}{ID, secret, grantType})
	fake.recordInvocation("Authenticate", []interface{}{ID, secret, grantType})
	fake.authenticateMutex.Unlock()
	if fake.AuthenticateStub != nil {
		return fake.AuthenticateStub(ID, secret, grantType)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.authenticateReturns.result1
}

func (fake *FakeAuthActor) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakeAuthActor) AuthenticateArgsForCall(i int) (string, string, constant.GrantType) {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return fake.authenticateArgsForCall[i].ID, fake.authenticateArgsForCall[i].secret, fake.authenticateArgsForCall[i].grantType
}

func (fake *FakeAuthActor) AuthenticateReturns(result1 error) {
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthActor) AuthenticateReturnsOnCall(i int, result1 error) {
	fake.AuthenticateStub = nil
	if fake.authenticateReturnsOnCall == nil {
		fake.authenticateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.authenticateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuthActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAuthActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.AuthActor = new(FakeAuthActor)
