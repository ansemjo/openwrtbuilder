# openwrtbuilder

A script to automate building firmware images for OpenWRT routers with the
imagebuilder SDK's provided for each target. Internally it generates a
Dockerfile on the fly and builds that with `img`, which is a rootless and
daemonless `buildkit` tool.

The container-based build enforces strict reproduceability and the layer caching
can accelerate similar builds for the same architecture.

## INSTALLATION

A release of `genuinetools/img` which includes commit
`0ec03c62e5146114e30e1a4721c3826ec9fba2b8` is required in your `PATH`. At the
time of this writing, you'll need to compile from `master` manually.

Head over to [genuinetools/img#installation][install] for installation
instructions.

[install]: https://github.com/genuinetools/img#installation

Apart from that it's just a shell script which uses a few basic system commands
that should be available on a common GNU/Linux system.

```
curl -LO https://raw.githubusercontent.com/ansemjo/openwrtbuilder/master/openwrtbuilder
bash ./openwrtbuilder --help
```

## USAGE

The build target and profile are configured with commandline flags. Additional
packages to be installed and an optional filesystem tree to be included can also
be specified. The options are based on the capabilities of the
[OpenWRT imagebuilder](https://openwrt.org/docs/guide-user/additional-software/imagebuilder).

You can use `-R`, `-T` and `-t ... -P` to list available releases, targets and
profiles for a particular target respectively.

| flag | description |
|------|-------------|
| `-r RELEASE` | the release version, can be `snapshot` or a version like `18.06.4` |
| `-t TARGET` | target architecture and chipset, e.g. `ar71xx/generic` |
| `-p PROFILE` | profile for a particular device, e.g. in the above target: `archer-c7-v2` |
| `-i PKGS` | include or exclude specific packages, e.g. `-ppp luci-ssl wireguard` |
| `-d DESTDIR` | destination directory for the build firmware archive |
| `-f DIRECTORY` | directory with a filesystem tree to include in firmware |
| `-c CONFIG` | source config options from this bash-compatible file |

When building a specific firmware image, the configuration is printed on the
terminal and can be copied to a file for later usage with the `-c` flag. See the
[configs](configs/) directory for examples.

## EXAMPLE

List available profiles for the `ar71xx/generic` target in the `18.06.4`
release:

```
$ openwrtbuilder -t ar71xx/generic -r 18.06.4 -P
available profiles for 18.06.4/ar71xx/generic:
 - A60
 - ALFAAP120C
 - ALFAAP96
 - ALFANX
/* ... */
 - archer-c60-v2
 - archer-c7-v1
 - archer-c7-v2
 - archer-c7-v2-il
/* ... */
 - zbt-we1526
 - ZCN1523H28
 - ZCN1523H516
```

Build a firmware image for the TP-Link Archer C7 v2:

```
$ openwrtbuilder -t ar71xx/generic -r 18.06.4 -p archer-c7-v2

# openwrtbuilder configuration
MIRROR='https://mirror.kumi.systems/openwrt/'
RELEASE='18.06.4'
TARGET='ar71xx/generic'
PROFILE='archer-c7-v2'
PACKAGES=
FILES=''

Building image
Setting up the rootfs... this may take a bit.
[+] Building 73.1s (13/13) FINISHED
 => [internal] load build definition from Dockerfile                                                 0.1s
 => => transferring dockerfile: 1.49kB                                                               0.0s
 => [internal] load .dockerignore                                                                    0.1s
 => => transferring context: 2B                                                                      0.0s
 => [internal] load metadata for docker.io/library/debian:stable                                     2.0s
 => [build 1/8] FROM docker.io/library/debian:stable@sha256:6a3ead8cbca86c3c28c5f32d250df9203f7cb93  0.0s
 => => resolve docker.io/library/debian:stable@sha256:6a3ead8cbca86c3c28c5f32d250df9203f7cb939ed07a  0.0s
 => CACHED [build 2/8] RUN apt-get update && apt-get install -y   build-essential libncurses5-dev z  0.0s
 => CACHED [build 3/8] WORKDIR /download                                                             0.0s
 => [build 4/8] RUN for file in   "https://mirror.kumi.systems/openwrt//releases/18.06.4/targets/a  11.8s
 => [build 5/8] RUN curl "https://git.openwrt.org/?p=keyring.git;a=blob_plain;f=gpg/626471F1.asc" |  5.8s
 => [build 6/8] WORKDIR /build                                                                       2.1s
 => [build 7/8] RUN tar xf "/download/openwrt-imagebuilder-18.06.4-ar71xx-generic.Linux-x86_64.tar.  9.4s
 => [build 8/8] RUN make image PROFILE='archer-c7-v2' PACKAGES=                                     40.9s
 => [stage-1 1/1] COPY --from=build /build/bin/targets/ar71xx/generic /                              0.1s
 => exporting to client                                                                              0.5s
 => => sending tarball                                                                               0.5s
Successfully built image
$ ls
openwrt-ar71xx-generic-archer-c7-v2-18.06.4.tar
$ tar tf openwrt-ar71xx-generic-archer-c7-v2-18.06.4.tar
openwrt-18.06.4-ar71xx-generic-archer-c7-v2-squashfs-factory-eu.bin
openwrt-18.06.4-ar71xx-generic-archer-c7-v2-squashfs-factory-us.bin
openwrt-18.06.4-ar71xx-generic-archer-c7-v2-squashfs-factory.bin
openwrt-18.06.4-ar71xx-generic-archer-c7-v2-squashfs-sysupgrade.bin
openwrt-18.06.4-ar71xx-generic-device-archer-c7-v2.manifest
openwrt-18.06.4-ar71xx-generic-root.squashfs
openwrt-18.06.4-ar71xx-generic-uImage-lzma.bin
openwrt-18.06.4-ar71xx-generic-vmlinux-lzma.elf
openwrt-18.06.4-ar71xx-generic-vmlinux.bin
openwrt-18.06.4-ar71xx-generic-vmlinux.elf
openwrt-18.06.4-ar71xx-generic-vmlinux.lzma
sha256sums
```
