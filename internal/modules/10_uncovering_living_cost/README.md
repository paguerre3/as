# E10: Uncovering the Cost of Living in the Galactic Empire

Code organization:

as/
│
├── internal/modules/10_uncovering_living_cost/domain/
│   ├── models.go          # Contains core domain entities.
│   ├── services.go        # Business logic related to H3 and predictions.
│
├── internal/modules/10_uncovering_living_cost/application/
│   ├── h3_service.go      # Logic for H3 index-based features.
│   ├── model_service.go   # Logic for machine learning model handling.
│
├── internal/modules/10_uncovering_living_cost/infrastructure/
│   ├── csv_repository.go  # Handles loading/saving CSV data.
│   ├── parquet_repository.go # Handles loading Parquet data.
│
├── internal/modules/10_uncovering_living_cost/interfaces/
│   ├── cli.go             # CLI for running the program.
│
├── cmd/10_uncovering_living_cost/main.go # Main entry point.
