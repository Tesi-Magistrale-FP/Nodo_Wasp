use wasmlib::*;
use gestioneapplicazioni::*;
use crate::*;

pub fn func_aggiungi_iscritto(ctx: &ScFuncContext, f: &AggiungiIscrittoContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let id_op: i32 = f.params.id_op().value();
	let did_produttore: String = f.params.did_produttore().value();
	let did_consumatore: String = f.params.did_consumatore().value();
	let iscritto_dati: String = f.params.iscritto_dati().value();
	let iscritto_log: String = f.params.iscritto_log().value();
	
	let iscritti: ArrayOfMutableIscritto = f.state.iscritti();
	
	if iscritti.length() > 0
	{
		for i in 0..iscritti.length() 
		{
			let iscritto: Iscritto = iscritti.get_iscritto(i).value();
			if iscritto.id_app == id_app && iscritto.id_op == id_op && iscritto.did_produttore == did_produttore && iscritto.did_consumatore == did_consumatore
			{
				f.events.iscritto_aggiunto();
				return;
			}
		}
	}
	
	let iscritto = Iscritto {
		id_app: id_app,
		id_op: id_op,
		did_produttore: did_produttore,
		did_consumatore: did_consumatore,
		iscritto_dati: iscritto_dati,
		iscritto_log: iscritto_log,
	};

	iscritti.append_iscritto().set_value(&iscritto);
	f.events.iscritto_aggiunto();
}

pub fn func_aggiungi_operazione(ctx: &ScFuncContext, f: &AggiungiOperazioneContext) 
{
    let id_app: i32 = f.params.id_app().value();
	let titolo: String = f.params.titolo().value();
	let descrizione: String = f.params.descrizione().value();
	let software: String = f.params.software().value();
	let tipo: i32 = f.params.tipo().value();
	
	let operazioni: ArrayOfMutableOperazione = f.state.operazioni();
	
	if operazioni.length() > 0
	{
		for i in 0..operazioni.length() 
		{
			let operazione: Operazione = operazioni.get_operazione(i).value();
			if operazione.id_app == id_app && operazione.titolo == titolo
			{
				f.events.operazione_aggiunta();
				return;
			}
		}
	}
	
	let operazione = Operazione {
		id_app: id_app,
		attiva_op: true,
		titolo: titolo,
		descrizione: descrizione,
		software: software,
		tipo: tipo,
	};
	
	operazioni.append_operazione().set_value(&operazione);
	f.events.operazione_aggiunta();
}

pub fn func_crea_canali(ctx: &ScFuncContext, f: &CreaCanaliContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let id_op: i32 = f.params.id_op().value();
	let did_produttore: String = f.params.did_produttore().value();
	let autore_dati: String = f.params.autore_dati().value();
	let annuncio_dati: String = f.params.annuncio_dati().value();
	let ind_ult_msg_dati: String = f.params.ind_ult_msg_dati().value();
	let autore_log: String = f.params.autore_log().value();
	let annuncio_log: String = f.params.annuncio_log().value();
	let ind_ult_msg_log: String = f.params.ind_ult_msg_log().value();
	
	let canali: ArrayOfMutableCanali = f.state.canali();
	
	if canali.length() > 0
	{
		for i in 0..canali.length() 
		{
			let canali_dati_log: Canali = canali.get_canali(i).value();
			if canali_dati_log.id_app == id_app && canali_dati_log.id_op == id_op && canali_dati_log.did_produttore == did_produttore
			{
				f.events.canali_creati();
				return;
			}
		}
	}
		
	let canali_dati_log = Canali {
		id_app: id_app,
		id_op: id_op,
		autore_dati: autore_dati,
		autore_log: autore_log,
		annuncio_dati: annuncio_dati,
		annuncio_log: annuncio_log,
		did_produttore: did_produttore,
		ind_ult_msg_dati: ind_ult_msg_dati,
		ind_ult_msg_log: ind_ult_msg_log,
	};
	
	canali.append_canali().set_value(&canali_dati_log);
	f.events.canali_creati();
}

