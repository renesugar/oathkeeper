# Go API client for swagger

ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: Latest
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://www.ory.am](https://www.ory.am)

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**GetWellKnown**](docs/DefaultApi.md#getwellknown) | **Get** /.well-known/jwks.json | Returns well known keys
*RuleApi* | [**CreateRule**](docs/RuleApi.md#createrule) | **Post** /rules | Create a rule
*RuleApi* | [**DeleteRule**](docs/RuleApi.md#deleterule) | **Delete** /rules/{id} | Delete a rule
*RuleApi* | [**GetRule**](docs/RuleApi.md#getrule) | **Get** /rules/{id} | Retrieve a rule
*RuleApi* | [**ListRules**](docs/RuleApi.md#listrules) | **Get** /rules | List all rules
*RuleApi* | [**UpdateRule**](docs/RuleApi.md#updaterule) | **Put** /rules/{id} | Update a rule


## Documentation For Models

 - [InlineResponse401](docs/InlineResponse401.md)
 - [JsonRule](docs/JsonRule.md)
 - [JsonWebKey](docs/JsonWebKey.md)
 - [JsonWebKeySet](docs/JsonWebKeySet.md)
 - [Rule](docs/Rule.md)
 - [Upstream](docs/Upstream.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author

hi@ory.am
