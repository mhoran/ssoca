---
title: ssoca openvpn create-onc-profile
aliases:
- /service/openvpn/create-onc-profile-cmd/
---

# `ssoca openvpn create-onc-profile ...`

Create an ONC profile

    Usage:
      ssoca [OPTIONS] openvpn [openvpn-OPTIONS] create-onc-profile [create-onc-profile-OPTIONS]
    
    Application Options:
          --config=              Configuration file path (default: ~/.config/ssoca/config) [$SSOCA_CONFIG]
      -e, --environment=         Environment name [$SSOCA_ENVIRONMENT]
          --log-level=           Log level (default: WARN) [$SSOCA_LOG_LEVEL]
    
    Help Options:
      -h, --help                 Show this help message
    
    [openvpn command options]
    
        Establish OpenVPN connections to remote servers:
          -s, --service=         Service name (default: openvpn) [$SSOCA_SERVICE]
    
    [create-onc-profile command options]
              --skip-auth-retry  Skip interactive authentication retries when logged out
              --name=            Specific network configuration name to use
    

## Usage Details

Chrome OS uses [Open Network Configuration](https://chromium.googlesource.com/chromium/src/+/master/components/onc/docs/onc_spec.md) files for VPN settings. Use the `create-onc-profile` command to generate a compatible file. Once generated and saved in a file:

 * Go to `chrome://net-internals/#chromeos`
 * From [Import ONC file], click [Choose File]
    * Select your `*.onc` file
    * Click [Open]
    * No confirmation of success or failure will be given
 * Go to `chrome://settings/networks?type=VPN` and verify the VPN is present


## Troubleshooting

If the VPN connection is not added or updated, visit `chrome://system` and review `Profile[0] chrome_user_log` for details (search for `onc_validator.cc` or `onc`).

If the VPN connection will not successfully connect, visit `chrome://system` and review `netlog` for `openvpn` for connection details (search for `openvpn`).
