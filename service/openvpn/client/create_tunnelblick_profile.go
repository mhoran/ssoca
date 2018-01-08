package client

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"os/exec"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type CreateTunnelblickProfileOpts struct {
	SsocaExec     string
	SkipAuthRetry bool

	Directory string
	FileName  string
}

func (s Service) CreateTunnelblickProfile(opts CreateTunnelblickProfileOpts) (string, error) {
	configManager, err := s.runtime.GetConfigManager()
	if err != nil {
		return "", bosherr.WrapError(err, "Getting config manager")
	}

	ssocaExec := opts.SsocaExec
	if ssocaExec == "" {
		ssocaExec = "ssoca" // TODO ssoca.exe
	}

	ssocaExec, err = exec.LookPath(ssocaExec)
	if err != nil {
		return "", bosherr.WrapError(err, "Resolving ssoca executable")
	}

	dir := opts.Directory
	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			return "", bosherr.WrapError(err, "Getting working directory")
		}
	}

	file := opts.FileName
	if file == "" {
		file = s.runtime.GetEnvironmentName()

		if s.name != "openvpn" {
			file = fmt.Sprintf("%s-%s", file, s.name)
		}
	}

	dir = fmt.Sprintf("%s/%s.tblk", dir, file)

	dirAbs, err := s.fs.ExpandPath(dir)
	if err != nil {
		return "", bosherr.WrapError(err, "Expanding path")
	}

	err = s.fs.MkdirAll(dirAbs, 0700)
	if err != nil {
		return "", bosherr.WrapError(err, "Creating target directory")
	}

	client, err := s.GetClient(opts.SkipAuthRetry)
	if err != nil {
		return "", bosherr.WrapError(err, "Getting client")
	}

	profile, err := client.BaseProfile()
	if err != nil {
		return "", bosherr.WrapError(err, "Getting base profile")
	}

	pathConfigOvpn := fmt.Sprintf("%s/config.ovpn", dirAbs)

	err = s.fs.WriteFileString(pathConfigOvpn, profile)
	if err != nil {
		return "", bosherr.WrapError(err, "Writing config.ovpn")
	}

	err = s.fs.Chmod(pathConfigOvpn, 0400)
	if err != nil {
		return "", bosherr.WrapError(err, "Chmoding config.ovpn")
	}

	pathPreConnect := fmt.Sprintf("%s/pre-connect.sh", dirAbs)

	var preconnectScriptBuf bytes.Buffer
	err = tunnelblickPreConnectScript.Execute(
		&preconnectScriptBuf,
		struct {
			Exec        string
			Config      string
			Environment string
			Service     string
		}{
			Exec:        ssocaExec,
			Config:      configManager.GetSource(),
			Environment: s.runtime.GetEnvironmentName(),
			Service:     s.name,
		},
	)
	if err != nil {
		return "", bosherr.WrapError(err, "Generating Tunnelblick pre-connect.sh")
	}

	preconnectScript := preconnectScriptBuf.String()

	err = s.fs.WriteFileString(pathPreConnect, preconnectScript)
	if err != nil {
		return "", bosherr.WrapError(err, "Writing pre-connect.sh")
	}

	err = s.fs.Chmod(pathPreConnect, 0500)
	if err != nil {
		return "", bosherr.WrapError(err, "Chmoding pre-connect.sh")
	}

	pathInstall := fmt.Sprintf("%s/ssoca-install.sh", dirAbs)

	err = s.fs.WriteFileString(pathInstall, tunnelblickInstallScript)
	if err != nil {
		return "", bosherr.WrapError(err, "Writing ssoca-install.sh")
	}

	err = s.fs.Chmod(pathInstall, 0500)
	if err != nil {
		return "", bosherr.WrapError(err, "Chmoding ssoca-install.sh")
	}

	return dir, nil
}

var tunnelblickPreConnectScript = template.Must(template.New("script").Parse(`#!/bin/bash

set -u

REM() { /bin/echo $( date -u +"%Y-%m-%dT%H:%M:%SZ" ) "$@"; }

name="$( basename "$( dirname "$( dirname "$( dirname "$0" )" )" )" )"
shadow="$( dirname "$0" )/config.ovpn"

file="$HOME/Library/Application Support/Tunnelblick/Configurations/$name/Contents/Resources/config.ovpn"

REM "renewing profile"
sudo -Hnu "$USER" -- {{.Exec}} --config "{{.Config}}" --environment "{{.Environment}}" openvpn create-profile --service "{{.Service}}" > "$file.tmp"
exit=$?

if [[ "0" != "$exit" ]]; then
  rm "$file.tmp"

  REM "exiting with failure"
  exit $exit
fi

set -e

REM "patching profile"
echo "remap-usr1 SIGTERM" >> "$file.tmp"

REM "installing profile"
mv -f "$file.tmp" "$file"

REM "installing shadow copy"
cat "$file" > "$shadow"

REM done
`))

var tunnelblickInstallScript = `#!/bin/bash

set -eu

[ -n "${SUDO_USER:-}" ] || ( echo "This install script be run with sudo" >&2 && exit 1 )

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

name="$( basename "$DIR" )"
profileDir="$HOME/Library/Application Support/Tunnelblick/Configurations/$name"
shadowDir="/Library/Application Support/Tunnelblick/Users/$SUDO_USER/$name"

#
# profile
#

mkdir -p "$profileDir/Contents/Resources"

chown "$SUDO_USER":admin "$profileDir"
chmod 750 "$profileDir"

chown "$SUDO_USER":admin "$profileDir/Contents"
chmod 750 "$profileDir/Contents"

chown "$SUDO_USER":admin "$profileDir/Contents/Resources"
chmod 750 "$profileDir/Contents/Resources"

cp "$DIR/config.ovpn" "$profileDir/Contents/Resources/config.ovpn"
chown "$SUDO_USER":admin "$profileDir/Contents/Resources/config.ovpn"
chmod 740 "$profileDir/Contents/Resources/config.ovpn"

cp "$DIR/pre-connect.sh" "$profileDir/Contents/Resources/pre-connect.sh"
chown "$SUDO_USER":admin "$profileDir/Contents/Resources/pre-connect.sh"
chmod 750 "$profileDir/Contents/Resources/pre-connect.sh"

#
# shadow
#

mkdir -p "$shadowDir/Contents/Resources"

chown root:wheel "$shadowDir"
chmod 755 "$shadowDir"

chown root:wheel "$shadowDir/Contents"
chmod 755 "$shadowDir/Contents"

chown root:wheel "$shadowDir/Contents/Resources"
chmod 755 "$shadowDir/Contents/Resources"

cp "$DIR/config.ovpn" "$shadowDir/Contents/Resources/config.ovpn"
chown root:wheel "$shadowDir/Contents/Resources/config.ovpn"
chmod 700 "$shadowDir/Contents/Resources/config.ovpn"

cp "$DIR/pre-connect.sh" "$shadowDir/Contents/Resources/pre-connect.sh"
chown root:wheel "$shadowDir/Contents/Resources/pre-connect.sh"
chmod 700 "$shadowDir/Contents/Resources/pre-connect.sh"

#
# fyi
#

echo "The profile '${name%.tblk}' has successfully been installed."
echo "If you do not see the profile listed, try restarting Tunnelblick."
`
