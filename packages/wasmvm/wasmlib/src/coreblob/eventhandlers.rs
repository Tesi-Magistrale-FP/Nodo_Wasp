// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

use std::collections::HashMap;

use crate::*;

use crate::*;

pub struct CoreBlobEventHandlers {
    my_id: u32,
    core_blob_handlers: HashMap<&'static str, fn(evt: &CoreBlobEventHandlers, dec: &mut WasmDecoder)>,

    store: Box<dyn Fn(&EventStore)>,
}

impl IEventHandlers for CoreBlobEventHandlers {
    fn call_handler(&self, topic: &str, dec: &mut WasmDecoder) {
        if let Some(handler) = self.core_blob_handlers.get(topic) {
            handler(self, dec);
        }
    }

    fn id(&self) -> u32 {
        self.my_id
    }
}

unsafe impl Send for CoreBlobEventHandlers {}
unsafe impl Sync for CoreBlobEventHandlers {}

impl CoreBlobEventHandlers {
    pub fn new() -> CoreBlobEventHandlers {
        let mut handlers: HashMap<&str, fn(evt: &CoreBlobEventHandlers, dec: &mut WasmDecoder)> = HashMap::new();
        handlers.insert("coreblob.store", |e, m| { (e.store)(&EventStore::new(m)); });
        return CoreBlobEventHandlers {
            my_id: EventHandlers::generate_id(),
            core_blob_handlers: handlers,
            store: Box::new(|_e| {}),
        };
    }

    pub fn on_core_blob_store<F>(&mut self, handler: F)
        where F: Fn(&EventStore) + 'static {
        self.store = Box::new(handler);
    }
}

pub struct EventStore {
    pub timestamp: u64,
    pub blob_hash: ScHash,
}

impl EventStore {
    pub fn new(dec: &mut WasmDecoder) -> EventStore {
        EventStore {
            timestamp: uint64_decode(dec),
            blob_hash: hash_decode(dec),
        }
    }
}
