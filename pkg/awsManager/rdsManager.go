package awsManager

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/service/rds"
	clientpkg "github.com/openshift/aws-account-shredder/pkg/aws"
)

func DescribeDBInstances(client clientpkg.Client) error {
	dbInstances, err := client.DescribeDBInstances(&rds.DescribeDBInstancesInput{})
	if err != nil {
		fmt.Println("ERROR", err)
		return errors.New("ERROR")
	}
	fmt.Println("DB INSTANCES")
	fmt.Println(dbInstances)

	return nil
}

func DescribeDBClusters(client clientpkg.Client) error {
	dbClusters, err := client.DescribeDBClusters(&rds.DescribeDBClustersInput{})
	if err != nil {
		fmt.Println("ERROR", err)
		return errors.New("ERROR")
	}

	fmt.Println("DB CLUSTERS")
	fmt.Println(dbClusters)

	return nil
}
