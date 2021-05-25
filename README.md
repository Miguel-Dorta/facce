# facce
Firebase Auth Custom-Claims Editor.

facce is a CLI tool for editing the custom claims of the Firebase Auth users.

## How to use it
### Prerequisites
You need to download a file with the credentials for your project. You can do it [by clicking here](https://console.firebase.google.com/project/_/settings/serviceaccounts/adminsdk) or going to Firebase Console > Settings > Service Accounts.

### Download and unzip the binary
You can do it [by clicking here](https://github.com/Miguel-Dorta/facce/releases/).

## Example commands
``` bash
# Print all the custom claims of the user
./facce get -c credentials-file.json --uid 0123456789abcdefghijklmnopqr

# Print the field "setting" of the object "example"
./facce get example.setting -c credentials-file.json --uid 0123456789abcdefghijklmnopqr

# Set a field "admin" to true
./facce set admin=true -c credentials-file.json --uid 0123456789abcdefghijklmnopqr

# Remove the field "role"
./facce unset role -c credentials-file.json --uid 0123456789abcdefghijklmnopqr
```
