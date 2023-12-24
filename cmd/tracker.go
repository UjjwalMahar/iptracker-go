package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var trackerCmd = &cobra.Command{
	Use:   "tracker", //use this work in the CLI to run this
	Short: "Track your IP",
	Long:  `Tracker your TP`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			for _, ip := range args {
				showIp(ip, os.Getenv("TOKEN"))

			}

		} else {
			fmt.Println("Please enter your IP address")
		}
	},
}

type Ip struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
	TimeZone string `json:"timezone"`
}

func showIp(ip string, token string) {

	//used ipinfo api

	url := "http://ipinfo.io/" + ip + "?token=" + token + ""
	responseInByte := fetchIp(url)

	data := Ip{}

	err := json.Unmarshal(responseInByte, &data)
	if err != nil {
		fmt.Printf("Unable to Unmarshal the response")
	}
	c := color.New(color.FgBlue).Add(color.Underline).Add(color.Bold)
	c.Printf("IP INFORMATION : \n")
	fmt.Printf("IP :%s\nCITY :%s\nREGION :%s\nCOUNTRY NAME :%s\nLOCATION :%s\nTIME ZONE :%s\n", data.Ip, data.City, data.Region, data.Country, data.Location, data.TimeZone)
}

func fetchIp(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Unable to get the response")
	}

	responseInBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Unable to read response")
	}
	return responseInBytes
}

// This is important!!
func init() {
	rootCmd.AddCommand(trackerCmd)
}
