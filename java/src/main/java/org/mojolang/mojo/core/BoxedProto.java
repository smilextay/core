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
    internal_static_mojo_core_BoxedBool_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_BoxedBool_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_BoxedInt_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_BoxedInt_fieldAccessorTable;
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_BoxedString_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_BoxedString_fieldAccessorTable;
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
      "\n\025mojo/core/boxed.proto\022\tmojo.core\"\032\n\tBo" +
      "xedBool\022\r\n\005value\030\001 \001(\010\"\031\n\010BoxedInt\022\r\n\005va" +
      "lue\030\001 \001(\003\"\034\n\013BoxedString\022\r\n\005value\030\001 \001(\t\"" +
      "\031\n\007Strings\022\016\n\006values\030\001 \003(\t\"\032\n\010Integers\022\016" +
      "\n\006values\030\001 \003(\003\"\031\n\007Doubles\022\016\n\006values\030\001 \003(" +
      "\001BW\n\026org.mojolang.mojo.coreB\nBoxedProtoP" +
      "\001Z/github.com/mojo-lang/core/go/pkg/mojo" +
      "/core;coreb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_core_BoxedBool_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_BoxedBool_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_BoxedBool_descriptor,
        new java.lang.String[] { "Value", });
    internal_static_mojo_core_BoxedInt_descriptor =
      getDescriptor().getMessageTypes().get(1);
    internal_static_mojo_core_BoxedInt_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_BoxedInt_descriptor,
        new java.lang.String[] { "Value", });
    internal_static_mojo_core_BoxedString_descriptor =
      getDescriptor().getMessageTypes().get(2);
    internal_static_mojo_core_BoxedString_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_BoxedString_descriptor,
        new java.lang.String[] { "Value", });
    internal_static_mojo_core_Strings_descriptor =
      getDescriptor().getMessageTypes().get(3);
    internal_static_mojo_core_Strings_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Strings_descriptor,
        new java.lang.String[] { "Values", });
    internal_static_mojo_core_Integers_descriptor =
      getDescriptor().getMessageTypes().get(4);
    internal_static_mojo_core_Integers_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Integers_descriptor,
        new java.lang.String[] { "Values", });
    internal_static_mojo_core_Doubles_descriptor =
      getDescriptor().getMessageTypes().get(5);
    internal_static_mojo_core_Doubles_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Doubles_descriptor,
        new java.lang.String[] { "Values", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
