# Track my day - gotth stack
A simple, modern stack for building fast web applications.
-   Go - Backend
-   Tailwind - CSS
-   Templ - Templating
-   HTMX - Interactivity

# Install
`
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
`

# Command Example
# Create a new migration in the development environment
make migrate-create name=create_users_table ENV=development

# Apply migrations in the test environment
make migrate-up ENV=test

# Roll back the latest migration in production
make migrate-down ENV=production

## Resources
[GoTTH](https://github.com/TomDoesTech/GOTTH/tree/main) - A simple, modern stack for building fast web applications.

[GoTTH Example](https://github.com/sigrdrifa/gotth-example/tree/main) - An example GOTTH WebApp application
