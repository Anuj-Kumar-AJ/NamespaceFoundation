// @generated by protobuf-ts 2.7.0 with parameter force_disable_services,add_pb_suffix
// @generated from protobuf file "languages/nodejs/testdata/services/simple/service.proto" (package "languages.nodejs.testdata.services.simple", syntax proto3)
// tslint:disable
//
// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation
//
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message languages.nodejs.testdata.services.simple.PostRequest
 */
export interface PostRequest {
    /**
     * @generated from protobuf field: string input = 1;
     */
    input: string;
}
/**
 * @generated from protobuf message languages.nodejs.testdata.services.simple.PostResponse
 */
export interface PostResponse {
    /**
     * @generated from protobuf field: string output = 1;
     */
    output: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class PostRequest$Type extends MessageType<PostRequest> {
    constructor() {
        super("languages.nodejs.testdata.services.simple.PostRequest", [
            { no: 1, name: "input", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<PostRequest>): PostRequest {
        const message = { input: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<PostRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PostRequest): PostRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string input */ 1:
                    message.input = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: PostRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string input = 1; */
        if (message.input !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.input);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message languages.nodejs.testdata.services.simple.PostRequest
 */
export const PostRequest = new PostRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PostResponse$Type extends MessageType<PostResponse> {
    constructor() {
        super("languages.nodejs.testdata.services.simple.PostResponse", [
            { no: 1, name: "output", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<PostResponse>): PostResponse {
        const message = { output: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<PostResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PostResponse): PostResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string output */ 1:
                    message.output = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: PostResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string output = 1; */
        if (message.output !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.output);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message languages.nodejs.testdata.services.simple.PostResponse
 */
export const PostResponse = new PostResponse$Type();