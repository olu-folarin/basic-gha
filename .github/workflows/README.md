# Reusable CI Pipeline Workflow

This workflow is designed to be a reusable GitHub Actions pipeline for the **basic-gha** repository. It is set up to perform several tasks including:

- **Testing**: It sets up the Go environment, updates dependencies in the `codegen` directory, builds the project, and runs tests.
- **Security Scanning**: It runs Gitleaks and Semgrep to detect potential secrets and security issues, uploading SARIF reports for further analysis.
- **Docker Image Build and Push**: It sets up Docker Buildx, configures AWS credentials, builds a Docker image from the `codegen` directory, and pushes the image to Amazon ECR.
- **Vulnerability Scanning**: It uses Trivy to scan the Docker image for known vulnerabilities.

## Setup

1. **Workflow Call Trigger**: The workflow has been converted to a reusable workflow by replacing the standard event triggers with a `workflow_call` trigger. This means it does not run automatically on push, pull request, or schedule events, but only when invoked from another workflow.

2. **Secrets and Environment Variables**: Make sure that the following secrets are defined in your repository or organization settings:
   - `AWS_ACCOUNT_ID`
   - `AWS_REGION`
   - `AWS_ACCESS_KEY_ID`
   - `AWS_SECRET_ACCESS_KEY`
   - `ECR_REPOSITORY`
   - `PAT_FOR_WORKFLOW` (if needed for GitHub token usage)
   
   Also ensure any required environment variables are set, such as `GITHUB_WORKSPACE` and others as needed for your setup.

## How to Reference This Workflow

To reference this reusable workflow from another repository (e.g., your microservice repo), add a workflow file (e.g., `.github/workflows/use-ci.yaml`) with the following content:

```yaml
name: Use Reusable CI Pipeline

on:
  push:
    branches: [ main ]

jobs:
  call-reusable-workflow:
    uses: <owner>/<repo>/.github/workflows/ci.yaml@<ref>
    with:
      # Provide any required inputs here
    secrets:
      AWS_ACCOUNT_ID: ${{ secrets.AWS_ACCOUNT_ID }}
      AWS_REGION: ${{ secrets.AWS_REGION }}
      AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
      AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
      ECR_REPOSITORY: ${{ secrets.ECR_REPOSITORY }}
      PAT_FOR_WORKFLOW: ${{ secrets.PAT_FOR_WORKFLOW }}
```

Replace `<owner>/<repo>` with the appropriate repository owner and name, and `<ref>` with the branch, tag, or commit that you want to use (e.g., `main`).

This will allow you to reuse the CI pipeline defined in this workflow in other repositories, ensuring consistent build, test, and deployment processes across your projects.

## Additional Notes

- **Customization**: You can further customize the reusable workflow by adding inputs and modifying steps as needed. Refer to the [GitHub Actions documentation on reusable workflows](https://docs.github.com/en/actions/using-workflows/reusing-workflows) for more details.
- **Dependencies**: The workflow assumes that certain directories (like `codegen`) exist and contain a valid Go project. Adjust the workflow steps if your project structure differs.
