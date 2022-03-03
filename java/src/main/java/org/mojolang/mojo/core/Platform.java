// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/platform.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.Platform}
 */
public final class Platform extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Platform)
    PlatformOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Platform.newBuilder() to construct.
  private Platform(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Platform() {
    arch_ = 0;
    os_ = 0;
    osVersion_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Platform();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Platform(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 8: {
            int rawValue = input.readEnum();

            arch_ = rawValue;
            break;
          }
          case 16: {
            int rawValue = input.readEnum();

            os_ = rawValue;
            break;
          }
          case 26: {
            java.lang.String s = input.readStringRequireUtf8();

            osVersion_ = s;
            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return org.mojolang.mojo.core.PlatformProto.internal_static_mojo_core_Platform_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.PlatformProto.internal_static_mojo_core_Platform_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Platform.class, org.mojolang.mojo.core.Platform.Builder.class);
  }

  public static final int ARCH_FIELD_NUMBER = 1;
  private int arch_;
  /**
   * <code>.mojo.core.Arch arch = 1;</code>
   * @return The enum numeric value on the wire for arch.
   */
  @java.lang.Override public int getArchValue() {
    return arch_;
  }
  /**
   * <code>.mojo.core.Arch arch = 1;</code>
   * @return The arch.
   */
  @java.lang.Override public org.mojolang.mojo.core.Arch getArch() {
    @SuppressWarnings("deprecation")
    org.mojolang.mojo.core.Arch result = org.mojolang.mojo.core.Arch.valueOf(arch_);
    return result == null ? org.mojolang.mojo.core.Arch.UNRECOGNIZED : result;
  }

  public static final int OS_FIELD_NUMBER = 2;
  private int os_;
  /**
   * <code>.mojo.core.OS os = 2;</code>
   * @return The enum numeric value on the wire for os.
   */
  @java.lang.Override public int getOsValue() {
    return os_;
  }
  /**
   * <code>.mojo.core.OS os = 2;</code>
   * @return The os.
   */
  @java.lang.Override public org.mojolang.mojo.core.OS getOs() {
    @SuppressWarnings("deprecation")
    org.mojolang.mojo.core.OS result = org.mojolang.mojo.core.OS.valueOf(os_);
    return result == null ? org.mojolang.mojo.core.OS.UNRECOGNIZED : result;
  }

  public static final int OS_VERSION_FIELD_NUMBER = 3;
  private volatile java.lang.Object osVersion_;
  /**
   * <code>string os_version = 3;</code>
   * @return The osVersion.
   */
  @java.lang.Override
  public java.lang.String getOsVersion() {
    java.lang.Object ref = osVersion_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      osVersion_ = s;
      return s;
    }
  }
  /**
   * <code>string os_version = 3;</code>
   * @return The bytes for osVersion.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getOsVersionBytes() {
    java.lang.Object ref = osVersion_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      osVersion_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (arch_ != org.mojolang.mojo.core.Arch.ARCH_UNSPECIFIED.getNumber()) {
      output.writeEnum(1, arch_);
    }
    if (os_ != org.mojolang.mojo.core.OS.OS_UNSPECIFIED.getNumber()) {
      output.writeEnum(2, os_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(osVersion_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, osVersion_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (arch_ != org.mojolang.mojo.core.Arch.ARCH_UNSPECIFIED.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(1, arch_);
    }
    if (os_ != org.mojolang.mojo.core.OS.OS_UNSPECIFIED.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(2, os_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(osVersion_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, osVersion_);
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof org.mojolang.mojo.core.Platform)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Platform other = (org.mojolang.mojo.core.Platform) obj;

    if (arch_ != other.arch_) return false;
    if (os_ != other.os_) return false;
    if (!getOsVersion()
        .equals(other.getOsVersion())) return false;
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + ARCH_FIELD_NUMBER;
    hash = (53 * hash) + arch_;
    hash = (37 * hash) + OS_FIELD_NUMBER;
    hash = (53 * hash) + os_;
    hash = (37 * hash) + OS_VERSION_FIELD_NUMBER;
    hash = (53 * hash) + getOsVersion().hashCode();
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Platform parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Platform parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Platform parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Platform parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(org.mojolang.mojo.core.Platform prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code mojo.core.Platform}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Platform)
      org.mojolang.mojo.core.PlatformOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.PlatformProto.internal_static_mojo_core_Platform_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.PlatformProto.internal_static_mojo_core_Platform_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Platform.class, org.mojolang.mojo.core.Platform.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Platform.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      arch_ = 0;

      os_ = 0;

      osVersion_ = "";

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.PlatformProto.internal_static_mojo_core_Platform_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Platform getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Platform.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Platform build() {
      org.mojolang.mojo.core.Platform result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Platform buildPartial() {
      org.mojolang.mojo.core.Platform result = new org.mojolang.mojo.core.Platform(this);
      result.arch_ = arch_;
      result.os_ = os_;
      result.osVersion_ = osVersion_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof org.mojolang.mojo.core.Platform) {
        return mergeFrom((org.mojolang.mojo.core.Platform)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Platform other) {
      if (other == org.mojolang.mojo.core.Platform.getDefaultInstance()) return this;
      if (other.arch_ != 0) {
        setArchValue(other.getArchValue());
      }
      if (other.os_ != 0) {
        setOsValue(other.getOsValue());
      }
      if (!other.getOsVersion().isEmpty()) {
        osVersion_ = other.osVersion_;
        onChanged();
      }
      this.mergeUnknownFields(other.unknownFields);
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      org.mojolang.mojo.core.Platform parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.Platform) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private int arch_ = 0;
    /**
     * <code>.mojo.core.Arch arch = 1;</code>
     * @return The enum numeric value on the wire for arch.
     */
    @java.lang.Override public int getArchValue() {
      return arch_;
    }
    /**
     * <code>.mojo.core.Arch arch = 1;</code>
     * @param value The enum numeric value on the wire for arch to set.
     * @return This builder for chaining.
     */
    public Builder setArchValue(int value) {
      
      arch_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.core.Arch arch = 1;</code>
     * @return The arch.
     */
    @java.lang.Override
    public org.mojolang.mojo.core.Arch getArch() {
      @SuppressWarnings("deprecation")
      org.mojolang.mojo.core.Arch result = org.mojolang.mojo.core.Arch.valueOf(arch_);
      return result == null ? org.mojolang.mojo.core.Arch.UNRECOGNIZED : result;
    }
    /**
     * <code>.mojo.core.Arch arch = 1;</code>
     * @param value The arch to set.
     * @return This builder for chaining.
     */
    public Builder setArch(org.mojolang.mojo.core.Arch value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      arch_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.core.Arch arch = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearArch() {
      
      arch_ = 0;
      onChanged();
      return this;
    }

    private int os_ = 0;
    /**
     * <code>.mojo.core.OS os = 2;</code>
     * @return The enum numeric value on the wire for os.
     */
    @java.lang.Override public int getOsValue() {
      return os_;
    }
    /**
     * <code>.mojo.core.OS os = 2;</code>
     * @param value The enum numeric value on the wire for os to set.
     * @return This builder for chaining.
     */
    public Builder setOsValue(int value) {
      
      os_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.core.OS os = 2;</code>
     * @return The os.
     */
    @java.lang.Override
    public org.mojolang.mojo.core.OS getOs() {
      @SuppressWarnings("deprecation")
      org.mojolang.mojo.core.OS result = org.mojolang.mojo.core.OS.valueOf(os_);
      return result == null ? org.mojolang.mojo.core.OS.UNRECOGNIZED : result;
    }
    /**
     * <code>.mojo.core.OS os = 2;</code>
     * @param value The os to set.
     * @return This builder for chaining.
     */
    public Builder setOs(org.mojolang.mojo.core.OS value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      os_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <code>.mojo.core.OS os = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearOs() {
      
      os_ = 0;
      onChanged();
      return this;
    }

    private java.lang.Object osVersion_ = "";
    /**
     * <code>string os_version = 3;</code>
     * @return The osVersion.
     */
    public java.lang.String getOsVersion() {
      java.lang.Object ref = osVersion_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        osVersion_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string os_version = 3;</code>
     * @return The bytes for osVersion.
     */
    public com.google.protobuf.ByteString
        getOsVersionBytes() {
      java.lang.Object ref = osVersion_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        osVersion_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string os_version = 3;</code>
     * @param value The osVersion to set.
     * @return This builder for chaining.
     */
    public Builder setOsVersion(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      osVersion_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string os_version = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearOsVersion() {
      
      osVersion_ = getDefaultInstance().getOsVersion();
      onChanged();
      return this;
    }
    /**
     * <code>string os_version = 3;</code>
     * @param value The bytes for osVersion to set.
     * @return This builder for chaining.
     */
    public Builder setOsVersionBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      osVersion_ = value;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:mojo.core.Platform)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Platform)
  private static final org.mojolang.mojo.core.Platform DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Platform();
  }

  public static org.mojolang.mojo.core.Platform getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Platform>
      PARSER = new com.google.protobuf.AbstractParser<Platform>() {
    @java.lang.Override
    public Platform parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Platform(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Platform> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Platform> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Platform getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

