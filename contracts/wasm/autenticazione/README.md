# ISC - Autenticazione
ISC usato per registrare e autenticare gli utenti nell'ecosistema.

### Diagramma UML
![Diagramma UML - ISC Autenticazione](https://github.com/Tesi-Magistrale-FP/Nodo_Wasp/tree/main/diagrammi/autenticazione.png)

### Comandi
Comando per eseguire il deployment:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain deploy-contract wasmtime autenticazione "autenticazione SC" ..\..\contracts\wasm\autenticazione\rs\autenticazionewasm\pkg\autenticazionewasm_bg.wasm --chain=mychain
```
<br><br>
Comandi per funzioni Func:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain post-request autenticazione registrazione String did String "did:prova" String password String "hash_pwd" --chain=mychain -s
wasp-cli chain post-request autenticazione eliminazione String did String "did:prova" String password String "hash_pwd" --chain=mychain -s
```
<br><br>
Comandi per funzioni View:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain call-view autenticazione login String did String "did:prova" String password String "hash_pwd" --chain=mychain | wasp-cli decode string esitoL bool
```