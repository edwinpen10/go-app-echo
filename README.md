<a href="https://echo.labstack.com"><img height="80" src="https://cdn.labstack.com/images/echo-logo.svg"></a>



## How to runnnig app

- install migrate database for detail visit (https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
- migrate structure database on directory databases/migration. For detail how to run migrate check this (https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
- go mod tidy
- go run server.go








# Structure directory

Maybe will be update

| Directory                                                                                           | Description                                                                                                                                                                                              |
|------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| app                                     | main directory                                                                     |
| config                             | setup database                                                                                                               |
| database/migrations                                          |  database file migration                                                                                                           |
| helper                                         |  lib or helper config                                                                                                        |
| app/controllers                                         | controller file                                                                                                       |
| app/middlewares| middleware file |
| app/models | everything about query transactation to database in here                                                                                                                                    |
| app/routes | everything route to call function controllers                                                                                                                                                                  |

Let's code and happy coding thankyou (^_-)

## Help

- [Forum](https://github.com/labstack/echo/discussions)
