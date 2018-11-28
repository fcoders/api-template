# API Template

Template project for Go applications. It has:

- HTTP configured with [Gin Gonic](https://github.com/gin-gonic/gin)
- Database access defined for MySQL (integrated with [GORM](http://gorm.io/))
- Basic logger with integration with `Slack`

## Usage

Use this project as a baseline for creating new ones. Update the namespaces to match the new project's name.

To run the project, download all the dependencies, copy `settings.example.yml` to `settings.yml` and configure the database setttings on it, as well as other settings you might want to change, like the HTTP port where the server runs.

Use `database/Storage` interface to define new DB access functions and `database/mysql/Storage` structure to write each implementation. Use `settings.DefaultStorage()` to get the storage instance. Using an interface approach like this makes it easier to write integration testings.