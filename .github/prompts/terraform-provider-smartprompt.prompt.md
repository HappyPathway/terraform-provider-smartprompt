Create a Terraform provider in Golang that leverages the Golang client library to integrate prompt refinement into Terraform configurations. The provider should include the following features:

- A resource or data source (choose one) that accepts a "lazy_prompt" as an input parameter.
- In the resource’s (or data source’s) Read or Create function, call the `RefinePrompt` function from the client library to generate the refined prompt.
- Expose an attribute (e.g., "refined_prompt") that holds the result of the transformation, making it available for use in other Terraform resources.
- Implement the necessary schema for the resource/data source with appropriate validations and descriptions.
- Provide clear logging and error handling to help diagnose issues during plan and apply.
- Structure the provider code according to Terraform Plugin SDK best practices, ensuring that the provider is easy to build and maintain.

Include a README with instructions on how to build and test the provider.
