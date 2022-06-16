package main

import "os"

func main() {
	os.Setenv("host", "localhost")
	os.Setenv("user", "postgres")
	os.Setenv("password", "postgres")
	os.Setenv("dbname", "fnd_comp_db")
	os.Setenv("posrt", "5432")
}
