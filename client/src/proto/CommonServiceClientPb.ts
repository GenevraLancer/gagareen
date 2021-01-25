/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as common_pb from './common_pb';


export class CommonServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoAddField = new grpcWeb.AbstractClientBase.MethodInfo(
    common_pb.FieldObject,
    (request: common_pb.AddFieldParams) => {
      return request.serializeBinary();
    },
    common_pb.FieldObject.deserializeBinary
  );

  addField(
    request: common_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null): Promise<common_pb.FieldObject>;

  addField(
    request: common_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: common_pb.FieldObject) => void): grpcWeb.ClientReadableStream<common_pb.FieldObject>;

  addField(
    request: common_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: common_pb.FieldObject) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/CommonService/AddField',
        request,
        metadata || {},
        this.methodInfoAddField,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/CommonService/AddField',
    request,
    metadata || {},
    this.methodInfoAddField);
  }

  methodInfoGetField = new grpcWeb.AbstractClientBase.MethodInfo(
    common_pb.FieldObject,
    (request: common_pb.GetFieldParams) => {
      return request.serializeBinary();
    },
    common_pb.FieldObject.deserializeBinary
  );

  getField(
    request: common_pb.GetFieldParams,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/CommonService/GetField',
      request,
      metadata || {},
      this.methodInfoGetField);
  }

}

