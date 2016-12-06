// +build windows

package sdk

import (
	"net"

	"github.com/docker/go-connections/sockets"
	"github.com/Microsoft/go-winio"
)

func newWindowsListener(address, pluginName string, pipeConfig *WindowsPipeConfig) func() (net.Listener, string, string, error) {
	return func() (net.Listener, string, string, error) {
		winioPipeConfig := winio.PipeConfig{
			SecurityDescriptor: pipeConfig.SecurityDescriptor,
			MessageMode:        true,
			InputBufferSize:    pipeConfig.InBufferSize,
			OutputBufferSize:   pipeConfig.OutBufferSize,
		}
		listener, err := sockets.NewWindowsSocket(address, &winioPipeConfig)
		if err != nil {
			return nil, "", "", err
		}
		spec, err := writeSpec(pluginName, listener.Addr().String(), ProtoNamedPipe)
		if err != nil {
			return nil, "", "", err
		}
		return listener, address, spec, nil
	}
}
