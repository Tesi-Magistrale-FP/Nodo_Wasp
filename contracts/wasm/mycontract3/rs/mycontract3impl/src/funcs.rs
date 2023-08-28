
use wasmlib::*;
use mycontract3::*;
use crate::*;

pub fn func_eliminazione(ctx: &ScFuncContext, f: &EliminazioneContext) 
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
            f.results.esito_e().set_value(true);
        }
    }
    else
    {
        f.results.esito_e().set_value(false);
    }
}

pub fn func_registrazione(ctx: &ScFuncContext, f: &RegistrazioneContext) 
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
			f.results.esito_r().set_value(&String::from("Utente registrato con successo"));
		}
		else
		{
			f.results.esito_r().set_value(&String::from("Utente già registrato"));
		}
    }
    else
    {
        pwd_utente.set_value(&password);
        f.results.esito_r().set_value(&String::from("Utente registrato con successo"));
    }
}

pub fn view_login(ctx: &ScViewContext, f: &LoginContext) 
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
