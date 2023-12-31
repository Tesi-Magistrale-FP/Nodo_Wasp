// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import * as wasmlib from '../index';
import * as wasmtypes from '../wasmtypes';

export class CoreRootEventHandlers implements wasmlib.IEventHandlers {
    private myID: u32;
    private coreRootHandlers: Map<string, (evt: CoreRootEventHandlers, dec: wasmlib.WasmDecoder) => void> = new Map();

    /* eslint-disable @typescript-eslint/no-empty-function */
    deploy: (evt: EventDeploy) => void = () => {};
    grant: (evt: EventGrant) => void = () => {};
    revoke: (evt: EventRevoke) => void = () => {};
    /* eslint-enable @typescript-eslint/no-empty-function */

    public constructor() {
        this.myID = wasmlib.eventHandlersGenerateID();
        this.coreRootHandlers.set('coreroot.deploy', (evt: CoreRootEventHandlers, dec: wasmlib.WasmDecoder) => evt.deploy(new EventDeploy(dec)));
        this.coreRootHandlers.set('coreroot.grant', (evt: CoreRootEventHandlers, dec: wasmlib.WasmDecoder) => evt.grant(new EventGrant(dec)));
        this.coreRootHandlers.set('coreroot.revoke', (evt: CoreRootEventHandlers, dec: wasmlib.WasmDecoder) => evt.revoke(new EventRevoke(dec)));
    }

    public callHandler(topic: string, dec: wasmlib.WasmDecoder): void {
        const handler = this.coreRootHandlers.get(topic);
        if (handler) {
            handler(this, dec);
        }
    }

    public id(): u32 {
        return this.myID;
    }

    public onCoreRootDeploy(handler: (evt: EventDeploy) => void): void {
        this.deploy = handler;
    }

    public onCoreRootGrant(handler: (evt: EventGrant) => void): void {
        this.grant = handler;
    }

    public onCoreRootRevoke(handler: (evt: EventRevoke) => void): void {
        this.revoke = handler;
    }
}

export class EventDeploy {
    public readonly timestamp: u64;
    public readonly name: string;
    public readonly progHash: wasmtypes.ScHash;

    public constructor(dec: wasmlib.WasmDecoder) {
        this.timestamp = wasmtypes.uint64Decode(dec);
        this.name = wasmtypes.stringDecode(dec);
        this.progHash = wasmtypes.hashDecode(dec);
        dec.close();
    }
}

export class EventGrant {
    public readonly timestamp: u64;
    public readonly deployer: wasmtypes.ScAgentID;

    public constructor(dec: wasmlib.WasmDecoder) {
        this.timestamp = wasmtypes.uint64Decode(dec);
        this.deployer = wasmtypes.agentIDDecode(dec);
        dec.close();
    }
}

export class EventRevoke {
    public readonly timestamp: u64;
    public readonly deployer: wasmtypes.ScAgentID;

    public constructor(dec: wasmlib.WasmDecoder) {
        this.timestamp = wasmtypes.uint64Decode(dec);
        this.deployer = wasmtypes.agentIDDecode(dec);
        dec.close();
    }
}
