// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/email_address.proto

package org.mojolang.mojo.core;

public final class EmailAddressProto {
  private EmailAddressProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_EmailAddress_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_EmailAddress_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\035mojo/core/email_address.proto\022\tmojo.co" +
      "re\032\026mojo/core/domain.proto\"E\n\014EmailAddre" +
      "ss\022\022\n\nlocal_part\030\001 \001(\t\022!\n\006domain\030\002 \001(\0132\021" +
      ".mojo.core.DomainB^\n\026org.mojolang.mojo.c" +
      "oreB\021EmailAddressProtoP\001Z/github.com/moj" +
      "o-lang/core/go/pkg/mojo/core;coreb\006proto" +
      "3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
          org.mojolang.mojo.core.DomainProto.getDescriptor(),
        });
    internal_static_mojo_core_EmailAddress_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_EmailAddress_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_EmailAddress_descriptor,
        new java.lang.String[] { "LocalPart", "Domain", });
    org.mojolang.mojo.core.DomainProto.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
