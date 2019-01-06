# go-log-to-elasticsearch

Send logs to Elasticseach 

This go software is designed to send prowl Request to you prowl token.
It can also store the prown events in your elasticsearch if you have one set.

* Prerequisits

Having an ElasticSearch running


* Using the app.

The app will be running on the port set in the configuration.yaml file.
you need to set an environment variable to specify the full path to you configuration file.

the environment variable is : LOGGER_CONFIGURATION_FILE

If no env variable is set, the application will search for it in the current file from 
which is started the application or the following path : 
```
/home/pi/go/src/go-log-to-elasticsearch/configuration.yaml
```

* RUN : to run the application either : 
```
go run main.go
```
or
```
go build  // to build the application
```
then 
```
./go-log-to-elasticsearch // to run the build
```

You will probably be missing dependencies
```
	"github.com/op/go-logging"
	"gopkg.in/yaml.v2"
```
to add a dependency run 

```
go get <name of dependency>
```

for example : 
```
go get github.com/Mimerel/go-logger-client
```

* using the app

send a post request to the server running this app with the configured port

```
http://<ip>:<port>/<collection>/<type of log>/<log message>
```

for example : 
```
http://192.168.0.100/mylogs/error/<log_message>

of

http://192.168.0.100/mylogs/info/<log_message>

```

* Configuration file

```
elasticSearch:
    url: <ip to your elasticsearch>:9200
port: 9999  // port on which will run this application
```