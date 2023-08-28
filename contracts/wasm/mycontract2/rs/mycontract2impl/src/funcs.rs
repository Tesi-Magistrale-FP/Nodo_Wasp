
use wasmlib::*;

use crate::*;

pub fn func_init(ctx: &ScFuncContext, f: &InitContext) 
{
    if f.params.prova().exists() 
    {
        f.state.prova().set_value(&f.params.prova().value());
        return;
    }
    f.state.prova().set_value(&String::from("Valore 6"));
}

pub fn func_set_prova(ctx: &ScFuncContext, f: &SetProvaContext) 
{
	f.state.prova().set_value(&f.params.prova().value());
}

pub fn view_get_prova(ctx: &ScViewContext, f: &GetProvaContext) 
{
	let prova = f.state.prova();
    if prova.exists() 
    {
        f.results.prova().set_value(&prova.value());
    }
}
