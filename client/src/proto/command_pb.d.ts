import * as jspb from 'google-protobuf'

import * as protoc$gen$openapiv2_options_annotations_pb from './protoc-gen-openapiv2/options/annotations_pb';
import * as google_api_annotations_pb from './google/api/annotations_pb';
import * as google_protobuf_struct_pb from 'google-protobuf/google/protobuf/struct_pb';


export class AddFieldParams extends jspb.Message {
  getId(): string;
  setId(value: string): AddFieldParams;

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
    id: string,
    systemname: string,
    engfullname: string,
    engshortname: string,
    label: string,
    comment: string,
    uielem: string,
  }
}

export class FieldObject1 extends jspb.Message {
  getId(): string;
  setId(value: string): FieldObject1;

  getSystemname(): string;
  setSystemname(value: string): FieldObject1;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): FieldObject1.AsObject;
  static toObject(includeInstance: boolean, msg: FieldObject1): FieldObject1.AsObject;
  static serializeBinaryToWriter(message: FieldObject1, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): FieldObject1;
  static deserializeBinaryFromReader(message: FieldObject1, reader: jspb.BinaryReader): FieldObject1;
}

export namespace FieldObject1 {
  export type AsObject = {
    id: string,
    systemname: string,
  }
}

