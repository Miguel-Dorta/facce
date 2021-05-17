package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/Miguel-Dorta/facce/pkg/claims"
	"github.com/Miguel-Dorta/facce/pkg/types"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
	"os"
	"strings"
)

var (
	Version = "Unknown"
	credentialPath, uid string

	cmdRoot = &cobra.Command{
		Use: "facce",
		Run: root,
	}
	cmdGet = &cobra.Command{
		Use: "get",
		Short: "Get Custom-Claims",
		Run: get,
	}
	cmdSet = &cobra.Command{
		Use: "set",
		Short: "Set new Custom-Claims",
		Run: set,
	}
	cmdVersion = &cobra.Command{
		Use: "version",
		Short: "Print version and exit",
		Run: version,
	}
)

func main() {
	cmdRoot.PersistentFlags().StringVarP(&credentialPath, "credentials-file", "c", "", "App credentials. See more info in https://firebase.google.com/docs/admin/setup")
	cmdRoot.PersistentFlags().StringVar(&uid, "uid", "", "User UID")
	cmdRoot.AddCommand(cmdGet, cmdSet, cmdVersion)
	if err := cmdRoot.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func root(cmd *cobra.Command, _ []string) {
	_ = cmd.Help()
}

func get(_ *cobra.Command, args []string) {
	path := ""
	if len(args) != 0 {
		path = args[0]
	}

	result, err := claims.Get(getAuthClient(), uid, path)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(result)
}

func set(_ *cobra.Command, args []string) {
	m := make(map[string]interface{}, len(args))
	for _, arg := range args {
		parts := strings.SplitN(arg, "=", 2)
		if len(parts) != 2 {
			panic("error parsing key-value pair: " + arg)
		}
		m[parts[0]] = types.DetectType(parts[1])
	}

	if err := claims.Set(getAuthClient(), uid, m); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func version(_ *cobra.Command, _ []string)  {
	fmt.Println("facce - Firebase Auth Custom-Claims Editor")
	fmt.Printf("Version: %s\n", Version)
	fmt.Println("Author: Miguel Dorta <contact@migueldorta.com>")
}

func getAuthClient() *auth.Client {
	app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile(credentialPath))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error initializing app (%s)\nCheck your credentials and connection\n", err)
		os.Exit(1)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error getting auth client: %s\n", err)
		os.Exit(1)
	}
	return authClient
}
