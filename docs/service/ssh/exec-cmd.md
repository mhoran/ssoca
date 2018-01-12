# `ssoca ssh exec ...`

Connect to a remote SSH server

    Usage:
      ssoca [OPTIONS] ssh exec [exec-OPTIONS] [HOST]
    
    Application Options:
          --config=              Configuration file path (default: ~/.config/ssoca/config) [$SSOCA_CONFIG]
      -e, --environment=         Environment name [$SSOCA_ENVIRONMENT]
          --log-level=           Log level (default: WARN) [$SSOCA_LOG_LEVEL]
    
    Help Options:
      -h, --help                 Show this help message
    
    [exec command options]
          -s, --service=         Service name (default: ssh) [$SSOCA_SERVICE]
              --skip-auth-retry  Skip interactive authentication retries when logged out
              --exec=            Path to the ssh binary
              --opt=             Additional option to pass to ssh
    