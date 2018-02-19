// Generated by go-rpcgen. Do not modify.

package stm

import (
	"net/rpc"
	"time"
)

// UserInterfaceService is generated service for UserInterface interface.
type UserInterfaceService struct {
	impl UserInterface
}

// NewUserInterfaceService creates a new UserInterfaceService instance.
func NewUserInterfaceService(impl UserInterface) *UserInterfaceService {
	return &UserInterfaceService{impl}
}

// RegisterUserInterfaceService registers impl in server.
func RegisterUserInterfaceService(server *rpc.Server, impl UserInterface) error {
	return server.RegisterName("Interface", NewUserInterfaceService(impl))
}

// UserInterfacePowerTickRequest is a helper structure for PowerTick method.
type UserInterfacePowerTickRequest struct {
	D time.Duration
}

// UserInterfacePowerTickResponse is a helper structure for PowerTick method.
type UserInterfacePowerTickResponse struct {
}

// PowerTick is RPC implementation of PowerTick calling it.
func (s *UserInterfaceService) PowerTick(request *UserInterfacePowerTickRequest, response *UserInterfacePowerTickResponse) (err error) {
	err = s.impl.PowerTick(request.D)
	return
}

// UserInterfaceDUTRequest is a helper structure for DUT method.
type UserInterfaceDUTRequest struct {
}

// UserInterfaceDUTResponse is a helper structure for DUT method.
type UserInterfaceDUTResponse struct {
}

// DUT is RPC implementation of DUT calling it.
func (s *UserInterfaceService) DUT(request *UserInterfaceDUTRequest, response *UserInterfaceDUTResponse) (err error) {
	err = s.impl.DUT()
	return
}

// UserInterfaceTSRequest is a helper structure for TS method.
type UserInterfaceTSRequest struct {
}

// UserInterfaceTSResponse is a helper structure for TS method.
type UserInterfaceTSResponse struct {
}

// TS is RPC implementation of TS calling it.
func (s *UserInterfaceService) TS(request *UserInterfaceTSRequest, response *UserInterfaceTSResponse) (err error) {
	err = s.impl.TS()
	return
}

// UserInterfaceGetCurrentRequest is a helper structure for GetCurrent method.
type UserInterfaceGetCurrentRequest struct {
}

// UserInterfaceGetCurrentResponse is a helper structure for GetCurrent method.
type UserInterfaceGetCurrentResponse struct {
	Value int
}

// GetCurrent is RPC implementation of GetCurrent calling it.
func (s *UserInterfaceService) GetCurrent(request *UserInterfaceGetCurrentRequest, response *UserInterfaceGetCurrentResponse) (err error) {
	response.Value, err = s.impl.GetCurrent()
	return
}

// UserInterfaceStartCurrentRecordRequest is a helper structure for StartCurrentRecord method.
type UserInterfaceStartCurrentRecordRequest struct {
	Samples  int
	Interval time.Duration
}

// UserInterfaceStartCurrentRecordResponse is a helper structure for StartCurrentRecord method.
type UserInterfaceStartCurrentRecordResponse struct {
}

// StartCurrentRecord is RPC implementation of StartCurrentRecord calling it.
func (s *UserInterfaceService) StartCurrentRecord(request *UserInterfaceStartCurrentRecordRequest, response *UserInterfaceStartCurrentRecordResponse) (err error) {
	err = s.impl.StartCurrentRecord(request.Samples, request.Interval)
	return
}

// UserInterfaceStopCurrentRecordRequest is a helper structure for StopCurrentRecord method.
type UserInterfaceStopCurrentRecordRequest struct {
}

// UserInterfaceStopCurrentRecordResponse is a helper structure for StopCurrentRecord method.
type UserInterfaceStopCurrentRecordResponse struct {
}

// StopCurrentRecord is RPC implementation of StopCurrentRecord calling it.
func (s *UserInterfaceService) StopCurrentRecord(request *UserInterfaceStopCurrentRecordRequest, response *UserInterfaceStopCurrentRecordResponse) (err error) {
	err = s.impl.StopCurrentRecord()
	return
}

