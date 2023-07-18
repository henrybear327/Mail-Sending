# README

> In memory of our time at Taipei office - with Anson Chen and Steven Lin

Original source is [here](https://github.com/xxi511/mailSender).

# Set up

- Set up `.bashrc` or whatever shell config you are using.
  - If you want to log in with a username and password, then set `PROTON_API_BRIDGE_TEST_USE_REUSABLE_LOGIN` to 0. After one successful login, you will see a `.credential` file, which you can use to fill `PROTON_API_BRIDGE_TEST_UID`, `PROTON_API_BRIDGE_TEST_ACCESS_TOKEN`, `PROTON_API_BRIDGE_TEST_REFRESH_TOKEN`, and `PROTON_API_BRIDGE_TEST_SALTEDKEYPASS` (don't forget to set `PROTON_API_BRIDGE_TEST_USE_REUSABLE_LOGIN` to 1 if you want to use access token for log in).
  - Don't forget to `source ~/.bashrc` after making changes :) 
```bash
export PROTON_API_BRIDGE_TEST_USERNAME="a@proton.me"
export PROTON_API_BRIDGE_TEST_PASSWORD='password'
export PROTON_API_BRIDGE_TEST_TWOFA="123456"
export PROTON_API_BRIDGE_APP_VERSION="macos-drive@5.0.14.2"
export PROTON_API_BRIDGE_USER_AGENT=""
export PROTON_API_BRIDGE_TEST_USE_REUSABLE_LOGIN=0
export PROTON_API_BRIDGE_TEST_UID=""
export PROTON_API_BRIDGE_TEST_ACCESS_TOKEN=""
export PROTON_API_BRIDGE_TEST_REFRESH_TOKEN=""
export PROTON_API_BRIDGE_TEST_SALTEDKEYPASS=""
```
- `go run .`
