// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/version_range.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.VersionRange}
 */
public final class VersionRange extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.VersionRange)
    VersionRangeOrBuilder {
private static final long serialVersionUID = 0L;
  // Use VersionRange.newBuilder() to construct.
  private VersionRange(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private VersionRange() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new VersionRange();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private VersionRange(
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
            org.mojolang.mojo.core.Version.Builder subBuilder = null;
            if (min_ != null) {
              subBuilder = min_.toBuilder();
            }
            min_ = input.readMessage(org.mojolang.mojo.core.Version.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(min_);
              min_ = subBuilder.buildPartial();
            }

            break;
          }
          case 18: {
            org.mojolang.mojo.core.Version.Builder subBuilder = null;
            if (max_ != null) {
              subBuilder = max_.toBuilder();
            }
            max_ = input.readMessage(org.mojolang.mojo.core.Version.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(max_);
              max_ = subBuilder.buildPartial();
            }

            break;
          }
          case 24: {

            minExcluded_ = input.readBool();
            break;
          }
          case 32: {

            maxExcluded_ = input.readBool();
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
    return org.mojolang.mojo.core.VersionRangeProto.internal_static_mojo_core_VersionRange_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.VersionRangeProto.internal_static_mojo_core_VersionRange_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.VersionRange.class, org.mojolang.mojo.core.VersionRange.Builder.class);
  }

  public static final int MIN_FIELD_NUMBER = 1;
  private org.mojolang.mojo.core.Version min_;
  /**
   * <code>.mojo.core.Version min = 1;</code>
   * @return Whether the min field is set.
   */
  @java.lang.Override
  public boolean hasMin() {
    return min_ != null;
  }
  /**
   * <code>.mojo.core.Version min = 1;</code>
   * @return The min.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Version getMin() {
    return min_ == null ? org.mojolang.mojo.core.Version.getDefaultInstance() : min_;
  }
  /**
   * <code>.mojo.core.Version min = 1;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.VersionOrBuilder getMinOrBuilder() {
    return getMin();
  }

  public static final int MAX_FIELD_NUMBER = 2;
  private org.mojolang.mojo.core.Version max_;
  /**
   * <code>.mojo.core.Version max = 2;</code>
   * @return Whether the max field is set.
   */
  @java.lang.Override
  public boolean hasMax() {
    return max_ != null;
  }
  /**
   * <code>.mojo.core.Version max = 2;</code>
   * @return The max.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.Version getMax() {
    return max_ == null ? org.mojolang.mojo.core.Version.getDefaultInstance() : max_;
  }
  /**
   * <code>.mojo.core.Version max = 2;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.VersionOrBuilder getMaxOrBuilder() {
    return getMax();
  }

  public static final int MIN_EXCLUDED_FIELD_NUMBER = 3;
  private boolean minExcluded_;
  /**
   * <code>bool min_excluded = 3;</code>
   * @return The minExcluded.
   */
  @java.lang.Override
  public boolean getMinExcluded() {
    return minExcluded_;
  }

  public static final int MAX_EXCLUDED_FIELD_NUMBER = 4;
  private boolean maxExcluded_;
  /**
   * <code>bool max_excluded = 4;</code>
   * @return The maxExcluded.
   */
  @java.lang.Override
  public boolean getMaxExcluded() {
    return maxExcluded_;
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
    if (min_ != null) {
      output.writeMessage(1, getMin());
    }
    if (max_ != null) {
      output.writeMessage(2, getMax());
    }
    if (minExcluded_ != false) {
      output.writeBool(3, minExcluded_);
    }
    if (maxExcluded_ != false) {
      output.writeBool(4, maxExcluded_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (min_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getMin());
    }
    if (max_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getMax());
    }
    if (minExcluded_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(3, minExcluded_);
    }
    if (maxExcluded_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(4, maxExcluded_);
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
    if (!(obj instanceof org.mojolang.mojo.core.VersionRange)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.VersionRange other = (org.mojolang.mojo.core.VersionRange) obj;

    if (hasMin() != other.hasMin()) return false;
    if (hasMin()) {
      if (!getMin()
          .equals(other.getMin())) return false;
    }
    if (hasMax() != other.hasMax()) return false;
    if (hasMax()) {
      if (!getMax()
          .equals(other.getMax())) return false;
    }
    if (getMinExcluded()
        != other.getMinExcluded()) return false;
    if (getMaxExcluded()
        != other.getMaxExcluded()) return false;
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
    if (hasMin()) {
      hash = (37 * hash) + MIN_FIELD_NUMBER;
      hash = (53 * hash) + getMin().hashCode();
    }
    if (hasMax()) {
      hash = (37 * hash) + MAX_FIELD_NUMBER;
      hash = (53 * hash) + getMax().hashCode();
    }
    hash = (37 * hash) + MIN_EXCLUDED_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getMinExcluded());
    hash = (37 * hash) + MAX_EXCLUDED_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getMaxExcluded());
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.VersionRange parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.VersionRange parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.VersionRange parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.VersionRange parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.VersionRange prototype) {
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
   * Protobuf type {@code mojo.core.VersionRange}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.VersionRange)
      org.mojolang.mojo.core.VersionRangeOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.VersionRangeProto.internal_static_mojo_core_VersionRange_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.VersionRangeProto.internal_static_mojo_core_VersionRange_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.VersionRange.class, org.mojolang.mojo.core.VersionRange.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.VersionRange.newBuilder()
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
      if (minBuilder_ == null) {
        min_ = null;
      } else {
        min_ = null;
        minBuilder_ = null;
      }
      if (maxBuilder_ == null) {
        max_ = null;
      } else {
        max_ = null;
        maxBuilder_ = null;
      }
      minExcluded_ = false;

      maxExcluded_ = false;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.VersionRangeProto.internal_static_mojo_core_VersionRange_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.VersionRange getDefaultInstanceForType() {
      return org.mojolang.mojo.core.VersionRange.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.VersionRange build() {
      org.mojolang.mojo.core.VersionRange result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.VersionRange buildPartial() {
      org.mojolang.mojo.core.VersionRange result = new org.mojolang.mojo.core.VersionRange(this);
      if (minBuilder_ == null) {
        result.min_ = min_;
      } else {
        result.min_ = minBuilder_.build();
      }
      if (maxBuilder_ == null) {
        result.max_ = max_;
      } else {
        result.max_ = maxBuilder_.build();
      }
      result.minExcluded_ = minExcluded_;
      result.maxExcluded_ = maxExcluded_;
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
      if (other instanceof org.mojolang.mojo.core.VersionRange) {
        return mergeFrom((org.mojolang.mojo.core.VersionRange)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.VersionRange other) {
      if (other == org.mojolang.mojo.core.VersionRange.getDefaultInstance()) return this;
      if (other.hasMin()) {
        mergeMin(other.getMin());
      }
      if (other.hasMax()) {
        mergeMax(other.getMax());
      }
      if (other.getMinExcluded() != false) {
        setMinExcluded(other.getMinExcluded());
      }
      if (other.getMaxExcluded() != false) {
        setMaxExcluded(other.getMaxExcluded());
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
      org.mojolang.mojo.core.VersionRange parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.VersionRange) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private org.mojolang.mojo.core.Version min_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Version, org.mojolang.mojo.core.Version.Builder, org.mojolang.mojo.core.VersionOrBuilder> minBuilder_;
    /**
     * <code>.mojo.core.Version min = 1;</code>
     * @return Whether the min field is set.
     */
    public boolean hasMin() {
      return minBuilder_ != null || min_ != null;
    }
    /**
     * <code>.mojo.core.Version min = 1;</code>
     * @return The min.
     */
    public org.mojolang.mojo.core.Version getMin() {
      if (minBuilder_ == null) {
        return min_ == null ? org.mojolang.mojo.core.Version.getDefaultInstance() : min_;
      } else {
        return minBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.Version min = 1;</code>
     */
    public Builder setMin(org.mojolang.mojo.core.Version value) {
      if (minBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        min_ = value;
        onChanged();
      } else {
        minBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Version min = 1;</code>
     */
    public Builder setMin(
        org.mojolang.mojo.core.Version.Builder builderForValue) {
      if (minBuilder_ == null) {
        min_ = builderForValue.build();
        onChanged();
      } else {
        minBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.Version min = 1;</code>
     */
    public Builder mergeMin(org.mojolang.mojo.core.Version value) {
      if (minBuilder_ == null) {
        if (min_ != null) {
          min_ =
            org.mojolang.mojo.core.Version.newBuilder(min_).mergeFrom(value).buildPartial();
        } else {
          min_ = value;
        }
        onChanged();
      } else {
        minBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Version min = 1;</code>
     */
    public Builder clearMin() {
      if (minBuilder_ == null) {
        min_ = null;
        onChanged();
      } else {
        min_ = null;
        minBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.Version min = 1;</code>
     */
    public org.mojolang.mojo.core.Version.Builder getMinBuilder() {
      
      onChanged();
      return getMinFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.Version min = 1;</code>
     */
    public org.mojolang.mojo.core.VersionOrBuilder getMinOrBuilder() {
      if (minBuilder_ != null) {
        return minBuilder_.getMessageOrBuilder();
      } else {
        return min_ == null ?
            org.mojolang.mojo.core.Version.getDefaultInstance() : min_;
      }
    }
    /**
     * <code>.mojo.core.Version min = 1;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Version, org.mojolang.mojo.core.Version.Builder, org.mojolang.mojo.core.VersionOrBuilder> 
        getMinFieldBuilder() {
      if (minBuilder_ == null) {
        minBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.Version, org.mojolang.mojo.core.Version.Builder, org.mojolang.mojo.core.VersionOrBuilder>(
                getMin(),
                getParentForChildren(),
                isClean());
        min_ = null;
      }
      return minBuilder_;
    }

    private org.mojolang.mojo.core.Version max_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Version, org.mojolang.mojo.core.Version.Builder, org.mojolang.mojo.core.VersionOrBuilder> maxBuilder_;
    /**
     * <code>.mojo.core.Version max = 2;</code>
     * @return Whether the max field is set.
     */
    public boolean hasMax() {
      return maxBuilder_ != null || max_ != null;
    }
    /**
     * <code>.mojo.core.Version max = 2;</code>
     * @return The max.
     */
    public org.mojolang.mojo.core.Version getMax() {
      if (maxBuilder_ == null) {
        return max_ == null ? org.mojolang.mojo.core.Version.getDefaultInstance() : max_;
      } else {
        return maxBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.Version max = 2;</code>
     */
    public Builder setMax(org.mojolang.mojo.core.Version value) {
      if (maxBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        max_ = value;
        onChanged();
      } else {
        maxBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Version max = 2;</code>
     */
    public Builder setMax(
        org.mojolang.mojo.core.Version.Builder builderForValue) {
      if (maxBuilder_ == null) {
        max_ = builderForValue.build();
        onChanged();
      } else {
        maxBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.Version max = 2;</code>
     */
    public Builder mergeMax(org.mojolang.mojo.core.Version value) {
      if (maxBuilder_ == null) {
        if (max_ != null) {
          max_ =
            org.mojolang.mojo.core.Version.newBuilder(max_).mergeFrom(value).buildPartial();
        } else {
          max_ = value;
        }
        onChanged();
      } else {
        maxBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.Version max = 2;</code>
     */
    public Builder clearMax() {
      if (maxBuilder_ == null) {
        max_ = null;
        onChanged();
      } else {
        max_ = null;
        maxBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.Version max = 2;</code>
     */
    public org.mojolang.mojo.core.Version.Builder getMaxBuilder() {
      
      onChanged();
      return getMaxFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.Version max = 2;</code>
     */
    public org.mojolang.mojo.core.VersionOrBuilder getMaxOrBuilder() {
      if (maxBuilder_ != null) {
        return maxBuilder_.getMessageOrBuilder();
      } else {
        return max_ == null ?
            org.mojolang.mojo.core.Version.getDefaultInstance() : max_;
      }
    }
    /**
     * <code>.mojo.core.Version max = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.Version, org.mojolang.mojo.core.Version.Builder, org.mojolang.mojo.core.VersionOrBuilder> 
        getMaxFieldBuilder() {
      if (maxBuilder_ == null) {
        maxBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.Version, org.mojolang.mojo.core.Version.Builder, org.mojolang.mojo.core.VersionOrBuilder>(
                getMax(),
                getParentForChildren(),
                isClean());
        max_ = null;
      }
      return maxBuilder_;
    }

    private boolean minExcluded_ ;
    /**
     * <code>bool min_excluded = 3;</code>
     * @return The minExcluded.
     */
    @java.lang.Override
    public boolean getMinExcluded() {
      return minExcluded_;
    }
    /**
     * <code>bool min_excluded = 3;</code>
     * @param value The minExcluded to set.
     * @return This builder for chaining.
     */
    public Builder setMinExcluded(boolean value) {
      
      minExcluded_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>bool min_excluded = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearMinExcluded() {
      
      minExcluded_ = false;
      onChanged();
      return this;
    }

    private boolean maxExcluded_ ;
    /**
     * <code>bool max_excluded = 4;</code>
     * @return The maxExcluded.
     */
    @java.lang.Override
    public boolean getMaxExcluded() {
      return maxExcluded_;
    }
    /**
     * <code>bool max_excluded = 4;</code>
     * @param value The maxExcluded to set.
     * @return This builder for chaining.
     */
    public Builder setMaxExcluded(boolean value) {
      
      maxExcluded_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>bool max_excluded = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearMaxExcluded() {
      
      maxExcluded_ = false;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.VersionRange)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.VersionRange)
  private static final org.mojolang.mojo.core.VersionRange DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.VersionRange();
  }

  public static org.mojolang.mojo.core.VersionRange getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<VersionRange>
      PARSER = new com.google.protobuf.AbstractParser<VersionRange>() {
    @java.lang.Override
    public VersionRange parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new VersionRange(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<VersionRange> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<VersionRange> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.VersionRange getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

