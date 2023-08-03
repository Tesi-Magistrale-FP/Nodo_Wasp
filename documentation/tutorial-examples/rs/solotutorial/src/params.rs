// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(unused_imports)]

use wasmlib::*;

use crate::*;

#[derive(Clone)]
pub struct ImmutableStoreStringParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableStoreStringParams {
    pub fn new() -> ImmutableStoreStringParams {
        ImmutableStoreStringParams {
            proxy: params_proxy(),
        }
    }

    pub fn str(&self) -> ScImmutableString {
        ScImmutableString::new(self.proxy.root(PARAM_STR))
    }
}

#[derive(Clone)]
pub struct MutableStoreStringParams {
    pub(crate) proxy: Proxy,
}

impl MutableStoreStringParams {
    pub fn str(&self) -> ScMutableString {
        ScMutableString::new(self.proxy.root(PARAM_STR))
    }
}
