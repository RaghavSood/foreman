package supervisor

import (
	"errors"
	"fmt"
	"github.com/kolo/xmlrpc"
)

const (
	apiVersion string = "3.0"
)

type Client struct {
	RpcClient  *xmlrpc.Client
	ApiVersion string
}

// NewClient creates a new supervisor RPC client.
func NewClient(url string) (client Client, err error) {
	var rpc *xmlrpc.Client
	if rpc, err = xmlrpc.NewClient(url, nil); err != nil {
		return
	}
	client = Client{rpc, apiVersion}

	version := ""
	version, err = client.APIVersion()
	if err != nil {
		return
	}
	if version != apiVersion {
		err = errors.New(fmt.Sprintf("expected Supervisor API version %s, got %s instead", apiVersion, version))
		return
	}
	return
}

// Close the client.
func (client Client) Close() error {
	return client.RpcClient.Close()
}

func (client Client) Version() (version string, err error) {
	err = client.RpcClient.Call("supervisor.getSupervisorVersion", nil, &version)
	return
}

func (client Client) APIVersion() (version string, err error) {
	err = client.RpcClient.Call("supervisor.getAPIVersion", nil, &version)
	return
}

func (client Client) PID() (pid int64, err error) {
	err = client.RpcClient.Call("supervisor.getPID", nil, &pid)
	return
}

func (client Client) ID() (id string, err error) {
	err = client.RpcClient.Call("supervisor.getIdentification", nil, &id)
	return
}

func (client Client) StartProcess(name string, wait bool) (result bool, err error) {
	params := xmlrpcParams(name, wait)
	err = client.RpcClient.Call("supervisor.startProcess", params, &result)
	return
}

func (client Client) StopProcess(name string, wait bool) (result bool, err error) {
	params := xmlrpcParams(name, wait)
	err = client.RpcClient.Call("supervisor.stopProcess", params, &result)
	return
}

func xmlrpcParams(params ...interface{}) []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(params))
	for i, v := range params {
		interfaceSlice[i] = v
	}
	return interfaceSlice
}
