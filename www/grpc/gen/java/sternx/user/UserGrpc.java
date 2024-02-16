package sternx.user;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.50.2)",
    comments = "Source: user.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class UserGrpc {

  private UserGrpc() {}

  public static final String SERVICE_NAME = "sternx.User";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<sternx.user.UserOuterClass.CreateUserRequest,
      sternx.user.UserOuterClass.UserResponse> getCreateUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "CreateUser",
      requestType = sternx.user.UserOuterClass.CreateUserRequest.class,
      responseType = sternx.user.UserOuterClass.UserResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<sternx.user.UserOuterClass.CreateUserRequest,
      sternx.user.UserOuterClass.UserResponse> getCreateUserMethod() {
    io.grpc.MethodDescriptor<sternx.user.UserOuterClass.CreateUserRequest, sternx.user.UserOuterClass.UserResponse> getCreateUserMethod;
    if ((getCreateUserMethod = UserGrpc.getCreateUserMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getCreateUserMethod = UserGrpc.getCreateUserMethod) == null) {
          UserGrpc.getCreateUserMethod = getCreateUserMethod =
              io.grpc.MethodDescriptor.<sternx.user.UserOuterClass.CreateUserRequest, sternx.user.UserOuterClass.UserResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "CreateUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.CreateUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.UserResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("CreateUser"))
              .build();
        }
      }
    }
    return getCreateUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<sternx.user.UserOuterClass.GetUserRequest,
      sternx.user.UserOuterClass.UserResponse> getGetUserByIDMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUserByID",
      requestType = sternx.user.UserOuterClass.GetUserRequest.class,
      responseType = sternx.user.UserOuterClass.UserResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<sternx.user.UserOuterClass.GetUserRequest,
      sternx.user.UserOuterClass.UserResponse> getGetUserByIDMethod() {
    io.grpc.MethodDescriptor<sternx.user.UserOuterClass.GetUserRequest, sternx.user.UserOuterClass.UserResponse> getGetUserByIDMethod;
    if ((getGetUserByIDMethod = UserGrpc.getGetUserByIDMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getGetUserByIDMethod = UserGrpc.getGetUserByIDMethod) == null) {
          UserGrpc.getGetUserByIDMethod = getGetUserByIDMethod =
              io.grpc.MethodDescriptor.<sternx.user.UserOuterClass.GetUserRequest, sternx.user.UserOuterClass.UserResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUserByID"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.GetUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.UserResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("GetUserByID"))
              .build();
        }
      }
    }
    return getGetUserByIDMethod;
  }

  private static volatile io.grpc.MethodDescriptor<sternx.user.UserOuterClass.GetUsersRequest,
      sternx.user.UserOuterClass.GetUsersResponse> getGetUsersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetUsers",
      requestType = sternx.user.UserOuterClass.GetUsersRequest.class,
      responseType = sternx.user.UserOuterClass.GetUsersResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<sternx.user.UserOuterClass.GetUsersRequest,
      sternx.user.UserOuterClass.GetUsersResponse> getGetUsersMethod() {
    io.grpc.MethodDescriptor<sternx.user.UserOuterClass.GetUsersRequest, sternx.user.UserOuterClass.GetUsersResponse> getGetUsersMethod;
    if ((getGetUsersMethod = UserGrpc.getGetUsersMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getGetUsersMethod = UserGrpc.getGetUsersMethod) == null) {
          UserGrpc.getGetUsersMethod = getGetUsersMethod =
              io.grpc.MethodDescriptor.<sternx.user.UserOuterClass.GetUsersRequest, sternx.user.UserOuterClass.GetUsersResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetUsers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.GetUsersRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.GetUsersResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("GetUsers"))
              .build();
        }
      }
    }
    return getGetUsersMethod;
  }

  private static volatile io.grpc.MethodDescriptor<sternx.user.UserOuterClass.UpdateUserRequest,
      sternx.user.UserOuterClass.UserResponse> getUpdateUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "UpdateUser",
      requestType = sternx.user.UserOuterClass.UpdateUserRequest.class,
      responseType = sternx.user.UserOuterClass.UserResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<sternx.user.UserOuterClass.UpdateUserRequest,
      sternx.user.UserOuterClass.UserResponse> getUpdateUserMethod() {
    io.grpc.MethodDescriptor<sternx.user.UserOuterClass.UpdateUserRequest, sternx.user.UserOuterClass.UserResponse> getUpdateUserMethod;
    if ((getUpdateUserMethod = UserGrpc.getUpdateUserMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getUpdateUserMethod = UserGrpc.getUpdateUserMethod) == null) {
          UserGrpc.getUpdateUserMethod = getUpdateUserMethod =
              io.grpc.MethodDescriptor.<sternx.user.UserOuterClass.UpdateUserRequest, sternx.user.UserOuterClass.UserResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "UpdateUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.UpdateUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.UserResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("UpdateUser"))
              .build();
        }
      }
    }
    return getUpdateUserMethod;
  }

  private static volatile io.grpc.MethodDescriptor<sternx.user.UserOuterClass.DeleteUserRequest,
      sternx.user.UserOuterClass.DeleteUserResponse> getDeleteUserMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "DeleteUser",
      requestType = sternx.user.UserOuterClass.DeleteUserRequest.class,
      responseType = sternx.user.UserOuterClass.DeleteUserResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<sternx.user.UserOuterClass.DeleteUserRequest,
      sternx.user.UserOuterClass.DeleteUserResponse> getDeleteUserMethod() {
    io.grpc.MethodDescriptor<sternx.user.UserOuterClass.DeleteUserRequest, sternx.user.UserOuterClass.DeleteUserResponse> getDeleteUserMethod;
    if ((getDeleteUserMethod = UserGrpc.getDeleteUserMethod) == null) {
      synchronized (UserGrpc.class) {
        if ((getDeleteUserMethod = UserGrpc.getDeleteUserMethod) == null) {
          UserGrpc.getDeleteUserMethod = getDeleteUserMethod =
              io.grpc.MethodDescriptor.<sternx.user.UserOuterClass.DeleteUserRequest, sternx.user.UserOuterClass.DeleteUserResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "DeleteUser"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.DeleteUserRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  sternx.user.UserOuterClass.DeleteUserResponse.getDefaultInstance()))
              .setSchemaDescriptor(new UserMethodDescriptorSupplier("DeleteUser"))
              .build();
        }
      }
    }
    return getDeleteUserMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserStub>() {
        @java.lang.Override
        public UserStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserStub(channel, callOptions);
        }
      };
    return UserStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserBlockingStub>() {
        @java.lang.Override
        public UserBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserBlockingStub(channel, callOptions);
        }
      };
    return UserBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UserFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserFutureStub>() {
        @java.lang.Override
        public UserFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserFutureStub(channel, callOptions);
        }
      };
    return UserFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class UserImplBase implements io.grpc.BindableService {

    /**
     */
    public void createUser(sternx.user.UserOuterClass.CreateUserRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getCreateUserMethod(), responseObserver);
    }

    /**
     */
    public void getUserByID(sternx.user.UserOuterClass.GetUserRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUserByIDMethod(), responseObserver);
    }

    /**
     */
    public void getUsers(sternx.user.UserOuterClass.GetUsersRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.GetUsersResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetUsersMethod(), responseObserver);
    }

    /**
     */
    public void updateUser(sternx.user.UserOuterClass.UpdateUserRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getUpdateUserMethod(), responseObserver);
    }

    /**
     */
    public void deleteUser(sternx.user.UserOuterClass.DeleteUserRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.DeleteUserResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getDeleteUserMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getCreateUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                sternx.user.UserOuterClass.CreateUserRequest,
                sternx.user.UserOuterClass.UserResponse>(
                  this, METHODID_CREATE_USER)))
          .addMethod(
            getGetUserByIDMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                sternx.user.UserOuterClass.GetUserRequest,
                sternx.user.UserOuterClass.UserResponse>(
                  this, METHODID_GET_USER_BY_ID)))
          .addMethod(
            getGetUsersMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                sternx.user.UserOuterClass.GetUsersRequest,
                sternx.user.UserOuterClass.GetUsersResponse>(
                  this, METHODID_GET_USERS)))
          .addMethod(
            getUpdateUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                sternx.user.UserOuterClass.UpdateUserRequest,
                sternx.user.UserOuterClass.UserResponse>(
                  this, METHODID_UPDATE_USER)))
          .addMethod(
            getDeleteUserMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                sternx.user.UserOuterClass.DeleteUserRequest,
                sternx.user.UserOuterClass.DeleteUserResponse>(
                  this, METHODID_DELETE_USER)))
          .build();
    }
  }

  /**
   */
  public static final class UserStub extends io.grpc.stub.AbstractAsyncStub<UserStub> {
    private UserStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserStub(channel, callOptions);
    }

    /**
     */
    public void createUser(sternx.user.UserOuterClass.CreateUserRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUserByID(sternx.user.UserOuterClass.GetUserRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUserByIDMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getUsers(sternx.user.UserOuterClass.GetUsersRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.GetUsersResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetUsersMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void updateUser(sternx.user.UserOuterClass.UpdateUserRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getUpdateUserMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void deleteUser(sternx.user.UserOuterClass.DeleteUserRequest request,
        io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.DeleteUserResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getDeleteUserMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class UserBlockingStub extends io.grpc.stub.AbstractBlockingStub<UserBlockingStub> {
    private UserBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserBlockingStub(channel, callOptions);
    }

    /**
     */
    public sternx.user.UserOuterClass.UserResponse createUser(sternx.user.UserOuterClass.CreateUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getCreateUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public sternx.user.UserOuterClass.UserResponse getUserByID(sternx.user.UserOuterClass.GetUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUserByIDMethod(), getCallOptions(), request);
    }

    /**
     */
    public sternx.user.UserOuterClass.GetUsersResponse getUsers(sternx.user.UserOuterClass.GetUsersRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetUsersMethod(), getCallOptions(), request);
    }

    /**
     */
    public sternx.user.UserOuterClass.UserResponse updateUser(sternx.user.UserOuterClass.UpdateUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getUpdateUserMethod(), getCallOptions(), request);
    }

    /**
     */
    public sternx.user.UserOuterClass.DeleteUserResponse deleteUser(sternx.user.UserOuterClass.DeleteUserRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getDeleteUserMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class UserFutureStub extends io.grpc.stub.AbstractFutureStub<UserFutureStub> {
    private UserFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<sternx.user.UserOuterClass.UserResponse> createUser(
        sternx.user.UserOuterClass.CreateUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getCreateUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<sternx.user.UserOuterClass.UserResponse> getUserByID(
        sternx.user.UserOuterClass.GetUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUserByIDMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<sternx.user.UserOuterClass.GetUsersResponse> getUsers(
        sternx.user.UserOuterClass.GetUsersRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetUsersMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<sternx.user.UserOuterClass.UserResponse> updateUser(
        sternx.user.UserOuterClass.UpdateUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getUpdateUserMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<sternx.user.UserOuterClass.DeleteUserResponse> deleteUser(
        sternx.user.UserOuterClass.DeleteUserRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getDeleteUserMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_CREATE_USER = 0;
  private static final int METHODID_GET_USER_BY_ID = 1;
  private static final int METHODID_GET_USERS = 2;
  private static final int METHODID_UPDATE_USER = 3;
  private static final int METHODID_DELETE_USER = 4;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final UserImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(UserImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_CREATE_USER:
          serviceImpl.createUser((sternx.user.UserOuterClass.CreateUserRequest) request,
              (io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse>) responseObserver);
          break;
        case METHODID_GET_USER_BY_ID:
          serviceImpl.getUserByID((sternx.user.UserOuterClass.GetUserRequest) request,
              (io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse>) responseObserver);
          break;
        case METHODID_GET_USERS:
          serviceImpl.getUsers((sternx.user.UserOuterClass.GetUsersRequest) request,
              (io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.GetUsersResponse>) responseObserver);
          break;
        case METHODID_UPDATE_USER:
          serviceImpl.updateUser((sternx.user.UserOuterClass.UpdateUserRequest) request,
              (io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.UserResponse>) responseObserver);
          break;
        case METHODID_DELETE_USER:
          serviceImpl.deleteUser((sternx.user.UserOuterClass.DeleteUserRequest) request,
              (io.grpc.stub.StreamObserver<sternx.user.UserOuterClass.DeleteUserResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class UserBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return sternx.user.UserOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("User");
    }
  }

  private static final class UserFileDescriptorSupplier
      extends UserBaseDescriptorSupplier {
    UserFileDescriptorSupplier() {}
  }

  private static final class UserMethodDescriptorSupplier
      extends UserBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    UserMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (UserGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserFileDescriptorSupplier())
              .addMethod(getCreateUserMethod())
              .addMethod(getGetUserByIDMethod())
              .addMethod(getGetUsersMethod())
              .addMethod(getUpdateUserMethod())
              .addMethod(getDeleteUserMethod())
              .build();
        }
      }
    }
    return result;
  }
}
