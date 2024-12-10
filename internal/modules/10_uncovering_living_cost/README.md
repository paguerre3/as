# E10: Uncovering the Cost of Living in the Galactic Empire

Code organization:

as/
│
├── internal/modules/10_uncovering_living_cost/domain/
│   ├── models.go          # Contains core domain entities.
│   ├── services.go        # Business logic related to H3 and predictions.
│
├── internal/modules/10_uncovering_living_cost/application/
│   ├── h3_use_case.go      # Logic for H3 index-based features.
│   ├── model_use_cases.go   # Logic for machine learning model handling.
│
├── internal/modules/10_uncovering_living_cost/infrastructure/
│   ├── csv_repository.go  # Handles loading/saving CSV data.
│   ├── parquet_repository.go # Handles loading Parquet data.
│
├── internal/modules/10_uncovering_living_cost/api/
│   ├── cli.go             # CLI for running the program.
|   |-- **cli_test.go        # Unit tests for the CLI. (*) // No CMD entry point, it's just run via integration test under "iteration 1"**
|
|--- internal/modules/10_uncovering_living_cost/assets/
│   ├── README.md          # Instructions for handling large files.
│   ├── mobility_data.parquet # Original Parquet file (must be uncompressed from split LFS files).
|   |-- mobility_data.z01
|   |-- mobility_data.z02
|   |-- mobility_data.zip
|   |-- test.csv
|   |-- train.csv
|   |-- /output/predictions_{random_seed}.csv (based on test.csv (empty data), using train.csv and mobility_data.parquet)    
|   

# Objective

Objective
Participants are challenged to estimate the cost of living (a value between 0 and 1) in different regions of LATAM based on mobility data. You will be provided with a dataset that contains anonymized device-level information, including geographic coordinates and timestamps. Your task is to transform this data into meaningful features that can accurately predict the cost of living.

# Instructions

You’ll face plenty of uncertainty, and the options aren’t endless, so you have to be resourceful. It’s about making sense of raw information, turning complexity into something useful, and finding that balance between creativity and accuracy. Every choice matters as you adjust your models, fine-tune your features, and learn as you go. It’s a chance to show how you handle real-world problems where the path to the answer isn’t always clear.

Dataset Provided
Columns:
device_id: An anonymized identifier for each mobile device.
lat: Latitude of the recorded location.
lon: Longitude of the recorded location.
timestamp: The time of the record.
Participants are encouraged to utilize external datasets and variables to enrich their predictions. You have complete freedom to incorporate any additional sources of information you think could improve your model's performance. ***(2nd iteration)***

**Submission Format**

Your submission should be a CSV file with the following columns:

hex_id: The H3 hexagon index that represents a specific geographic area.
cost_of_living: Your predicted value for the cost of living in that hexagon, ranging between 0 and 1.
Guidelines



---
# Pre-requisite for running with `cli_test.go` outside docker container

```bash
sudo apt update
sudo apt install build-essential
```

This installs **gcc**, make, and other essential development tools.

```bash
gcc --version

# inside VSCode:
which gcc
```

Ensure to rebuild:
```bash
go clean
go build
```

