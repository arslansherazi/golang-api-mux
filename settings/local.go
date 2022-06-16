package main

import "os"

func main() {
	// Database settings
	os.Setenv("host", "localhost")
	os.Setenv("user", "postgres")
	os.Setenv("password", "postgres")
	os.Setenv("dbname", "fnd_comp_db")
	os.Setenv("posrt", "5432")

	// AWS configurations
	os.Setenv("AWS_REGION", "us-east-2")
	os.Setenv("AWS_FND_COMP_BUCKET", "fnd-cmp-files")
	os.Setenv("BUCKET_BASE_URL", "https://fnd-cmp-files.us-east-2.amazon.aws.com/")
}
