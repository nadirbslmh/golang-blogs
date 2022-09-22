## golang-blogs

Simple REST API implementation with standard libraries.

## How to use

1. Clone this repository.

2. Create a new database called `golangblog`.

```sql
CREATE DATABASE golangblog;
```

3. Create a new table called `blog` in `golangblog` database.

```sql
USE golangblog;

CREATE TABLE blog(
    id INT PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    content VARCHAR(255) NOT NULL
);

```

4. Run the application.

```sh
go run main.go
```