pub fn func_modifica_applicazione(ctx: &ScFuncContext, f: &ModificaApplicazioneContext) 
{
    let nome_attuale: String = f.params.nome_attuale().value();
	let nome: String = f.params.nome().value();
	let descrizione: String = f.params.descrizione().value();
	let url: String = f.params.url().value();
	let did_proprietario: String = f.params.did_proprietario().value();
	let attiva: bool = f.params.attiva().value();
	
	let applicazioni: ArrayOfMutableApplicazione = f.state.applicazioni();
	
	if applicazioni.length() > 0
	{
		if nome_attuale != nome
		{
			for i in 0..applicazioni.length() 
			{
				let applicazione: Applicazione = applicazioni.get_applicazione(i).value();
				
				if applicazione.nome == nome
				{
					f.events.nome_app_usato();
					return;
				}
			}
		}
		
		for i in 0..applicazioni.length() 
		{
			let mut applicazione: Applicazione = applicazioni.get_applicazione(i).value();
			
			if applicazione.nome == nome_attuale
			{
				applicazione.nome = nome;
				applicazione.descrizione = descrizione;
				applicazione.url = url;
				applicazione.did_proprietario = did_proprietario;
				applicazione.attiva_app = attiva;
				
				applicazioni.get_applicazione(i).set_value(&applicazione);
				f.events.applicazione_modificata();
				return;
			}
		}
	}
	
	f.events.applicazione_non_modificata();
}

pub fn func_modifica_ind_ult_msg(ctx: &ScFuncContext, f: &ModificaIndUltMsgContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let id_op: i32 = f.params.id_op().value();
	let did_produttore: String = f.params.did_produttore().value();
	let ind_ult_msg_dati: String = f.params.ind_ult_msg_dati().value();
	let ind_ult_msg_log: String = f.params.ind_ult_msg_log().value();

	let canali: ArrayOfMutableCanali = f.state.canali();
	
	if canali.length() > 0
	{
		for i in 0..canali.length() 
		{
			let mut canali_dati_log: Canali = canali.get_canali(i).value();
			if canali_dati_log.id_app == id_app && canali_dati_log.id_op == id_op && canali_dati_log.did_produttore == did_produttore
			{
				canali_dati_log.ind_ult_msg_dati = ind_ult_msg_dati;
				canali_dati_log.ind_ult_msg_log = ind_ult_msg_log;
				canali.get_canali(i).set_value(&canali_dati_log);
				
				f.events.indirizzi_modificati();
				return;
			}
		}
	}
}

pub fn func_modifica_stato_operazione(ctx: &ScFuncContext, f: &ModificaStatoOperazioneContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let titolo: String = f.params.titolo().value();
	let nuovo_stato: bool = f.params.nuovo_stato().value();
	
	let operazioni: ArrayOfMutableOperazione = f.state.operazioni();
	
	if operazioni.length() > 0
	{
		for i in 0..operazioni.length() 
		{
			let mut operazione: Operazione = operazioni.get_operazione(i).value();
			if operazione.id_app == id_app && operazione.titolo == titolo
			{
				operazione.attiva_op = nuovo_stato;
				operazioni.get_operazione(i).set_value(&operazione);
				
				if nuovo_stato == true
				{
					f.events.operazione_attivata();
				}
				else
				{
					f.events.operazione_disattivata();
				}
				
				return;
			}
		}
	}
	
	f.events.operazione_inesistente();
}

pub fn func_pubblica_applicazione(ctx: &ScFuncContext, f: &PubblicaApplicazioneContext) 
{
	let nome: String = f.params.nome().value();
	let descrizione: String = f.params.descrizione().value();
	let url: String = f.params.url().value();
	let did_proprietario: String = f.params.did_proprietario().value();
	
	let applicazioni: ArrayOfMutableApplicazione = f.state.applicazioni();
	
	if applicazioni.length() > 0
	{
		for i in 0..applicazioni.length() 
		{
			let applicazione: Applicazione = applicazioni.get_applicazione(i).value();
			if applicazione.nome == nome && applicazione.did_proprietario == did_proprietario
			{
				f.events.applicazione_pubblicata();
				return;
			}
		}
	}
		
	let applicazione = Applicazione {
		attiva_app: true,
		descrizione: descrizione,
		did_proprietario: did_proprietario,
		nome: nome,
		url: url,
	};
	
	applicazioni.append_applicazione().set_value(&applicazione);
	f.events.applicazione_pubblicata();
}

