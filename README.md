# version
[![go report](https://goreportcard.com/badge/github.com/ivan-dyachenko/version)](https://goreportcard.com/report/github.com/ivan-dyachenko/version)

CLI command to verify versions and version constraints build based [hashicorp/go-version](https://github.com/hashicorp/go-version).

Versions used with go-version must follow [SemVer](http://semver.org/).

# Install
If you are OSX user, you can use Homebrew:

```
$ brew tap ivan-dyachenko/version
$ brew install version
```

For other operating systems check [releases](https://github.com/ivan-dyachenko/version/releases)

# Basic usage

```
version ">=1.0, <2.0" "1.7"
go version | version ">=1.9"
```

# Issues and Contributing
If you find an issue with this library, please report an issue. If you'd like, we welcome any contributions. Fork this library and submit a pull request.
