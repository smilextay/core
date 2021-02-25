// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/email_address.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.EmailAddress}
 */
public final class EmailAddress extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.EmailAddress)
    EmailAddressOrBuilder {
private static final long serialVersionUID = 0L;
  // Use EmailAddress.newBuilder() to construct.
  private EmailAddress(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private EmailAddress() {
    localPart_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new EmailAddress();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private EmailAddress(
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
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();

            localPart_ = s;
            break;
          }
          case 18: {
            org.mojolang.mojo.core.Domain.Builder subBuilder = null;
            if (domain_ != null) {
              subBuilder = domain_.toBuilder();
            }
            domain_ = input.readMessage(org.mojolang.mojo.core.Domain.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(domain_);
              domain_ = subBuilder.buildPartial();
            }

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
    return org.mojolang.mojo.core.EmailAddressProto.internal_static_mojo_core_EmailAddress_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.EmailAddressProto.internal_static_mojo_core_EmailAddress_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.EmailAddress.class, org.mojolang.mojo.core.EmailAddress.Builder.class);
  }

  public static final int LOCAL_PART_FIELD_NUMBER = 1;
  private volatile java.lang.Object localPart_;
  /**
   * <code>string local_part = 1;</code>
   * @return The localPart.
   */
  @java.lang.Override
  public java.lang.String getLocalPart() {
    java.lang.Object ref = localPart_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      localPart_ = s;
      return s;
    }
  }
  /**
   * <code>string local_part = 1;</code>
   * @return The bytes for localPart.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getLocalPartBytes() {
    java.lang.Object ref = localPart_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      localPart_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int DOMAIN_FIELD_NUMBER = 2;
  private org.mojolang.mojo.core.Domain domain_;
  /**
   * <code>.mojo.core.Domain domain = 2;</code>
   * @return Whether the domain field is set.
   */
  @java.lang.Override
  public boolean hasDomain() {
    return domain_ != null;
  }
  /**
   * <code>.mojo.core.Domain domain = 2;</code>
   * @return The domain.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Domain getDomain() {
    return domain_ == null ? org.mojolang.mojo.core.Domain.getDefaultInstance() : domain_;
  }
  /**
   * <code>.mojo.core.Domain domain = 2;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.DomainOrBuilder getDomainOrBuilder() {
    return getDomain();
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
    if (!getLocalPartBytes().isEmpty()) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, localPart_);
    }
    if (domain_ != null) {
      output.writeMessage(2, getDomain());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!getLocalPartBytes().isEmpty()) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, localPart_);
    }
    if (domain_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getDomain());
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
    if (!(obj instanceof org.mojolang.mojo.core.EmailAddress)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.EmailAddress other = (org.mojolang.mojo.core.EmailAddress) obj;

    if (!getLocalPart()
        .equals(other.getLocalPart())) return false;
    if (hasDomain() != other.hasDomain()) return false;
    if (hasDomain()) {
      if (!getDomain()
          .equals(other.getDomain())) return false;
    }
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
    hash = (37 * hash) + LOCAL_PART_FIELD_NUMBER;
    hash = (53 * hash) + getLocalPart().hashCode();
    if (hasDomain()) {
      hash = (37 * hash) + DOMAIN_FIELD_NUMBER;
      hash = (53 * hash) + getDomain().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.EmailAddress parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.EmailAddress parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.EmailAddress parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.EmailAddress parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.EmailAddress prototype) {
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
   * Protobuf type {@code mojo.core.EmailAddress}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.EmailAddress)
      org.mojolang.mojo.core.EmailAddressOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.EmailAddressProto.internal_static_mojo_core_EmailAddress_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.EmailAddressProto.internal_static_mojo_core_EmailAddress_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.EmailAddress.class, org.mojolang.mojo.core.EmailAddress.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.EmailAddress.newBuilder()
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
      localPart_ = "";

      if (domainBuilder_ == null) {
        domain_ = null;
      } else {
        domain_ = null;
        domainBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.EmailAddressProto.internal_static_mojo_core_EmailAddress_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.EmailAddress getDefaultInstanceForType() {
      return org.mojolang.mojo.core.EmailAddress.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.EmailAddress build() {
      org.mojolang.mojo.core.EmailAddress result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.EmailAddress buildPartial() {
      org.mojolang.mojo.core.EmailAddress result = new org.mojolang.mojo.core.EmailAddress(this);
      result.localPart_ = localPart_;
      if (domainBuilder_ == null) {
        result.domain_ = domain_;
      } else {
        result.domain_ = domainBuilder_.build();
      }
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
      if (other instanceof org.mojolang.mojo.core.EmailAddress) {
        return mergeFrom((org.mojolang.mojo.core.EmailAddress)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.EmailAddress other) {
      if (other == org.mojolang.mojo.core.EmailAddress.getDefaultInstance()) return this;
      if (!other.getLocalPart().isEmpty()) {
        localPart_ = other.localPart_;
        onChanged();
      }
      if (other.hasDomain()) {
        mergeDomain(other.getDomain());
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
      org.mojolang.mojo.core.EmailAddress parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.EmailAddress) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private java.lang.Object localPart_ = "";
    /**
     * <code>string local_part = 1;</code>
     * @return The localPart.
     */
    public java.lang.String getLocalPart() {
      java.lang.Object ref = localPart_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        localPart_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string local_part = 1;</code>
     * @return The bytes for localPart.
     */
    public com.google.protobuf.ByteString
        getLocalPartBytes() {
      java.lang.Object ref = localPart_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        localPart_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string local_part = 1;</code>
     * @param value The localPart to set.
     * @return This builder for chaining.
     */
    public Builder setLocalPart(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      localPart_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string local_part = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearLocalPart() {
      
      localPart_ = getDefaultInstance().getLocalPart();
      onChanged();
      return this;
    }
    /**
     * <code>string local_part = 1;</code>
     * @param value The bytes for localPart to set.
     * @return This builder for chaining.
     */
    public Builder setLocalPartBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      localPart_ = value;
      onChanged();
      return this;
    }

    private org.mojolang.mojo.core.Domain domain_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Domain, org.mojolang.mojo.core.Domain.Builder, org.mojolang.mojo.core.DomainOrBuilder> domainBuilder_;
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     * @return Whether the domain field is set.
     */
    public boolean hasDomain() {
      return domainBuilder_ != null || domain_ != null;
    }
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     * @return The domain.
     */
    public org.mojolang.mojo.core.Domain getDomain() {
      if (domainBuilder_ == null) {
        return domain_ == null ? org.mojolang.mojo.core.Domain.getDefaultInstance() : domain_;
      } else {
        return domainBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     */
    public Builder setDomain(org.mojolang.mojo.core.Domain value) {
      if (domainBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        domain_ = value;
        onChanged();
      } else {
        domainBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     */
    public Builder setDomain(
        org.mojolang.mojo.core.Domain.Builder builderForValue) {
      if (domainBuilder_ == null) {
        domain_ = builderForValue.build();
        onChanged();
      } else {
        domainBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     */
    public Builder mergeDomain(org.mojolang.mojo.core.Domain value) {
      if (domainBuilder_ == null) {
        if (domain_ != null) {
          domain_ =
            org.mojolang.mojo.core.Domain.newBuilder(domain_).mergeFrom(value).buildPartial();
        } else {
          domain_ = value;
        }
        onChanged();
      } else {
        domainBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     */
    public Builder clearDomain() {
      if (domainBuilder_ == null) {
        domain_ = null;
        onChanged();
      } else {
        domain_ = null;
        domainBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     */
    public org.mojolang.mojo.core.Domain.Builder getDomainBuilder() {
      
      onChanged();
      return getDomainFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     */
    public org.mojolang.mojo.core.DomainOrBuilder getDomainOrBuilder() {
      if (domainBuilder_ != null) {
        return domainBuilder_.getMessageOrBuilder();
      } else {
        return domain_ == null ?
            org.mojolang.mojo.core.Domain.getDefaultInstance() : domain_;
      }
    }
    /**
     * <code>.mojo.core.Domain domain = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Domain, org.mojolang.mojo.core.Domain.Builder, org.mojolang.mojo.core.DomainOrBuilder> 
        getDomainFieldBuilder() {
      if (domainBuilder_ == null) {
        domainBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.Domain, org.mojolang.mojo.core.Domain.Builder, org.mojolang.mojo.core.DomainOrBuilder>(
                getDomain(),
                getParentForChildren(),
                isClean());
        domain_ = null;
      }
      return domainBuilder_;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.EmailAddress)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.EmailAddress)
  private static final org.mojolang.mojo.core.EmailAddress DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.EmailAddress();
  }

  public static org.mojolang.mojo.core.EmailAddress getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<EmailAddress>
      PARSER = new com.google.protobuf.AbstractParser<EmailAddress>() {
    @java.lang.Override
    public EmailAddress parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new EmailAddress(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<EmailAddress> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<EmailAddress> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.EmailAddress getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

