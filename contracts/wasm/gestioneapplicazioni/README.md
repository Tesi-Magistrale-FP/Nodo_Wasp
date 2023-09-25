# ISC - GestioneApplicazioni
ISC usato per pubblicare le applicazioni, aggiungere operazioni sui dati, gestire le loro modifiche e tutti i riferimenti ai canali dati, log, oggetti Author e Subscriber.

## Funzioni 
### Funzioni Func
- <strong>pubblicaApplicazione</strong>: pubblica l’applicazione con i relativi dati nello store.
- <strong>modificaApplicazione</strong>: modifica i dati di un’applicazione pubblicata.
- <strong>registraUtenteVC</strong>: registra le credenziali verificabili di un utente.
- <strong>aggiungiOperazione</strong>: aggiunge un’operazione sui dati per un’app.
- <strong>modificaStatoOperazione</strong>: attiva o disattiva un’operazione sui dati.
- <strong>creaCanali</strong>: crea i canali dati e log quando un utente fornisce l’autorizzazione a un’operazione.
- <strong>modificaIndUltMsg</strong>: aggiorna gli indirizzi degli ultimi messaggi pubblicati sui canali dati e log a seguito di un’operazione di scrittura sui canali stessi.
- <strong>aggiungiIscritto</strong>: aggiunge gli oggetti Subscriber usati da un certo consumatore per accedere ai canali dati e log associati a un’autorizzazione ottenuta.

### Funzioni View
- <strong>idApplicazione</strong>: restituisce l’ID univoco associato a un’applicazione.
- <strong>datiApplicazione</strong>: restituisce i dati di un’applicazione pubblicata.
- <strong>elencoApplicazioni</strong>: restituisce il numero totale delle applicazioni pubblicate.
- <strong>elencoAppsUtente</strong>: restituisce le applicazioni pubblicate da uno specifico utente.
- <strong>verificaUtenteVC</strong>: verifica le credenziali verificabili di un utente.
- <strong>idOperazione</strong>: restituisce l’ID univoco associato a un’operazione.
- <strong>datiOperazione</strong>: restituisce i dati di un’operazione.
- <strong>elencoOperazioni</strong>: elenca tutte le operazioni di una determinata app.
- <strong>indUltMsg</strong>: restituisce gli indirizzi degli ultimi messaggi pubblicati sui canali dati e log associati a un’autorizzazione concessa.
- <strong>ottieniAutore</strong>: restituisce gli oggetti Author usati da un certo produttore per accedere ai canali dati e log associati a un’autorizzazione concessa.
- <strong>ottieniIscritto</strong>: restituisce gli oggetti Subscriber usati da un certo consumatore per accedere ai canali dati e log associati a un’autorizzazione ottenuta.

## Comandi Wasp-cli
Comando per eseguire il deployment:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain deploy-contract wasmtime gestioneapplicazioni "gestione applicazioni SC" ..\..\contracts\wasm\gestioneapplicazioni\rs\gestioneapplicazioniwasm\pkg\gestioneapplicazioniwasm_bg.wasm --chain=mychain
```
Comandi per funzioni Func:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain post-request gestioneapplicazioni pubblicaApplicazione String nome String "GMaps0" String descrizione String "GMaps per viaggi0" String url String "https://gmaps.com0" String didProprietario String "did:iota:2000" --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni modificaApplicazione String nomeAttuale String "GMaps0" String nome String "GMaps" String descrizione String "GMaps per viaggi" String url String "https://gmaps.com" String didProprietario String "did:iota:200" String attiva Bool false --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni modificaApplicazione String nomeAttuale String "GMaps" String nome String "GMaps" String descrizione String "GMaps per viaggi" String url String "https://gmaps.com" String didProprietario String "did:iota:200" String attiva Bool true --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni registraUtenteVC String rifUtente String "0|did:iota:100" String vcUtente String "Francesco|Paglia" --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni aggiungiOperazione String idApp Int32 0 String titolo String "Raccolta" String descrizione String "Raccoglie dati utente" String software String "input:dati" String tipo Int32 2 --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni modificaStatoOperazione String idApp Int32 0 String titolo String "Raccolta" String nuovoStato Bool false --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni modificaStatoOperazione String idApp Int32 0 String titolo String "Raccolta" String nuovoStato Bool true --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni creaCanali String idApp Int32 0 String idOp Int32 0 String didProduttore String "did:iota:100" String autoreDati String "autore dati" String annuncioDati String "annuncio dati" String indUltMsgDati String "indirizzo dati 1" String autoreLog String "autore log" String annuncioLog String "annuncio log" String indUltMsgLog String "indirizzo log 1" --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni modificaIndUltMsg String idApp Int32 0 String idOp Int32 0 String didProduttore String "did:iota:100" String indUltMsgDati String "indirizzo dati 2" String indUltMsgLog String "indirizzo log 2" --chain=mychain -s
wasp-cli chain post-request gestioneapplicazioni aggiungiIscritto String idApp Int32 0 String idOp Int32 0 String didProduttore String "did:iota:100" String didConsumatore String "did:iota:200" String iscrittoDati String "iscritto dati" String iscrittoLog String "iscritto log" --chain=mychain -s
```
Comandi per funzioni View:
```
cd Nodo_Wasp\tools\local-setup
wasp-cli chain call-view gestioneapplicazioni idApplicazione String nome String "GMaps" String didProprietario String "did:iota:200" --chain=mychain | wasp-cli decode string idApp int32
wasp-cli chain call-view gestioneapplicazioni datiApplicazione String idApp Int32 0 --chain=mychain | wasp-cli decode string datiApp string
wasp-cli chain call-view gestioneapplicazioni elencoApplicazioni --chain=mychain | wasp-cli decode string numeroApps int32
wasp-cli chain call-view gestioneapplicazioni elencoAppsUtente String didUtente String "did:iota:200" --chain=mychain | wasp-cli decode string idApps string
wasp-cli chain call-view gestioneapplicazioni verificaUtenteVC String rifUtente String "0|did:iota:100" String vcUtente String "Francesco|Paglia" --chain=mychain | wasp-cli decode string esitoC bool
wasp-cli chain call-view gestioneapplicazioni idOperazione String idApp Int32 0 String titolo String "Raccolta" --chain=mychain | wasp-cli decode string idOp int32
wasp-cli chain call-view gestioneapplicazioni datiOperazione String idOp Int32 0 --chain=mychain | wasp-cli decode string datiOp string
wasp-cli chain call-view gestioneapplicazioni elencoOperazioni String idApp Int32 0 --chain=mychain | wasp-cli decode string idOps string
wasp-cli chain call-view gestioneapplicazioni indUltMsg String idApp Int32 0 String idOp Int32 0 String didProduttore String "did:iota:100" --chain=mychain | wasp-cli decode string indirizzi string
wasp-cli chain call-view gestioneapplicazioni ottieniAutore String idApp Int32 0 String idOp Int32 0 String didProduttore String "did:iota:100" --chain=mychain | wasp-cli decode string autori string
wasp-cli chain call-view gestioneapplicazioni ottieniIscritto String idApp Int32 0 String idOp Int32 0 String didProduttore String "did:iota:100" String didConsumatore String "did:iota:200" --chain=mychain | wasp-cli decode string iscritti string
```