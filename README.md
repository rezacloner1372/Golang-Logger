# Logging Repository for Go

This repository provides a flexible logging solution in Go, allowing seamless integration with various logging packages such as Zap or ZeroLog. The design emphasizes modularity, enabling easy swapping between different logging implementations while maintaining a structured logging format.

## Overview

The logging interface implemented here serves as a bridge between application code and specific logging packages. By adhering to this interface, developers can abstract away the underlying logging implementation, facilitating easy transitions between different logging libraries without needing to modify the application code extensively.

## Features

- **Modularity**: Developers can easily switch between different logging implementations (e.g., Zap, Zero Log) without significant code changes.
- **Structured Logging**: Logs are structured and adhere to a predefined format, ensuring consistency and enabling integration with visualization tools like ELK (Elasticsearch, Logstash, Kibana).
- **Category-based Logging**: Logs are organized into categories and subcategories, allowing for better organization and management.
- **Customizable**: Users can define additional fields (ExtraKey) to be included in the log entries, enforcing consistency in log data.

## Directory Structure

The repository structure is organized as follows:

```
├── pkg
    └── logging
        ├── category.go
        ├── logger.go
        ├── logger_helper.go
        ├── zap_logger.go
        └── zero_logger.go
```

- `category.go`: Defines categories and subcategories for logging, along with constants for ExtraKeys.
- `logger.go`: Declares the logging interface and provides a `NewLogger` function for creating logger instances.
- `zap_logger.go`: Implements the logger interface for the Zap logging package.

## Usage

1. **Define Categories and Subcategories**: Define logging categories and subcategories in `category.go` based on your application's needs.

2. **Implement Logging Interface**: Implement the logging interface defined in `logger.go` for the desired logging package (e.g., Zap).

3. **Integration**: Integrate the chosen logger implementation into your application by utilizing the `NewLogger` function provided in `logger.go`.

4. **Customization**: Customize logging configurations and define any additional fields (ExtraKeys) required for logging.

## Example

```go
package main

import (
	"github.com/your-username/your-repository/pkg/logging"
)

func main() {
	// Initialize logger
	loggerConfig := logging.LoggerConfig{
		// Define logger configuration
	}
	logger := logging.NewLogger(loggerConfig)

	// Log example
	logger.Info("Application started", logging.CategoryGeneral, logging.SubCategoryApplication, map[string]interface{}{
		"version": "1.0",
	})
}
```

## Contribution

Contributions to this repository are welcome! If you have any suggestions, feature requests, or bug reports, feel free to open an issue or submit a pull request.
