// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/target_attribute.proto

package org.mojolang.mojo.core;

/**
 * Protobuf enum {@code mojo.core.DeclType}
 */
public enum DeclType
    implements com.google.protobuf.ProtocolMessageEnum {
  /**
   * <code>DECL_TYPE_UNSPECIFIED = 0;</code>
   */
  DECL_TYPE_UNSPECIFIED(0),
  /**
   * <code>DECL_TYPE_TYPE = 1;</code>
   */
  DECL_TYPE_TYPE(1),
  /**
   * <code>DECL_TYPE_VALUE = 2;</code>
   */
  DECL_TYPE_VALUE(2),
  /**
   * <code>DECL_TYPE_FUNCTION = 3;</code>
   */
  DECL_TYPE_FUNCTION(3),
  /**
   * <code>DECL_TYPE_CONSTRUCTOR = 4;</code>
   */
  DECL_TYPE_CONSTRUCTOR(4),
  /**
   * <code>DECL_TYPE_ATTRIBUTE = 5;</code>
   */
  DECL_TYPE_ATTRIBUTE(5),
  /**
   * <code>DECL_TYPE_PACKAGE = 6;</code>
   */
  DECL_TYPE_PACKAGE(6),
  /**
   * <code>DECL_TYPE_GENERIC = 7;</code>
   */
  DECL_TYPE_GENERIC(7),
  UNRECOGNIZED(-1),
  ;

  /**
   * <code>DECL_TYPE_UNSPECIFIED = 0;</code>
   */
  public static final int DECL_TYPE_UNSPECIFIED_VALUE = 0;
  /**
   * <code>DECL_TYPE_TYPE = 1;</code>
   */
  public static final int DECL_TYPE_TYPE_VALUE = 1;
  /**
   * <code>DECL_TYPE_VALUE = 2;</code>
   */
  public static final int DECL_TYPE_VALUE_VALUE = 2;
  /**
   * <code>DECL_TYPE_FUNCTION = 3;</code>
   */
  public static final int DECL_TYPE_FUNCTION_VALUE = 3;
  /**
   * <code>DECL_TYPE_CONSTRUCTOR = 4;</code>
   */
  public static final int DECL_TYPE_CONSTRUCTOR_VALUE = 4;
  /**
   * <code>DECL_TYPE_ATTRIBUTE = 5;</code>
   */
  public static final int DECL_TYPE_ATTRIBUTE_VALUE = 5;
  /**
   * <code>DECL_TYPE_PACKAGE = 6;</code>
   */
  public static final int DECL_TYPE_PACKAGE_VALUE = 6;
  /**
   * <code>DECL_TYPE_GENERIC = 7;</code>
   */
  public static final int DECL_TYPE_GENERIC_VALUE = 7;


  public final int getNumber() {
    if (this == UNRECOGNIZED) {
      throw new java.lang.IllegalArgumentException(
          "Can't get the number of an unknown enum value.");
    }
    return value;
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   * @deprecated Use {@link #forNumber(int)} instead.
   */
  @java.lang.Deprecated
  public static DeclType valueOf(int value) {
    return forNumber(value);
  }

  /**
   * @param value The numeric wire value of the corresponding enum entry.
   * @return The enum associated with the given numeric wire value.
   */
  public static DeclType forNumber(int value) {
    switch (value) {
      case 0: return DECL_TYPE_UNSPECIFIED;
      case 1: return DECL_TYPE_TYPE;
      case 2: return DECL_TYPE_VALUE;
      case 3: return DECL_TYPE_FUNCTION;
      case 4: return DECL_TYPE_CONSTRUCTOR;
      case 5: return DECL_TYPE_ATTRIBUTE;
      case 6: return DECL_TYPE_PACKAGE;
      case 7: return DECL_TYPE_GENERIC;
      default: return null;
    }
  }

  public static com.google.protobuf.Internal.EnumLiteMap<DeclType>
      internalGetValueMap() {
    return internalValueMap;
  }
  private static final com.google.protobuf.Internal.EnumLiteMap<
      DeclType> internalValueMap =
        new com.google.protobuf.Internal.EnumLiteMap<DeclType>() {
          public DeclType findValueByNumber(int number) {
            return DeclType.forNumber(number);
          }
        };

  public final com.google.protobuf.Descriptors.EnumValueDescriptor
      getValueDescriptor() {
    if (this == UNRECOGNIZED) {
      throw new java.lang.IllegalStateException(
          "Can't get the descriptor of an unrecognized enum value.");
    }
    return getDescriptor().getValues().get(ordinal());
  }
  public final com.google.protobuf.Descriptors.EnumDescriptor
      getDescriptorForType() {
    return getDescriptor();
  }
  public static final com.google.protobuf.Descriptors.EnumDescriptor
      getDescriptor() {
    return org.mojolang.mojo.core.TargetAttributeProto.getDescriptor().getEnumTypes().get(0);
  }

  private static final DeclType[] VALUES = values();

  public static DeclType valueOf(
      com.google.protobuf.Descriptors.EnumValueDescriptor desc) {
    if (desc.getType() != getDescriptor()) {
      throw new java.lang.IllegalArgumentException(
        "EnumValueDescriptor is not for this type.");
    }
    if (desc.getIndex() == -1) {
      return UNRECOGNIZED;
    }
    return VALUES[desc.getIndex()];
  }

  private final int value;

  private DeclType(int value) {
    this.value = value;
  }

  // @@protoc_insertion_point(enum_scope:mojo.core.DeclType)
}

