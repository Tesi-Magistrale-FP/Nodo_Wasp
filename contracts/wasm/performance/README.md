# ISC - Performance
ISC usato per misurare le performance di scrittura e lettura sugli smart contract.

## Funzioni 
### Funzioni Func
- <strong>memorizzaValore</strong>: memorizza un valore nello stato interno dello ISC.

### Funzioni View
- <strong>ottieniValore</strong>: ottiene un valore identificato da un indice.

## Comandi Wasp-cli
Comando per eseguire il deployment:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain deploy-contract wasmtime performance "Performance SC" ..\..\contracts\wasm\performance\rs\performancewasm\pkg\performancewasm_bg.wasm --chain=mychain
```
Comandi per funzioni Func:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain post-request performance memorizzaValore String valore String "ciao" --chain=mychain -s
```
Comandi per funzioni View:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain call-view performance ottieniValore String indice Int32 0 --chain=mychain | wasp-cli decode string valore string
```