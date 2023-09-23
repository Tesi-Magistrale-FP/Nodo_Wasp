use wasmlib::*;
use autenticazione::*;
use crate::*;

pub fn func_eliminazione(ctx: &ScFuncContext, f: &EliminazioneContext)									// Eliminazione utente
{
	let did: String = f.params.did().value();															// Input DID												
    let password: String = f.params.password().value();													// Input password
    let utenti: MapStringToMutableString = f.state.utenti();											// Recupero utenti registrati all'ecosistema

    let pwd_utente: ScMutableString = utenti.get_string(&did);											// Ottiene la password associata alla DID

    if pwd_utente.exists() 																				// Se l'utente è registrato
    {
		let pwd_utente_c: String = utenti.get_string(&did).to_string();									// Ottiene la password associata alla DID in formato String
		
        if pwd_utente_c == password																		// Se la password recuperata e quella ricevuta in input sono uguali
        {
            pwd_utente.set_value(&String::from(""));													// Elimino l'utente
			f.events.eliminazione_successo();															// Evento che indica il successo dell'eliminazione
			return;
        }
    }
	
	f.events.eliminazione_fallimento();																	// Evento che indica il fallimento dell'eliminazione
}

pub fn func_registrazione(ctx: &ScFuncContext, f: &RegistrazioneContext) 								// Registrazione utente
{
	let did: String = f.params.did().value();
    let password: String = f.params.password().value();
    let utenti: MapStringToMutableString = f.state.utenti();

    let pwd_utente: ScMutableString = utenti.get_string(&did);

    if pwd_utente.exists() 
    {
		let pwd_utente_c: String = utenti.get_string(&did).to_string();
		
        if pwd_utente_c == String::from("")
        {
			pwd_utente.set_value(&password);
			f.events.registrazione_successo();
			return;
		}
    }
    else
    {
        pwd_utente.set_value(&password);
		f.events.registrazione_successo();
		return;
    }
	
	f.events.registrazione_fallimento();
}

pub fn view_login(ctx: &ScViewContext, f: &LoginContext) 												// Login utente
{
	let did: String = f.params.did().value();
    let password: String = f.params.password().value();
    let utenti: MapStringToImmutableString = f.state.utenti();

    let pwd_utente: ScImmutableString = utenti.get_string(&did);

    if pwd_utente.exists() 
    {
		let pwd_utente_c: String = utenti.get_string(&did).to_string();
		
        if pwd_utente_c == password && pwd_utente_c != String::from("")
        {
            f.results.esito_l().set_value(true);														// Login eseguito con successo
            return;
        }
    }
    f.results.esito_l().set_value(false);																// Login fallito
}