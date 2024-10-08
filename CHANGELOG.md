<!-- Code generated for API Clients. DO NOT EDIT. -->

## 0.4.0

ENHANCEMENTS:

* Added support for Cloud Endpoints (currently in private beta).

## 0.3.0

ENHANCEMENTS:

* Added support for the Bot User API. The Bot User API allows you to manage the bots that are registered to your ngrok account. You can automate the creation, management, and deletion of bot users in your account.

## 0.2.0

FIXED:

* The following `optional` fields are now also marked `computed`:
  * APIKeys: `owner_id`
  * HTTPResponseBackends: `status_code`
  * ReservedDomains: `domain`
  * SSHHostCertificates: `valid_after`, `valid_until`
  * SSHUserCertificates: `extensions`, `valid_after`, `valid_until`
* The following read-only `computed` fields are no longer marked `optional`, as they cannot be set on Create or Update.
  * `id` and `uri` on all resources
  * APIKeys: `token`
  * FailoverBackends, HTTPResponseBackends: `created_at`
  * ReservedAddrs: `addr`
  * ReservedDomains: `acme_challenge_cname_target`, `cname_target`
  * SSHCertificateAuthorities: `key_type`, `public_key`
  * SSHHostCertificates, SSHUserCertificates: `certificate`, `key_type`
  * TLSCertificates: `dns_names`, `ips`, `subject_alternative_names`
  * TunnelGroupBackends: `created_at`, `tunnels`

## 0.1.5

ENHANCEMENTS:

* Added `owner_id` field to the `api_key`, `credential`, and `ssh_credential` resources. If supplied at credential creation, ownership will be assigned to the specified User or Bot. Only admins may specify an owner other than themselves. Defaults to the authenticated User or Bot.
* Added `failover_backend`, `http_response_backend`, and `tunnel_group_backend` resources. A Failover backend defines failover behavior within a list of referenced backends. Traffic is sent to the first backend in the list. If that backend is offline or no connection can be established, ngrok attempts to connect to the next backend in the list until one is successful.

CHANGED:

* The `domain` field of the ReservedDomains schema is marked `optional` as it can now be set on Create. This field was no longer marked `computed` in error, this is fixed in version `0.2.0`.

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
