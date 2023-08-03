// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

#![allow(dead_code)]
#![allow(unused_imports)]

use crate::*;
use crate::coreerrors::*;

#[derive(Clone)]
pub struct ImmutableRegisterErrorParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableRegisterErrorParams {
    pub fn new() -> ImmutableRegisterErrorParams {
        ImmutableRegisterErrorParams {
            proxy: params_proxy(),
        }
    }

    // error message template string
    pub fn template(&self) -> ScImmutableString {
        ScImmutableString::new(self.proxy.root(PARAM_TEMPLATE))
    }
}

#[derive(Clone)]
pub struct MutableRegisterErrorParams {
    pub(crate) proxy: Proxy,
}

impl MutableRegisterErrorParams {
    // error message template string
    pub fn template(&self) -> ScMutableString {
        ScMutableString::new(self.proxy.root(PARAM_TEMPLATE))
    }
}

#[derive(Clone)]
pub struct ImmutableGetErrorMessageFormatParams {
    pub(crate) proxy: Proxy,
}

impl ImmutableGetErrorMessageFormatParams {
    pub fn new() -> ImmutableGetErrorMessageFormatParams {
        ImmutableGetErrorMessageFormatParams {
            proxy: params_proxy(),
        }
    }

    // serialized error code
    pub fn error_code(&self) -> ScImmutableBytes {
        ScImmutableBytes::new(self.proxy.root(PARAM_ERROR_CODE))
    }
}

#[derive(Clone)]
pub struct MutableGetErrorMessageFormatParams {
    pub(crate) proxy: Proxy,
}

impl MutableGetErrorMessageFormatParams {
    // serialized error code
    pub fn error_code(&self) -> ScMutableBytes {
        ScMutableBytes::new(self.proxy.root(PARAM_ERROR_CODE))
    }
}
