# Ollama Proxy
## Overview

This project integrates with various AI models through specified APIs, allowing developers to leverage powerful AI capabilities. It supports any OpenAI compatible API, including models from OpenRouter and OpenAI itself.

## Technologies Used

- Go SDK 1.24.3
- Go Programming Language 1.24

## Configuration

The configuration for API integration is defined in a YAML format. Below is an example configuration and its description.

### Example Configuration

```yaml
- identifier: openrouter
  implementation: OPENAI_API_V1
  models:
    - mistralai/devstral-small
    - qwen/qwen3-235b-a22b
  endpoint: https://openrouter.ai/api/v1/
  key: $OPENROUTER_API_KEY

- identifier: openai
  implementation: OPENAI_API_V1
  models:
    - gpt-4o-mini
  endpoint: https://api.openai.com/v1/
  key: $OPENAI_API_KEY
```


### Configuration Fields

- **identifier**: A unique name for the API integration (e.g., `openrouter`, `openai`).

- **implementation**: The version of the API being implemented. This is typically set to `OPENAI_API_V1` for compatibility with OpenAI APIs.

- **models**: A list of AI models that can be accessed through this API. Each model is typically identified by its name (e.g., `mistralai/devstral-small`, `gpt-4o-mini`).

- **endpoint**: The URL of the API endpoint where requests will be sent. This should point to the respective service.

- **key**: The environment variable name that holds the API key for authentication (e.g., `$OPENROUTER_API_KEY`, `$OPENAI_API_KEY`).

## How to Configure

1. **Create a Configuration File**:
Create a YAML configuration file in your project directory (e.g., `config.yaml`).

2. **Define Your APIs**:
Populate the configuration file with the necessary APIs by using the provided example format. Ensure you replace any placeholder values as necessary.

3. **Set Environment Variables**:
Make sure to set your API keys as environment variables in your terminal or shell configuration:
```shell script
export OPENROUTER_API_KEY='your_openrouter_api_key'
export OPENAI_API_KEY='your_openai_api_key'
```

4. **Load Configuration in Your Application**:
Ensure your application is set up to read the configuration from the YAML file to initialize the API integrations.

## Getting Started

1. **Clone the repository**:
```shell script
git clone <repository-url>
cd <repository-directory>
```


2. **Set Up Environment Variables**:
Ensure that you have your API keys set as environment variables:
```shell script
export OPENROUTER_API_KEY='your_openrouter_api_key'
export OPENAI_API_KEY='your_openai_api_key'
```


3. **Build and Run**:
Use the Go tools to build and run the project:
```shell script
go build
./your-executable
```

## Docker Usage

This project can be containerized with Docker. Below are the steps to build and run the Docker image:

1. **Build the Docker Image**:
   Navigate to the directory containing your `Dockerfile` and run the following command:
   ```shell script
   docker build -t ollama-proxy .
   ```

2. **Run the Docker Container**:
   Once the image is built, you can run it using the following command:
   ```shell script
   docker run -d -p 11434:11434 --env OPENROUTER_API_KEY='your_openrouter_api_key' --env OPENAI_API_KEY='your_openai_api_key' ollama-proxy
   ```
   This command maps port `11434` inside the container to port `11434` on your host machine, allowing you to access the service from your browser or API client.

3. **Access the Application**:
   After running the container, you can access the application at `http://localhost:11434/`.

Ensure that your environment variables for the API keys are correctly set when running the container.


## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Feel free to add any additional sections such as "Usage", "Examples", or "FAQs" as needed!
