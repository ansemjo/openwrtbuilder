module github.com/ansemjo/openwrt-imagebuilder

go 1.14

replace github.com/containerd/containerd => github.com/containerd/containerd v1.3.1-0.20200512144102-f13ba8f2f2fd

replace github.com/docker/docker => github.com/docker/docker v17.12.0-ce-rc1.0.20200310163718-4634ce647cf2+incompatible

require (
	github.com/moby/buildkit v0.7.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.6.0
	gopkg.in/yaml.v2 v2.2.4
)
