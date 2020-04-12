// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func ListNatGateway(client *Client) {
	req := client.ec2conn.DescribeNatGatewaysRequest(&ec2.DescribeNatGatewaysInput{})

	p := ec2.NewDescribeNatGatewaysPaginator(req)
	fmt.Println("")
	fmt.Println("aws_nat_gateway:")
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.NatGateways {
			fmt.Println(*r.NatGatewayId)
			for _, t := range r.Tags {
				fmt.Printf("\t%s: %s\n", *t.Key, *t.Value)
			}
		}
	}

	if err := p.Err(); err != nil {
		log.Printf("aws_nat_gateway: %s", err)
	}

}