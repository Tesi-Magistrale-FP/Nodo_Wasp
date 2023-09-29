use wasmlib::*;
use performance::*;
use crate::*;

pub fn func_memorizza_valore(ctx: &ScFuncContext, f: &MemorizzaValoreContext) 
{
	let valore: String = f.params.valore().value();
	let valori: ArrayOfMutableString = f.state.valori();
	valori.append_string().set_value(&valore);
	f.events.valore_registrato();
}

pub fn view_ottieni_valore(ctx: &ScViewContext, f: &OttieniValoreContext) 
{
	let indice: i32 = f.params.indice().value();
	let valori: ArrayOfImmutableString = f.state.valori();
	let valore: String = valori.get_string(indice as u32).to_string();
	f.results.valore().set_value(&valore);
}