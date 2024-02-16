// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var user_pb = require('./user_pb.js');

function serialize_sternx_CreateUserRequest(arg) {
  if (!(arg instanceof user_pb.CreateUserRequest)) {
    throw new Error('Expected argument of type sternx.CreateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sternx_CreateUserRequest(buffer_arg) {
  return user_pb.CreateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sternx_DeleteUserRequest(arg) {
  if (!(arg instanceof user_pb.DeleteUserRequest)) {
    throw new Error('Expected argument of type sternx.DeleteUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sternx_DeleteUserRequest(buffer_arg) {
  return user_pb.DeleteUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sternx_DeleteUserResponse(arg) {
  if (!(arg instanceof user_pb.DeleteUserResponse)) {
    throw new Error('Expected argument of type sternx.DeleteUserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sternx_DeleteUserResponse(buffer_arg) {
  return user_pb.DeleteUserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sternx_GetUserRequest(arg) {
  if (!(arg instanceof user_pb.GetUserRequest)) {
    throw new Error('Expected argument of type sternx.GetUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sternx_GetUserRequest(buffer_arg) {
  return user_pb.GetUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sternx_GetUsersRequest(arg) {
  if (!(arg instanceof user_pb.GetUsersRequest)) {
    throw new Error('Expected argument of type sternx.GetUsersRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sternx_GetUsersRequest(buffer_arg) {
  return user_pb.GetUsersRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sternx_GetUsersResponse(arg) {
  if (!(arg instanceof user_pb.GetUsersResponse)) {
    throw new Error('Expected argument of type sternx.GetUsersResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sternx_GetUsersResponse(buffer_arg) {
  return user_pb.GetUsersResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sternx_UpdateUserRequest(arg) {
  if (!(arg instanceof user_pb.UpdateUserRequest)) {
    throw new Error('Expected argument of type sternx.UpdateUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sternx_UpdateUserRequest(buffer_arg) {
  return user_pb.UpdateUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_sternx_UserResponse(arg) {
  if (!(arg instanceof user_pb.UserResponse)) {
    throw new Error('Expected argument of type sternx.UserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_sternx_UserResponse(buffer_arg) {
  return user_pb.UserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var UserService = exports.UserService = {
  createUser: {
    path: '/sternx.User/CreateUser',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.CreateUserRequest,
    responseType: user_pb.UserResponse,
    requestSerialize: serialize_sternx_CreateUserRequest,
    requestDeserialize: deserialize_sternx_CreateUserRequest,
    responseSerialize: serialize_sternx_UserResponse,
    responseDeserialize: deserialize_sternx_UserResponse,
  },
  getUserById: {
    path: '/sternx.User/GetUserById',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.GetUserRequest,
    responseType: user_pb.UserResponse,
    requestSerialize: serialize_sternx_GetUserRequest,
    requestDeserialize: deserialize_sternx_GetUserRequest,
    responseSerialize: serialize_sternx_UserResponse,
    responseDeserialize: deserialize_sternx_UserResponse,
  },
  getUsers: {
    path: '/sternx.User/GetUsers',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.GetUsersRequest,
    responseType: user_pb.GetUsersResponse,
    requestSerialize: serialize_sternx_GetUsersRequest,
    requestDeserialize: deserialize_sternx_GetUsersRequest,
    responseSerialize: serialize_sternx_GetUsersResponse,
    responseDeserialize: deserialize_sternx_GetUsersResponse,
  },
  updateUser: {
    path: '/sternx.User/UpdateUser',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.UpdateUserRequest,
    responseType: user_pb.UserResponse,
    requestSerialize: serialize_sternx_UpdateUserRequest,
    requestDeserialize: deserialize_sternx_UpdateUserRequest,
    responseSerialize: serialize_sternx_UserResponse,
    responseDeserialize: deserialize_sternx_UserResponse,
  },
  deleteUser: {
    path: '/sternx.User/DeleteUser',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.DeleteUserRequest,
    responseType: user_pb.DeleteUserResponse,
    requestSerialize: serialize_sternx_DeleteUserRequest,
    requestDeserialize: deserialize_sternx_DeleteUserRequest,
    responseSerialize: serialize_sternx_DeleteUserResponse,
    responseDeserialize: deserialize_sternx_DeleteUserResponse,
  },
};

exports.UserClient = grpc.makeGenericClientConstructor(UserService);