pub fn func_registra_utente_vc(ctx: &ScFuncContext, f: &RegistraUtenteVCContext) 
{
	let rif_utente: String = f.params.rif_utente().value();
	let vc_utente: String = f.params.vc_utente().value();
	
	let utenti: MapStringToMutableString = f.state.utenti();
	let vc_utente_reg: ScMutableString = utenti.get_string(&rif_utente);
	
	if ! vc_utente_reg.exists()
	{
		vc_utente_reg.set_value(&vc_utente);
	}
	
	f.events.utente_registrato();
}

pub fn view_dati_applicazione(ctx: &ScViewContext, f: &DatiApplicazioneContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let applicazioni: ArrayOfImmutableApplicazione = f.state.applicazioni();
	let mut risultato: String = String::from("");
	
	if (applicazioni.length() as i32) > id_app
	{
		let applicazione: Applicazione = applicazioni.get_applicazione(id_app as u32).value();
		
		if applicazione.attiva_app == true
		{
			risultato = format!("{0}|{1}|{2}|{3}", applicazione.nome, applicazione.descrizione, applicazione.url, applicazione.did_proprietario);
		}
	}
	
	f.results.dati_app().set_value(&risultato);
}

pub fn view_dati_operazione(ctx: &ScViewContext, f: &DatiOperazioneContext) 
{
	let id_op: i32 = f.params.id_op().value();
	let operazioni: ArrayOfImmutableOperazione = f.state.operazioni();
	let mut risultato: String = String::from("");
	
	if (operazioni.length() as i32) > id_op
	{
		let operazione: Operazione = operazioni.get_operazione(id_op as u32).value();

		if operazione.attiva_op == true
		{
			risultato = format!("{0}|{1}|{2}|{3}|{4}", operazione.id_app, operazione.titolo, operazione.descrizione, operazione.software, operazione.tipo);
		}
	}
	
	f.results.dati_op().set_value(&risultato);
}

pub fn view_elenco_applicazioni(ctx: &ScViewContext, f: &ElencoApplicazioniContext) 
{
	let applicazioni: ArrayOfImmutableApplicazione = f.state.applicazioni();
	
	f.results.numero_apps().set_value(applicazioni.length() as i32);
}

pub fn view_elenco_apps_utente(ctx: &ScViewContext, f: &ElencoAppsUtenteContext) 
{
	let did_utente: String = f.params.did_utente().value();
	let applicazioni: ArrayOfImmutableApplicazione = f.state.applicazioni();
	let mut risultato: String = String::from("");
	
	if applicazioni.length() > 0
	{
		for i in 0..applicazioni.length() 
		{
			let applicazione: Applicazione = applicazioni.get_applicazione(i).value();
			if applicazione.did_proprietario == did_utente
			{
				if risultato.len() == 0
				{
					risultato = format!("{0}", i);
				}
				else
				{
					risultato = format!("{0}|{1}", risultato, i);
				}
			}
		}
	}
	
	f.results.id_apps().set_value(&risultato);
}

pub fn view_elenco_operazioni(ctx: &ScViewContext, f: &ElencoOperazioniContext) 
{
	let idApp: i32 = f.params.id_app().value();
	let operazioni: ArrayOfImmutableOperazione = f.state.operazioni();
	let mut risultato: String = String::from("");
	
	if operazioni.length() > 0
	{
		for i in 0..operazioni.length() 
		{
			let operazione: Operazione = operazioni.get_operazione(i).value();
			if operazione.id_app == idApp
			{
				if risultato.len() == 0
				{
					risultato = format!("{0}", i);
				}
				else
				{
					risultato = format!("{0}|{1}", risultato, i);
				}
			}
		}
	}
	
	f.results.id_ops().set_value(&risultato);
}

pub fn view_id_applicazione(ctx: &ScViewContext, f: &IDApplicazioneContext) 
{
	let nome: String = f.params.nome().value();
	let did_proprietario: String = f.params.did_proprietario().value();
	
	let applicazioni: ArrayOfImmutableApplicazione = f.state.applicazioni();
	
	if applicazioni.length() > 0
	{
		for i in 0..applicazioni.length() 
		{
			let applicazione: Applicazione = applicazioni.get_applicazione(i).value();
			if applicazione.nome == nome && applicazione.did_proprietario == did_proprietario
			{
				f.results.id_app().set_value(i as i32);
				return;
			}
		}
	}
	
	f.results.id_app().set_value(-1);
}

