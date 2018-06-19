# Guardian Open API
This is a test project for tinkering with the
[Guardian Open API](https://open-platform.theguardian.com/documentation/)
and the
[Go language](https://golang.org/).

It requires the following environment variables to be set:

* `AWS_SDK_LOAD_CONFIG` set to a "truthy" value causes the AWS SDK to load its configuration from 
  the `config` and `credentials` files in the user's `.aws` directory 
* `AWS_PROFILE` set to the name of a user in the credentials file (if not set,the default user will be used)
* `GUARDIAN_API_KEY` set to an access key for the Guardian API
