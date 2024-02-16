///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;
import 'dart:convert' as $convert;
import 'dart:typed_data' as $typed_data;
@$core.Deprecated('Use createUserRequestDescriptor instead')
const CreateUserRequest$json = const {
  '1': 'CreateUserRequest',
  '2': const [
    const {'1': 'email', '3': 1, '4': 1, '5': 9, '10': 'email'},
    const {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `CreateUserRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List createUserRequestDescriptor = $convert.base64Decode('ChFDcmVhdGVVc2VyUmVxdWVzdBIUCgVlbWFpbBgBIAEoCVIFZW1haWwSGgoIcGFzc3dvcmQYAiABKAlSCHBhc3N3b3Jk');
@$core.Deprecated('Use getUserRequestDescriptor instead')
const GetUserRequest$json = const {
  '1': 'GetUserRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `GetUserRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUserRequestDescriptor = $convert.base64Decode('Cg5HZXRVc2VyUmVxdWVzdBIOCgJpZBgBIAEoCVICaWQ=');
@$core.Deprecated('Use userResponseDescriptor instead')
const UserResponse$json = const {
  '1': 'UserResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'email', '3': 2, '4': 1, '5': 9, '10': 'email'},
  ],
};

/// Descriptor for `UserResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List userResponseDescriptor = $convert.base64Decode('CgxVc2VyUmVzcG9uc2USDgoCaWQYASABKAlSAmlkEhQKBWVtYWlsGAIgASgJUgVlbWFpbA==');
@$core.Deprecated('Use getUsersRequestDescriptor instead')
const GetUsersRequest$json = const {
  '1': 'GetUsersRequest',
  '2': const [
    const {'1': 'page_index', '3': 1, '4': 1, '5': 13, '10': 'pageIndex'},
    const {'1': 'page_size', '3': 2, '4': 1, '5': 13, '10': 'pageSize'},
  ],
};

/// Descriptor for `GetUsersRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUsersRequestDescriptor = $convert.base64Decode('Cg9HZXRVc2Vyc1JlcXVlc3QSHQoKcGFnZV9pbmRleBgBIAEoDVIJcGFnZUluZGV4EhsKCXBhZ2Vfc2l6ZRgCIAEoDVIIcGFnZVNpemU=');
@$core.Deprecated('Use getUsersResponseDescriptor instead')
const GetUsersResponse$json = const {
  '1': 'GetUsersResponse',
  '2': const [
    const {'1': 'users', '3': 1, '4': 3, '5': 11, '6': '.sternx.UserResponse', '10': 'users'},
  ],
};

/// Descriptor for `GetUsersResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List getUsersResponseDescriptor = $convert.base64Decode('ChBHZXRVc2Vyc1Jlc3BvbnNlEioKBXVzZXJzGAEgAygLMhQuc3Rlcm54LlVzZXJSZXNwb25zZVIFdXNlcnM=');
@$core.Deprecated('Use updateUserRequestDescriptor instead')
const UpdateUserRequest$json = const {
  '1': 'UpdateUserRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
    const {'1': 'email', '3': 2, '4': 1, '5': 9, '10': 'email'},
    const {'1': 'password', '3': 3, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `UpdateUserRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List updateUserRequestDescriptor = $convert.base64Decode('ChFVcGRhdGVVc2VyUmVxdWVzdBIOCgJpZBgBIAEoCVICaWQSFAoFZW1haWwYAiABKAlSBWVtYWlsEhoKCHBhc3N3b3JkGAMgASgJUghwYXNzd29yZA==');
@$core.Deprecated('Use deleteUserRequestDescriptor instead')
const DeleteUserRequest$json = const {
  '1': 'DeleteUserRequest',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `DeleteUserRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteUserRequestDescriptor = $convert.base64Decode('ChFEZWxldGVVc2VyUmVxdWVzdBIOCgJpZBgBIAEoCVICaWQ=');
@$core.Deprecated('Use deleteUserResponseDescriptor instead')
const DeleteUserResponse$json = const {
  '1': 'DeleteUserResponse',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 9, '10': 'id'},
  ],
};

/// Descriptor for `DeleteUserResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List deleteUserResponseDescriptor = $convert.base64Decode('ChJEZWxldGVVc2VyUmVzcG9uc2USDgoCaWQYASABKAlSAmlk');
const $core.Map<$core.String, $core.dynamic> UserServiceBase$json = const {
  '1': 'User',
  '2': const [
    const {'1': 'CreateUser', '2': '.sternx.CreateUserRequest', '3': '.sternx.UserResponse'},
    const {'1': 'GetUserById', '2': '.sternx.GetUserRequest', '3': '.sternx.UserResponse'},
    const {'1': 'GetUsers', '2': '.sternx.GetUsersRequest', '3': '.sternx.GetUsersResponse'},
    const {'1': 'UpdateUser', '2': '.sternx.UpdateUserRequest', '3': '.sternx.UserResponse'},
    const {'1': 'DeleteUser', '2': '.sternx.DeleteUserRequest', '3': '.sternx.DeleteUserResponse'},
  ],
};

@$core.Deprecated('Use userServiceDescriptor instead')
const $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> UserServiceBase$messageJson = const {
  '.sternx.CreateUserRequest': CreateUserRequest$json,
  '.sternx.UserResponse': UserResponse$json,
  '.sternx.GetUserRequest': GetUserRequest$json,
  '.sternx.GetUsersRequest': GetUsersRequest$json,
  '.sternx.GetUsersResponse': GetUsersResponse$json,
  '.sternx.UpdateUserRequest': UpdateUserRequest$json,
  '.sternx.DeleteUserRequest': DeleteUserRequest$json,
  '.sternx.DeleteUserResponse': DeleteUserResponse$json,
};

/// Descriptor for `User`. Decode as a `google.protobuf.ServiceDescriptorProto`.
final $typed_data.Uint8List userServiceDescriptor = $convert.base64Decode('CgRVc2VyEj0KCkNyZWF0ZVVzZXISGS5zdGVybnguQ3JlYXRlVXNlclJlcXVlc3QaFC5zdGVybnguVXNlclJlc3BvbnNlEjsKC0dldFVzZXJCeUlkEhYuc3Rlcm54LkdldFVzZXJSZXF1ZXN0GhQuc3Rlcm54LlVzZXJSZXNwb25zZRI9CghHZXRVc2VycxIXLnN0ZXJueC5HZXRVc2Vyc1JlcXVlc3QaGC5zdGVybnguR2V0VXNlcnNSZXNwb25zZRI9CgpVcGRhdGVVc2VyEhkuc3Rlcm54LlVwZGF0ZVVzZXJSZXF1ZXN0GhQuc3Rlcm54LlVzZXJSZXNwb25zZRJDCgpEZWxldGVVc2VyEhkuc3Rlcm54LkRlbGV0ZVVzZXJSZXF1ZXN0Ghouc3Rlcm54LkRlbGV0ZVVzZXJSZXNwb25zZQ==');
