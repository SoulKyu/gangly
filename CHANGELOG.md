# Changelog


## v3.4.0 (Unreleased)

### UI improments

Adds Windows Powershell instructions, cleans up templates and bakes in the required 
css and js assets, allowing Gangly to function without an internet connection. Shows 
cluster name on homepage. Mostly based on PR https://github.com/vmware-archive/gangly/pull/137

Fixes:
- https://github.com/vmware-archive/gangly/issues/102 
- https://github.com/vmware-archive/gangly/issues/135 
- https://github.com/vmware-archive/gangly/issues/136
- https://github.com/vmware-archive/gangly/issues/177
- https://github.com/vmware-archive/gangly/issues/189

### Support self-signed CA

Adds an `idp-certificate-authority-data` to the kubectl config to allow it to renew tokens
when the IDP uses a self-signed CA. Merged from https://github.com/vmware-archive/gangly/pull/94


### Allow custom assets to be used in custom templates

Adds the `customAssetsDir` config option to override the contents of /assets/ for use in 
custom templates.


### Override hard-coded session salt

Adds the `sessionSalt` config option to override the hard-coded salt. Min. length is 8 characters.
Fixes https://github.com/vmware-archive/gangly/issues/71

### todo

...

### Other minor stuff

* Update to Go 1.21
* Update dependencies
* Replace esc with go:embed
* Root path check (https://github.com/vmware-archive/gangly/pull/143)
* Fix URL encoding (https://github.com/vmware-archive/gangly/pull/179)
* BREAKING - corrected ENV variable name of `customHTMLTemplatesDir` to `CUSTOM_HTML_TEMPLATES_DIR`,
  this was (incorrectly) `CUSTOM_HTTP_TEMPLATES_DIR`
* Config option `showClaims` to show/hide received claims
* Validate JWT signatures
* Change `authorizeURL` to `ProviderURL` such as `https://accounts.google.com`

## v3.3.0 (2021-07-15)

All the stuff it did before, see https://github.com/vmware-archive/gangly