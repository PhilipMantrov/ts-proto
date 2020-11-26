"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.addGrpcWebMisc = exports.generateGrpcMethodDesc = exports.generateGrpcServiceDesc = exports.generateGrpcClientImpl = void 0;
const types_1 = require("./types");
const ts_poet_1 = require("ts-poet");
const grpc = ts_poet_1.TypeNames.anyType('grpc@@improbable-eng/grpc-web');
const share = ts_poet_1.TypeNames.anyType('share@rxjs/operators');
const take = ts_poet_1.TypeNames.anyType('take@rxjs/operators');
const BrowserHeaders = ts_poet_1.TypeNames.anyType('BrowserHeaders@browser-headers');
const Code = ts_poet_1.TypeNames.anyType('Code@@improbable-eng/grpc-web/dist/typings/Code');
/** Generates a client that uses the `@improbable-web/grpc-web` library. */
function generateGrpcClientImpl(typeMap, fileDesc, serviceDesc, options) {
    // Define the FooServiceImpl class
    let client = ts_poet_1.ClassSpec.create(`${serviceDesc.name}ClientImpl`)
        .addModifiers(ts_poet_1.Modifier.EXPORT)
        .addInterface(serviceDesc.name);
    // Create the constructor(rpc: Rpc)
    client = client.addFunction(ts_poet_1.FunctionSpec.createConstructor().addParameter('rpc', 'Rpc').addStatement('this.rpc = rpc'));
    client = client.addProperty('rpc', 'Rpc', { modifiers: [ts_poet_1.Modifier.PRIVATE, ts_poet_1.Modifier.READONLY] });
    // Create a method for each FooService method
    for (const methodDesc of serviceDesc.method) {
        client = client.addFunction(generateRpcMethod(options, typeMap, serviceDesc, methodDesc));
    }
    return client;
}
exports.generateGrpcClientImpl = generateGrpcClientImpl;
/** Creates the RPC methods that client code actually calls. */
function generateRpcMethod(options, typeMap, serviceDesc, methodDesc) {
    const requestFn = ts_poet_1.FunctionSpec.create(methodDesc.name);
    const inputType = types_1.requestType(typeMap, methodDesc, options);
    const partialInputType = ts_poet_1.TypeNames.parameterizedType(ts_poet_1.TypeNames.anyType('DeepPartial'), inputType);
    return requestFn
        .addParameter('request', partialInputType)
        .addParameter('metadata?', ts_poet_1.TypeNames.anyType('grpc.Metadata'))
        .addStatement(methodDesc.serverStreaming
        ? 'return this.rpc.invoke(%L, %T.fromPartial(request), metadata)'
        : 'return this.rpc.unary(%L, %T.fromPartial(request), metadata)', methodDescName(serviceDesc, methodDesc), inputType)
        .returns(options.returnObservable || methodDesc.serverStreaming
        ? types_1.responseObservable(typeMap, methodDesc, options)
        : types_1.responsePromise(typeMap, methodDesc, options));
}
/** Creates the service descriptor that grpc-web needs at runtime. */
function generateGrpcServiceDesc(fileDesc, serviceDesc) {
    return ts_poet_1.CodeBlock.empty()
        .add('export const %LDesc = ', serviceDesc.name)
        .beginHash()
        .addHashEntry('serviceName', ts_poet_1.CodeBlock.empty().add('%S', `${fileDesc.package}.${serviceDesc.name}`))
        .endHash();
}
exports.generateGrpcServiceDesc = generateGrpcServiceDesc;
/**
 * Creates the method descriptor that grpc-web needs at runtime to make `unary` calls.
 *
 * Note that we take a few liberties in the implementation give we don't 100% match
 * what grpc-web's existing output is, but it works out; see comments in the method
 * implementation.
 */
