#CSV-Upload
1) File upload Feature that accept formatted ".csv" and ".xlsx" extension file

2) Order of csv file format (id,name,email,mobile,country)

3) "id" is unique field. Base on "id" will update and insert database table.

4) dependency for handling ".xlsx" file "http://github.com/tealeg/xlsx" 


## To Start Project
1. Clone the repository
2. Set the environment variables using your copy of `setenv.sh`
3. Run `go generate` to execute the build pipeline.
4. Run the app using `go run cmd/srv/main.go`
5. Your application will come up at `localhost:9999`


### Installing Dependencies 

Initialize your project:
1. ``` dep init ```

The only thing to follow as a rule of thumb when a new dependency is added:
1. ``` dep ensure -add dependency_name  ```
2. ``` dep ensure ```
