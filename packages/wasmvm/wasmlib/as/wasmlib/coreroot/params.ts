// Code generated by schema tool; DO NOT EDIT.

// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

import * as wasmtypes from '../wasmtypes';
import * as sc from './index';

export class MapStringToImmutableBytes extends wasmtypes.ScProxy {

    getBytes(key: string): wasmtypes.ScImmutableBytes {
        return new wasmtypes.ScImmutableBytes(this.proxy.key(wasmtypes.stringToBytes(key)));
    }
}

export class ImmutableDeployContractParams extends wasmtypes.ScProxy {
    // additional params for smart contract init function
    initParams(): sc.MapStringToImmutableBytes {
        return new sc.MapStringToImmutableBytes(this.proxy);
    }

    // The name of the contract to be deployed, used to calculate the contract's hname.
    // The hname must be unique among all contract hnames in the chain.
    name(): wasmtypes.ScImmutableString {
        return new wasmtypes.ScImmutableString(this.proxy.root(sc.ParamName));
    }

    // hash of blob that has been previously stored in blob contract
    programHash(): wasmtypes.ScImmutableHash {
        return new wasmtypes.ScImmutableHash(this.proxy.root(sc.ParamProgramHash));
    }
}

export class MapStringToMutableBytes extends wasmtypes.ScProxy {

    clear(): void {
        this.proxy.clearMap();
    }

    getBytes(key: string): wasmtypes.ScMutableBytes {
        return new wasmtypes.ScMutableBytes(this.proxy.key(wasmtypes.stringToBytes(key)));
    }
}

export class MutableDeployContractParams extends wasmtypes.ScProxy {
    // additional params for smart contract init function
    initParams(): sc.MapStringToMutableBytes {
        return new sc.MapStringToMutableBytes(this.proxy);
    }

    // The name of the contract to be deployed, used to calculate the contract's hname.
    // The hname must be unique among all contract hnames in the chain.
    name(): wasmtypes.ScMutableString {
        return new wasmtypes.ScMutableString(this.proxy.root(sc.ParamName));
    }

    // hash of blob that has been previously stored in blob contract
    programHash(): wasmtypes.ScMutableHash {
        return new wasmtypes.ScMutableHash(this.proxy.root(sc.ParamProgramHash));
    }
}

export class ImmutableGrantDeployPermissionParams extends wasmtypes.ScProxy {
    // agent to grant deploy permission to
    deployer(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamDeployer));
    }
}

export class MutableGrantDeployPermissionParams extends wasmtypes.ScProxy {
    // agent to grant deploy permission to
    deployer(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamDeployer));
    }
}

export class ImmutableRequireDeployPermissionsParams extends wasmtypes.ScProxy {
    // turns permission check on or off
    deployPermissionsEnabled(): wasmtypes.ScImmutableBool {
        return new wasmtypes.ScImmutableBool(this.proxy.root(sc.ParamDeployPermissionsEnabled));
    }
}

export class MutableRequireDeployPermissionsParams extends wasmtypes.ScProxy {
    // turns permission check on or off
    deployPermissionsEnabled(): wasmtypes.ScMutableBool {
        return new wasmtypes.ScMutableBool(this.proxy.root(sc.ParamDeployPermissionsEnabled));
    }
}

export class ImmutableRevokeDeployPermissionParams extends wasmtypes.ScProxy {
    // agent to revoke deploy permission for
    deployer(): wasmtypes.ScImmutableAgentID {
        return new wasmtypes.ScImmutableAgentID(this.proxy.root(sc.ParamDeployer));
    }
}

export class MutableRevokeDeployPermissionParams extends wasmtypes.ScProxy {
    // agent to revoke deploy permission for
    deployer(): wasmtypes.ScMutableAgentID {
        return new wasmtypes.ScMutableAgentID(this.proxy.root(sc.ParamDeployer));
    }
}

export class ImmutableFindContractParams extends wasmtypes.ScProxy {
    // The smart contract’s Hname
    hname(): wasmtypes.ScImmutableHname {
        return new wasmtypes.ScImmutableHname(this.proxy.root(sc.ParamHname));
    }
}

export class MutableFindContractParams extends wasmtypes.ScProxy {
    // The smart contract’s Hname
    hname(): wasmtypes.ScMutableHname {
        return new wasmtypes.ScMutableHname(this.proxy.root(sc.ParamHname));
    }
}