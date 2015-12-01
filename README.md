#Go API Utils [![Build Status](https://travis-ci.org/cfmobile/goapiutils.svg?branch=master)](https://travis-ci.org/cfmobile/goapiutils)

This repo contains simple classes for creating Api Handlers to use with an http router like Gorilla Mux or chaining with Alice:
- **ApiHandler** interface, containing just a `Handle()` method
- **ApiResponse** wrapping a valid API response
- **ApiError** wrapping an API error with an http status code
- **NotFoundHandler** with a default handler for not found
- **Response Writer utilities** to serialize the response into an http.ResponseWriter object

###Usage

*TODO*

###Testing

Tests written with Ginkgo, running on Travis CI.
