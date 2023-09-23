use wasmlib::*;
use autorizzazioni::*;
use crate::*;

pub fn func_fornisci_autorizzazione(ctx: &ScFuncContext, f: &FornisciAutorizzazioneContext) 
{
	let did: String = f.params.did_produttore().value();
	let applicazione: i32 = f.params.applicazione().value();
	let operazione: i32 = f.params.operazione().value();

	let autorizzazione = Autorizzazione {
		concessa: true,
		applicazione: applicazione, 
		operazione: operazione,
	};
	
	let elenco_autorizzazioni: MapStringToMutableElencoAutorizzazioni = f.state.autorizzazioni();
	let autorizzazioni_utente: ArrayOfMutableAutorizzazione = elenco_autorizzazioni.get_elenco_autorizzazioni(&did);

	if autorizzazioni_utente.length() > 0
	{
		for i in 0..autorizzazioni_utente.length() 
		{
			let autorizzazione: Autorizzazione = autorizzazioni_utente.get_autorizzazione(i).value();
			if autorizzazione.applicazione == applicazione && autorizzazione.operazione == operazione
			{
				autorizzazioni_utente.get_autorizzazione(i).set_value(&autorizzazione);
				f.events.autorizzazione_fornita();
				return;
			}
		}
	}

	autorizzazioni_utente.append_autorizzazione().set_value(&autorizzazione);
	f.events.autorizzazione_fornita();
}

pub fn func_rimuovi_autorizzazione(ctx: &ScFuncContext, f: &RimuoviAutorizzazioneContext) 
{
	let did: String = f.params.did_produttore().value();
	let applicazione: i32 = f.params.applicazione().value();
	let operazione: i32 = f.params.operazione().value();

	let elenco_autorizzazioni: MapStringToMutableElencoAutorizzazioni = f.state.autorizzazioni();
	let autorizzazioni_utente: ArrayOfMutableAutorizzazione = elenco_autorizzazioni.get_elenco_autorizzazioni(&did);

	if autorizzazioni_utente.length() > 0
	{
		for i in 0..autorizzazioni_utente.length() 
		{
			let autorizzazione: Autorizzazione = autorizzazioni_utente.get_autorizzazione(i).value();
			if autorizzazione.applicazione == applicazione && autorizzazione.operazione == operazione
			{
				let autorizzazione = Autorizzazione {
					concessa: false,
					applicazione: applicazione, 
					operazione: operazione,
				};

				autorizzazioni_utente.get_autorizzazione(i).set_value(&autorizzazione);
				f.events.autorizzazione_rimossa();
				return;
			}
		}
	}
	
	f.events.autorizzazione_rimossa();
}

pub fn view_controlla_autorizzazione(ctx: &ScViewContext, f: &ControllaAutorizzazioneContext) 
{
	let did: String = f.params.did_produttore().value();
	let applicazione: i32 = f.params.applicazione().value();
	let operazione: i32 = f.params.operazione().value();

	let elenco_autorizzazioni: MapStringToImmutableElencoAutorizzazioni = f.state.autorizzazioni();
	let autorizzazioni_utente: ArrayOfImmutableAutorizzazione = elenco_autorizzazioni.get_elenco_autorizzazioni(&did);

	if autorizzazioni_utente.length() > 0
	{
		for i in 0..autorizzazioni_utente.length() 
		{
			let autorizzazione: Autorizzazione = autorizzazioni_utente.get_autorizzazione(i).value();
			if autorizzazione.applicazione == applicazione && autorizzazione.operazione == operazione
			{
				f.results.esito_c().set_value(autorizzazione.concessa);
				return;
			}
		}
	}

	f.results.esito_c().set_value(false);
}
