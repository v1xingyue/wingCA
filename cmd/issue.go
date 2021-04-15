package cmd

import (
	"log"
	"net"
	"strings"
	"wingCA/rootCA"

	"github.com/spf13/cobra"
)

var (
	issueCmd = &cobra.Command{
		Use:   "issue",
		Short: "Issue Site Or Client By Your CA",
		Run:   issueStart,
	}
)

func init() {
	issueCmd.Flags().StringVarP(&issueType, "type", "t", "", "Issue Type site Or client.")
	issueCmd.Flags().StringVarP(&commonName, "common", "", "", "Common Name In You Cert.")
	issueCmd.Flags().StringVarP(&email, "email", "", "", "Email In You Cert.")
	issueCmd.Flags().StringArrayVarP(&siteNames, "site", "", []string{}, "Domains In Your Cert.")
	issueCmd.Flags().StringArrayVarP(&siteIPStr, "ip", "", []string{}, "IP Addrs In Your Cert.")
	issueCmd.Flags().BoolVarP(&withP12, "withp12", "", false, "if export p12 file")
	issueCmd.Flags().StringVarP(&password, "password", "p", "", "Password for your p12 file.")
}

func issueStart(cmd *cobra.Command, args []string) {
	switch issueType {
	case "site":
		issueSite()
	case "client":
		issueClient()
	default:
		log.Println("--type must be site or client ")
	}
}

func issueSite() {

	var (
		err error
	)

	log.Println("Issue Site : \n Names : ", siteNames, "\nips:", siteIPStr, "\nemail:", email, "\ncommonName:", commonName)
	log.Println("..")

	if commonName == "" {
		log.Println("commonName can not be empty")
		return
	}

	for _, ip := range siteIPStr {
		t := net.ParseIP(ip)
		log.Println(ip, t)
		if t != nil {
			siteIPs = append(siteIPs, t)
		}
	}

	if err = rootCA.IssueSite(commonName, siteIPs, siteNames, email); err == nil {
		log.Printf(
			"Issue Success !\nCert Path : %s \n Key Path : %s\n",
			rootCA.SiteCertPath(commonName),
			rootCA.PrivateKeyPath(commonName),
		)
	}

	if err != nil {
		log.Println(err)
	}
}

func issueClient() {

	commonName := strings.Split(email, "@")[0]
	log.Println("Issue Client : \n email : ", email, " common name : ", commonName)

	if rootCA.IssueClient(commonName, email) == nil {
		log.Printf(
			"Issue Success !\nCert Path : %s \n Key Path : %s\n",
			rootCA.ClientCertPath(commonName),
			rootCA.PrivateKeyPath(commonName),
		)
	}

	if withP12 {
		log.Println("Make p12 file ")
		if password == "" {
			log.Println("password must required ..")
		} else {
			err := rootCA.MakePKCS12(rootCA.ClientCertPath(commonName), rootCA.PrivateKeyPath(commonName), password)
			if err == nil {
				log.Println("p12 file path : ", rootCA.P12Path(commonName))
			} else {
				log.Println("issue client error : ", err)
			}

		}
	}
}
