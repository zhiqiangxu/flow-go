package fvm_test

import (
	"fmt"
	"sync"
	"testing"
	"github.com/onflow/cadence/runtime"
	"github.com/onflow/cadence/runtime/ast"
	"github.com/onflow/cadence/runtime/common"
	"github.com/onflow/cadence/runtime/sema"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/engine/execution/testutil"
	"github.com/onflow/flow-go/fvm"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/model/hash"
	"github.com/onflow/flow-go/utils/unittest"
)

const CacheSize = 256

func TestTransactionASTCache(t *testing.T) {
	rt := runtime.NewInterpreterRuntime()

	chain := flow.Testnet.Chain()

	vm := fvm.New(rt)

	cache, err := fvm.NewLRUASTCache(CacheSize)
	require.NoError(t, err)

	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain), fvm.WithASTCache(cache))

	t.Run("transaction execution results in cached program", func(t *testing.T) {
		txBody := flow.NewTransactionBody().
			SetScript([]byte(`
                transaction {
                  prepare(signer: AuthAccount) {}
                }
            `)).
			AddAuthorizer(unittest.AddressFixture())

		err := testutil.SignTransactionAsServiceAccount(txBody, 0, chain)
		require.NoError(t, err)

		ledger := testutil.RootBootstrappedLedger(vm, ctx)

		tx := fvm.Transaction(txBody, 0)

		err = vm.Run(ctx, tx, ledger)
		require.NoError(t, err)

		assert.NoError(t, tx.Err)

		// Determine location of transaction
		txID := txBody.ID()
		location := common.TransactionLocation(txID[:])

		// Get cached program
		program, err := cache.GetProgram(location)
		require.NotNil(t, program)
		require.NoError(t, err)
	})
}

func TestScriptASTCache(t *testing.T) {
	rt := runtime.NewInterpreterRuntime()

	chain := flow.Testnet.Chain()

	vm := fvm.New(rt)

	cache, err := fvm.NewLRUASTCache(CacheSize)
	require.NoError(t, err)

	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain), fvm.WithASTCache(cache))

	t.Run("script execution results in cached program", func(t *testing.T) {
		code := []byte(`
			pub fun main(): Int {
				return 42
			}
		`)

		ledger := testutil.RootBootstrappedLedger(vm, ctx)

		script := fvm.Script(code)

		err := vm.Run(ctx, script, ledger)
		require.NoError(t, err)

		assert.NoError(t, script.Err)

		// Determine location
		scriptHash := hash.DefaultHasher.ComputeHash(code)
		location := common.ScriptLocation(scriptHash)

		// Get cached program
		program, err := cache.GetProgram(location)
		require.NotNil(t, program)
		require.NoError(t, err)

	})
}

func TestTransactionWithProgramASTCache(t *testing.T) {
	rt := runtime.NewInterpreterRuntime()

	chain := flow.Testnet.Chain()

	vm := fvm.New(rt)

	cache, err := fvm.NewLRUASTCache(CacheSize)
	require.NoError(t, err)

	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain), fvm.WithASTCache(cache))

	// Create a number of account private keys.
	privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
	require.NoError(t, err)

	// Bootstrap a ledger, creating accounts with the provided private keys and the root account.
	ledger := testutil.RootBootstrappedLedger(vm, ctx)
	accounts, err := testutil.CreateAccounts(vm, ledger, privateKeys, chain)
	require.NoError(t, err)

	// Create deployment transaction that imports the FlowToken contract
	txBody := flow.NewTransactionBody().
		SetScript([]byte(fmt.Sprintf(`
				import FlowToken from 0x%s
				transaction {
					prepare(signer: AuthAccount) {}
					execute {
						let v <- FlowToken.createEmptyVault()
						destroy v
					}
				}
			`, fvm.FlowTokenAddress(chain))),
		).
		AddAuthorizer(accounts[0]).
		SetProposalKey(accounts[0], 0, 0).
		SetPayer(chain.ServiceAddress())

	err = testutil.SignPayload(txBody, accounts[0], privateKeys[0])
	require.NoError(t, err)

	err = testutil.SignEnvelope(txBody, chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
	require.NoError(t, err)

	// Run the Use import (FT Vault resource) transaction

	tx := fvm.Transaction(txBody, 0)

	err = vm.Run(ctx, tx, ledger)
	require.NoError(t, err)

	assert.NoError(t, tx.Err)

	// Determine location of transaction
	txID := txBody.ID()
	location := common.TransactionLocation(txID[:])

	// Get cached program
	program, err := cache.GetProgram(location)
	require.NotNil(t, program)
	require.NoError(t, err)
}

