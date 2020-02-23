package cmd

import (
	"errors"
	"fmt"
	"github.com/aiziyuer/registryV2/impl/util"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/xlab/treeprint"
)

func init() {

	var imageManifestCmd = &cobra.Command{
		Use:          "manifest",
		Short:        "manifest for image",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 1 {
				cmd.SilenceUsage = false
				return errors.New("image name require, eg: ubuntu:latest ")
			}

			registryHost := util.GetEnvAnyWithDefault("registry-1.docker.io", "REGISTRY_V2_HOST")
			func() {
				// support for private registry
				m := util.RegexNamedMatch(args[0], `(?:(?P<Host>[^.]+\.[^/]+)/)?(?P<RepoName>[a-z0-9]+(?:[._\-/][a-z0-9]+)*):(?P<TagName>[a-z0-9]+(?:[._\-/][a-z0-9]+)*)`)
				if len(m) != 0 && m["Host"] != "" {
					registryHost = m["Host"]
				}
			}()
			c, err := getClient(registryHost)
			if err != nil {
				return err
			}

			manifestV2, err := c.ManifestV2(args[0])
			if err != nil {
				logrus.Errorf("can't get manifestV2 from image(%s)", args[0])
				return err
			}
			switch outputFormat {
			case "table":
				tree := treeprint.New()
				rootTree := tree.AddBranch(fmt.Sprintf("[D] %s %d", manifestV2.Digest, manifestV2.Size))

				for _, m := range manifestV2.Manifests {
					subTree := rootTree.AddBranch(fmt.Sprintf("[P %s/%s] %s %d", m.Platform.OS, m.Platform.Architecture, m.Digest, m.Size))
					subTree.AddNode(fmt.Sprintf("[C] %s %d", m.Config.Digest, m.Config.Size))
					//size := len(m.Layers)
					for i, layer := range m.Layers {
						subTree.AddNode(fmt.Sprintf("[L %3d] %s %d", i+1, layer.Digest, layer.Size))
					}
				}

				fmt.Println(tree.String())
			case "json":
				json, err := util.Object2PrettyJson(manifestV2)
				if err != nil {
					return err
				}
				fmt.Println(json)
			case "yaml":
				// TODO
			}

			return nil
		},
	}

	rootCmd.AddCommand(imageManifestCmd)

}
