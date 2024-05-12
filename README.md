# Snippetcode

Snippetcode is a web application built with Golang that facilitates the sharing of code snippets or text, providing users with a platform similar to Pastebin or GitHub's Gists for efficient collaboration and showcasing of work.

![](https://github.com/vhsenna/snippetcode/assets/34524951/e3206fcb-14b7-4dcb-9460-80e524d8f727)

## Features
- **Paste and Share**: Quickly share code snippets or text via unique URLs.
- **Easy Sharing**: Share snippets through unique URLs.
- **User Profiles**: Create an account to manage snippets and track your activity.

## Contributing
Contributions to Snippetcode are welcome! To contribute, follow these steps:

1. Fork the Snippetcode repository on GitHub.
3. Create a new branch for your feature or bug fix.
4. Make your changes to the codebase.
5. Ensure that your changes are thoroughly tested and functional.
6. Submit a pull request with your changes, providing a clear description of the modifications.

## Prerequisites
Before getting started, ensure you have Go installed on your system. If you don't already have it, you can download and install it by following the instructions at [https://go.dev/dl](https://go.dev/dl).

Additionally, you'll need to have MySQL installed. If you haven't already done so, you can download and install it by following the instructions at [https://dev.mysql.com/downloads/mysql](https://dev.mysql.com/downloads/mysql).

Finally, you'll need to create a user in MySQL for our application. You can do this by following the instructions below in your terminal. Replace `'newuser'` and `'user_password'` with whatever you'd like:

**Note:** Please execute the following commands in your MySQL terminal:

```bash
mysql
mysql-> CREATE USER `newuser`@`localhost` IDENTIFIED BY `user_password`;
mysql-> CREATE DATABASE `snippetcode`;
mysql-> GRANT ALL PRIVILEGES ON `snippetcode`.* TO `newuser`@`localhost`;
mysql-> exit
```
This will create a new user named `'newuser'` with the password `'user_password'` and grant them all privileges to access the `'snippetcode'` database.

## Running the Project
Before running the project, you need to set environment variables for the MySQL username and password:

```bash
# For Linux/MacOS:
export MySQLUser="username"
export MySQLPass="password"

# For Windows (PowerShell):
$env:MySQLUser="username"
$env:MySQLPass="password"
```

Finally, you can run the project using the following command:

```bash
go run ./cmd/web
```

## Testing

To run tests for the project, execute the following command:


```bash
go test ./...
```

## Upcoming Features

- [ ] **Syntax Highlighting**: Enhance snippet readability by displaying code with syntax highlighting.
- [ ] **Share Functionality**: Implement a "Share" feature to generate unique URLs for snippet sharing.
- [ ] **Pagination**: Navigate through multiple pages for improved browsing.
- [ ] **Username Instead of Name**: Replace "name" with "username" for consistency and clarity.
- [ ] **Author Display**: Show the author of each snippet to provide attribution and context.
- [ ] **Snippet Deletion**: Allow snippet authors to delete their own snippets for content management.
