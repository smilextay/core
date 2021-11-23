// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/error_details.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.RetryInfo}
 */
public final class RetryInfo extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.RetryInfo)
    RetryInfoOrBuilder {
private static final long serialVersionUID = 0L;
  // Use RetryInfo.newBuilder() to construct.
  private RetryInfo(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private RetryInfo() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new RetryInfo();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private RetryInfo(
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
            org.mojolang.mojo.core.Duration.Builder subBuilder = null;
            if (retryDelay_ != null) {
              subBuilder = retryDelay_.toBuilder();
            }
            retryDelay_ = input.readMessage(org.mojolang.mojo.core.Duration.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(retryDelay_);
              retryDelay_ = subBuilder.buildPartial();
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
    return org.mojolang.mojo.core.ErrorDetailsProto.internal_static_mojo_core_RetryInfo_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.ErrorDetailsProto.internal_static_mojo_core_RetryInfo_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.RetryInfo.class, org.mojolang.mojo.core.RetryInfo.Builder.class);
  }

  public static final int RETRY_DELAY_FIELD_NUMBER = 1;
  private org.mojolang.mojo.core.Duration retryDelay_;
  /**
   * <code>.mojo.core.Duration retry_delay = 1;</code>
   * @return Whether the retryDelay field is set.
   */
  @java.lang.Override
  public boolean hasRetryDelay() {
    return retryDelay_ != null;
  }
  /**
   * <code>.mojo.core.Duration retry_delay = 1;</code>
   * @return The retryDelay.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Duration getRetryDelay() {
    return retryDelay_ == null ? org.mojolang.mojo.core.Duration.getDefaultInstance() : retryDelay_;
  }
  /**
   * <code>.mojo.core.Duration retry_delay = 1;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.DurationOrBuilder getRetryDelayOrBuilder() {
    return getRetryDelay();
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
    if (retryDelay_ != null) {
      output.writeMessage(1, getRetryDelay());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (retryDelay_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getRetryDelay());
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
    if (!(obj instanceof org.mojolang.mojo.core.RetryInfo)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.RetryInfo other = (org.mojolang.mojo.core.RetryInfo) obj;

    if (hasRetryDelay() != other.hasRetryDelay()) return false;
    if (hasRetryDelay()) {
      if (!getRetryDelay()
          .equals(other.getRetryDelay())) return false;
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
    if (hasRetryDelay()) {
      hash = (37 * hash) + RETRY_DELAY_FIELD_NUMBER;
      hash = (53 * hash) + getRetryDelay().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.RetryInfo parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.RetryInfo parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.RetryInfo parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.RetryInfo parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.RetryInfo prototype) {
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
   * Protobuf type {@code mojo.core.RetryInfo}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.RetryInfo)
      org.mojolang.mojo.core.RetryInfoOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.ErrorDetailsProto.internal_static_mojo_core_RetryInfo_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.ErrorDetailsProto.internal_static_mojo_core_RetryInfo_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.RetryInfo.class, org.mojolang.mojo.core.RetryInfo.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.RetryInfo.newBuilder()
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
      if (retryDelayBuilder_ == null) {
        retryDelay_ = null;
      } else {
        retryDelay_ = null;
        retryDelayBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.ErrorDetailsProto.internal_static_mojo_core_RetryInfo_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.RetryInfo getDefaultInstanceForType() {
      return org.mojolang.mojo.core.RetryInfo.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.RetryInfo build() {
      org.mojolang.mojo.core.RetryInfo result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.RetryInfo buildPartial() {
      org.mojolang.mojo.core.RetryInfo result = new org.mojolang.mojo.core.RetryInfo(this);
      if (retryDelayBuilder_ == null) {
        result.retryDelay_ = retryDelay_;
      } else {
        result.retryDelay_ = retryDelayBuilder_.build();
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
      if (other instanceof org.mojolang.mojo.core.RetryInfo) {
        return mergeFrom((org.mojolang.mojo.core.RetryInfo)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.RetryInfo other) {
      if (other == org.mojolang.mojo.core.RetryInfo.getDefaultInstance()) return this;
      if (other.hasRetryDelay()) {
        mergeRetryDelay(other.getRetryDelay());
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
      org.mojolang.mojo.core.RetryInfo parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.RetryInfo) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private org.mojolang.mojo.core.Duration retryDelay_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Duration, org.mojolang.mojo.core.Duration.Builder, org.mojolang.mojo.core.DurationOrBuilder> retryDelayBuilder_;
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     * @return Whether the retryDelay field is set.
     */
    public boolean hasRetryDelay() {
      return retryDelayBuilder_ != null || retryDelay_ != null;
    }
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     * @return The retryDelay.
     */
    public org.mojolang.mojo.core.Duration getRetryDelay() {
      if (retryDelayBuilder_ == null) {
        return retryDelay_ == null ? org.mojolang.mojo.core.Duration.getDefaultInstance() : retryDelay_;
      } else {
        return retryDelayBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     */
    public Builder setRetryDelay(org.mojolang.mojo.core.Duration value) {
      if (retryDelayBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        retryDelay_ = value;
        onChanged();
      } else {
        retryDelayBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     */
    public Builder setRetryDelay(
        org.mojolang.mojo.core.Duration.Builder builderForValue) {
      if (retryDelayBuilder_ == null) {
        retryDelay_ = builderForValue.build();
        onChanged();
      } else {
        retryDelayBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     */
    public Builder mergeRetryDelay(org.mojolang.mojo.core.Duration value) {
      if (retryDelayBuilder_ == null) {
        if (retryDelay_ != null) {
          retryDelay_ =
            org.mojolang.mojo.core.Duration.newBuilder(retryDelay_).mergeFrom(value).buildPartial();
        } else {
          retryDelay_ = value;
        }
        onChanged();
      } else {
        retryDelayBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     */
    public Builder clearRetryDelay() {
      if (retryDelayBuilder_ == null) {
        retryDelay_ = null;
        onChanged();
      } else {
        retryDelay_ = null;
        retryDelayBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     */
    public org.mojolang.mojo.core.Duration.Builder getRetryDelayBuilder() {
      
      onChanged();
      return getRetryDelayFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     */
    public org.mojolang.mojo.core.DurationOrBuilder getRetryDelayOrBuilder() {
      if (retryDelayBuilder_ != null) {
        return retryDelayBuilder_.getMessageOrBuilder();
      } else {
        return retryDelay_ == null ?
            org.mojolang.mojo.core.Duration.getDefaultInstance() : retryDelay_;
      }
    }
    /**
     * <code>.mojo.core.Duration retry_delay = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Duration, org.mojolang.mojo.core.Duration.Builder, org.mojolang.mojo.core.DurationOrBuilder> 
        getRetryDelayFieldBuilder() {
      if (retryDelayBuilder_ == null) {
        retryDelayBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.Duration, org.mojolang.mojo.core.Duration.Builder, org.mojolang.mojo.core.DurationOrBuilder>(
                getRetryDelay(),
                getParentForChildren(),
                isClean());
        retryDelay_ = null;
      }
      return retryDelayBuilder_;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.RetryInfo)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.RetryInfo)
  private static final org.mojolang.mojo.core.RetryInfo DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.RetryInfo();
  }

  public static org.mojolang.mojo.core.RetryInfo getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<RetryInfo>
      PARSER = new com.google.protobuf.AbstractParser<RetryInfo>() {
    @java.lang.Override
    public RetryInfo parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new RetryInfo(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<RetryInfo> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<RetryInfo> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.RetryInfo getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

