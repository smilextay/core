// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/phone_number.proto

package org.mojolang.mojo.core;

public final class PhoneNumberProto {
  private PhoneNumberProto() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  static final com.google.protobuf.Descriptors.Descriptor
    internal_static_mojo_core_PhoneNumber_descriptor;
  static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_mojo_core_PhoneNumber_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\034mojo/core/phone_number.proto\022\tmojo.cor" +
      "e\"\223\004\n\013PhoneNumber\022\024\n\014country_code\030\001 \001(\005\022" +
      "\027\n\017national_number\030\002 \001(\004\022\021\n\textension\030\003 " +
      "\001(\t\022\034\n\024italian_leading_zero\030\004 \001(\010\022\037\n\027num" +
      "ber_of_leading_zeros\030\010 \001(\005\022\021\n\traw_input\030" +
      "\005 \001(\t\022E\n\023country_code_source\030\006 \001(\0162(.moj" +
      "o.core.PhoneNumber.CountryCodeSource\022\'\n\037" +
      "preferred_domestic_carrier_code\030\007 \001(\t\"\377\001" +
      "\n\021CountryCodeSource\022#\n\037COUNTRY_CODE_SOUR" +
      "CE_UNSPECIFIED\020\000\0222\n.COUNTRY_CODE_SOURCE_" +
      "FROM_NUMBER_WITH_PLUS_SIGN\020\001\022,\n(COUNTRY_" +
      "CODE_SOURCE_FROM_NUMBER_WITH_IDD\020\002\0225\n1CO" +
      "UNTRY_CODE_SOURCE_FROM_NUMBER_WITHOUT_PL" +
      "US_SIGN\020\003\022,\n(COUNTRY_CODE_SOURCE_FROM_DE" +
      "FAULT_COUNTRY\020\004B]\n\026org.mojolang.mojo.cor" +
      "eB\020PhoneNumberProtoP\001Z/github.com/mojo-l" +
      "ang/core/go/pkg/mojo/core;coreb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_mojo_core_PhoneNumber_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_mojo_core_PhoneNumber_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_mojo_core_PhoneNumber_descriptor,
        new java.lang.String[] { "CountryCode", "NationalNumber", "Extension", "ItalianLeadingZero", "NumberOfLeadingZeros", "RawInput", "CountryCodeSource", "PreferredDomesticCarrierCode", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
