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
- error
  - invalid_request
  - unauthorized_client
  - invalid_grant
  - access_denied
  - unsupported_grant_type
  - unsupported_response_type
  - invalid_scope
  - server_error
  - temporarily_unavailable
- error_description
- error_uri
- state
