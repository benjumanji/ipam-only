package main

import (
  "fmt"
  "github.com/containernetworking/cni/pkg/skel"
  "github.com/containernetworking/cni/pkg/types"
  "github.com/containernetworking/cni/pkg/types/current"
  "github.com/containernetworking/cni/pkg/version"
  "github.com/containernetworking/plugins/pkg/ipam"
  "github.com/containernetworking/plugins/pkg/ns"
)

type IpamOnlyConf struct {
	types.NetConf
	PrevResult *current.Result `json:"prevResult"`
}

func cmdAdd(args *skel.CmdArgs) error {
	netns, err := ns.GetNS(args.Netns)
	if err != nil {
    		return fmt.Errorf("failed to open netns %q: %v", args.Netns, err)
	}
	n := &IpamOnlyConf{}
	r, err := ipam.ExecAdd(n.IPAM.Type, args.StdinData)
	return
}

func cmdDel(args *skel.CmdArgs) error {
	return fmt.Errorf("not implemented")
}

func main() {
	skel.PluginMain(cmdAdd, cmdDel, version.All)
}
