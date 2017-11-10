# Predix-Go-Sample-App

Sample web app in Go pushed to Predix Cloud reading data from PostgresDB on Predix.

## Prerequisistes

- Predix.io account
- PostgreSQL instance on Predix

On your local machine:

- CF CLI
- Git
- Go (v.1.9.x)
- GoVendor: install it with `go get -u github.com/kardianos/govendor`

## Get the app, configure and push to Predix

1. Clone the repos:

  ```shell
  $ go get -u https://github.com/indaco/predix-go-sample-app
  ```

2. Move to the app folder:

  ```shell
  $ cd $GOPATH/src/github.com/indaco/predix-go-sample-app
  ```

3. Sync dependencies:

  ```shell
  $ $GOPATH/bin/govendor sync
  ```

4. Edit _config.json_ and update it with Predix Postgres credentials
5. Use the sample `db.sql` file to create a dummy table with some entries on PostgreSQL or adapt the main.go code to reflect the data table on PostgreSQL
6. Edit _manifest.yml_ file and update the application name
7. Push the app to Predix Cloud:

  ```shell
  $ cf push
  ```