func TestTransactionWithProgramASTCacheConsistentRegTouches(t *testing.T) {
	createLedgerOps := func(withCache bool) map[string]bool {
		rt := runtime.NewInterpreterRuntime()
		h := unittest.BlockHeaderFixture()

		chain := flow.Testnet.Chain()

		vm := fvm.New(rt)

		cache, err := fvm.NewLRUASTCache(CacheSize)
		require.NoError(t, err)

		options := []fvm.Option{
			fvm.WithChain(chain),
			fvm.WithBlockHeader(&h),
		}

		if withCache {
			options = append(options, fvm.WithASTCache(cache))
		}

		ctx := fvm.NewContext(zerolog.Nop(), options...)

		// Create a number of account private keys.
		privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
		require.NoError(t, err)

		// Bootstrap a ledger, creating accounts with the provided private keys and the root account.
		ledger := testutil.RootBootstrappedLedger(vm, ctx)
		accounts, err := testutil.CreateAccounts(vm, ledger, privateKeys, chain)
		require.NoError(t, err)

		// Create deployment transaction that imports the FlowToken contract
		txBody := flow.NewTransactionBody().
			SetScript([]byte(fmt.Sprintf(`
					import FlowToken from 0x%s
					transaction {
						prepare(signer: AuthAccount) {}
						execute {
							let v <- FlowToken.createEmptyVault()
							destroy v
						}
					}
				`, fvm.FlowTokenAddress(chain))),
			).
			AddAuthorizer(accounts[0]).
			SetProposalKey(accounts[0], 0, 0).
			SetPayer(chain.ServiceAddress())

		err = testutil.SignPayload(txBody, accounts[0], privateKeys[0])
		require.NoError(t, err)

		err = testutil.SignEnvelope(txBody, chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
		require.NoError(t, err)

		// Run the Use import (FT Vault resource) transaction

		tx := fvm.Transaction(txBody, 0)

		err = vm.Run(ctx, tx, ledger)
		require.NoError(t, err)

		assert.NoError(t, tx.Err)

		return ledger.RegisterTouches
	}

	assert.Equal(t,
		createLedgerOps(true),
		createLedgerOps(false),
	)
}

func BenchmarkTransactionWithProgramASTCache(b *testing.B) {
	rt := runtime.NewInterpreterRuntime()

	chain := flow.Testnet.Chain()

	vm := fvm.New(rt)

	cache, err := fvm.NewLRUASTCache(CacheSize)
	require.NoError(b, err)

	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain), fvm.WithASTCache(cache))

	// Create a number of account private keys.
	privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
	require.NoError(b, err)

	// Bootstrap a ledger, creating accounts with the provided private keys and the root account.
	ledger := testutil.RootBootstrappedLedger(vm, ctx)
	accounts, err := testutil.CreateAccounts(vm, ledger, privateKeys, chain)
	require.NoError(b, err)

	// Create many transactions that import the FlowToken contract.
	var txs []*flow.TransactionBody

	for i := 0; i < 1000; i++ {
		tx := flow.NewTransactionBody().
			SetScript([]byte(fmt.Sprintf(`
				import FlowToken from 0x%s
				transaction {
					prepare(signer: AuthAccount) {}
					execute {
						log("Transaction %d")
						let v <- FlowToken.createEmptyVault()
						destroy v
					}
				}
			`, fvm.FlowTokenAddress(chain), i)),
			).
			AddAuthorizer(accounts[0]).
			SetProposalKey(accounts[0], 0, uint64(i)).
			SetPayer(chain.ServiceAddress())

		err = testutil.SignPayload(tx, accounts[0], privateKeys[0])
		require.NoError(b, err)

		err = testutil.SignEnvelope(tx, chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
		require.NoError(b, err)

		txs = append(txs, tx)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j, txBody := range txs {
			// Run the Use import (FT Vault resource) transaction.
			tx := fvm.Transaction(txBody, uint32(j))

			err := vm.Run(ctx, tx, ledger)
			assert.NoError(b, err)

			assert.NoError(b, tx.Err)
		}
	}

}

type nonFunctioningCache struct{}

func (cache *nonFunctioningCache) GetProgram(_ common.Location) (*ast.Program, error) {
	return nil, nil
}

func (cache *nonFunctioningCache) SetProgram(_ common.Location, _ *ast.Program) error {
	return nil
}

