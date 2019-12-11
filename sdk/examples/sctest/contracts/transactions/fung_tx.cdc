import FungibleToken from 0x01

transaction {
  prepare(acct: Account) {
    var minterRef = acct.storage[&FungibleToken.VaultMinter] ?? panic("missing minter reference!")

    let recipient = getAccount(0x01)

    let receiverRef = recipient.published[&FungibleToken.Receiver] ?? panic("missing receiver ref!")
    
    //minterRef.mintTokens(amount: 30, recipient: receiverRef)
  
  }
}