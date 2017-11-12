swivel-authenticator
====================
An alternative to the Swivel Mobile authenticator from [SwivelSecure](https://swivelsecure.com/), which uses the officially documented API.
Currently this project does not cover every functionality of the Swivel Mobile App but it satisfies my use case.
Feel free to extend it for yours.

I am using it to connect to a VPN, which asks for the OTC. The call looks like this and connects instantly to the VPN.
```Shell
swivelt $SWIVEL_PROVISION_ID $SWIVEL_USER_PIN | vpnc $VPNC_CONFIG_LOCATION
```

Settings and keys are stored in:
`$HOME/swivel-$SWIVEL_SERVER_ID`

## Arguments ##
* Provision
```Shell
swivelp $SWIVEL_SERVER_ID $SWIVEL_USERNAME $SWIVEL_PROVISION_CODE
```
* OTC
```Shell
swivelt $SWIVEL_SERVER_ID $SWIVEL_USER_PIN 
```

## Build from source ##
1. install `go`
2. run `go {build,install} $GIT_ROOT/src/github.com/berlam/swivel-authenticator/{domain,cmd/swivelp,cmd/swivelt}`

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
3. Get the number at each number position of the PIN. Example: PIN is 9132. Security key is 0123456789. Result is: 0821.
4. Append the index as two digits number. Example: Index is 6, Result is: 082106.

Note: The keys will be updated, when the index reaches zero. You will get 100 new security keys.

## Resources ##
[Swivel Knowledge Base Authentication API](https://kb.swivelsecure.com/wiki/index.php/AuthenticationAPI)

## Credits ##

This project was created by [@berlam](https://github.com/berlam).

## License ##

See [LICENSE](LICENSE).
