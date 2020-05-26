package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Run the awesome app
func Run(source string, output string) {
	if output == "" {
		output = "output.tf"
	}

	fmt.Println("source: ", source)
	fmt.Println("output: ", output)

	var data JsonSource
	file, _ := ioutil.ReadFile(source)

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		f.Close()
	}

	for _, item := range data.ResourceRecordSets {

		if item.Type == "NS" || item.Type == "SOA" {
			continue
		}

		resourceName := strings.TrimRight(item.Name, ".")
		resourceName = strings.ReplaceAll(resourceName, ".", "_")
		if item.Type == "MX" || item.Type == "SPF" || item.Type == "TXT" {
			resourceName = resourceName + "_" + strings.ToLower(item.Type)
		}

		fmt.Fprintf(f, "resource \"aws_route53_record\" \"%s\" {\n", resourceName)
		fmt.Fprintf(f, "  zone_id = \"ZONEID\"\n")
		fmt.Fprintf(f, "  name    = \"%s\"\n", item.Name)
		fmt.Fprintf(f, "  type    = \"%s\"\n", item.Type)

		if item.AliasTarget.HostedZoneID != "" {
			fmt.Fprintf(f, "\n")
			fmt.Fprintf(f, "  alias {\n")
			fmt.Fprintf(f, "    name                   = \"%s\"\n", item.AliasTarget.DNSName)
			fmt.Fprintf(f, "    zone_id                = \"%s\"\n", item.AliasTarget.HostedZoneID)
			fmt.Fprintf(f, "    evaluate_target_health = \"%t\"\n", item.AliasTarget.EvaluateTargetHealth)
			fmt.Fprintf(f, "  }\n")
		} else {
			fmt.Fprintf(f, "  ttl     = \"%d\"\n", item.TTL)
			fmt.Fprintf(f, "  records = [\n")

			length := len(item.ResourceRecords)

			for index, record := range item.ResourceRecords {

				value := strings.ReplaceAll(record.Value, "\"", "")

				if (index + 1) == length {
					fmt.Fprintf(f, "    \"%s\"\n", value)
				} else {
					fmt.Fprintf(f, "    \"%s\",\n", value)
				}

			}

			fmt.Fprintf(f, "  ]\n")
		}

		fmt.Fprintf(f, "}\n\n")
	}

	fmt.Println("generated successfully:", output)
}

