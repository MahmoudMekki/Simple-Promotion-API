# Go Promo

Simple CRUDing for a promotion APi


## Introduction

Back-End development for Promotion REST API

## Requirements
* MySQL installed
* Go installed

## Installation

* Clone this repo 

```bash
git clone https://github.com/MahmoudMekki/Simple-Promotion-API.git
```

* Change Directory

```bash
cd Simple-Promotion-API
```

* Modify `.env` file with your correct database credentials and desired Port

## Usage

To run this application, execute:

```bash
go run main.go
```

You should be able to access this application at `localhost:8080`

>**NOTE**<br>
>If you modified the port in the `.env` file, you should access the application for the port you set

## Usage 101
on Postman ,this API provids Methods like:
.POST -> localhost:8080/api/{user or company}/promotions with Raw json{"title": "ur promo title","description":"ur promo description","end_date":"","start_date":"}

.GET -> localhost:8080/api/{user or company}/promotions ->  to get all the promotions

.Put -> localhost:8080/api/{user or company}/promotions/:id -> to update the promo with the specified id
.GET -> localhost:8080/api/{user or company}/promotions/:id -> to GET the promo with the specified id
.Delete -> localhost:8080/api/{user or company}/promotions/:id ->  to Delete the promo with the specified id

## Reflections yet to come!

* implementing it with GRPC 
* Adding AUth with Outh2








