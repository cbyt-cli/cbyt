package main

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var version = "0.0.0-dev" // overwritten by goreleaser during build

func main() {
	rootCmd := &cobra.Command{
		Use:   "cbyt",
		Short: "cbyt — compressed encrypted packing for IPFS",
		Long:  "cbyt compresses, encrypts, chunks, and publishes encrypted blobs; manifest contains wrapped CEK + locators.",
		Version: version,
	}

	rootCmd.AddCommand(packCmd())
	rootCmd.AddCommand(unpackCmd())
	rootCmd.AddCommand(versionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
	}
}

func packCmd() *cobra.Command {
	var input, outManifest, pinService, recipient string
	cmd := &cobra.Command{
		Use:   "pack",
		Short: "pack a file and upload chunks",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: wire your packer implementation here (compress -> encrypt -> upload -> manifest)
			fmt.Printf("pack: input=%s out=%s pin=%s rec=%s\n", input, outManifest, pinService, recipient)
			return nil
		},
	}
	cmd.Flags().StringVar(&input, "input", "", "input file to pack (required)")
	cmd.Flags().StringVar(&outManifest, "out", "cbyt-manifest.json", "manifest output path")
	cmd.Flags().StringVar(&pinService, "pin-service", "", "pinning service to use (pinata|web3|infura)")
	cmd.Flags().StringVar(&recipient, "recipient-age", "", "recipient age public key")
	cmd.MarkFlagRequired("input")
	return cmd
}

func unpackCmd() *cobra.Command {
	var manifestPath, ageKey string
	cmd := &cobra.Command{
		Use:   "unpack",
		Short: "unpack a manifest and reconstruct the file",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: implement unpack logic: unwrap CEK -> fetch from ipfs/gateway -> decrypt -> decompress
			fmt.Printf("unpack: manifest=%s key=%s\n", manifestPath, ageKey)
			return nil
		},
	}
	cmd.Flags().StringVar(&manifestPath, "manifest", "", "manifest file path (required)")
	cmd.Flags().StringVar(&ageKey, "age-key-file", "", "recipient private key (age) for unwrap")
	cmd.MarkFlagRequired("manifest")
	return cmd
}

