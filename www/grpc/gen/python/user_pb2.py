# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nuser.proto\x12\x06sternx\"E\n\x11\x43reateUserRequest\x12\x14\n\x05\x65mail\x18\x01 \x01(\tR\x05\x65mail\x12\x1a\n\x08password\x18\x02 \x01(\tR\x08password\" \n\x0eGetUserRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"4\n\x0cUserResponse\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n\x05\x65mail\x18\x02 \x01(\tR\x05\x65mail\"M\n\x0fGetUsersRequest\x12\x1d\n\npage_index\x18\x01 \x01(\rR\tpageIndex\x12\x1b\n\tpage_size\x18\x02 \x01(\rR\x08pageSize\">\n\x10GetUsersResponse\x12*\n\x05users\x18\x01 \x03(\x0b\x32\x14.sternx.UserResponseR\x05users\"U\n\x11UpdateUserRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n\x05\x65mail\x18\x02 \x01(\tR\x05\x65mail\x12\x1a\n\x08password\x18\x03 \x01(\tR\x08password\"#\n\x11\x44\x65leteUserRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"$\n\x12\x44\x65leteUserResponse\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id2\xc5\x02\n\x04User\x12=\n\nCreateUser\x12\x19.sternx.CreateUserRequest\x1a\x14.sternx.UserResponse\x12;\n\x0bGetUserByID\x12\x16.sternx.GetUserRequest\x1a\x14.sternx.UserResponse\x12=\n\x08GetUsers\x12\x17.sternx.GetUsersRequest\x1a\x18.sternx.GetUsersResponse\x12=\n\nUpdateUser\x12\x19.sternx.UpdateUserRequest\x1a\x14.sternx.UserResponse\x12\x43\n\nDeleteUser\x12\x19.sternx.DeleteUserRequest\x1a\x1a.sternx.DeleteUserResponseB%\n\x0bsternx.userZ\x16sternx/www/grpc/sternxb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'user_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\013sternx.userZ\026sternx/www/grpc/sternx'
  _CREATEUSERREQUEST._serialized_start=22
  _CREATEUSERREQUEST._serialized_end=91
  _GETUSERREQUEST._serialized_start=93
  _GETUSERREQUEST._serialized_end=125
  _USERRESPONSE._serialized_start=127
  _USERRESPONSE._serialized_end=179
  _GETUSERSREQUEST._serialized_start=181
  _GETUSERSREQUEST._serialized_end=258
  _GETUSERSRESPONSE._serialized_start=260
  _GETUSERSRESPONSE._serialized_end=322
  _UPDATEUSERREQUEST._serialized_start=324
  _UPDATEUSERREQUEST._serialized_end=409
  _DELETEUSERREQUEST._serialized_start=411
  _DELETEUSERREQUEST._serialized_end=446
  _DELETEUSERRESPONSE._serialized_start=448
  _DELETEUSERRESPONSE._serialized_end=484
  _USER._serialized_start=487
  _USER._serialized_end=812
# @@protoc_insertion_point(module_scope)
