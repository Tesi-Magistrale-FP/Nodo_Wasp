# ISC - Autenticazione
ISC usato per registrare e autenticare gli utenti nell'ecosistema.

## Funzioni 
### Funzioni Func
- <strong>registrazione</strong>: registrazione utente nell'ecosistema.
- <strong>eliminazione</strong>: registrazione utente nell'ecosistema.

### Funzioni View
- <strong>login</strong>: login utente nell'ecosistema.

## Comandi Wasp-cli
Comando per eseguire il deployment:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain deploy-contract wasmtime autenticazione "autenticazione SC" ..\..\contracts\wasm\autenticazione\rs\autenticazionewasm\pkg\autenticazionewasm_bg.wasm --chain=mychain
```
Comandi per funzioni Func:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain post-request autenticazione registrazione String did String "did:prova" String password String "hash_pwd" --chain=mychain -s
wasp-cli chain post-request autenticazione eliminazione String did String "did:prova" String password String "hash_pwd" --chain=mychain -s
```
Comandi per funzioni View:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain call-view autenticazione login String did String "did:prova" String password String "hash_pwd" --chain=mychain | wasp-cli decode string esitoL bool
```