function generateGrpcMethodDesc(options, typeMap, serviceDesc, methodDesc) {
    let inputType = types_1.requestType(typeMap, methodDesc, options);
    let outputType = types_1.responseType(typeMap, methodDesc, options);
    return (ts_poet_1.CodeBlock.empty()
        .add('export const %L: UnaryMethodDefinitionish = ', methodDescName(serviceDesc, methodDesc))
        .beginHash()
        .addHashEntry('methodName', ts_poet_1.CodeBlock.empty().add('%S', methodDesc.name))
        .addHashEntry('service', `${serviceDesc.name}Desc`)
        .addHashEntry('requestStream', 'false')
        .addHashEntry('responseStream', methodDesc.serverStreaming ? 'true' : 'false')
        // grpc-web expects this to be a class, but the ts-proto messages are just interfaces.
        //
        // That said, grpc-web's runtime doesn't really use this (at least so far for what ts-proto
        // does), so we could potentially set it to `null!`.
        //
        // However, grpc-web does want messages to have a `.serializeBinary()` method, which again
        // due to the class-less nature of ts-proto's messages, we don't have. So we appropriate
        // this `requestType` as a placeholder for our GrpcWebImpl to Object.assign-in this request
        // message's `serializeBinary` method into the data before handing it off to grpc-web.
        //
        // This makes our data look enough like an object/class that grpc-web works just fine.
        .addHashEntry('requestType', ts_poet_1.CodeBlock.empty()
        .beginHash()
        .addHashEntry('serializeBinary', ts_poet_1.FunctionSpec.create('serializeBinary').addStatement('return %T.encode(this).finish()', inputType))
        .endHash()
        .add(' as any'))
        // grpc-web also expects this to be a class, but with a static `deserializeBinary` method to
        // create new instances of messages. We again don't have an actual class constructor/symbol
        // to pass to it, but we can make up a lambda that has a `deserializeBinary` that does what
        // we want/what grpc-web's runtime needs.
        .addHashEntry('responseType', ts_poet_1.CodeBlock.empty()
        .beginHash()
        .addHashEntry('deserializeBinary', ts_poet_1.FunctionSpec.create('deserializeBinary')
        .addParameter('data', 'Uint8Array')
        .addStatement('return { ...%T.decode(data), toObject() { return this; } }', outputType))
        .endHash()
        .add(' as any'))
        .endHash());
}
exports.generateGrpcMethodDesc = generateGrpcMethodDesc;
function methodDescName(serviceDesc, methodDesc) {
    return `${serviceDesc.name}${methodDesc.name}Desc`;
}
/** Adds misc top-level definitions for grpc-web functionality. */
function addGrpcWebMisc(options, _file) {
    let file = _file;
    file = file.addCode(ts_poet_1.CodeBlock.empty()
        .addStatement('import UnaryMethodDefinition = grpc.UnaryMethodDefinition')
        .addStatement('interface UnaryMethodDefinitionishR extends UnaryMethodDefinition<any, any> { requestStream: any; responseStream: any; }')
        .addStatement('type UnaryMethodDefinitionish = UnaryMethodDefinitionishR'));
    file = file.addInterface(generateGrpcWebRpcType(options.returnObservable));
    file = file.addClass(options.returnObservable ? generateGrpcWebImplObservable() : generateGrpcWebImplPromise());
    return file;
}
exports.addGrpcWebMisc = addGrpcWebMisc;
/** Makes an `Rpc` interface to decouple from the low-level grpc-web `grpc.invoke and grpc.unary`/etc. methods. */
function generateGrpcWebRpcType(returnObservable) {
    let rpc = ts_poet_1.InterfaceSpec.create('Rpc');
    let fnU = ts_poet_1.FunctionSpec.create('unary');
    let fnI = ts_poet_1.FunctionSpec.create('invoke');
    const t = ts_poet_1.TypeNames.typeVariable('T', ts_poet_1.TypeNames.bound('UnaryMethodDefinitionish'));
    fnU = fnU
        .addTypeVariable(t)
        .addParameter('methodDesc', t)
        .addParameter('request', ts_poet_1.TypeNames.ANY)
        .addParameter('metadata', ts_poet_1.TypeNames.unionType(ts_poet_1.TypeNames.anyType('grpc.Metadata'), ts_poet_1.TypeNames.UNDEFINED))
        .returns(returnObservable
        ? ts_poet_1.TypeNames.anyType('Observable@rxjs').param(ts_poet_1.TypeNames.ANY)
        : ts_poet_1.TypeNames.PROMISE.param(ts_poet_1.TypeNames.ANY));
    fnI = fnI
        .addTypeVariable(t)
        .addParameter('methodDesc', t)
        .addParameter('request', ts_poet_1.TypeNames.ANY)
        .addParameter('metadata', ts_poet_1.TypeNames.unionType(ts_poet_1.TypeNames.anyType('grpc.Metadata'), ts_poet_1.TypeNames.UNDEFINED))
        .returns(ts_poet_1.TypeNames.anyType('Observable@rxjs').param(ts_poet_1.TypeNames.ANY));
    rpc = rpc.addFunction(fnU).addFunction(fnI);
    return rpc;
}
/** Implements the `Rpc` interface by making calls using the `grpc.unary` method. */
function generateGrpcWebImplPromise() {
    const maybeMetadata = ts_poet_1.TypeNames.unionType(ts_poet_1.TypeNames.anyType('grpc.Metadata'), ts_poet_1.TypeNames.UNDEFINED);
    const optionsParam = ts_poet_1.TypeNames.anonymousType(['transport?', ts_poet_1.TypeNames.anyType('grpc.TransportFactory')], ['debug?', ts_poet_1.TypeNames.BOOLEAN], ['metadata?', maybeMetadata]);
    const t = ts_poet_1.TypeNames.typeVariable('T', ts_poet_1.TypeNames.bound('UnaryMethodDefinitionish'));
    return ts_poet_1.ClassSpec.create('GrpcWebImpl')
        .addModifiers(ts_poet_1.Modifier.EXPORT)
        .addProperty(ts_poet_1.PropertySpec.create('host', ts_poet_1.TypeNames.STRING).addModifiers(ts_poet_1.Modifier.PRIVATE))
        .addProperty(ts_poet_1.PropertySpec.create('options', optionsParam).addModifiers(ts_poet_1.Modifier.PRIVATE))
        .addInterface('Rpc')
        .addFunction(ts_poet_1.FunctionSpec.createConstructor()
        .addParameter('host', 'string')
        .addParameter('options', optionsParam)
        .addStatement('this.host = host')
        .addStatement('this.options = options'))
        .addFunction(ts_poet_1.FunctionSpec.create('unary')
        .addTypeVariable(t)
        .addParameter('methodDesc', t)
        .addParameter('_request', ts_poet_1.TypeNames.ANY)
        .addParameter('metadata', maybeMetadata)
        .returns(ts_poet_1.TypeNames.PROMISE.param(ts_poet_1.TypeNames.ANY))
        .addCodeBlock(ts_poet_1.CodeBlock.empty().add(`const request = { ..._request, ...methodDesc.requestType };
            const maybeCombinedMetadata =
    metadata && this.options.metadata
      ? new %T({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata || this.options.metadata;
return new Promise((resolve, reject) => {
  %T.unary(methodDesc, {
    request,
    host: this.host,
    metadata: maybeCombinedMetadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (response.status === grpc.Code.OK) {
        resolve(response.message);
      } else {
        const err = new Error(response.statusMessage) as any;
        err.code = response.status;
        err.metadata = response.trailers;
        reject(err);
      }
    },
  });
});
`, BrowserHeaders, grpc)))
        .addFunction(ts_poet_1.FunctionSpec.create('invoke')
        .addTypeVariable(t)
        .addParameter('methodDesc', t)
        .addParameter('_request', ts_poet_1.TypeNames.ANY)
        .addParameter('metadata', maybeMetadata)
        .returns(ts_poet_1.TypeNames.anyType('Observable@rxjs').param(ts_poet_1.TypeNames.ANY))
        .addCodeBlock(ts_poet_1.CodeBlock.empty().add(`const upStreamCodes = [2, 4, 8, 9, 10, 13, 14, 15]; /* Status Response Codes (https://developers.google.com/maps-booking/reference/grpc-api/status_codes) */
            const DEFAULT_TIMEOUT_TIME: number = 3 /* seconds */ * 1000 /* ms */;
            const request = { ..._request, ...methodDesc.requestType };
            const maybeCombinedMetadata =
    metadata && this.options.metadata
      ? new %T({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata || this.options.metadata;
return new Observable(observer => {
      const upStream = (() => {
        %T.invoke(methodDesc, {
          host: this.host,
          request,
          metadata: maybeCombinedMetadata,
          debug: this.options.debug,
          onMessage: (next) => {
            observer.next(next as any);
          },
           onEnd: (code: %T) => {
            if (upStreamCodes.find(upStreamCode => code === upStreamCode)) {
              setTimeout(() => {
                upStream();
              }, DEFAULT_TIMEOUT_TIME);
            }
          },
        });
      });

      upStream();
    }).pipe(%T());
`, BrowserHeaders, grpc, Code, share)));
}
function generateGrpcWebImplObservable() {
    const maybeMetadata = ts_poet_1.TypeNames.unionType(ts_poet_1.TypeNames.anyType('grpc.Metadata'), ts_poet_1.TypeNames.UNDEFINED);
    const optionsParam = ts_poet_1.TypeNames.anonymousType(['transport?', ts_poet_1.TypeNames.anyType('grpc.TransportFactory')], ['debug?', ts_poet_1.TypeNames.BOOLEAN], ['metadata?', maybeMetadata]);
    const t = ts_poet_1.TypeNames.typeVariable('T', ts_poet_1.TypeNames.bound('UnaryMethodDefinitionish'));
    return ts_poet_1.ClassSpec.create('GrpcWebImpl')
        .addModifiers(ts_poet_1.Modifier.EXPORT)
        .addProperty(ts_poet_1.PropertySpec.create('host', ts_poet_1.TypeNames.STRING).addModifiers(ts_poet_1.Modifier.PRIVATE))
        .addProperty(ts_poet_1.PropertySpec.create('options', optionsParam).addModifiers(ts_poet_1.Modifier.PRIVATE))
        .addInterface('Rpc')
        .addFunction(ts_poet_1.FunctionSpec.createConstructor()
        .addParameter('host', 'string')
        .addParameter('options', optionsParam)
        .addStatement('this.host = host')
        .addStatement('this.options = options'))
        .addFunction(ts_poet_1.FunctionSpec.create('unary')
        .addTypeVariable(t)
        .addParameter('methodDesc', t)
        .addParameter('_request', ts_poet_1.TypeNames.ANY)
        .addParameter('metadata', maybeMetadata)
        .returns(ts_poet_1.TypeNames.anyType('Observable@rxjs').param(ts_poet_1.TypeNames.ANY))
        .addCodeBlock(ts_poet_1.CodeBlock.empty().add(`const request = { ..._request, ...methodDesc.requestType };
            const maybeCombinedMetadata =
    metadata && this.options.metadata
      ? new %T({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata || this.options.metadata;
return new Observable(observer => {
  %T.unary(methodDesc, {
      request,
      host: this.host,
      metadata: maybeCombinedMetadata,
      transport: this.options.transport,
      debug: this.options.debug,
      onEnd: (next) => {
        if (next.status !== 0) {
          observer.error({
            code: next.status,
            message: next.statusMessage,
          });
        } else {
          observer.next(next.message as any);
          observer.complete();
        }
      },
    });
  }).pipe(%T(1));
`, BrowserHeaders, grpc, take)))
        .addFunction(ts_poet_1.FunctionSpec.create('invoke')
        .addTypeVariable(t)
        .addParameter('methodDesc', t)
        .addParameter('_request', ts_poet_1.TypeNames.ANY)
        .addParameter('metadata', maybeMetadata)
        .returns(ts_poet_1.TypeNames.anyType('Observable@rxjs').param(ts_poet_1.TypeNames.ANY))
        .addCodeBlock(ts_poet_1.CodeBlock.empty().add(`const upStreamCodes = [2, 4, 8, 9, 10, 13, 14, 15]; /* Status Response Codes (https://developers.google.com/maps-booking/reference/grpc-api/status_codes) */
            const DEFAULT_TIMEOUT_TIME: number = 3 /* seconds */ * 1000 /* ms */;
            const request = { ..._request, ...methodDesc.requestType };
            const maybeCombinedMetadata =
    metadata && this.options.metadata
      ? new %T({ ...this.options?.metadata.headersMap, ...metadata?.headersMap })
      : metadata || this.options.metadata;
return new Observable(observer => {
      const upStream = (() => {
        %T.invoke(methodDesc, {
          host: this.host,
          request,
          metadata: maybeCombinedMetadata,
          transport: grpc.WebsocketTransport(),
          debug: this.options.debug,
          onMessage: (next) => {
            observer.next(next as any);
          },
          onEnd: (code: %T) => {
            if (upStreamCodes.find(upStreamCode => code === upStreamCode)) {
              setTimeout(() => {
                upStream();
              }, DEFAULT_TIMEOUT_TIME);
            }
          },
        });
      });

      upStream();
    }).pipe(%T());
`, BrowserHeaders, grpc, Code, share)));
}
