import * as jspb from 'google-protobuf'

import * as protoc$gen$openapiv2_options_annotations_pb from './protoc-gen-openapiv2/options/annotations_pb';
import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';


export class BulkFields extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BulkFields.AsObject;
  static toObject(includeInstance: boolean, msg: BulkFields): BulkFields.AsObject;
  static serializeBinaryToWriter(message: BulkFields, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BulkFields;
  static deserializeBinaryFromReader(message: BulkFields, reader: jspb.BinaryReader): BulkFields;
}

export namespace BulkFields {
  export type AsObject = {
  }
}

export class GetFieldParams extends jspb.Message {
  getId(): string;
  setId(value: string): GetFieldParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetFieldParams.AsObject;
  static toObject(includeInstance: boolean, msg: GetFieldParams): GetFieldParams.AsObject;
  static serializeBinaryToWriter(message: GetFieldParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetFieldParams;
  static deserializeBinaryFromReader(message: GetFieldParams, reader: jspb.BinaryReader): GetFieldParams;
}

export namespace GetFieldParams {
  export type AsObject = {
    id: string,
  }
}

export class AddFieldParams extends jspb.Message {
  getSystemname(): string;
  setSystemname(value: string): AddFieldParams;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddFieldParams.AsObject;
  static toObject(includeInstance: boolean, msg: AddFieldParams): AddFieldParams.AsObject;
  static serializeBinaryToWriter(message: AddFieldParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddFieldParams;
  static deserializeBinaryFromReader(message: AddFieldParams, reader: jspb.BinaryReader): AddFieldParams;
}

export namespace AddFieldParams {
  export type AsObject = {
    systemname: string,
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

