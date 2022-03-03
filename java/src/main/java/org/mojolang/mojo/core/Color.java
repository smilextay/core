// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/color.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.Color}
 */
public final class Color extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.Color)
    ColorOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Color.newBuilder() to construct.
  private Color(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Color() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Color();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private Color(
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

            red_ = input.readInt32();
            break;
          }
          case 16: {

            green_ = input.readInt32();
            break;
          }
          case 24: {

            blue_ = input.readInt32();
            break;
          }
          case 37: {

            alpha_ = input.readFloat();
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
    return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.Color.class, org.mojolang.mojo.core.Color.Builder.class);
  }

  public static final int RED_FIELD_NUMBER = 1;
  private int red_;
  /**
   * <code>int32 red = 1;</code>
   * @return The red.
   */
  @java.lang.Override
  public int getRed() {
    return red_;
  }

  public static final int GREEN_FIELD_NUMBER = 2;
  private int green_;
  /**
   * <code>int32 green = 2;</code>
   * @return The green.
   */
  @java.lang.Override
  public int getGreen() {
    return green_;
  }

  public static final int BLUE_FIELD_NUMBER = 3;
  private int blue_;
  /**
   * <code>int32 blue = 3;</code>
   * @return The blue.
   */
  @java.lang.Override
  public int getBlue() {
    return blue_;
  }

  public static final int ALPHA_FIELD_NUMBER = 4;
  private float alpha_;
  /**
   * <code>float alpha = 4;</code>
   * @return The alpha.
   */
  @java.lang.Override
  public float getAlpha() {
    return alpha_;
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
    if (red_ != 0) {
      output.writeInt32(1, red_);
    }
    if (green_ != 0) {
      output.writeInt32(2, green_);
    }
    if (blue_ != 0) {
      output.writeInt32(3, blue_);
    }
    if (java.lang.Float.floatToRawIntBits(alpha_) != 0) {
      output.writeFloat(4, alpha_);
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (red_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(1, red_);
    }
    if (green_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, green_);
    }
    if (blue_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(3, blue_);
    }
    if (java.lang.Float.floatToRawIntBits(alpha_) != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeFloatSize(4, alpha_);
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
    if (!(obj instanceof org.mojolang.mojo.core.Color)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.Color other = (org.mojolang.mojo.core.Color) obj;

    if (getRed()
        != other.getRed()) return false;
    if (getGreen()
        != other.getGreen()) return false;
    if (getBlue()
        != other.getBlue()) return false;
    if (java.lang.Float.floatToIntBits(getAlpha())
        != java.lang.Float.floatToIntBits(
            other.getAlpha())) return false;
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
    hash = (37 * hash) + RED_FIELD_NUMBER;
    hash = (53 * hash) + getRed();
    hash = (37 * hash) + GREEN_FIELD_NUMBER;
    hash = (53 * hash) + getGreen();
    hash = (37 * hash) + BLUE_FIELD_NUMBER;
    hash = (53 * hash) + getBlue();
    hash = (37 * hash) + ALPHA_FIELD_NUMBER;
    hash = (53 * hash) + java.lang.Float.floatToIntBits(
        getAlpha());
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.Color parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Color parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.Color parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.Color prototype) {
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
   * Protobuf type {@code mojo.core.Color}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.Color)
      org.mojolang.mojo.core.ColorOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.Color.class, org.mojolang.mojo.core.Color.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.Color.newBuilder()
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
      red_ = 0;

      green_ = 0;

      blue_ = 0;

      alpha_ = 0F;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.ColorProto.internal_static_mojo_core_Color_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Color getDefaultInstanceForType() {
      return org.mojolang.mojo.core.Color.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Color build() {
      org.mojolang.mojo.core.Color result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.Color buildPartial() {
      org.mojolang.mojo.core.Color result = new org.mojolang.mojo.core.Color(this);
      result.red_ = red_;
      result.green_ = green_;
      result.blue_ = blue_;
      result.alpha_ = alpha_;
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
      if (other instanceof org.mojolang.mojo.core.Color) {
        return mergeFrom((org.mojolang.mojo.core.Color)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.Color other) {
      if (other == org.mojolang.mojo.core.Color.getDefaultInstance()) return this;
      if (other.getRed() != 0) {
        setRed(other.getRed());
      }
      if (other.getGreen() != 0) {
        setGreen(other.getGreen());
      }
      if (other.getBlue() != 0) {
        setBlue(other.getBlue());
      }
      if (other.getAlpha() != 0F) {
        setAlpha(other.getAlpha());
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
      org.mojolang.mojo.core.Color parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.Color) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private int red_ ;
    /**
     * <code>int32 red = 1;</code>
     * @return The red.
     */
    @java.lang.Override
    public int getRed() {
      return red_;
    }
    /**
     * <code>int32 red = 1;</code>
     * @param value The red to set.
     * @return This builder for chaining.
     */
    public Builder setRed(int value) {
      
      red_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 red = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearRed() {
      
      red_ = 0;
      onChanged();
      return this;
    }

    private int green_ ;
    /**
     * <code>int32 green = 2;</code>
     * @return The green.
     */
    @java.lang.Override
    public int getGreen() {
      return green_;
    }
    /**
     * <code>int32 green = 2;</code>
     * @param value The green to set.
     * @return This builder for chaining.
     */
    public Builder setGreen(int value) {
      
      green_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 green = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearGreen() {
      
      green_ = 0;
      onChanged();
      return this;
    }

    private int blue_ ;
    /**
     * <code>int32 blue = 3;</code>
     * @return The blue.
     */
    @java.lang.Override
    public int getBlue() {
      return blue_;
    }
    /**
     * <code>int32 blue = 3;</code>
     * @param value The blue to set.
     * @return This builder for chaining.
     */
    public Builder setBlue(int value) {
      
      blue_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 blue = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearBlue() {
      
      blue_ = 0;
      onChanged();
      return this;
    }

    private float alpha_ ;
    /**
     * <code>float alpha = 4;</code>
     * @return The alpha.
     */
    @java.lang.Override
    public float getAlpha() {
      return alpha_;
    }
    /**
     * <code>float alpha = 4;</code>
     * @param value The alpha to set.
     * @return This builder for chaining.
     */
    public Builder setAlpha(float value) {
      
      alpha_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>float alpha = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearAlpha() {
      
      alpha_ = 0F;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.Color)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.Color)
  private static final org.mojolang.mojo.core.Color DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.Color();
  }

  public static org.mojolang.mojo.core.Color getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Color>
      PARSER = new com.google.protobuf.AbstractParser<Color>() {
    @java.lang.Override
    public Color parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new Color(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<Color> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Color> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.Color getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

