
×General Crypto
Add item member
Review of Distributed Ledger Technologies
Bitcoin

Link: https://bitcoin.org/bitcoin.pdf 

Description: The original blockchain

The algorithm: Longest chain wins

Consensus Mechanism: PoW

Limitations: Power usage and fees are very high. Transaction rates aren't high enough.

How BTC Transactions are signed.

 



 

Bitcoin Cash
Link: https://www.bitcoincash.org/bitcoin.pdf  

Description: The original blockchain with one difference: 8mb block sizes.

The algorithm: Longest chain wins

Consensus Mechanism: PoW

Limitations: Power usage. Transaction rates aren't high enough.

 

Bitcoin Gold
Link: https://bitcoingold.org/wp-content/uploads/2017/10/BitcoinGold-Roadmap.pdf 

Description: Bitcoin Gold is a fork of Bitcoin with replay protection, 10 min blocks, DigiShield protection, EQUIHASH PoW, and funds for community outreach. There has been complaints about Bitcoin Gold because the founders effectively premined a lot of it to pay for *development* but ultimately it was a cash grab.

Consensus Mechanism: EQUIHASH

Limitations: 1 MB blocks at 10 min. Transaction rates aren't high enough.

 

GHOST
Link: https://eprint.iacr.org/2013/881.pdf 

Description: A modification of the 'longest chain' algorithm to ensure security in systems that have fast transaction times. Blockchains with high block creation rates suffer from the time it takes blocks to propigate through the network. Therefore, new valid blocks become stale quickly allowing strong miners to gain an unfair advantage or pool an unfair advantage.

The algorithm: Orphaned (or stale) blocks are included in the longest chain 'weight'. 

Consensus Mechanism: PoW

Limitations: The GHOST protocol doesn't deal with increased centralization in high transaction blockchains.

 

"If miner A is a mining pool with 30% hashpower and B has 10% hashpower, A will have a risk of producing a stale block 70% of the time (since the other 30% of the time A produced the last block and so will get mining data immediately) whereas B will have a risk of producing a stale block 90% of the time."

 

It is proposed to fix this issue by using a DAG in the Includive Protocols white paper: http://fc15.ifca.ai/preproceedings/paper_101.pdf

 

Ethereum
Link: http://www.the-blockchain.com/docs/Ethereum_white_paper-a_next_generation_smart_contract_and_decentralized_application_platform-vitalik-buterin.pdf 

Description: The original smart contract blockchain that uses GAS to charge contracts to execute.

The algorithm: Longest chain wins (similar to Bitcoin) however a modified GHOST protocol is used. The modified version allows stale blocks to receive rewards to compensate for centralization bias.

Consensus Mechanism: PoW until Casper is released.

Limitations: High transaction volume, GAS can be expensive, turing completeness allows people to be 'too clever by half'.

 

Casper
Link: https://blog.ethereum.org/2015/08/01/introducing-casper-friendly-ghost/

Description: A proof of stake system for Ethereum. Casper bonds validators who either borrower or purchase their way into their positions. They then benefit from being a validator by being paid.

The algorithm: Everyone can vote on each block at each level. Votes will determine the order in which transactions are validated by choosing the block to be mined first. If validators vote improperly or validate invalid blocks they forfeit their security deposit. Validators are incentivized to ensure other validators will vote with them ensuring cooperation. 

Consensus Mechanism: PoS

Limitations: Super high throughput is not possible with 250 geographically distributed validators voting in rounds. 

 

LightningNetwork
Link: https://lightning.network/lightning-network-paper.pdf 

Description: The LightningNetwork allows off-chain transactions which can be executed immediately. When the time comes, the off-chain transactions can be settled on the blockchain.

The algorithm: Susie and Joe want to transact with each other. They both multi-sig some BTC into an address. Then, as Susie and Joe pay one to the other they author new txns but to not broadcast them. After a few weeks it may be that Susie wants to settle with Joe so she broadcasts the txns onto the blockchain and the money is distributed appropriately. The BTC that is in the network or a specific payment channel is 'locked' and cannot be double spent. 

Consensus Mechanism: None until settled on Bitcoin's blockchain. PoW indirectly.

Limitations: Off chain transactions require that the multi-sig wallet has BTC in advance. Requires a payment to join and leave the network. Leaves avenue for centralization as third parties manage a series of payment channels. Lastly, to be able to reach everyone on the lightning network, one would have to deposit 10x the amount of BTC for a max payment (to anyone). Here for more info.

 

