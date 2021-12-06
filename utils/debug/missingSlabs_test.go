package debug

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/model/flow"
)

func TestDebugger_RunTestTransaction(t *testing.T) {

	// this code is mostly a sample code so we skip by default
	//t.Skip()

	grpcAddress := "35.209.130.97:9000"
	chain := flow.Emulator.Chain()
	debugger := NewRemoteDebugger(grpcAddress, chain, zerolog.New(os.Stdout).With().Logger())

	const scriptTemplate = `
        import FanTopToken from 0x48602d8056ff9d93
        import FanTopPermission from 0x48602d8056ff9d93

        transaction() {
            let minterRef: &FanTopPermission.Minter

            prepare(account: AuthAccount) {
                self.minterRef = account.borrow<&FanTopPermission.Holder>(from: /storage/FanTopPermission)?.borrowMinter(by: account)
                    ?? panic("No minter in storage")
            }
	    }
	`

	// Failing account
	address := flow.HexToAddress("434a1f199a7ae3ba")

	// Other account
	//address := flow.HexToAddress("b9047a906effe57b")

	script := []byte(scriptTemplate)
	txBody := flow.NewTransactionBody().
		SetGasLimit(9999).
		SetScript([]byte(script)).
		SetPayer(address).
		SetProposalKey(address, 0, 0)
	txBody.Authorizers = []flow.Address{address}

	// Run at the latest blockID
	txErr, err := debugger.RunTransaction(txBody)
	require.NoError(t, txErr)
	require.NoError(t, err)
}
