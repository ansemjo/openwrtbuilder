# openwrtbuilder

A script to automate building OpenWRT firmware images for routers with the imagebuilder SDKs provided for each target. Internally it generates a Dockerfile on the fly and builds that with a buildkit-capable Docker runtime. Due to efficient layer caching, many firmwares for similar targets can be built quickly, because the imagebuilder only needs to be downloaded once.

My personal use-case was that I always wanted to install a handful of packages like `luci-ssl` and `wireguard-tools`. And honestly .. it was easier to just learn to use the imagebuilder than reinstall them each time after a sysupgrade.

## INSTALLATION

In a previous version, this script required `genuinetools/img` because it was one of the only good possibilities for a rootless and buildkit-capable builder. Nowadays any modern `docker` binary will do. Whether you chose to use a rootless flavour is up to you but I strongly recommend it. So head over to [docs.docker.com](https://docs.docker.com/engine/install/ubuntu/) for installation instructions.

Apart from that, it's just a shell script which uses a few basic system commands that should readily be available on a common Linux system with GNU coreutils.

```
curl -LO https://raw.githubusercontent.com/ansemjo/openwrtbuilder/main/openwrtbuilder
bash ./openwrtbuilder --help
```

## USAGE

The build target and profile are configured with commandline flags. Additional packages to be installed and an optional filesystem tree to be included can also be specified. The options are based on the capabilities of the [OpenWRT imagebuilder](https://openwrt.org/docs/guide-user/additional-software/imagebuilder).

You can use `-R`, `-T` and `-t ... -P` to list available releases, targets and profiles for a particular target respectively.

| flag | description |
|------|-------------|
| `-R` | list available releases |
| `-r RELEASE` | the release version, can be `snapshot` or a version like `18.06.4` |
| `-T` | list available targets for chosen release (latest by default) |
| `-t TARGET` | target architecture and chipset, e.g. `ath79/generic` |
| `-t TARGET -P` | list available profiles (specific devices) for a given `TARGET` |
| `-p PROFILE` | profile for a particular device, e.g. in the above target: `tplink_archer-c7-v2` |
| `-i PKGS` | include or exclude specific packages,<br />e.g. exclude `ppp` and include `luci-ssl` with `-i "-ppp luci-ssl"` |
| `-d DESTDIR` | destination directory for the built firmware archive<br />a concatenation of release version, target, profile and timestamp by default |
| `-f DIRECTORY` | directory with a filesystem tree to include in firmware<br />can be used for configuration files like `/etc/profile` |
| `-c CONFIG` | source config options from this bash-compatible file<br />by default, `owrtbuildconf` in the current directory is used automatically, if found |

When building a specific firmware image, the configuration is printed on the terminal and can be copied to a file for later usage with the `-c` flag. See the [configs](configs/) directory for examples.

## EXAMPLES

List available profiles for the `ath79/generic` target in the `latest` release:

```
$ openwrtbuilder -r latest -t ath79/generic -P
available profiles for 22.03.3/ath79/generic:
 - 8dev_carambola2
 - 8dev_lima
/* ... */
 - tplink_archer-c7-v1
 - tplink_archer-c7-v2
/* ... */
 - ziking_cpe46b
 - zyxel_nbg6616
```

Build a firmware image for the TP-Link Archer C7 v2:

```
$ openwrtbuilder -r latest -t ath79/generic -p tplink_archer-c7-v2

# openwrtbuilder configuration
MIRROR='https://downloads.openwrt.org'
RELEASE='22.03.3'
TARGET='ath79/generic'
PROFILE='tplink_archer-c7-v2'
PACKAGES=
FILES=''

# build openwrtbuilder image ...
[+] Building 18.5s (5/6)                                                                                        
 => [internal] load build definition from Dockerfile                                                       0.0s
 => => transferring dockerfile: 497B                                                                       0.0s
 => [internal] load .dockerignore                                                                          0.0s
 => => transferring context: 2B                                                                            0.0s
 => [internal] load metadata for docker.io/library/debian:stable                                           2.3s
 => [1/3] FROM docker.io/library/debian:stable@sha256:12931ad2bfd4a9609cf8ef7898f113d67dce8058f0c27f01c9  10.3s
 => => resolve docker.io/library/debian:stable@sha256:12931ad2bfd4a9609cf8ef7898f113d67dce8058f0c27f01c90  0.0s
/* ... */
```

This will result in a directory with factory and sysupgrade images, depending on the specific profile:

```
$ ls openwrt-ath79-generic-tplink_archer-c7-v2-22.03.3/
openwrt-22.03.3-ath79-generic-tplink_archer-c7-v2.manifest
openwrt-22.03.3-ath79-generic-tplink_archer-c7-v2-squashfs-factory.bin
openwrt-22.03.3-ath79-generic-tplink_archer-c7-v2-squashfs-factory-eu.bin
openwrt-22.03.3-ath79-generic-tplink_archer-c7-v2-squashfs-factory-us.bin
openwrt-22.03.3-ath79-generic-tplink_archer-c7-v2-squashfs-sysupgrade.bin
profiles.json
sha256sums
```

