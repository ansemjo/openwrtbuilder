Not supported anymore.
OpenWRT 19.07 was EoL in April 2022.

It is still possible to build this config but you'll need to use
an older openwrtbuilder script (<= r44-ge462a2e) and edit the script
to pin it to a debian:bullseye base image, which still contains
packages for python2.7 required by this old OpenWRT release.
