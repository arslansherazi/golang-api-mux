# Find Competitor
## Important Commands
Build code
~~~
go build main.go
~~~
Run server
~~~
go run main.go
~~~
Migrate models
~~~
go run scripts/migrate_models.go
~~~
Remove unused dependencies from the project
~~~
go mod tidy -v
~~~

## Handle AWS Configurations
install aws cli
~~~
sudo apt install awscli -y
~~~
add configurations
~~~
aws configure
~~~