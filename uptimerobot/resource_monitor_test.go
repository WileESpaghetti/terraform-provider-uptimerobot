package uptimerobot

import (
	"testing"
)
func TestAccAWSS3Bucket_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckUptimerobotMonitorDestroy,
		Steps: []resource.TestStep{
			{
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAWSS3BucketExists("aws_s3_bucket.bucket"),
					resource.TestCheckResourceAttr(
						"aws_s3_bucket.bucket", "hosted_zone_id", HostedZoneIDForRegion("us-west-2")),
					resource.TestCheckResourceAttr(
						"aws_s3_bucket.bucket", "region", "us-west-2"),
					resource.TestCheckNoResourceAttr(
						"aws_s3_bucket.bucket", "website_endpoint"),
					resource.TestMatchResourceAttr(
						"aws_s3_bucket.bucket", "arn", arnRegexp),
					resource.TestCheckResourceAttr(
						"aws_s3_bucket.bucket", "bucket", testAccBucketName(rInt)),
					resource.TestCheckResourceAttr(
						"aws_s3_bucket.bucket", "bucket_domain_name", testAccBucketDomainName(rInt)),
				),
			},
		},
	})
}

func testAccCheckUptimerobotMonitorDestroy() {
}
