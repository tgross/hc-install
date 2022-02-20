package product

const vaultName = "vault"

var Vault = Product{
	Name:              vaultName,
	BinaryName:        binaryName(vaultName),
	GetVersion:        nil,
	BuildInstructions: nil,
}
