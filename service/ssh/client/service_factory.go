package client

import (
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	"github.com/dpb587/ssoca/client"
)

type ServiceFactory struct {
	runtime   client.Runtime
	fs        boshsys.FileSystem
	cmdRunner boshsys.CmdRunner
}

func NewServiceFactory(runtime client.Runtime, fs boshsys.FileSystem, cmdRunner boshsys.CmdRunner) ServiceFactory {
	return ServiceFactory{
		runtime:   runtime,
		fs:        fs,
		cmdRunner: cmdRunner,
	}
}

func (sf ServiceFactory) New(name string) Service {
	return NewService(name, sf.runtime, sf.fs, sf.cmdRunner)
}
