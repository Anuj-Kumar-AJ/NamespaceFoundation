// This file was automatically generated.

import * as impl from "./impl";
import { DependencyGraph, Initializer, Registrar } from "@namespacelabs/foundation";



export const Package = {
  name: "namespacelabs.dev/foundation/languages/nodejs/testdata/services/simplehttp",
};

export const TransitiveInitializers: Initializer[] = [
];

export type WireService = (registrar: Registrar) => void;
export const wireService: WireService = impl.wireService;