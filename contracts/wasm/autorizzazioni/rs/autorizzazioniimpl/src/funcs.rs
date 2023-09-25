use wasmlib::*;
use autorizzazioni::*;
use crate::*;

pub fn func_fornisci_autorizzazione(ctx: &ScFuncContext, f: &FornisciAutorizzazioneContext) 
{
	let did_produttore: String = f.params.did_produttore().value();
	let did_consumatore: String = f.params.did_consumatore().value();
	let id_applicazione: i32 = f.params.id_applicazione().value();
	let id_operazione: i32 = f.params.id_operazione().value();

	let elenco_autorizzazioni: MapStringToMutableElencoAutorizzazioni = f.state.autorizzazioni();
	let autorizzazioni_utente: ArrayOfMutableAutorizzazione = elenco_autorizzazioni.get_elenco_autorizzazioni(&did_produttore);

	if autorizzazioni_utente.length() > 0
	{
		for i in 0..autorizzazioni_utente.length() 
		{
			let mut autorizzazione: Autorizzazione = autorizzazioni_utente.get_autorizzazione(i).value();
			if autorizzazione.id_applicazione == id_applicazione && autorizzazione.id_operazione == id_operazione && autorizzazione.did_consumatore == did_consumatore
			{
				autorizzazione.concessa = true;
				autorizzazioni_utente.get_autorizzazione(i).set_value(&autorizzazione);
				f.events.autorizzazione_fornita();
				return;
			}
		}
	}

	let autorizzazione = Autorizzazione {
		concessa: true,
		id_applicazione: id_applicazione, 
		id_operazione: id_operazione,
		did_consumatore: did_consumatore,
	};

	autorizzazioni_utente.append_autorizzazione().set_value(&autorizzazione);
	f.events.autorizzazione_fornita();
}

pub fn func_rimuovi_autorizzazione(ctx: &ScFuncContext, f: &RimuoviAutorizzazioneContext) 
{
	let did_produttore: String = f.params.did_produttore().value();
	let did_consumatore: String = f.params.did_consumatore().value();
	let id_applicazione: i32 = f.params.id_applicazione().value();
	let id_operazione: i32 = f.params.id_operazione().value();

	let elenco_autorizzazioni: MapStringToMutableElencoAutorizzazioni = f.state.autorizzazioni();
	let autorizzazioni_utente: ArrayOfMutableAutorizzazione = elenco_autorizzazioni.get_elenco_autorizzazioni(&did_produttore);

	if autorizzazioni_utente.length() > 0
	{
		for i in 0..autorizzazioni_utente.length() 
		{
			let mut autorizzazione: Autorizzazione = autorizzazioni_utente.get_autorizzazione(i).value();
			if autorizzazione.id_applicazione == id_applicazione && autorizzazione.id_operazione == id_operazione && autorizzazione.did_consumatore == did_consumatore
			{
				autorizzazione.concessa = false;
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
	let did_produttore: String = f.params.did_produttore().value();
	let did_consumatore: String = f.params.did_consumatore().value();
	let id_applicazione: i32 = f.params.id_applicazione().value();
	let id_operazione: i32 = f.params.id_operazione().value();

	let elenco_autorizzazioni: MapStringToImmutableElencoAutorizzazioni = f.state.autorizzazioni();
	let autorizzazioni_utente: ArrayOfImmutableAutorizzazione = elenco_autorizzazioni.get_elenco_autorizzazioni(&did_produttore);

	if autorizzazioni_utente.length() > 0
	{
		for i in 0..autorizzazioni_utente.length() 
		{
			let autorizzazione: Autorizzazione = autorizzazioni_utente.get_autorizzazione(i).value();
			if autorizzazione.id_applicazione == id_applicazione && autorizzazione.id_operazione == id_operazione && autorizzazione.did_consumatore == did_consumatore
			{
				f.results.esito_c().set_value(autorizzazione.concessa);
				return;
			}
		}
	}

	f.results.esito_c().set_value(false);
}
