# Release tools issue 

Originally found trying to run `go run mage.go` in a checkout of [sigs.k8s.io/bom](https://github.com/kubernetes-sigs/bom) 

## Reproducer setup

1. Initialized this repo and added the deps for mage and release-utils

```sh
rm -f "${GOPATH}/bin/zeitgeist"
go mod init github.com/pnasrat/magedep
go get github.com/uwu-tools/magex@v0.10.0
go get sigs.k8s.io/release-utils@v0.7.6
```

2. Created a [zero install](https://magefile.org/zeroinstall/) `mage.go` and `magefile.go`
3. Create a empty dependencies list `echo dependencies: > dependencies.yaml`
4. Run `go run mage.go` twice

## Expected outcome 

zeitgeist gets downloaded and the mage command returns no errors

## Actual result

Error that version does not contain a 3-part semver value

```
go run mage.go  
Running external dependency checks...
 __  /   ____|  _ _|  __ __|    ___|   ____|  _ _|    ___|   __ __|
    /    __|      |      |     |       __|      |   \___ \      |
   /     |        |      |     |   |   |        |         |     |
 ____|  _____|  ___|    _|    \____|  _____|  ___|  _____/     _|
zeitgeist: Zeitgeist is a language-agnostic dependency checker

GitVersion:    v0.4.1
GitCommit:     unknown
GitTreeState:  unknown
BuildDate:     unknown
GoVersion:     go1.21.3
Compiler:      gc
Platform:      linux/amd64

Error: ensuring zeitgeist is installed: ensuring package: the output of /home/pnasrat/workspace/bin/zeitgeist version did not include a 3-part semver value: 
exit status 1
```


## Analysis

zeitgeist@0.4.1 version writes it's version info to stderr which is not what [uwu-tools/magex]( github.com/uwu-tools/magex) expects in `GetCommandVersion` see https://github.com/uwu-tools/magex/blob/v0.10.0/pkg/install.go#L288

This can be demonstrated with `$GOPATH/bin/zeitgeist version 2>/dev/null`

This is fixed in zeitgeist from [PR 544](https://github.com/kubernetes-sigs/zeitgeist/pull/544/commits) onwards commit: https://github.com/kubernetes-sigs/zeitgeist/commit/5108cb4e034ebefd17b9e827243174344c010379

This can be verified 

```
git checkout 5108cb4
go run main.go version 2>/dev/null
 __  /   ____|  _ _|  __ __|    ___|   ____|  _ _|    ___|   __ __|
    /    __|      |      |     |       __|      |   \___ \      |
   /     |        |      |     |   |   |        |         |     |
 ____|  _____|  ___|    _|    \____|  _____|  ___|  _____/     _|
zeitgeist: Zeitgeist is a language-agnostic dependency checker

GitVersion:    devel
GitCommit:     unknown
GitTreeState:  unknown
BuildDate:     unknown
GoVersion:     go1.21.3
Compiler:      gc
Platform:      linux/amd64
```

Versus the revision  before https://github.com/kubernetes-sigs/zeitgeist/commit/a98a760d114d56f73f974794627fdd24132ce2df which outputs to stderr

```
git checkout a98a760
go run main.go 
```