SPECTRE
Link: https://eprint.iacr.org/2016/1159.pdf

Description: A DAG (Directed Acyclic Graph) of blocks that are added to the graph asynchronously. It allows for extremely high throughput while maintaining increased security/protection from a 51% style attack.

The algorithm: Given a block x there are votes for blocks in past(x) and also future(x) (i.e by the time my node is voting others have already voted for x and I see that as future(x)). An honest node only votes for valid blocks and for who's blocks are a descendant of the genesis block. If 50% of past(x) + future(x) + the node itself vote for a block it is confirmed. To be able to vote, each node needs to solve a simple PoW problem.

 

To be more precise: With high probability, if a node accepts your transaction it is extremely likely it will be accepted by the network as long as it is honest. This is explained in detail in the Appendix.

Consensus Mechanism: PoW and Pairwise Voting (described briefly above and in detail in the Appendix).

Limitations: Vulnerable to a 51% attack even though transaction speeds don't increase the threat.

 

IOTA
Link: https://iota.org/IOTA_Whitepaper.pdf 

Description: Very similar to Spectre however based on nodes for transactions rather than blocks. DAG of all txns is called the Tangle. 

The algorithm: For devices to issue a txn they must do a small amount of PoW and approve two other transactions. By issuing txns (i.e approving 2 others) nodes contribute to the security of the Tangle. Once an acceptable set of nodes have approved a transaction is becomes irrefutable. Transactions have weights and their cumulative weight is the weight of all txns that have approved it as well as itself. The cumulative weight is used to prevent double spending attacks since honest nodes won't approve if these txns. Therefore, those txns will always have a lower cumulative weight than honest txns.

Consensus Mechanism: PoW and Cumulative Weighting or txns.

 Limitations: Centralized implementation. Vulnerable to double spends. Vulnerable to a 51% attack. Since the Tangle is a graph of txns and not blocks it may have storage and network issues due to volume.

 

ZCash
Link: http://zerocash-project.org/media/pdf/zerocash-extended-20140518.pdf

Description: zCash is a optionally private cryptocurrency where senders, recipients, and amounts are hidden from view. It uses zk-SNARKs to obscure all details of transactions. zk-SNARKs is a cryptographic system that can prove hidden details of a transaction while revealing some public details. Using zk-SNARKs, zCash can prove that there is a specific valid transaction in the ledger while hiding all other details from anyone but the recipient (also hidden to all but the recipient).

The algorithm: Ultimately it is a zk-SNARKs that makes a Pour operation possible. Pouring allows users to subdivide coins into smaller denominations, merge coins, and transfer ownership of anonymous coins, or make public payments. It also hides the identity of the sender and the amount. The recipient publishes a public key which is to receive the anonymous coins. The end result is that there is no way to know whom set what to a public key that can be owned by anyone and anonymously re-spent.

Consensus Mechanism: ASIC resistant PoW. 

Limitations: zk-SNARKs proofs are CPU to intensive to generate and, therefore, unusable on mobile devices. Additionally, completely private systems hide all information even from those that would like to maintain the network. It is impossible for those who administrate the network, the zCash foundation, or their developers to know what is happening. Thus, they cannot fix problems that might arise such as counterfeiting.

 

BTC Private
Link: https://btcprivate.org/whitepaper.pdf 

Description: BTC Private is a fork of ZClassic which is a fork of ZCash. The fork doesn't have the 20% miner reward back to the foundation. What's interesting is that BTC Private is replaying BTCs TXNs and ZClassic's to give 1:1 coins for the other coins. It also employs replay attack protection by denoting the SIGHASH_FORKID. BTC Private also uses BIP 9 to make democratic decisions. EQUIHASH is used to be more democratic and to not have ASIC farms take it over. Blocks are 2MB and block times are 2.5 minutes. Contributions to the foundation are voluntary from the miners. 

Consensus Mechanism: EQUIHASH

Limitations: Same limitations as zCash. zkSNARKs are CPU intensive. Private chains give you no insight into what's happening. There is an additional limitation in that BTC Private relies on zCash for innovation.

 

Monero
Link: https://downloads.getmonero.org/whitepaper_annotated.pdf 

