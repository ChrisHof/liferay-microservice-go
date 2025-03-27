# Liferay Microservice Client Extension in Go

Liferay Microservice Client Extension starter project using Go.

## Features

- **JWT Authentication:** Securely validates incoming requests using JWT.
- **OAuth2 Application Validation:** Ensures that the JWT client_id matches the configured OAuth2 application.
- **Environment Configuration:** Use `.env` files or OS environment variables (e.g. via Dockerfile) for easy configuration.
- **Object Action Endpoint:** Sample Object Action endpoint that outputs JSON body from POST data.
- **Liveness Endpoint:** Health check endpoint for monitoring.

## Prerequisites

- Liferay CE/DXP 7.4+ (see [note](#note) at bottom for Liferay Cloud deployments)
- Go 1.24 or higher

## Usage

1. Set environment variables in the `.env` file or OS environment variables for your setup:

    ```
    HTTP_SERVER_PORT=18080
    LIFERAY_BASE_URL=http://liferay:8080
    OAUTH2_APPLICATION_REFERENCE_CODE=microservice-go-oauth-application-user-agent
    ```

- **HTTP_SERVER_PORT:** Port for the Microservice HTTP Server. Must also be set in `client-extension.yaml`.
- **LIFERAY_BASE_URL:** Liferay Url reachable from the Microservice.
- **OAUTH2_APPLICATION_REFERENCE_CODE:** Defined in `client-extension.yaml`.


2. Refer to the [Liferay Documentation](https://learn.liferay.com/w/dxp/liferay-development/client-extensions/working-with-client-extensions) for Client Extension configuration, but `.serviceAddress` + `.serviceScheme` must be set in `client-extension.yaml` to set the URL your Microservice will be reachable from for your environment.

3. Endpoints must be defined in `client-extension.yaml` and `go/endpoints.go`.

4. Depending on your Liferay environment (Cloud or self-hosted) the deployment and runtime environment of the Client Extension will vary, [see the docs](https://learn.liferay.com/w/dxp/liferay-development/client-extensions/working-with-client-extensions#deploying-to-your-liferay-instance).

5. [Configure the Action](https://learn.liferay.com/w/dxp/liferay-development/objects/creating-and-managing-objects/actions/defining-object-actions) on the desired Liferay Object.

## Note

If deploying to Liferay Cloud, configure LCP.json according to the [documentation](https://learn.liferay.com/w/liferay-cloud/reference/configuration-via-lcp-json). 