func (cache *nonFunctioningCache) GetElaboration(_ common.Location) (*sema.Checker, error) {
	return nil, nil
}

func (cache *nonFunctioningCache) SetElaboration(_ common.Location, _ *sema.Checker) error {
	return nil
}

func (cache *nonFunctioningCache) Clear() {
}

func BenchmarkTransactionWithoutProgramASTCache(b *testing.B) {
	rt := runtime.NewInterpreterRuntime()

	chain := flow.Testnet.Chain()

	vm := fvm.New(rt)

	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain), fvm.WithASTCache(&nonFunctioningCache{}))

	// Create a number of account private keys.
	privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
	require.NoError(b, err)

	// Bootstrap a ledger, creating accounts with the provided private keys and the root account.
	ledger := testutil.RootBootstrappedLedger(vm, ctx)
	accounts, err := testutil.CreateAccounts(vm, ledger, privateKeys, chain)
	require.NoError(b, err)

	// Create many transactions that import the FlowToken contract.
	var txs []*flow.TransactionBody

	for i := 0; i < 1000; i++ {
		tx := flow.NewTransactionBody().
			SetScript([]byte(fmt.Sprintf(`
				import FlowToken from 0x%s
				transaction {
					prepare(signer: AuthAccount) {}
					execute {
						log("Transaction %d")
						let v <- FlowToken.createEmptyVault()
						destroy v
					}
				}
			`, fvm.FlowTokenAddress(chain), i)),
			).
			AddAuthorizer(accounts[0]).
			SetPayer(accounts[0]).
			SetProposalKey(accounts[0], 0, uint64(i))

		err = testutil.SignEnvelope(tx, accounts[0], privateKeys[0])
		require.NoError(b, err)

		txs = append(txs, tx)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j, txBody := range txs {
			// Run the Use import (FT Vault resource) transaction.
			tx := fvm.Transaction(txBody, uint32(j))

			err := vm.Run(ctx, tx, ledger)
			assert.NoError(b, err)

			assert.NoError(b, tx.Err)
		}
	}
}

func TestProgramASTCacheAvoidRaceCondition(t *testing.T) {
	rt := runtime.NewInterpreterRuntime()

	chain := flow.Testnet.Chain()

	vm := fvm.New(rt)

	cache, err := fvm.NewLRUASTCache(CacheSize)
	require.NoError(t, err)

	ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain), fvm.WithASTCache(cache))

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)

		ledger := testutil.RootBootstrappedLedger(vm, ctx)

		go func(id int, wg *sync.WaitGroup) {
			defer wg.Done()

			code := []byte(fmt.Sprintf(`
				import FlowToken from 0x%s
				pub fun main() {
					log("Script %d")
					let v <- FlowToken.createEmptyVault()
					destroy v
				}
			`, fvm.FlowTokenAddress(chain), id))

			script := fvm.Script(code)

			err := vm.Run(ctx, script, ledger)
			require.NoError(t, err)

			assert.NoError(t, script.Err)
		}(i, &wg)
	}
	wg.Wait()

	location := common.AddressLocation{Address: runtime.Address(fvm.FlowTokenAddress(chain)), Name: "FlowToken"}

	// Get cached program
	var program *ast.Program
	program, err = cache.GetProgram(location)
	require.NotNil(t, program)
	require.NoError(t, err)
}

