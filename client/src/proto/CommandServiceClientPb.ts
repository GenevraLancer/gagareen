/**
 * @fileoverview gRPC-Web generated client stub for 
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as command_pb from './command_pb';


export class CommandServiceClient {
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
    command_pb.FieldObject1,
    (request: command_pb.AddFieldParams) => {
      return request.serializeBinary();
    },
    command_pb.FieldObject1.deserializeBinary
  );

  addField(
    request: command_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null): Promise<command_pb.FieldObject1>;

  addField(
    request: command_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: command_pb.FieldObject1) => void): grpcWeb.ClientReadableStream<command_pb.FieldObject1>;

  addField(
    request: command_pb.AddFieldParams,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: command_pb.FieldObject1) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/CommandService/AddField',
        request,
        metadata || {},
        this.methodInfoAddField,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/CommandService/AddField',
    request,
    metadata || {},
    this.methodInfoAddField);
  }

}

