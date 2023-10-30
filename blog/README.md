# Blog system

## Overview

This service works as a backend for blog system. It has functionalities like creating article, fetching all articles as well as specific article by id.
Total three endpoints are there which is list below:

1. create an article
2. fetch article by article id
3. fetch all articles

## Setup in local and run service

1. Install Golang Version => 1.20
2. Install Docker
3. Install mysql database
4. run `go run main.go` to run service

## Run test cases and code coverage report

1. To run test `go test ./...`
2. To run test with coverage `go test -cover ./...`