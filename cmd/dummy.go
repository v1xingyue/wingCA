package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	dummpyCmd = &cobra.Command{
		Use:   "dummy",
		Short: "Some Dummy Command",
		Run: func(cmd *cobra.Command, args []string) {
			prefix := "./wingCA"
			lines := []string{
				"init --confirm",
				"init --confirm -l Beijing -n SimpleRootCA -o \"Big CA Center\" --postcode 100093 -p Beijing -s \"NoWhere Road 9+3/4 Site Corner\" ",
				"issue --type site -c a.b.ssl.com.cn --email xingyue@ssl.com.cn --ip 127.0.0.1 --ip 10.41.13.133 --site a.b.ssl.com.cn --site \"*.d.ssl.com.cn\" --site localhost ",
				"issue --type client --email xingyue@ssl.com.cn",
				"issue --type client --email xingyue@ssl.com.cn --withp12 --password super",
				"sample --common a.b.ssl.com.cn --double",
				"sample --common a.b.ssl.com.cn ",
			}

			log.Println(cmd.Short + "\n")
			for _, line := range lines {
				fmt.Println(prefix, line)
			}
			fmt.Println("")
		},
	}
)
