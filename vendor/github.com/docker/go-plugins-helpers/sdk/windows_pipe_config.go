package sdk

// WindowsPipeConfig is a helper structure for configuring named pipe parameters on Windows.
type WindowsPipeConfig struct {
	SecurityDescriptor string
	InBufferSize int32
	OutBufferSize int32
}