func TestASTCacheInvalidation(t *testing.T) {
	rt := runtime.NewInterpreterRuntime()
	vm := fvm.New(rt)
	chain := flow.Testnet.Chain()

	lruCache, err := fvm.NewLRUASTCache(CacheSize)
	require.NoError(t, err)
	internalCache, err := fvm.NewLRUASTCache(CacheSize)
	require.NoError(t, err)
	nonClearingCache := &nonClearingCache{internalCache: internalCache}

	// Run the test with two types of caches. A one that properly invalidates
	// the cache (LRUASTCache), and a one that does not invalidate the cache.
	caches := []fvm.ASTCache{lruCache, nonClearingCache}

	var txErrors []fvm.Error

	for _, cache := range caches {

		ctx := fvm.NewContext(zerolog.Nop(), fvm.WithChain(chain), fvm.WithASTCache(cache))

		// Create an account private key.
		privateKeys, err := testutil.GenerateAccountPrivateKeys(1)
		require.NoError(t, err)

		// Bootstrap a ledger, creating accounts with the provided private keys and the root account.
		ledger := testutil.RootBootstrappedLedger(vm, ctx)
		accounts, err := testutil.CreateAccounts(vm, ledger, privateKeys, chain)
		require.NoError(t, err)

		deployContract := func(
			script string,
			name string,
			sequenceNum uint64,
			update bool) {

			var txBody *flow.TransactionBody
			if update {
				txBody = testutil.CreateContractUpdateTransaction(name, script, accounts[0], chain)
			} else {
				txBody = testutil.CreateContractDeploymentTransaction(name, script, accounts[0], chain)
			}

			txBody.SetProposalKey(chain.ServiceAddress(), 0, sequenceNum)
			txBody.SetPayer(chain.ServiceAddress())

			err := testutil.SignPayload(txBody, accounts[0], privateKeys[0])
			require.NoError(t, err)

			err = testutil.SignEnvelope(txBody, chain.ServiceAddress(), unittest.ServiceAccountPrivateKey)
			require.NoError(t, err)

			tx := fvm.Transaction(txBody, 0)
			err = vm.Run(ctx, tx, ledger)
			assert.NoError(t, err)
			assert.NoError(t, tx.Err)
		}

		executeTransaction := func(sequenceNum uint64) fvm.Error {
			account := accounts[0]
			script := fmt.Sprintf(`
				import Foo from 0x%s
				transaction {
					prepare(acct: AuthAccount) {}
					execute {
						log(Foo.sayHello())
					}
				}
				`, account)

			txBody := flow.NewTransactionBody().
				SetScript([]byte(script)).
				AddAuthorizer(account).
				SetPayer(account).
				SetProposalKey(account, 0, sequenceNum)
			err := testutil.SignEnvelope(txBody, account, privateKeys[0])
			require.NoError(t, err)

			tx := fvm.Transaction(txBody, 0)
			err = vm.Run(ctx, tx, ledger)

			assert.NoError(t, err)
			return tx.Err
		}

		// Create a dummy contract 'Bar' for importing purpose
		script := fmt.Sprintf(`
		access(all) contract Bar {
			pub fun getValue(): Int {
				return 5
			}
		}`)
		deployContract(script, "Bar", 0, false)

		// Deploy another contract 'Foo' that imports 'Bar'
		script = fmt.Sprintf(`
		import Bar from 0x%s
		access(all) contract Foo {
			pub fun sayHello(): String {
				return "Hello"
			}

			// An unused method which uses 'sayHello' coming from 'Bar'.
			// Thus, removing 'sayHello' would only affect the elaboration
			// but not the runtime/intepreter.
			priv fun unusedMethod() {
				Bar.getValue()
			}
		}`, accounts[0])
		deployContract(script, "Foo", 1, false)

		// Execute a transaction that imports 'Foo'.
		txError := executeTransaction(0)
		assert.NoError(t, txError)

		// Update the transitive dependency 'Bar' contract, and remove the 'sayHello' method.
		script = fmt.Sprintf(`access(all) contract Bar { }`)
		deployContract(script, "Bar", 2, true)

		// Re-run the transaction that imports Foo.
		txError = executeTransaction(1)
		txErrors = append(txErrors, txError)
	}

	// Case I (LRUASTCache):
	// The cache is invalidated, and the checker will re-check all imports. Since a method
	// is removed in a transitive import, the transaction should fail with a checking error.
	require.NotNil(t, txErrors[0], "Transaction expected to be failed")
	require.Error(t, txErrors[0])
	assert.Contains(t, txErrors[0].Error(), "error: value of type `Bar` has no member `getValue`")

	// Case II (nonClearingCache):
	// The caching does not get invalidated, and the old 'Foo' contract is picked from the cache.
	// Thus, the transaction will eventually succeed despite having an invalid code.
	assert.NoError(t, txErrors[1], "Transaction expected to be succeed")
}

type nonClearingCache struct {
	internalCache fvm.ASTCache
}

func (cache *nonClearingCache) GetProgram(location common.Location) (*ast.Program, error) {
	return cache.internalCache.GetProgram(location)
}

func (cache *nonClearingCache) SetProgram(location common.Location, elaboration *ast.Program) error {
	return cache.internalCache.SetProgram(location, elaboration)
}

func (cache *nonClearingCache) GetElaboration(location common.Location) (*sema.Checker, error) {
	return cache.internalCache.GetElaboration(location)
}

func (cache *nonClearingCache) SetElaboration(location common.Location, elaboration *sema.Checker) error {
	return cache.internalCache.SetElaboration(location, elaboration)
}

func (cache *nonClearingCache) Clear() {
	// Do not clear the cache
}
