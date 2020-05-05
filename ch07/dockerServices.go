package dockerServices

type Infos interface {
	// docker version
	Docker() string

	// uname -r
	Kernel() string

	//memory
	Memory() float64
}