pub fn view_id_operazione(ctx: &ScViewContext, f: &IDOperazioneContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let titolo: String = f.params.titolo().value();
	
	let operazioni: ArrayOfImmutableOperazione = f.state.operazioni();
	
	if operazioni.length() > 0
	{
		for i in 0..operazioni.length() 
		{
			let operazione: Operazione = operazioni.get_operazione(i).value();
			if operazione.id_app == id_app && operazione.titolo == titolo
			{
				f.results.id_op().set_value(i as i32);
				return;
			}
		}
	}
	
	f.results.id_op().set_value(-1);
}

pub fn view_ind_ult_msg(ctx: &ScViewContext, f: &IndUltMsgContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let id_op: i32 = f.params.id_op().value();
	let did_produttore: String = f.params.did_produttore().value();
	
	let canali: ArrayOfImmutableCanali = f.state.canali();
	let mut risultato: String = String::from("");
	
	if canali.length() > 0
	{
		for i in 0..canali.length() 
		{
			let mut canali_dati_log: Canali = canali.get_canali(i).value();
			if canali_dati_log.id_app == id_app && canali_dati_log.id_op == id_op && canali_dati_log.did_produttore == did_produttore
			{
				risultato = format!("{0}|{1}", canali_dati_log.ind_ult_msg_dati, canali_dati_log.ind_ult_msg_log);
				f.results.indirizzi().set_value(&risultato);
				return;
			}
		}
	}
	
	f.results.indirizzi().set_value(&risultato);
}

pub fn view_ottieni_autore(ctx: &ScViewContext, f: &OttieniAutoreContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let id_op: i32 = f.params.id_op().value();
	let did_produttore: String = f.params.did_produttore().value();
	
	let canali: ArrayOfImmutableCanali = f.state.canali();
	let mut risultato: String = String::from("");
	
	if canali.length() > 0
	{
		for i in 0..canali.length() 
		{
			let mut canali_dati_log: Canali = canali.get_canali(i).value();
			if canali_dati_log.id_app == id_app && canali_dati_log.id_op == id_op && canali_dati_log.did_produttore == did_produttore
			{
				risultato = format!("{0}|{1}", canali_dati_log.autore_dati, canali_dati_log.autore_log);
				f.results.autori().set_value(&risultato);
				return;
			}
		}
	}
	
	f.results.autori().set_value(&risultato);
}

pub fn view_ottieni_iscritto(ctx: &ScViewContext, f: &OttieniIscrittoContext) 
{
	let id_app: i32 = f.params.id_app().value();
	let id_op: i32 = f.params.id_op().value();
	let did_produttore: String = f.params.did_produttore().value();
	let did_consumatore: String = f.params.did_consumatore().value();
	
	let iscritti: ArrayOfImmutableIscritto = f.state.iscritti();
	let mut risultato: String = String::from("");
	
	if iscritti.length() > 0
	{
		for i in 0..iscritti.length() 
		{
			let iscritto: Iscritto = iscritti.get_iscritto(i).value();
			if iscritto.id_app == id_app && iscritto.id_op == id_op && iscritto.did_produttore == did_produttore && iscritto.did_consumatore == did_consumatore
			{
				risultato = format!("{0}|{1}", iscritto.iscritto_dati, iscritto.iscritto_log);
				f.results.iscritti().set_value(&risultato);
				return;
			}
		}
	}
	
	f.results.iscritti().set_value(&risultato);
}

pub fn view_verifica_utente_vc(ctx: &ScViewContext, f: &VerificaUtenteVCContext) 
{
	let rif_utente: String = f.params.rif_utente().value();
	let vc_utente: String = f.params.vc_utente().value();
	
	let utenti: MapStringToImmutableString = f.state.utenti();
	let vc_utente_reg: ScImmutableString = utenti.get_string(&rif_utente);
	
	if vc_utente_reg.exists()
	{
		let vc_utente_reg_s: String = utenti.get_string(&rif_utente).to_string();
		
		if vc_utente_reg_s == vc_utente
		{
			f.results.esito_c().set_value(true);
		}
		else
		{
			f.results.esito_c().set_value(false);
		}
		
		return;
	}
	
	f.results.esito_c().set_value(false);
}
