<!-- Code generated for API Clients. DO NOT EDIT. -->

## 0.1.5

ENHANCEMENTS:

* Added `owner_id` field to the `api_key`, `credential`, and `ssh_credential` resources. If supplied at credential creation, ownership will be assigned to the specified User or Bot. Only admins may specify an owner other than themselves. Defaults to the authenticated User or Bot.
* Added `failover_backend`, `http_response_backend`, and `tunnel_group_backend` resources. A Failover backend defines failover behavior within a list of referenced backends. Traffic is sent to the first backend in the list. If that backend is offline or no connection can be established, ngrok attempts to connect to the next backend in the list until one is successful.

## 0.1.4

ENHANCEMENTS:

* Added `nameid_format` field to the SAML endpoint configuration.

## 0.1.3

FIXED:

* Suppress capitalization diff on webhook validation provider
* Mark fields transmitting secrets as sensitive
* Suppress ID diff on automatically provisioned certificates
* Additional create fields marked as required

## 0.1.2

FIXED:

* Remove many unnecessary fields.

## 0.1.1

FIXED:

* Correctly handle referenced resources by id.

## 0.1.0

ENHANCEMENTS:

* Add a resource to manage IP Policy Rules.

## 0.0.27

ENHANCEMENTS:

* Improve error reporting from lower-level API client.
* Add CHANGELOG and top-level README.

## 0.0.26

ENHANCEMENTS:

* Replace documentation generation with [`tfplugindocs`](https://github.com/hashicorp/terraform-plugin-docs)
