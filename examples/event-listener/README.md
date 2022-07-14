### Event Listener

#### Required environment variables
- **MSP_ID** - MSP identifier
- **CONFIG_PATH** - path to configuration file of SDK _(eg. config.yaml)_
- **CERT_PATH** - path to certificate file of identity _(eg. signcerts/cert.pem)_
- **KEY_PATH** - path to private key of identity _(e.g. keystore/your_pk)_
- **CHANNEL** - channel used for listening new events
- **CHAINCODE** - chaincode used for used listening chaincode events

#### Usage
Run with environment variables described above:
```bash
go run main.go
```