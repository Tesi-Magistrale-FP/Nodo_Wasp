
use wasmlib::*;
use iota_client::Client;
use iota_client::Result;
use crate::*;

pub async fn view_get_info(ctx: &ScViewContext, f: &GetInfoContext) -> Result<()>
{
	let iota = Client::builder()
        .with_node(&net_url)
        .unwrap()
        .finish()
        .await
        .unwrap();

    let info = iota.get_info().await.unwrap();
    let risultato: String = format!("- Nome: {}\n- Versione: {}\n- Attivo: {}\n- Rete: {}\n- Url: {}\n- Min. PoW score: {}\n- Messaggi ricevuti al secondo: {}\n- Messaggi referenziati al secondo: {}\n- Features: {}, {}", info.nodeinfo.name, info.nodeinfo.version, info.nodeinfo.is_healthy, info.nodeinfo.network_id, info.url, info.nodeinfo.min_pow_score, info.nodeinfo.messages_per_second, info.nodeinfo.referenced_messages_per_second, info.nodeinfo.features[0], info.nodeinfo.features[1]);

	f.results.info().set_value(info);
	
	Ok(())
}

pub fn view_get_info(ctx: &ScViewContext, f: &GetInfoContext) {
}
