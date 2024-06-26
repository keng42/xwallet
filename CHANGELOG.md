<a name="unreleased"></a>
## [Unreleased]


<a name="v0.2.2"></a>
## [v0.2.2] - 2024-06-05
### Chore
- **deps:** bump github.com/ethereum/go-ethereum from 1.12.1 to 1.13.15
- **deps:** bump github.com/btcsuite/btcd from 0.23.4 to 0.24.0
- **deps:** bump golang.org/x/crypto from 0.9.0 to 0.17.0

### Pull Requests
- Merge pull request [#11](https://github.com/keng42/xwallet/issues/11) from keng42/dependabot/go_modules/github.com/ethereum/go-ethereum-1.13.15
- Merge pull request [#10](https://github.com/keng42/xwallet/issues/10) from keng42/dependabot/go_modules/github.com/btcsuite/btcd-0.24.0
- Merge pull request [#9](https://github.com/keng42/xwallet/issues/9) from keng42/dependabot/go_modules/golang.org/x/crypto-0.17.0


<a name="v0.2.1"></a>
## [v0.2.1] - 2023-09-07
### Chore
- **deps:** bump github.com/ethereum/go-ethereum from 1.11.2 to 1.12.1

### Pull Requests
- Merge pull request [#7](https://github.com/keng42/xwallet/issues/7) from keng42/dependabot/go_modules/github.com/ethereum/go-ethereum-1.12.1


<a name="v0.2.0"></a>
## [v0.2.0] - 2023-09-05
### Chore
- **deps:** update go to 1.20
- **deps:** bump github.com/ethereum/go-ethereum

### Docs
- **readme:** add badges

### Feat
- cache Bip32ExtendedKey to generate addresses in bulk faster
- upgraded btcd package

### Pull Requests
- Merge pull request [#1](https://github.com/keng42/xwallet/issues/1) from keng42/dependabot/go_modules/github.com/ethereum/go-ethereum-1.10.17


<a name="v0.1.0"></a>
## v0.1.0 - 2022-03-25
### Chore
- add git-chg config
- **gh-action:** add release workflow
- **scripts:** add build and release scripts

### Docs
- **readme:** add installation and usage examples
- **readme:** add networks support status and some links of bip

### Feat
- init commit
- **account:** add package account
- **chains:** add multiple utility functions for many chains
- **cmd:** add cli entry
- **mnemonic:** add package mnemonic
- **mod:** add multiple dependencies
- **network:** add Symbol in Network
- **network:** add package network
- **utilities:** add package hash
- **utilities:** add package info
- **wallet:** check if the mnemonic is valid in NewAddress
- **wallet:** add package wallet


[Unreleased]: https://github.com/keng42/xwallet/compare/v0.2.2...HEAD
[v0.2.2]: https://github.com/keng42/xwallet/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/keng42/xwallet/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/keng42/xwallet/compare/v0.1.0...v0.2.0
