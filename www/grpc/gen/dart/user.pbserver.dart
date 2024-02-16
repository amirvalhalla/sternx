///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'user.pb.dart' as $0;
import 'user.pbjson.dart';

export 'user.pb.dart';

abstract class UserServiceBase extends $pb.GeneratedService {
  $async.Future<$0.UserResponse> createUser($pb.ServerContext ctx, $0.CreateUserRequest request);
  $async.Future<$0.UserResponse> getUserByID($pb.ServerContext ctx, $0.GetUserRequest request);
  $async.Future<$0.GetUsersResponse> getUsers($pb.ServerContext ctx, $0.GetUsersRequest request);
  $async.Future<$0.UserResponse> updateUser($pb.ServerContext ctx, $0.UpdateUserRequest request);
  $async.Future<$0.DeleteUserResponse> deleteUser($pb.ServerContext ctx, $0.DeleteUserRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'CreateUser': return $0.CreateUserRequest();
      case 'GetUserByID': return $0.GetUserRequest();
      case 'GetUsers': return $0.GetUsersRequest();
      case 'UpdateUser': return $0.UpdateUserRequest();
      case 'DeleteUser': return $0.DeleteUserRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'CreateUser': return this.createUser(ctx, request as $0.CreateUserRequest);
      case 'GetUserByID': return this.getUserByID(ctx, request as $0.GetUserRequest);
      case 'GetUsers': return this.getUsers(ctx, request as $0.GetUsersRequest);
      case 'UpdateUser': return this.updateUser(ctx, request as $0.UpdateUserRequest);
      case 'DeleteUser': return this.deleteUser(ctx, request as $0.DeleteUserRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => UserServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => UserServiceBase$messageJson;
}

