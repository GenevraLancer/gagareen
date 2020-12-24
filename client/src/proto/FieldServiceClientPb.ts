/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as field_pb from './field_pb';


export class FieldServiceClient {
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
    field_pb.FieldObject,
    (request: field_pb.AddFieldParams) => {
      return request.serializeBinary();
    },
    field_pb.FieldObject.deserializeBinary
  );

  addField(
    request: field_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null): Promise<field_pb.FieldObject>;

  addField(
    request: field_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: field_pb.FieldObject) => void): grpcWeb.ClientReadableStream<field_pb.FieldObject>;

  addField(
    request: field_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: field_pb.FieldObject) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FieldService/AddField',
        request,
        metadata || {},
        this.methodInfoAddField,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FieldService/AddField',
    request,
    metadata || {},
    this.methodInfoAddField);
  }

  methodInfoDeleteField = new grpcWeb.AbstractClientBase.MethodInfo(
    field_pb.deleteFieldResponse,
    (request: field_pb.DeleteFieldParams) => {
      return request.serializeBinary();
    },
    field_pb.deleteFieldResponse.deserializeBinary
  );

  deleteField(
    request: field_pb.DeleteFieldParams,
    metadata: grpcWeb.Metadata | null): Promise<field_pb.deleteFieldResponse>;

  deleteField(
    request: field_pb.DeleteFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: field_pb.deleteFieldResponse) => void): grpcWeb.ClientReadableStream<field_pb.deleteFieldResponse>;

  deleteField(
    request: field_pb.DeleteFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: field_pb.deleteFieldResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FieldService/DeleteField',
        request,
        metadata || {},
        this.methodInfoDeleteField,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FieldService/DeleteField',
    request,
    metadata || {},
    this.methodInfoDeleteField);
  }

  methodInfoGetField = new grpcWeb.AbstractClientBase.MethodInfo(
    field_pb.FieldResponse,
    (request: field_pb.GetFieldParams) => {
      return request.serializeBinary();
    },
    field_pb.FieldResponse.deserializeBinary
  );

  getField(
    request: field_pb.GetFieldParams,
    metadata: grpcWeb.Metadata | null): Promise<field_pb.FieldResponse>;

  getField(
    request: field_pb.GetFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: field_pb.FieldResponse) => void): grpcWeb.ClientReadableStream<field_pb.FieldResponse>;

  getField(
    request: field_pb.GetFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: field_pb.FieldResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/FieldService/GetField',
        request,
        metadata || {},
        this.methodInfoGetField,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/FieldService/GetField',
    request,
    metadata || {},
    this.methodInfoGetField);
  }

  methodInfoSubscribe = new grpcWeb.AbstractClientBase.MethodInfo(
    field_pb.SubscribeResponse,
    (request: field_pb.SubscribeRequest) => {
      return request.serializeBinary();
    },
    field_pb.SubscribeResponse.deserializeBinary
  );

  subscribe(
    request: field_pb.SubscribeRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/FieldService/Subscribe',
      request,
      metadata || {},
      this.methodInfoSubscribe);
  }

}

