import * as jspb from 'google-protobuf'

import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';


export class GetFieldParams extends jspb.Message {
  getField(): string;
  setField(value: string): GetFieldParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFieldParams.AsObject;
  static toObject(includeInstance: boolean, msg: GetFieldParams): GetFieldParams.AsObject;
  static serializeBinaryToWriter(message: GetFieldParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFieldParams;
  static deserializeBinaryFromReader(message: GetFieldParams, reader: jspb.BinaryReader): GetFieldParams;
}

export namespace GetFieldParams {
  export type AsObject = {
    field: string,
  }
}

export class AddFieldParams extends jspb.Message {
  getField(): string;
  setField(value: string): AddFieldParams;

  getSystemname(): string;
  setSystemname(value: string): AddFieldParams;

  getEngfullname(): string;
  setEngfullname(value: string): AddFieldParams;

  getEngshortname(): string;
  setEngshortname(value: string): AddFieldParams;

  getLabel(): string;
  setLabel(value: string): AddFieldParams;

  getComment(): string;
  setComment(value: string): AddFieldParams;

  getUielem(): string;
  setUielem(value: string): AddFieldParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddFieldParams.AsObject;
  static toObject(includeInstance: boolean, msg: AddFieldParams): AddFieldParams.AsObject;
  static serializeBinaryToWriter(message: AddFieldParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddFieldParams;
  static deserializeBinaryFromReader(message: AddFieldParams, reader: jspb.BinaryReader): AddFieldParams;
}

export namespace AddFieldParams {
  export type AsObject = {
    field: string,
    systemname: string,
    engfullname: string,
    engshortname: string,
    label: string,
    comment: string,
    uielem: string,
  }
}

export class DeleteFieldParams extends jspb.Message {
  getId(): string;
  setId(value: string): DeleteFieldParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteFieldParams.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteFieldParams): DeleteFieldParams.AsObject;
  static serializeBinaryToWriter(message: DeleteFieldParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteFieldParams;
  static deserializeBinaryFromReader(message: DeleteFieldParams, reader: jspb.BinaryReader): DeleteFieldParams;
}

export namespace DeleteFieldParams {
  export type AsObject = {
    id: string,
  }
}

export class FieldObject extends jspb.Message {
  getId(): string;
  setId(value: string): FieldObject;

  getSystemname(): string;
  setSystemname(value: string): FieldObject;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FieldObject.AsObject;
  static toObject(includeInstance: boolean, msg: FieldObject): FieldObject.AsObject;
  static serializeBinaryToWriter(message: FieldObject, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FieldObject;
  static deserializeBinaryFromReader(message: FieldObject, reader: jspb.BinaryReader): FieldObject;
}

export namespace FieldObject {
  export type AsObject = {
    id: string,
    systemname: string,
  }
}

export class FieldResponse extends jspb.Message {
  getFieldsList(): Array<FieldObject>;
  setFieldsList(value: Array<FieldObject>): FieldResponse;
  clearFieldsList(): FieldResponse;
  addFields(value?: FieldObject, index?: number): FieldObject;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FieldResponse.AsObject;
  static toObject(includeInstance: boolean, msg: FieldResponse): FieldResponse.AsObject;
  static serializeBinaryToWriter(message: FieldResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FieldResponse;
  static deserializeBinaryFromReader(message: FieldResponse, reader: jspb.BinaryReader): FieldResponse;
}

export namespace FieldResponse {
  export type AsObject = {
    fieldsList: Array<FieldObject.AsObject>,
  }
}

export class deleteFieldResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): deleteFieldResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): deleteFieldResponse.AsObject;
  static toObject(includeInstance: boolean, msg: deleteFieldResponse): deleteFieldResponse.AsObject;
  static serializeBinaryToWriter(message: deleteFieldResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): deleteFieldResponse;
  static deserializeBinaryFromReader(message: deleteFieldResponse, reader: jspb.BinaryReader): deleteFieldResponse;
}

export namespace deleteFieldResponse {
  export type AsObject = {
    message: string,
  }
}

export class SubscribeRequest extends jspb.Message {
  getQuery(): string;
  setQuery(value: string): SubscribeRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubscribeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SubscribeRequest): SubscribeRequest.AsObject;
  static serializeBinaryToWriter(message: SubscribeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubscribeRequest;
  static deserializeBinaryFromReader(message: SubscribeRequest, reader: jspb.BinaryReader): SubscribeRequest;
}

export namespace SubscribeRequest {
  export type AsObject = {
    query: string,
  }
}

export class SubscribeResponse extends jspb.Message {
  getQuery(): string;
  setQuery(value: string): SubscribeResponse;

  getData(): google_protobuf_struct_pb.Struct | undefined;
  setData(value?: google_protobuf_struct_pb.Struct): SubscribeResponse;
  hasData(): boolean;
  clearData(): SubscribeResponse;

  getEventsList(): Array<SubscribeResponse.Event>;
  setEventsList(value: Array<SubscribeResponse.Event>): SubscribeResponse;
  clearEventsList(): SubscribeResponse;
  addEvents(value?: SubscribeResponse.Event, index?: number): SubscribeResponse.Event;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SubscribeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SubscribeResponse): SubscribeResponse.AsObject;
  static serializeBinaryToWriter(message: SubscribeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SubscribeResponse;
  static deserializeBinaryFromReader(message: SubscribeResponse, reader: jspb.BinaryReader): SubscribeResponse;
}

export namespace SubscribeResponse {
  export type AsObject = {
    query: string,
    data?: google_protobuf_struct_pb.Struct.AsObject,
    eventsList: Array<SubscribeResponse.Event.AsObject>,
  }

  export class Event extends jspb.Message {
    getKey(): string;
    setKey(value: string): Event;

    getEventsList(): Array<string>;
    setEventsList(value: Array<string>): Event;
    clearEventsList(): Event;
    addEvents(value: string, index?: number): Event;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Event.AsObject;
    static toObject(includeInstance: boolean, msg: Event): Event.AsObject;
    static serializeBinaryToWriter(message: Event, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Event;
    static deserializeBinaryFromReader(message: Event, reader: jspb.BinaryReader): Event;
  }

  export namespace Event {
    export type AsObject = {
      key: string,
      eventsList: Array<string>,
    }
  }

}

