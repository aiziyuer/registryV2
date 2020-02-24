registryV2 ![Build Test](https://github.com/aiziyuer/registryV2/workflows/Build%20Test/badge.svg)
---

`registryV2` is different with docker, all func is via api, no image donwload any more,
it may improve the speed to operate orci image, just have fun.


## ⚙️ Installation

``` bash
CGO_ENABLED=0 \
GOBIN=/usr/bin \
go get -u -v github.com/aiziyuer/registryV2
```

## ⚡️ Quickstart

```
# view image manifest( more details: -o json )
➜  ~ registryV2 manifest centos:7
# .
# └── [D] {manifest sha256sum} {manifest size}
#    ├── [P {os platform}] {submanifest sha256sum} {submanifest size}
#    │   ├── [C] {config sha256sum} {config size}
#    │   └── [L   {layer index}] {layer sha256sum} {layer size}
#    ...
# eg:
# .
# └── [D] sha256:4a701376d03f6b39b8c2a8f4a8e499441b0d567f9ab9d58e4991de4472fb813c 405024774
#    ├── [P linux/amd64] sha256:285bc3161133ec01d8ca8680cd746eecbfdbc1faa6313bd863151c4b26d7e5a5 75782895
#    │   ├── [C] sha256:5e35e350aded98340bc8fcb0ba392d809c807bc3eb5c618d4a0674d98d88bccd 2183
#    │   └── [L   1] sha256:ab5ef0e5819490abe86106fd9f4381123e37a03e80e650be39f7938d30ecb530 75780712
#    ├── [P linux/arm] sha256:9fd67116449f225c6ef60d769b5219cf3daa831c5a0a6389bbdd7c952b7b352d 70031570
#    │   ├── [C] sha256:8c52f2d0416faa8009082cf3ebdea85b3bc1314d97925342be83bc9169178efe 2181
#    │   └── [L   1] sha256:193bcbf05ff9ae85ac1a58cacd9c07f8f4297dc648808c347cceb3797ae603af 70029389
#    ├── [P linux/arm64] sha256:fc5a0399d94336d15305d4d43754cd3c57808123cc67a578687748734af8f06b 103621812
#    │   ├── [C] sha256:4dfd99be812b186ee379da6f8e270b2eca37dca5a046d61c216c2a6b630712c7 2183
#    │   └── [L   1] sha256:3f2696f8166ff69dd0c116674b19eebd351ed3fc4111a42dbd57c673601c725d 103619629
#    ├── [P linux/386] sha256:1f832b4e3b9ddf67fd77831cdfb591ce5e968548a01581672e5f6b32ce1212fe 75656436
#    │   ├── [C] sha256:fe70670fcbec5e3b3081c6800cb531002474c36563689b450d678a34a89b62c3 2337
#    │   └── [L   1] sha256:39016a8400a36ce04799adba71f8678ae257d9d8dba638d81b8c5755f01fe213 75654099
#    └── [P linux/ppc64le] sha256:4b8a19661b7d770bbab54747541812b581bfd4944ef64b58653d4fc77f3e1ebc 79932061
#        ├── [C] sha256:ec71c93f9d8cfde5403701971e08e1f7c197ce6492977915b11ec4b65f63699a 2185
#        └── [L   1] sha256:23bd9eb8fdc010dbd36575046a8c42317f78a9926da949829722cfc815d46cf9 79929876
```

## 🤖 Benchmarks

## 🎯 Features

## ⭐️ FAQ