// UserInterfaceGetCurrentRecordRequest is a helper structure for GetCurrentRecord method.
type UserInterfaceGetCurrentRecordRequest struct {
}

// UserInterfaceGetCurrentRecordResponse is a helper structure for GetCurrentRecord method.
type UserInterfaceGetCurrentRecordResponse struct {
	Samples []int
}

// GetCurrentRecord is RPC implementation of GetCurrentRecord calling it.
func (s *UserInterfaceService) GetCurrentRecord(request *UserInterfaceGetCurrentRecordRequest, response *UserInterfaceGetCurrentRecordResponse) (err error) {
	response.Samples, err = s.impl.GetCurrentRecord()
	return
}

// UserInterfaceClient is generated client for UserInterface interface.
type UserInterfaceClient struct {
	client *rpc.Client
}

// DialUserInterfaceClient connects to addr and creates a new UserInterfaceClient instance.
func DialUserInterfaceClient(addr string) (*UserInterfaceClient, error) {
	client, err := rpc.Dial("tcp", addr)
	return &UserInterfaceClient{client}, err
}

// NewUserInterfaceClient creates a new UserInterfaceClient instance.
func NewUserInterfaceClient(client *rpc.Client) *UserInterfaceClient {
	return &UserInterfaceClient{client}
}

// Close terminates the connection.
func (_c *UserInterfaceClient) Close() error {
	return _c.client.Close()
}

// PowerTick is part of implementation of UserInterface calling corresponding method on RPC server.
func (_c *UserInterfaceClient) PowerTick(d time.Duration) (err error) {
	_request := &UserInterfacePowerTickRequest{d}
	_response := &UserInterfacePowerTickResponse{}
	err = _c.client.Call("Interface.PowerTick", _request, _response)
	return err
}

// DUT is part of implementation of UserInterface calling corresponding method on RPC server.
func (_c *UserInterfaceClient) DUT() (err error) {
	_request := &UserInterfaceDUTRequest{}
	_response := &UserInterfaceDUTResponse{}
	err = _c.client.Call("Interface.DUT", _request, _response)
	return err
}

// TS is part of implementation of UserInterface calling corresponding method on RPC server.
func (_c *UserInterfaceClient) TS() (err error) {
	_request := &UserInterfaceTSRequest{}
	_response := &UserInterfaceTSResponse{}
	err = _c.client.Call("Interface.TS", _request, _response)
	return err
}

// GetCurrent is part of implementation of UserInterface calling corresponding method on RPC server.
func (_c *UserInterfaceClient) GetCurrent() (value int, err error) {
	_request := &UserInterfaceGetCurrentRequest{}
	_response := &UserInterfaceGetCurrentResponse{}
	err = _c.client.Call("Interface.GetCurrent", _request, _response)
	return _response.Value, err
}

// StartCurrentRecord is part of implementation of UserInterface calling corresponding method on RPC server.
func (_c *UserInterfaceClient) StartCurrentRecord(samples int, interval time.Duration) (err error) {
	_request := &UserInterfaceStartCurrentRecordRequest{samples, interval}
	_response := &UserInterfaceStartCurrentRecordResponse{}
	err = _c.client.Call("Interface.StartCurrentRecord", _request, _response)
	return err
}

// StopCurrentRecord is part of implementation of UserInterface calling corresponding method on RPC server.
func (_c *UserInterfaceClient) StopCurrentRecord() (err error) {
	_request := &UserInterfaceStopCurrentRecordRequest{}
	_response := &UserInterfaceStopCurrentRecordResponse{}
	err = _c.client.Call("Interface.StopCurrentRecord", _request, _response)
	return err
}

// GetCurrentRecord is part of implementation of UserInterface calling corresponding method on RPC server.
func (_c *UserInterfaceClient) GetCurrentRecord() (samples []int, err error) {
	_request := &UserInterfaceGetCurrentRecordRequest{}
	_response := &UserInterfaceGetCurrentRecordResponse{}
	err = _c.client.Call("Interface.GetCurrentRecord", _request, _response)
	return _response.Samples, err
}
