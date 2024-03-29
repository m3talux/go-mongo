// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/luxation/go-mongo/v2 (interfaces: Client)

// Package mongo is a generated GoMock package.
package mongo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	mongo "github.com/luxation/go-mongo/v2"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	mongo0 "go.mongodb.org/mongo-driver/mongo"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockClient) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockClientMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClient)(nil).Connect))
}

// Delete mocks base method.
func (m *MockClient) Delete(arg0 mongo.Document) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockClientMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClient)(nil).Delete), arg0)
}

// DeleteWhere mocks base method.
func (m *MockClient) DeleteWhere(arg0 mongo.Document, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWhere", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWhere indicates an expected call of DeleteWhere.
func (mr *MockClientMockRecorder) DeleteWhere(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWhere", reflect.TypeOf((*MockClient)(nil).DeleteWhere), arg0, arg1, arg2)
}

// Disconnect mocks base method.
func (m *MockClient) Disconnect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockClientMockRecorder) Disconnect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockClient)(nil).Disconnect))
}

// FindAll mocks base method.
func (m *MockClient) FindAll(arg0 mongo.Document, arg1 primitive.M, arg2 mongo.ResultDecoder, arg3 ...*mongo.FindOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAll", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAll indicates an expected call of FindAll.
func (mr *MockClientMockRecorder) FindAll(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockClient)(nil).FindAll), varargs...)
}

// FindOne mocks base method.
func (m *MockClient) FindOne(arg0 mongo.Document, arg1 primitive.M, arg2 ...*mongo.FindOptions) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockClientMockRecorder) FindOne(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockClient)(nil).FindOne), varargs...)
}

// FindOneById mocks base method.
func (m *MockClient) FindOneById(arg0 mongo.Document, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindOneById indicates an expected call of FindOneById.
func (mr *MockClientMockRecorder) FindOneById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneById", reflect.TypeOf((*MockClient)(nil).FindOneById), arg0, arg1)
}

// GenerateUUID mocks base method.
func (m *MockClient) GenerateUUID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUUID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GenerateUUID indicates an expected call of GenerateUUID.
func (mr *MockClientMockRecorder) GenerateUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUUID", reflect.TypeOf((*MockClient)(nil).GenerateUUID))
}

// GetCollection mocks base method.
func (m *MockClient) GetCollection(arg0 mongo.Document) (*mongo0.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", arg0)
	ret0, _ := ret[0].(*mongo0.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockClientMockRecorder) GetCollection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockClient)(nil).GetCollection), arg0)
}

// GetCollectionByName mocks base method.
func (m *MockClient) GetCollectionByName(arg0 string) (*mongo0.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionByName", arg0)
	ret0, _ := ret[0].(*mongo0.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionByName indicates an expected call of GetCollectionByName.
func (mr *MockClientMockRecorder) GetCollectionByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionByName", reflect.TypeOf((*MockClient)(nil).GetCollectionByName), arg0)
}

// GetURI mocks base method.
func (m *MockClient) GetURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetURI indicates an expected call of GetURI.
func (mr *MockClientMockRecorder) GetURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURI", reflect.TypeOf((*MockClient)(nil).GetURI))
}

// HealthCheck mocks base method.
func (m *MockClient) HealthCheck() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockClientMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockClient)(nil).HealthCheck))
}

// Persist mocks base method.
func (m *MockClient) Persist(arg0 mongo.Document) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Persist indicates an expected call of Persist.
func (mr *MockClientMockRecorder) Persist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockClient)(nil).Persist), arg0)
}

// Replace mocks base method.
func (m *MockClient) Replace(arg0 mongo.Document) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replace", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replace indicates an expected call of Replace.
func (mr *MockClientMockRecorder) Replace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replace", reflect.TypeOf((*MockClient)(nil).Replace), arg0)
}

// ReplaceOrPersist mocks base method.
func (m *MockClient) ReplaceOrPersist(arg0 mongo.Document) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceOrPersist", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplaceOrPersist indicates an expected call of ReplaceOrPersist.
func (mr *MockClientMockRecorder) ReplaceOrPersist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceOrPersist", reflect.TypeOf((*MockClient)(nil).ReplaceOrPersist), arg0)
}

// Update mocks base method.
func (m *MockClient) Update(arg0 mongo.Document, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockClientMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClient)(nil).Update), arg0, arg1, arg2)
}

// UpdateMany mocks base method.
func (m *MockClient) UpdateMany(arg0 mongo.Document, arg1 primitive.M, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMany", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMany indicates an expected call of UpdateMany.
func (mr *MockClientMockRecorder) UpdateMany(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMany", reflect.TypeOf((*MockClient)(nil).UpdateMany), arg0, arg1, arg2)
}

// UpdateWhere mocks base method.
func (m *MockClient) UpdateWhere(arg0 mongo.Document, arg1 primitive.M, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWhere", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWhere indicates an expected call of UpdateWhere.
func (mr *MockClientMockRecorder) UpdateWhere(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWhere", reflect.TypeOf((*MockClient)(nil).UpdateWhere), arg0, arg1, arg2)
}

// WithContext mocks base method.
func (m *MockClient) WithContext(arg0 context.Context) mongo.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext", arg0)
	ret0, _ := ret[0].(mongo.Client)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockClientMockRecorder) WithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockClient)(nil).WithContext), arg0)
}
