# ISC - Autorizzazioni
ISC usato per gestire le autorizzazioni concesse dai produttori per i consumatori.

## Funzioni 
### Funzioni Func
- <strong>fornisciAutorizzazione</strong>: utente produttore fornisce la sua autorizzazione a un utente consumatore per un’operazione dati di una particolare applicazione.
- <strong>rimuoviAutorizzazione</strong>: rimuove un’autorizzazione precedentemente concessa.

### Funzioni View
- <strong>controllaAutorizzazione</strong>: controlla se un’autorizzazione è stata concessa o meno.

## Comandi Wasp-cli
Comando per eseguire il deployment:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain deploy-contract wasmtime autorizzazioni "autorizzazioni SC" ..\..\contracts\wasm\autorizzazioni\rs\autorizzazioniwasm\pkg\autorizzazioniwasm_bg.wasm --chain=mychain
```
Comandi per funzioni Func:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain post-request autorizzazioni fornisciAutorizzazione String didProduttore String "provaP" String didConsumatore String "provaC" String idApplicazione Int32 0 String idOperazione Int32 0 --chain=mychain -s
wasp-cli chain post-request autorizzazioni rimuoviAutorizzazione String didProduttore String "provaP" String didConsumatore String "provaC" String idApplicazione Int32 0 String idOperazione Int32 0 --chain=mychain -s
```
Comandi per funzioni View:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain call-view autorizzazioni controllaAutorizzazione String didProduttore String "provaP" String didConsumatore String "provaC" String idApplicazione Int32 0 String idOperazione Int32 0 --chain=mychain | wasp-cli decode string esitoC bool
```