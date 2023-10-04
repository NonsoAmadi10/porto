/*
Copyright © 2023 NAME HERE <nonsoamadi@aol.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var bucketName string
var files []string
var region string

// upload represents the command required to push files to s3 buckets
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload files to S3",
	Run: func(cmd *cobra.Command, args []string) {

		if bucketName == "" {
			log.Fatal("Please provide a bucket name using the -bucket flag")
		}

		if len(files) == 0 {
			log.Println("Please provide at least one file to upload")
			return
		}

		// Create an AWS SDK v2 config with the provided region or use the default region from ~/.aws/credentials.
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(region),
		)
		if err != nil {
			log.Fatalf("Failed to load AWS SDK config: %v", err)
		}

		// Initialize an S3 client
		client := s3.NewFromConfig(cfg)

		// Iterate over the files and upload each one to the specified S3 bucket
		for _, file := range files {
			err := uploadToS3(client, bucketName, file)
			if err != nil {
				log.Printf("Failed to upload %s: %v", file, err)
			}
		}
	},
}

func init() {
	uploadCmd.PersistentFlags().StringVar(&bucketName, "bucket", "", "S3 Bucket Name")
	uploadCmd.PersistentFlags().StringSliceVar(&files, "file", []string{}, "Files to upload")
	uploadCmd.PersistentFlags().StringVar(&region, "region", "", "AWS Region")
	rootCmd.AddCommand(uploadCmd)
}

func uploadToS3(client *s3.Client, bucketName, filePath string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	key := aws.String(filePath) // You can set the S3 object key to the same as the file path

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    key,
		Body:   file,
	})
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://%v.s3.amazonaws.com/%v", bucketName, key)
	log.Printf("File %s uploaded to %s\n", filePath, url)
	return nil
}
