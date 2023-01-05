/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "dataocean.dataocean";

export interface MsgCreateVideo {
  creator: string;
  title: string;
  description: string;
  coverLink: string;
  videoLink: string;
  priceMB: number;
}

export interface MsgCreateVideoResponse {
  id: number;
}

export interface MsgPlayVideo {
  creator: string;
  videoId: number;
}

export interface MsgPlayVideoResponse {
  url: string;
  exp: string;
}

export interface MsgPaySign {
  creator: string;
  videoId: number;
  receivedSizeMB: number;
  timestamp: number;
}

export interface MsgPaySignResponse {
}

export interface MsgSubmitPaySign {
  creator: string;
  paySign: string;
}

export interface MsgSubmitPaySignResponse {
}

function createBaseMsgCreateVideo(): MsgCreateVideo {
  return { creator: "", title: "", description: "", coverLink: "", videoLink: "", priceMB: 0 };
}

export const MsgCreateVideo = {
  encode(message: MsgCreateVideo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.coverLink !== "") {
      writer.uint32(34).string(message.coverLink);
    }
    if (message.videoLink !== "") {
      writer.uint32(42).string(message.videoLink);
    }
    if (message.priceMB !== 0) {
      writer.uint32(48).uint64(message.priceMB);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateVideo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateVideo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.coverLink = reader.string();
          break;
        case 5:
          message.videoLink = reader.string();
          break;
        case 6:
          message.priceMB = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVideo {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      title: isSet(object.title) ? String(object.title) : "",
      description: isSet(object.description) ? String(object.description) : "",
      coverLink: isSet(object.coverLink) ? String(object.coverLink) : "",
      videoLink: isSet(object.videoLink) ? String(object.videoLink) : "",
      priceMB: isSet(object.priceMB) ? Number(object.priceMB) : 0,
    };
  },

  toJSON(message: MsgCreateVideo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    message.description !== undefined && (obj.description = message.description);
    message.coverLink !== undefined && (obj.coverLink = message.coverLink);
    message.videoLink !== undefined && (obj.videoLink = message.videoLink);
    message.priceMB !== undefined && (obj.priceMB = Math.round(message.priceMB));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVideo>, I>>(object: I): MsgCreateVideo {
    const message = createBaseMsgCreateVideo();
    message.creator = object.creator ?? "";
    message.title = object.title ?? "";
    message.description = object.description ?? "";
    message.coverLink = object.coverLink ?? "";
    message.videoLink = object.videoLink ?? "";
    message.priceMB = object.priceMB ?? 0;
    return message;
  },
};

function createBaseMsgCreateVideoResponse(): MsgCreateVideoResponse {
  return { id: 0 };
}

export const MsgCreateVideoResponse = {
  encode(message: MsgCreateVideoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateVideoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateVideoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateVideoResponse {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: MsgCreateVideoResponse): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateVideoResponse>, I>>(object: I): MsgCreateVideoResponse {
    const message = createBaseMsgCreateVideoResponse();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgPlayVideo(): MsgPlayVideo {
  return { creator: "", videoId: 0 };
}

export const MsgPlayVideo = {
  encode(message: MsgPlayVideo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.videoId !== 0) {
      writer.uint32(16).uint64(message.videoId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlayVideo {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlayVideo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.videoId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlayVideo {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      videoId: isSet(object.videoId) ? Number(object.videoId) : 0,
    };
  },

  toJSON(message: MsgPlayVideo): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.videoId !== undefined && (obj.videoId = Math.round(message.videoId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPlayVideo>, I>>(object: I): MsgPlayVideo {
    const message = createBaseMsgPlayVideo();
    message.creator = object.creator ?? "";
    message.videoId = object.videoId ?? 0;
    return message;
  },
};

function createBaseMsgPlayVideoResponse(): MsgPlayVideoResponse {
  return { url: "", exp: "" };
}

export const MsgPlayVideoResponse = {
  encode(message: MsgPlayVideoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.url !== "") {
      writer.uint32(10).string(message.url);
    }
    if (message.exp !== "") {
      writer.uint32(18).string(message.exp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlayVideoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlayVideoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.url = reader.string();
          break;
        case 2:
          message.exp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPlayVideoResponse {
    return { url: isSet(object.url) ? String(object.url) : "", exp: isSet(object.exp) ? String(object.exp) : "" };
  },

  toJSON(message: MsgPlayVideoResponse): unknown {
    const obj: any = {};
    message.url !== undefined && (obj.url = message.url);
    message.exp !== undefined && (obj.exp = message.exp);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPlayVideoResponse>, I>>(object: I): MsgPlayVideoResponse {
    const message = createBaseMsgPlayVideoResponse();
    message.url = object.url ?? "";
    message.exp = object.exp ?? "";
    return message;
  },
};

function createBaseMsgPaySign(): MsgPaySign {
  return { creator: "", videoId: 0, receivedSizeMB: 0, timestamp: 0 };
}

export const MsgPaySign = {
  encode(message: MsgPaySign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.videoId !== 0) {
      writer.uint32(16).uint64(message.videoId);
    }
    if (message.receivedSizeMB !== 0) {
      writer.uint32(24).uint64(message.receivedSizeMB);
    }
    if (message.timestamp !== 0) {
      writer.uint32(32).uint64(message.timestamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPaySign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPaySign();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.videoId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.receivedSizeMB = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.timestamp = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPaySign {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      videoId: isSet(object.videoId) ? Number(object.videoId) : 0,
      receivedSizeMB: isSet(object.receivedSizeMB) ? Number(object.receivedSizeMB) : 0,
      timestamp: isSet(object.timestamp) ? Number(object.timestamp) : 0,
    };
  },

  toJSON(message: MsgPaySign): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.videoId !== undefined && (obj.videoId = Math.round(message.videoId));
    message.receivedSizeMB !== undefined && (obj.receivedSizeMB = Math.round(message.receivedSizeMB));
    message.timestamp !== undefined && (obj.timestamp = Math.round(message.timestamp));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPaySign>, I>>(object: I): MsgPaySign {
    const message = createBaseMsgPaySign();
    message.creator = object.creator ?? "";
    message.videoId = object.videoId ?? 0;
    message.receivedSizeMB = object.receivedSizeMB ?? 0;
    message.timestamp = object.timestamp ?? 0;
    return message;
  },
};

function createBaseMsgPaySignResponse(): MsgPaySignResponse {
  return {};
}

export const MsgPaySignResponse = {
  encode(_: MsgPaySignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPaySignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPaySignResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgPaySignResponse {
    return {};
  },

  toJSON(_: MsgPaySignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPaySignResponse>, I>>(_: I): MsgPaySignResponse {
    const message = createBaseMsgPaySignResponse();
    return message;
  },
};

function createBaseMsgSubmitPaySign(): MsgSubmitPaySign {
  return { creator: "", paySign: "" };
}

export const MsgSubmitPaySign = {
  encode(message: MsgSubmitPaySign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.paySign !== "") {
      writer.uint32(18).string(message.paySign);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitPaySign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitPaySign();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.paySign = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSubmitPaySign {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      paySign: isSet(object.paySign) ? String(object.paySign) : "",
    };
  },

  toJSON(message: MsgSubmitPaySign): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.paySign !== undefined && (obj.paySign = message.paySign);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitPaySign>, I>>(object: I): MsgSubmitPaySign {
    const message = createBaseMsgSubmitPaySign();
    message.creator = object.creator ?? "";
    message.paySign = object.paySign ?? "";
    return message;
  },
};

function createBaseMsgSubmitPaySignResponse(): MsgSubmitPaySignResponse {
  return {};
}

export const MsgSubmitPaySignResponse = {
  encode(_: MsgSubmitPaySignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSubmitPaySignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSubmitPaySignResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgSubmitPaySignResponse {
    return {};
  },

  toJSON(_: MsgSubmitPaySignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSubmitPaySignResponse>, I>>(_: I): MsgSubmitPaySignResponse {
    const message = createBaseMsgSubmitPaySignResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateVideo(request: MsgCreateVideo): Promise<MsgCreateVideoResponse>;
  PlayVideo(request: MsgPlayVideo): Promise<MsgPlayVideoResponse>;
  PaySign(request: MsgPaySign): Promise<MsgPaySignResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SubmitPaySign(request: MsgSubmitPaySign): Promise<MsgSubmitPaySignResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateVideo = this.CreateVideo.bind(this);
    this.PlayVideo = this.PlayVideo.bind(this);
    this.PaySign = this.PaySign.bind(this);
    this.SubmitPaySign = this.SubmitPaySign.bind(this);
  }
  CreateVideo(request: MsgCreateVideo): Promise<MsgCreateVideoResponse> {
    const data = MsgCreateVideo.encode(request).finish();
    const promise = this.rpc.request("dataocean.dataocean.Msg", "CreateVideo", data);
    return promise.then((data) => MsgCreateVideoResponse.decode(new _m0.Reader(data)));
  }

  PlayVideo(request: MsgPlayVideo): Promise<MsgPlayVideoResponse> {
    const data = MsgPlayVideo.encode(request).finish();
    const promise = this.rpc.request("dataocean.dataocean.Msg", "PlayVideo", data);
    return promise.then((data) => MsgPlayVideoResponse.decode(new _m0.Reader(data)));
  }

  PaySign(request: MsgPaySign): Promise<MsgPaySignResponse> {
    const data = MsgPaySign.encode(request).finish();
    const promise = this.rpc.request("dataocean.dataocean.Msg", "PaySign", data);
    return promise.then((data) => MsgPaySignResponse.decode(new _m0.Reader(data)));
  }

  SubmitPaySign(request: MsgSubmitPaySign): Promise<MsgSubmitPaySignResponse> {
    const data = MsgSubmitPaySign.encode(request).finish();
    const promise = this.rpc.request("dataocean.dataocean.Msg", "SubmitPaySign", data);
    return promise.then((data) => MsgSubmitPaySignResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
