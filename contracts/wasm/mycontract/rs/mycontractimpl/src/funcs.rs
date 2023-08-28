
use wasmlib::*;
use mycontract::*;
use crate::*;

pub fn func_init(ctx: &ScFuncContext, f: &InitContext) 
{
    if f.params.prova().exists() 
    {
        f.state.prova().set_value(&f.params.prova().value());
        return;
    }
    f.state.prova().set_value(&String::from("Valore 1"));
}

pub fn func_set_prova(ctx: &ScFuncContext, f: &SetProvaContext) 
{
	f.state.prova().set_value(&String::from("Valore 2"));
}

pub fn view_get_prova(ctx: &ScViewContext, f: &GetProvaContext) 
{
    let prova = f.state.prova();
    if prova.exists() 
    {
        f.results.prova().set_value(&prova.value());
    }
}
