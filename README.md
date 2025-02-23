# tf-smart-prompt

Repository for the gcp-kubernetes project in HappyPathway

## Getting Started

1. Clone this repository:
```bash
git clone git@github.com:HappyPathway/tf-smart-prompt.git
cd tf-smart-prompt
```

2. Set up Python environment and install dependencies:
```bash
python -m venv .venv
source .venv/bin/activate  # On Windows use: .venv\Scripts\activate
pip install -r scripts/requirements.txt
```

3. Run the initialization script:
```bash
python scripts/init.py
```

This will:
- Verify Git SSH access to GitHub
- Create the workspace directory structure
- Clone or update all project repositories
- Set up repository configurations

For debugging, you can run:
```bash
python scripts/init.py --debug
```

## Repository Structure

This project consists of multiple repositories:

- terraform-provider-smartprompt: tf-smart-prompt::terraform-provider-smartprompt
- smartprompt-api: tf-smart-prompt::smartprompt-api
- smartprompt-client: tf-smart-prompt::smartprompt-client
- smartprompt-infra-deployment: tf-smart-prompt::smartprompt-infra-deployment

## Development Environment

This repository includes:
- VS Code workspace configuration
- GitHub Copilot settings
- Project-specific documentation and guidelines
- Python-based initialization tools

## Contributing

Please see the [CONTRIBUTING.md](.github/CONTRIBUTING.md) file for guidelines.

# Terraform Provider for Smart Prompt

This provider allows you to use the Smart Prompt API to refine prompts as part of your Terraform configurations.

## Usage

```hcl
terraform {
  required_providers {
    smartprompt = {
      source = "yourusername/smartprompt"
    }
  }
}

provider "smartprompt" {
  api_url = "http://localhost:8000"  # URL of the Smart Prompt API
  timeout = 30  # Request timeout in seconds
}

data "smartprompt_refined" "example" {
  lazy_prompt = "tell me about terraform"
}

output "refined_prompt" {
  value = data.smartprompt_refined.example.refined_prompt
}
```

## Provider Configuration

- `api_url` (Required) - The base URL of the Smart Prompt API
- `timeout` (Optional) - Request timeout in seconds. Defaults to 30.

## Data Sources

### smartprompt_refined

Use this data source to refine a prompt using the Smart Prompt API.

#### Arguments
- `lazy_prompt` (Required) - The simple prompt to be refined

#### Attributes
- `refined_prompt` - The resulting refined prompt