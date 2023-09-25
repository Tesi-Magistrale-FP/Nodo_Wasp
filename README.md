# Nodo Wasp
[Wasp](https://github.com/iotaledger/wasp) è un nodo software sviluppato dalla [IOTA Foundation](http://iota.org) per eseguire gli <strong>IOTA Smart Contracts (ISC)</strong> sul Tangle di IOTA.

### IOTA Smart Contracts implementati
- [Autenticazione](https://github.com/Tesi-Magistrale-FP/Nodo_Wasp/tree/main/contracts/wasm/autenticazione): autentica gli utenti nell'ecosistema.
- [Autorizzazioni](https://github.com/Tesi-Magistrale-FP/Nodo_Wasp/tree/main/contracts/wasm/autorizzazioni): gestisce le autorizzazioni dei produttori per i consumatori.
- [Gestione applicazioni](https://github.com/Tesi-Magistrale-FP/Nodo_Wasp/tree/main/contracts/wasm/gestioneapplicazioni): gestisce le applicazioni pubblicate, le operazioni sui dati e l'accesso ai canali Streams.

### Requisiti
- [Go](https://golang.org/dl/)
- Gcc (o equivalente per Windows [(TDM-GCC)](https://jmeubank.github.io/tdm-gcc/))
- [Rust](https://www.rust-lang.org/tools/install)
- [Wasp-cli](https://github.com/iotaledger/wasp/releases)
- [Wasm-pack](https://rustwasm.github.io/wasm-pack/installer/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [Visual Studio Code](https://code.visualstudio.com/Download) (VSCode)
  - [Go Extension](https://marketplace.visualstudio.com/items?itemName=golang.Go)
  - [Rust Analyzer](https://marketplace.visualstudio.com/items?itemName=matklad.rust-analyzer)
  - [Even Better TOML](https://marketplace.visualstudio.com/items?itemName=tamasfe.even-better-toml)

### Nodo Wasp
- [Repository con versione più aggiornata](https://github.com/iotaledger/wasp)
- [Setup locale con Docker Desktop](https://github.com/iotaledger/wasp/tree/develop/tools/local-setup)
- [Guida di IOTA | ISC, node e chains](https://wiki.iota.org/wasp/running-a-node/)

## Comandi Wasp-cli
### Setup con Docker Desktop
Comandi per eseguire un nodo Wasp localmente:
```
git clone https://github.com/Tesi-Magistrale-FP/Nodo_Wasp.git
cd Nodo_Wasp\tools\local-setup
docker-compose pull
docker volume create --name hornet-nest-db
docker volume create --name wasp-db
docker-compose up -d
```

### Setup chain di test
Comandi per creare e attivare una chain di test su cui effettuare il deployment degli ISC:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli init
wasp-cli set l1.apiaddress http://localhost:14265
wasp-cli set l1.faucetaddress http://localhost:8091
wasp-cli wasp add 0 http://localhost:9090
wasp-cli request-funds
wasp-cli chain deploy --chain=mychain
wasp-cli chain add mychain <ID_CHAIN>
wasp-cli chain activate --chain=mychain
wasp-cli chain deposit base:500000000 --chain=mychain
```

### Generazione, deployment e interrograzione ISC
Comandi per generare l'[ISC Autenticazione](https://github.com/Tesi-Magistrale-FP/Nodo_Wasp/tree/main/contracts/wasm/autenticazione), effettuare il suo deployment e chiamare le funzioni Func e View:
```
cd Nodo-Wasp\contracts\wasm
schema -init autenticazione
cd autenticazione
```
Modificare il file [schema.yaml](https://github.com/Tesi-Magistrale-FP/Nodo_Wasp/blob/main/contracts/wasm/autenticazione/schema.yaml) <br><br>
`schema -rs`
<br><br>
Implementare la logica delle funzioni [func.rs](https://github.com/Tesi-Magistrale-FP/Nodo_Wasp/blob/main/contracts/wasm/autenticazione/rs/autenticazioneimpl/src/funcs.rs)
```
schema -rs -build
cd ..\..\tools\local-setup
wasp-cli chain deploy-contract wasmtime autenticazione "autenticazione SC" ..\..\contracts\wasm\autenticazione\rs\autenticazionewasm\pkg\autenticazionewasm_bg.wasm --chain=mychain
wasp-cli chain post-request autenticazione registrazione String did String "prova" String password String "prova" --chain=mychain -s
wasp-cli chain call-view autenticazione login String did String "prova" String password String "prova" --chain=mychain | wasp-cli decode string esitoL bool
wasp-cli chain post-request autenticazione eliminazione String did String "prova" String password String "prova" --chain=mychain -s
```