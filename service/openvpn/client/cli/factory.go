package cli

import (
	"os"

	"github.com/dpb587/ssoca/client"
	clientcmd "github.com/dpb587/ssoca/client/cmd"
	svc "github.com/dpb587/ssoca/service/openvpn/client"
)

type Commands struct {
	BaseProfile              BaseProfile              `command:"base-profile" description:"Show the base connection profile of the OpenVPN server"`
	Connect                  Connect                  `command:"connect" description:"Connect to a remote OpenVPN server"`
	CreateProfile            CreateProfile            `command:"create-profile" description:"Create and sign an OpenVPN configuration profile"`
	CreateTunnelblickProfile CreateTunnelblickProfile `command:"create-tunnelblick-profile" description:"Create a Tunnelblick profile"`
}

func CreateCommands(runtime client.Runtime, service svc.Service) *Commands {
	cmd := clientcmd.ServiceCommand{
		Runtime:     runtime,
		ServiceName: service.Type(),
	}

	return &Commands{
		BaseProfile: BaseProfile{
			ServiceCommand: cmd,
			Service:        service,
		},
		Connect: Connect{
			ServiceCommand: cmd,
			Service:        service,
		},
		CreateProfile: CreateProfile{
			ServiceCommand: cmd,
			Service:        service,
		},
		CreateTunnelblickProfile: CreateTunnelblickProfile{
			ServiceCommand: cmd,
			Service:        service,
			Name:           service.Type(),
			SsocaExec:      os.Args[0],
		},
	}
}