// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/pixel.proto

package org.mojolang.mojo.core;

public final class PixelProto {
  private PixelProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_Pixel_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_Pixel_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\025mojo/core/pixel.proto\022\tmojo.core\"\026\n\005Pi" +
      "xel\022\r\n\005value\030\001 \001(\003BW\n\026org.mojolang.mojo." +
      "coreB\nPixelProtoP\001Z/github.com/mojo-lang" +
      "/core/go/pkg/mojo/core;coreb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_core_Pixel_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_Pixel_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_Pixel_descriptor,
        new java.lang.String[] { "Value", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
