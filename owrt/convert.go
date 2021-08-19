package builder

import (
	"fmt"
	"strings"
	"time"

	"github.com/ansemjo/openwrt-imagebuilder/config"
	"github.com/moby/buildkit/client/llb"
)

func ConfigToLLB(conf config.Builder) llb.State {

	base := BuildBase()

	return base.File(llb.Mkfile("/myconfig", 0644, conf.ToJSON()),
		llb.WithCustomName("copy builder configuration to target"))

}

// BuildBase returns the llb.State for the builder base image. It contains all
// the software requirements necessary to create a custom OpenWRT image and
// signing keys to verify downloads.
func BuildBase() llb.State {

	// start from a base image
	s := llb.Image("docker.io/library/debian:stable",
		llb.WithCustomName("create openwrtbuilder base image"))

	// list of packages to install
	software := strings.Join([]string{
		"build-essential",
		"libncurses5-dev",
		"zlib1g-dev",
		"gawk",
		"git",
		"gettext",
		"libssl-dev",
		"xsltproc",
		"wget",
		"unzip",
		"python",
		"python3",
		"curl",
		"xxd",
		"signify-openbsd",
	}, " ")

	// run apt-get to install packages
	opts := []llb.RunOption{
		llb.Args([]string{"sh", "-c", "apt-get update && apt-get install -y " + software}),
		llb.WithCustomName("install required software packages"),
	}
	s = s.Run(opts...).Root()

	// write the public signing keys from SigningKeys to files
	const keydir = "/signingkeys"
	t, _ := time.Parse(time.RFC3339, SigningKeysTimestamp)
	fileop := llb.Mkdir(keydir, 0755, llb.WithCreatedTime(t))
	for id, key := range SigningKeys {
		filename := fmt.Sprintf("%s/%s", keydir, id)
		fileop = fileop.Mkfile(filename, 0644, []byte(key), llb.WithCreatedTime(t))
	}
	// create in scratch and copy to enable caching
	keys := llb.Scratch().File(fileop, llb.WithCustomName("write signing keys to scratch"))
	s = s.File(llb.Copy(keys, keydir, keydir),
		llb.WithCustomName("copy trusted signing keys"),
	)

	return s

}

// The SigningKeys map can be updated with `go generate`:
//go:generate sh -c "sh tools/get_signingkeys.sh | tee signingkeys.go"
