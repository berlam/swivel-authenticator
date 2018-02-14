swivel-authenticator
====================
An alternative to the Swivel Mobile authenticator from [SwivelSecure](https://swivelsecure.com/), which uses the officially documented API.
Currently this project does not cover every functionality of the Swivel Mobile App but it satisfies my use case.
Feel free to extend it for yours.

I am using it to connect to a VPN, which asks for the OTC. The call looks like this and connects instantly to the VPN.
```Shell
$ swivelt $SWIVEL_PROVISION_ID $SWIVEL_USER_PIN | vpnc $VPNC_CONFIG_LOCATION
```

Settings and keys are stored in:
`$HOME/swivel-$SWIVEL_SERVER_ID`

## Arguments ##
* Provision
```Shell
$ swivelp $SWIVEL_SERVER_ID $SWIVEL_USERNAME $SWIVEL_PROVISION_CODE
```
* OTC
```Shell
$ swivelt $SWIVEL_SERVER_ID $SWIVEL_USER_PIN 
```

You can disable certificate validation with '--no-verify' as last argument.

## Build from source ##
If you want to build Swivel Authenticator right away you need to have a working [Go environment](https://golang.org/doc/install).
```Shell
$ go get -d github.com/berlam/swivel-authenticator/cmd/{swivelp,swivelt}
$ go install github.com/berlam/swivel-authenticator/{pkg,cmd/swivelp,cmd/swivelt}
```

## Authentication API ##
Here are the Swivel APIs used for the authentication.

### GET Server Details with ServerId ###
`GET https://ssd.swivelsecure.net/ssdserver/getServerDetails?id=$SWIVEL_SERVER_ID`

### POST to get provisioning message ###
`POST $SCHEME://$HOSTNAME:$PORT/$CONNECTION_TYPE/AgentXML`
```xml
<?xml version='1.0' ?>
<SASRequest>
	<Version>3.6</Version>
	<Action>provision</Action>
	<Username>$SWIVEL_USERNAME</Username>
	<ProvisionCode>$SWIVEL_PROVISION_CODE</ProvisionCode>
</SASRequest>
```

### POST to get new security keys ###
`POST $SCHEME://$HOSTNAME:$PORT/$CONNECTION_TYPE/AgentXML`
```xml
<?xml version='1.0' ?>
<SASRequest>
	<Version>3.6</Version>
	<Action>SecurityStrings</Action>
	<Id>$SWIVEL_PROVISION_ID</Id>
</SASRequest>
```

## OTC ##
Security keys are always 10 digits.

The OTC is generated as follows:

1. Get the index of the last security key used. Decrease it by one.
2. Get the security key at position of the index.
3. Get the number at each number position of the PIN (0 is position 10).
4. Append the index as two digits number.

Example: PIN is 9132. Security key at index 6 is 0123456789. Result is: 802106. [See Test](pkg/token_test.go).

Note: The keys will be updated, when the index reaches zero. You will get 100 new security keys.

## Resources ##
[Swivel Knowledge Base Authentication API](https://kb.swivelsecure.com/wiki/index.php/AuthenticationAPI)

## Credits ##

This project was created by [@berlam](https://github.com/berlam).

## License ##

See [LICENSE](LICENSE).
