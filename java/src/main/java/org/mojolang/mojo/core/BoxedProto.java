// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/boxed.proto

package org.mojolang.mojo.core;

public final class BoxedProto {
  private BoxedProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Strings_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Strings_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Integers_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Integers_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Doubles_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Doubles_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\025mojo/core/boxed.proto\022\tmojo.core\"\031\n\007St" +
      "rings\022\016\n\006values\030\001 \003(\t\"\032\n\010Integers\022\016\n\006val" +
      "ues\030\001 \003(\003\"\031\n\007Doubles\022\016\n\006values\030\001 \003(\001BW\n\026" +
      "org.mojolang.mojo.coreB\nBoxedProtoP\001Z/gi" +
      "thub.com/mojo-lang/core/go/pkg/mojo/core" +
      ";coreb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_core_Strings_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_Strings_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Strings_descriptor,
        new java.lang.String[] { "Values", });
    internal_static_mojo_core_Integers_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_mojo_core_Integers_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Integers_descriptor,
        new java.lang.String[] { "Values", });
    internal_static_mojo_core_Doubles_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_mojo_core_Doubles_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Doubles_descriptor,
        new java.lang.String[] { "Values", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
