// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import * as wasmlib from '../index';
import * as wasmtypes from '../wasmtypes';

export class CoreAccountsEventHandlers implements wasmlib.IEventHandlers {
    private myID: u32;
    private coreAccountsHandlers: Map<string, (evt: CoreAccountsEventHandlers, dec: wasmlib.WasmDecoder) => void> = new Map();

    /* eslint-disable @typescript-eslint/no-empty-function */
    foundryCreated: (evt: EventFoundryCreated) => void = () => {};
    foundryDestroyed: (evt: EventFoundryDestroyed) => void = () => {};
    foundryModified: (evt: EventFoundryModified) => void = () => {};
    /* eslint-enable @typescript-eslint/no-empty-function */

    public constructor() {
        this.myID = wasmlib.eventHandlersGenerateID();
        this.coreAccountsHandlers.set('coreaccounts.foundryCreated', (evt: CoreAccountsEventHandlers, dec: wasmlib.WasmDecoder) => evt.foundryCreated(new EventFoundryCreated(dec)));
        this.coreAccountsHandlers.set('coreaccounts.foundryDestroyed', (evt: CoreAccountsEventHandlers, dec: wasmlib.WasmDecoder) => evt.foundryDestroyed(new EventFoundryDestroyed(dec)));
        this.coreAccountsHandlers.set('coreaccounts.foundryModified', (evt: CoreAccountsEventHandlers, dec: wasmlib.WasmDecoder) => evt.foundryModified(new EventFoundryModified(dec)));
    }

    public callHandler(topic: string, dec: wasmlib.WasmDecoder): void {
        const handler = this.coreAccountsHandlers.get(topic);
        if (handler) {
            handler(this, dec);
        }
    }

    public id(): u32 {
        return this.myID;
    }

    public onCoreAccountsFoundryCreated(handler: (evt: EventFoundryCreated) => void): void {
        this.foundryCreated = handler;
    }

    public onCoreAccountsFoundryDestroyed(handler: (evt: EventFoundryDestroyed) => void): void {
        this.foundryDestroyed = handler;
    }

    public onCoreAccountsFoundryModified(handler: (evt: EventFoundryModified) => void): void {
        this.foundryModified = handler;
    }
}

export class EventFoundryCreated {
    public readonly timestamp: u64;
    public readonly foundrySN: u32;

    public constructor(dec: wasmlib.WasmDecoder) {
        this.timestamp = wasmtypes.uint64Decode(dec);
        this.foundrySN = wasmtypes.uint32Decode(dec);
        dec.close();
    }
}

export class EventFoundryDestroyed {
    public readonly timestamp: u64;
    public readonly foundrySN: u32;

    public constructor(dec: wasmlib.WasmDecoder) {
        this.timestamp = wasmtypes.uint64Decode(dec);
        this.foundrySN = wasmtypes.uint32Decode(dec);
        dec.close();
    }
}

export class EventFoundryModified {
    public readonly timestamp: u64;
    public readonly foundrySN: u32;

    public constructor(dec: wasmlib.WasmDecoder) {
        this.timestamp = wasmtypes.uint64Decode(dec);
        this.foundrySN = wasmtypes.uint32Decode(dec);
        dec.close();
    }
}
