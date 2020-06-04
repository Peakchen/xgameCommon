/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.MSG_Rpc.CS_Rpc_Req', null, global);
goog.exportSymbol('proto.MSG_Rpc.ErrorCode', null, global);
goog.exportSymbol('proto.MSG_Rpc.SUBMSG', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.MSG_Rpc.CS_Rpc_Req = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.MSG_Rpc.CS_Rpc_Req, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.MSG_Rpc.CS_Rpc_Req.displayName = 'proto.MSG_Rpc.CS_Rpc_Req';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.MSG_Rpc.CS_Rpc_Req.prototype.toObject = function(opt_includeInstance) {
  return proto.MSG_Rpc.CS_Rpc_Req.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.MSG_Rpc.CS_Rpc_Req} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.MSG_Rpc.CS_Rpc_Req.toObject = function(includeInstance, msg) {
  var f, obj = {
    rpcmodule: jspb.Message.getFieldWithDefault(msg, 1, ""),
    rpcfunc: jspb.Message.getFieldWithDefault(msg, 2, ""),
    data: msg.getData_asB64()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.MSG_Rpc.CS_Rpc_Req}
 */
proto.MSG_Rpc.CS_Rpc_Req.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.MSG_Rpc.CS_Rpc_Req;
  return proto.MSG_Rpc.CS_Rpc_Req.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.MSG_Rpc.CS_Rpc_Req} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.MSG_Rpc.CS_Rpc_Req}
 */
proto.MSG_Rpc.CS_Rpc_Req.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRpcmodule(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setRpcfunc(value);
      break;
    case 3:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setData(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.MSG_Rpc.CS_Rpc_Req.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.MSG_Rpc.CS_Rpc_Req.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.MSG_Rpc.CS_Rpc_Req} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.MSG_Rpc.CS_Rpc_Req.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRpcmodule();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRpcfunc();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getData_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      3,
      f
    );
  }
};


/**
 * optional string rpcmodule = 1;
 * @return {string}
 */
proto.MSG_Rpc.CS_Rpc_Req.prototype.getRpcmodule = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.MSG_Rpc.CS_Rpc_Req.prototype.setRpcmodule = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string rpcfunc = 2;
 * @return {string}
 */
proto.MSG_Rpc.CS_Rpc_Req.prototype.getRpcfunc = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.MSG_Rpc.CS_Rpc_Req.prototype.setRpcfunc = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bytes data = 3;
 * @return {!(string|Uint8Array)}
 */
proto.MSG_Rpc.CS_Rpc_Req.prototype.getData = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * optional bytes data = 3;
 * This is a type-conversion wrapper around `getData()`
 * @return {string}
 */
proto.MSG_Rpc.CS_Rpc_Req.prototype.getData_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getData()));
};


/**
 * optional bytes data = 3;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getData()`
 * @return {!Uint8Array}
 */
proto.MSG_Rpc.CS_Rpc_Req.prototype.getData_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getData()));
};


/** @param {!(string|Uint8Array)} value */
proto.MSG_Rpc.CS_Rpc_Req.prototype.setData = function(value) {
  jspb.Message.setProto3BytesField(this, 3, value);
};


/**
 * @enum {number}
 */
proto.MSG_Rpc.SUBMSG = {
  BEGIN: 0,
  CS_RPC: 1,
  SC_RPC: 2
};

/**
 * @enum {number}
 */
proto.MSG_Rpc.ErrorCode = {
  INVALID: 0,
  SUCCESS: 1,
  FAIL: 2
};

goog.object.extend(exports, proto.MSG_Rpc);
