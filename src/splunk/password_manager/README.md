# Overview
This app is a utility for managing the passwords stored in `passwords.conf` on a Splunk server. The app is an example of creating a Splunk dashboard that utilizes WebAssembly. For more information about WebAssembly in general visit https://webassembly.org/. The WASM binary is written in Go and loaded by the `Password Manager` dashboard.

# Pre-requisites
In order to use the `Password Manager` dashboard the user must utilize a browser capable of executing WASM binaries. Most modern browsers support this capability but for a list of compatible browsers visit https://developer.mozilla.org/en-US/docs/WebAssembly#Browser_compatibility.  

# Installation
Installation is identical to all other Splunk apps. Extract the archive into the `$SPLUNK_HOME/etc/apps` directory. The WASM binary is pre-built and included in the app's `appserver/static` directory.

# Using Password Manager
The app contains a single dashboard called `Password Manager`. After opening the dashboard the user will be presented with a simple table of all passwords stored in the Splunk server, listing the app, realm, username, and masked password. Each found password will have two buttons to the right of the entry; `Reveal` and `Delete`. Clicking the `Reveal` button will remove the mask from the password and show the cleartext. Clicking the `Delete` button will delete the password from Splunk.

Below the table there is a series of form inputs for creating new passwords. Once the relevant fields of the form are complete the user must click the `Add` button in order to save the password. The page will reload, and the password will be visible in the table.

# Binary File Declaration
- password_manager.wasm: https://github.com/badarsebard/splunk-password-manager