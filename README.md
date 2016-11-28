# cloud-auth-go
A gRPC based authentication service(OpenID Connect)
=================================

##  General OAuth 2.0 models(refer to [RFC6749](https://tools.ietf.org/html/rfc6749)):
### 1. Access Token Request:
- grant_type (authorization_code, implicit, password and client_credentials)
- client_id (could be public accessible)
- client_secret (should be protected in motion and at rest)
- code (if grant_type := authorization_code)
- redirect_uri (if grant_type := authorization_code)
- username (if grant_type := password)
- password (if grant_type := password)
- refresh_token (if grant_type := refresh_token)
- scope
- state

### 2. Access Token Response:
- access_token
- token_type (e.g. basic, bearer, mac, hmac, etc.)
- expires_in (seconds)
- refresh_token
- scope
- state

### 3. Error Response:
- ### error 
  - #### invalid_request
    The request is missing a required parameter, includes an
    invalid parameter value, includes a parameter more than
    once, or is otherwise malformed.
  - #### unauthorized_client
    The client is not authorized to request an authorization
    code using this method.
  - #### invalid_grant
    The provided authorization grant (e.g., authorization
    code, resource owner credentials) or refresh token is
    invalid, expired, revoked, does not match the redirection
    URI used in the authorization request, or was issued to
    another client.
  - #### access_denied
    The resource owner or authorization server denied the 
    request. 
  - #### unsupported_grant_type
    The authorization grant type is not supported by the
    authorization server.
  - #### unsupported_response_type
    The authorization server does not support obtaining an 
    authorization code using this method.
  - #### invalid_scope 
    The requested scope is invalid, unknown, or malformed.
  - #### server_error
    The authorization server encountered an unexpected 
    condition that prevented it from fulfilling the request.
    (This error code is needed because a 500 Internal Server
    Error HTTP status code cannot be returned to the client
    via an HTTP redirect.)
  - #### temporarily_unavailable 
    The authorization server is currently unable to handle
    the request due to a temporary overloading or maintenance 
    of the server.  (This error code is needed because a 503
    Service Unavailable HTTP status code cannot be returned
    to the client via an HTTP redirect.)
- ### error_description
  Human-readable ASCII [USASCII] text providing additional 
  information, used to assist the client developer in 
  understanding the error that occurred.
- ### error_uri
  A URI identifying a human-readable web page with 
  information about the error, used to provide the client 
  developer with additional information about the error.
- ### state
  if a "state" parameter was present in the client 
  authorization request. The exact value received from the 
  client.