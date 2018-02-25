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

```bash
version ">=1.0, <2.0" "1.7"
go version | version ">=1.9"
```

# Usage in the bash scripts

`version -b "..."` returns `true|false` that can be used in bash scripts

Check `git` version by using [pipeline](https://en.wikipedia.org/wiki/Pipeline_(Unix):
```bash
#!/bin/bash

# version supports Pipeline (Unix)
if `git version | version -b ">2.15.0"`; then
  echo "git version > 2.15.0"
else
  echo "please install git > 2.15.0"
fi
```

Check `gcc` version:
```bash
#!/bin/bash

if `version -b ">=9.0.0" "$(gcc --version)"`; then
  echo "gcc version satisfies constraints >=9.0.0"
else
  echo "gcc version doesn't satisfies constraints >=9.0.0"
fi
```

# Issues and Contributing
If you find an issue with this library, please report an issue. If you'd like, we welcome any contributions. Fork this library and submit a pull request.