Description: Monero is an anonymous currency that uses Ring Signatures to hide senders of coins and, also, uses several heuristics to keep the recipient anonymous. Ring signatures are a group or randomly chosen public keys and one real/signing key grouped together so that you cannot know who actually signed the txn. However, you know that one of them has signed (but, once again, can't tell who). Furthermore, the recipient of a txn is completely obscured so there is no way to trace, through history, who is sending to whom. In Monero parlance, the Ring Signature is Anonymous and the receiving address is Unlinkable (to the actual recipient).

The algorithm: Imagine a txn is signed by 20 signatures but you can't tell which one. Further, imagine the recipient's address is a collision resistant hash that only the recipients private key can unlock. You now have all of the pieces to anonymously send money to the recipient. Sender cannot be determined and nobody can determine who the recipient is. Lastly, the recipient addresses' private key can only be determined by the recipient. Thus, only the recipient can spend the received funds. Meanwhile nobody knows who the recipient actually is.

Consensus Mechanism: ASIC resistant PoW

Limitations: Like zCash, Monero's wallets have to scan every txn to know what an address has received and if a txn was meant for her.

 

OmiseGo

Link: https://cdn.omise.co/omg/whitepaper.pdf 

Description: OmiseGo intends to be a multi-chain exchange system. For example, from ETH->BTC or, ultimately, between anything. It does this by providing a coin who's sole purpose is to fuel the transaction. Meanwhile, the coin has a market cap $1,607,741,629 USD. Since the coin is only used to pay minimal fees for helping facilitate a txn between two currencies (at this time only BTC and ETH are supported) it seems strange the market cap is that high. OmiseGo claims to enable payments between eWallets and the unbanked (presumably in developing countries), however as described in the white paper OmiseGo is not a payments system. Rather, it is simply a system to conduct cross chain transactions.

The algorithm: Conceptually the algorithm is to have an ETH smart contract that tells a clearing house (escrow really) to deliver funds in return for OMG. The clearing house is expected to operate honestly because if it doesn't it gets no OMG. Not only is there an ETH smart contract but also an OMG merkle tree that moves OMG and also functions as a data feed. The data feed is for third parties (presumably to measure market activity). If the txn is just ETH->ETH then there is no clearing house but just a smart contract. At the moment clearing houses are used for anything that isn't ETH->ETH and induce centralization.

Consensus Mechanism: PoS via OMG tokens. Attempts to behave improperly result in your token being burned. Validators have to bond themselves buy purchasing OMG tokens as collateral.

Limitations: Final settlement has to occur over Ethereum in an ETH smart contract. Therefore, people need to buy ETH and incur fees to trade in anything. Validators have to run an Ethereum node as a result. Developing nations need zero fees. If the exchange involves BTC, it uses the LightningNetwork which has its own limitations (see above). 

 

MakerDAO
Link: https://makerdao.com/whitepaper/DaiDec17WP.pdf 

Description: MakerDAO introduces two coins (MKR and DAI). MKR is used to pay fees on the network and DAI is the stable coin portion. Essentially, MakerDAO uses ETH deposits which are ether rewarded or forfeited to stabilize the DAI. At the end of the day, people in the Maker system are paying a fee to keep a stable coin that actually isn't completely stable. It is still affected by price shocks but they tend to be temporary. 

The algorithm: Maker uses Collateralized Debt Positions (CDPs) paid in ETH and used to generate Dai (the stable coin portion). The CDP accrues debt which has to be paid down with MKR and subsequently burned. Once all of the debt is paid off the CDP becomes debt free and the ETH is then unencumbered. Ultimately, CDPs become more or less expensive in relation to Dai being above or below the target price relatively. This causes more or less CDPs to be generated and this, in turn, changes the money supply. CDPs can be liquidated based upon their risk portfolio and their fees can rise and fall causing actors to hold them, get more, or reduce their holdings.

Consensus Mechanism: None specified however since Dai are issued on ETH, ETH might be the consensus mechanism. Meanwhile, there is none specified for MKR.

Limitations: People are paying ETH fees to use the system. One has to question the value. For example, to get a stable coin I have to pay $150 to get $100. 

 

BaseCoin
Link: http://www.getbasecoin.com/basecoin_whitepaper_0_99.pdf 

Description: BaseCoin introduces three coins into its stable coin system (base bonds, base shares, and base coins). Base coins are the stable portion of the system while base bonds are used to restrict the money supply and base shares (in addition to paying out base shares dividends) are used to increase the money supply. 

The algorithm: If the target price goes down they issue BaseBonds to dry up the money supply. If the price goes above the target price, they pay off issued BaseBonds and then issue dividends to BaseShares holders. BaseShares are sold to early adopters. BaseBonds can be made non-redeemable if they are 5 years old. An Oracle is used to determine what the target price for a BASE coin should be.

Consensus Mechanism: There is no discussion of a consensus mechanism.

Limitations: Basecoin relies on demand for its BaseConds to provide liquidity to the market. If that demand dries up, BaseCoin will not work. Also, BaseCoin, like MakerDAO, essentially charges users to ensure the BASE coin is stable. One has to pay to make the currency stable. Oracles are used to get the target price for BaseCoins. Those oracles are points of centralization. 

 

BaseCoin attempts to suggest buying and selling of bonds is how the US treasury affects the money supply and, therefore, the dollar value. This is incorrect. The fed uses:

 

1. Changes in the reserve requirement for fractional reserve lending by banks.

2. Changes in the discount rate at the discount window (basically changing the interest rate the fed charges financial institutions to borrow).

3. Opem market operations (buying and selling of government securities).

 

EOS
Link: https://github.com/EOSIO/Documentation/blob/master/TechnicalWhitePaper.md 

Description: EOS aims to be the Ethereum for high volume and low latency. it employs a permission system that's similar to org charts. EOS also executes DAPPs in parallel.

The algorithm: Not enough information to determine. There are Merkle proof's for Light Client Validation which lets clients verify someone without having to scan to the merkle root.

Consensus Mechanism: DPOS (delegated proof of stake). It is a proof of stake mechanism where Validators elect only one validator to produce blocks at any given time. To be able to vote for a validator you have to have stake in EOS. Thus, it's a delegated proof of stake. Selected block producers are shuffled so the same can't be voted for every time.

Limitations: Not enough information to determine.

 

Ripple
Link: https://ripple.com/files/ripple_consensus_whitepaper.pdf 

Description: The Ripple protocol (RPCA - Ripple Protocol Consensus Algorithm) is a consensus algorithm with pre-selected master nodes called the UNL (unique node list). Ripple also has ancillary software to facilitate cross currency trades using XRP (the native currency) as a go-between. 

The algorithm:  Operates via a ULN (master node list) with a heuristic to allow it to achieve 98.8889% probable consensus. If the UNL is not the same everywhere then its security guarantee does not apply. As long as the master nodes exhibit enough liveness and correctness (80% of the time) the system performs quickly and accurately. If there is disagreement among the UNL, 4 rounds of voting occur before the network shuts down so techs can remove the befouled nodes.

Consensus Mechanism: Federated Byzantine Agreement

Limitations: Centralized. FBA requires 80% validator agreement to be secure.

 

Stellar
Link: https://www.stellar.org/papers/stellar-consensus-protocol.pdf 

Description: Stellar aims to be an improvement upon Ripple. It is a federated byzantine agreement protocol where nodes can choose the validators they listen to. I think of it like a school of fish where you only see the fish around you however you also get to choose where in the school you exist. Ultimately, the school moves in the same direction even though every fish only sees the fish around them.

The algorithm: SCP (stellar consensus protocol) which is the same as RPCA except for one key difference: SCP will not allow disagreement in the network. If there is, the network halts. Therefore, if nodes cannot pass the byzantine (3f+1 rule) it all shuts down waiting for human intervention.

Consensus Mechanism: Modified FBA (sliced FBA)

Limitations: The network will halt if there isn't agreement. Nodes have to reach agreement in rounds which means ledgers don't get published until that occurs.

 

Telegram
Link: https://drive.google.com/file/d/1ucUeKg_NiR8RxNAonb8Q55jZha03WC0O/view 

Description: Telegram wants to build what appears to be Filecoin, SPECTRE, and LightningNetwork all in one but with one exception: They want to use hypercube routing which implies centralized control over nodes.

The algorithm: Not enough information, the white paper references another that isn't public.

Consensus Mechanism: Not enough information. It appears to be SPECTRE or IOS like but remains to be seen. 

Limitations: Telegram claims their userbase is enough to reach network effects however there is no strategy beyond that. It remains unclear if they can stand side by side with traditional currencies as a medium of exchange. Also, their white paper suggests they are trying to do Too Much which is always a bad sign.

 

STEEM
Link: https://steem.io/SteemWhitePaper.pdf 

Description: Steem is a content creator blockchain where creators are rewarded with Steem and consumers vote with Steem to reward the creators. There are three currencies, STEEM, SP (steem power), and SBD (Steem Dollars). SP has a 13 week vesting schedule. STEEM is liquid. and SBD are (basically) US Treasuries with variable payouts and timeframes. SBD are used to stabilize STEEM from wild fluctuations.

The algorithm: One STEEM, one vote. Those with the most STEEM on their balance sheet can have more an an effect on who receives STEEM rewards for their contributions. SP holders get interest and greater voting power. SBD holders get more interest because of increased risk. The interest rate and final price they can sell SBD for are flexible due to market conditions. 

Consensus Mechanism: Subjective Proof of Work. Basically, those that produce the best content get the most votes since they were rewarded the most Steem. 

Limitations: A lot of content is free or advertising based. Why pay for it? Under super high volume how do you get enough votes in time?

 

MimbleWimble
Link: https://download.wpsoftware.net/bitcoin/wizardry/mimblewimble.pdf

Description: This is a system that is very similar to Bitcoin but with non-interactive CoinJoin and the blockchain only contains enough information to verify its current state. 

The algorithm: As mentioned in the description, the algorithm is to implement a compact, verifiable blockchain and non-interactively CoinJoin all txns together for improved space and privacy. There are two terms that are useful: Proven Work which is a capped length of blocks which are all of the transaction and blockchain data. It serves to make it prohibitively expensive to forge a chain. Total Work is the compact chain that is only enough data to verify the chain state assuming the chain state is valid. Matching of inputs and outputs is not necessary so much of the data can be removed in the Total Work chain via something called a Cut Through. In Bitcoin the blockchain is the sum of all Proven Work and Total Work. Furthermore, MimbleWimble uses sinking signatures which effectively allow transactions to be combined. In fact, all transactions in a block are combined into a single transaction which completely obscures the identity of the parties sending coins. A good description is here: https://www.weusecoins.com/mimble-wimble-andrew-poelstra/ 

Consensus Mechanism: PoW

Limitations: Some of the scripting capabilities of Bitcoin had to be removed. Additionally, it is a PoW system and suffers from the mining issues Bitcoin suffers from.

 
Grin
Link: http://grin-tech.org/ 

Description: Grin is an implementation of the MimbleWimble transaction format with extensions required for a complete blockchain.

The algorithm: Aside from being MinbleWimble transactions, Grin implements the Cuckoo Cycle as well. In all other respects, it is the same as MimbleWimble.

Consensus Mechanism: Cuckoo Cycle which is Grin's Proof of Work system. The Cuckoo Cycle is a memory constrained ASIC resistant algorithm. Since memory is the constraining factor, it prevents the Bitcoin hardware 'arms race'.

Limitations: Unknown.

 
Paxos
Link: https://en.wikipedia.org/wiki/Paxos_(computer_science) 

Description: A distributed consensus algorithm that is one of the most cited. It comes in various forms (Byzantine, fast, cheap, and non-byzantine).

The algorithm: There are several variants however the basic form is this. Node A sends prepare to B and C. The nodes respond with the last value they have agreed to that wasn't 'decided'. A picks the highest sequence number from the values in A, B, and C. A sends 'accepts' to B and C. If everyone responds 'accepts' then A sends decided to B and C. If it is not accepted, then a member if the quorum sends another 'prepare' request and the round begins again. This is the Ripple algorithm except for one key difference. Ripple can repeat only 4 times before the network halts.

Consensus Mechanism: Paxos basic

Limitations: Centralized, not byzantine tolerant in the basic form however there is a byzantine version that ads a 'verify' step.

 

0x
Link: https://0xproject.com/pdfs/0x_white_paper.pdf

Description:  Ox is a protocol for a distributed exchange of ERC-20 tokens on the Ethereum protocol. The issue with decentralized exchanges on Ethereum is that people need to spend a lot of gas to manage their orders between ERC-20 tokens. 0x alleviates this problem by managing the orderbook of-fchain with instant and cheap on-chain settlement.

The algorithm: A hybrid implementation, they we refer to as “off-chain order relay with on-chain settlement,” combines the efficiency of state channels with the near instant settlement of on-chain order books. Prices are also adjusted with AMM: "The AMM smart contract replaces the order book with a price-adjustment model in which an asset’s spot price deterministically responds to market forces and market participants on either side of the market trade with the AMM rather than with each other."

Consensus Mechanism: PoW until Casper is launched on Ethereum

Limitations: Still needs gas to operate. Liquidity is an issue for DEXs (distributed exchanges).

 

dharma.io
Link: https://github.com/dharmaprotocol/WhitePaper/blob/master/English/WhitePaper%5BEN%5D.md 

Description: dharma.io is a loan chain that uses 0x as its exchange and is an ERC-20 token. The players include a RAA (risk assessment attester) or underwriter, a lender, and a borrower. The intent is to create a trustless environment where these parties can interact with each other, browse the loan 'order book', and resell loans.

The Algorithm: It is based on smart contracts that control how the various players interact with each other. Beyond that, there isn't an algorithm per-se. 

Consensus Mechanism: Pow until Casper is launched on Ethereum.

Limitations: dharma was written by people who have never worked in finance before. Therefore, they have made certain assumptions that don't always hold true.

 

For example: 

1. In private credit it’s the creditor who prices credit.

z2. The servicer can be the underwriter, could be the creditor or could be a third party.

 

For example, in a mortgage https://www.investopedia.com/terms/m/msr.asp

 

NuCypher
Link: https://cdn2.hubspot.net/hubfs/2807639/NuCypher%20KMS%20Technical%20White%20Paper.pdf 

Description: NuCypher is a distributed proxy re-encryption service. The premise is that I can give it encrypted text and delegate the decryption of that text to others without NuCypher's nodes ever knowing or being able to know the content of that text. There are several schemes they support: threshold of K of N, M of N, standard bilinear polynomials (RSA), and an actual symmetric cypher (old school!). There is an added benefit in that the re-encryption keys pseudo-anonymize the originator. 

The algorithm: Threshold, RSA, symmetric cipher, split-key re-encryption where the secrets needed are distributed on several nodes and have to be recombined (m-of-m), BBS98, a quantum resistant NTRU, and re-encryption keys that time out.

Consensus Mechanism: PoS where nodes can loose their collateral if they cheat.

Limitations: Not sure how popular this will be. Otherwise, seems GREAT!

 

IPFS
Link: https://ipfs.io/ipfs/QmR7GSQM93Cx5eAg6a6yRzNde1FQv7uL6X1o4k7zrJa3LX/ipfs.draft3.pdf 

Description: IPFS is a distributed files system on distributed nodes. It's called the permanent web in that files written in IPFS cannot be removed once broadcast. It is common for clients to encrypt files before writing to IPFS so that they can only be read by one or multiple parties (using k of n or m of n encryption schemes). Essentially it's a distributed hash table of files stored on the network.

The algorithm: It's a DHT (distributed hash table) that references files distributed via a protocol that's very close to bitTorrent.

Consensus Mechanism: There isn't one per-se. You share my files and I'll share yours. We exchange credits and if I have a debt then you won't share my file until I have a credit. FileCoin will eventually micro-pay people to host files.

Limitations: Files are permanent. Relies on bartering so what's the incentive to run a IPFS node?

Raft
Link: 

Description:

The algorithm:

Consensus Mechanism:

Limitations: 

 

Chia
Link: 

Description: Bitcoin competitor that claims to be 'green', based on hard drive space rather then computation.

The algorithm:

Consensus Mechanism:

Limitations:

  

Augur
Link: 

Description: ERC20 token based prediction market based on hivemind. Exists ontop of Ethereum. Augur is has a work token, REP, which allows the reporting of events to the prediction market.

The algorithm:

Consensus Mechanism:

Limitations:

 

Gnosis
Link: 

Description:

The algorithm:

Consensus Mechanism:

Limitations:

 

Hivemind
Link: 

Description: sidechain based prediction market, which augur is derived from

The algorithm: 

Consensus Mechanism:

Limitations:

 

ByteBall
Link: 

Description:

The algorithm: DAG

Consensus Mechanism:

Limitations:

 

RaiBlocks/Nano
Link: 

Description:

The algorithm:

Consensus Mechanism:

Limitations:


