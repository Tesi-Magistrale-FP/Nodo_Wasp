use wasmlib::*;
use mycontract::*;
use crate::*;

pub fn func_eliminazione(ctx: &ScFuncContext, f: &EliminazioneContext)									// Eliminazione utente
{
	let did: String = f.params.did().value();															
    let password: String = f.params.password().value();
    let utenti: MapStringToMutableString = f.state.utenti();

    let pwd_utente: ScMutableString = utenti.get_string(&did);

    if pwd_utente.exists() 
    {
		let pwd_utente_c: String = utenti.get_string(&did).to_string();
		
        if pwd_utente_c == password
        {
            pwd_utente.set_value(&String::from(""));
        }
    }
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
		}
    }
    else
    {
        pwd_utente.set_value(&password);
    }
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
		
        if pwd_utente_c == password
        {
            f.results.esito_l().set_value(true);
            return;
        }
    }
    f.results.esito_l().set_value(false);
}
