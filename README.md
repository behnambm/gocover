# GoCover - Simplifying Test Coverage for Go Projects

GoCover is a tool designed to simplify the process of fetching, cloning, executing tests, and generating comprehensive test coverage reports for Go projects. 
Whether you're dealing with a Git repository or a local codebase, GoCover empowers you with an effortless and efficient way to manage your test coverage.


## Features

- **Effortless Coverage:** GoCover automates the entire process of fetching a Git repository, running tests, and generating comprehensive coverage reports.

- **User-Friendly Interface:** GoCover provides a web page to show the coverage which simplifies the identification of uncovered segments within your codebase.

- **Remote and Local Support:** Whether you're dealing with a remote repository or working with local code, GoCover has you covered.


## Installation

```bash
go install github.com/behnambm/gocover@latest
```


## Usage

#### Remote repo:

```bash 
gocover -url https://github.com/labstack/echo.git 
```

#### Local code:

##### Relative path
```bash
gocover -path .
```

or

##### Absolute path
```bash
gocover -path /home/user/go/src/github.com/labstack/echo
```

## Visual Output

This is an example of what you will see in your browser:

![image](https://github.com/behnambm/gocover/assets/26994700/a22deb1e-072c-4e9a-9467-035cdd18ced3)


## Todo

- [ ] Refactor the code
- [ ] Add support for private repositories
