package main

import (
	"github.com/moby/buildkit/frontend/gateway/grpcclient"
	"github.com/moby/buildkit/util/appcontext"
	"github.com/sirupsen/logrus"
	ib "github.com/ansemjo/openwrt-imagebuilder"
)

func main() {
	if err := grpcclient.RunFromEnvironment(appcontext.Context(), ib.Build); err != nil {
		logrus.Errorf("fatal error: %+v", err)
		panic(err)
	}
}