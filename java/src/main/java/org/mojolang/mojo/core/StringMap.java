// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/boxed.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.StringMap}
 */
public final class StringMap extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.StringMap)
    StringMapOrBuilder {
private static final long serialVersionUID = 0L;
  // Use StringMap.newBuilder() to construct.
  private StringMap(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private StringMap() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new StringMap();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private StringMap(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    int mutable_bitField0_ = 0;
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
            if (!((mutable_bitField0_ & 0x00000001) != 0)) {
              vals_ = com.google.protobuf.MapField.newMapField(
                  ValsDefaultEntryHolder.defaultEntry);
              mutable_bitField0_ |= 0x00000001;
            }
            com.google.protobuf.MapEntry<java.lang.String, java.lang.String>
            vals__ = input.readMessage(
                ValsDefaultEntryHolder.defaultEntry.getParserForType(), extensionRegistry);
            vals_.getMutableMap().put(
                vals__.getKey(), vals__.getValue());
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
    return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringMap_descriptor;
  }

  @SuppressWarnings({"rawtypes"})
  @java.lang.Override
  protected com.google.protobuf.MapField internalGetMapField(
      int number) {
    switch (number) {
      case 1:
        return internalGetVals();
      default:
        throw new RuntimeException(
            "Invalid map field number: " + number);
    }
  }
  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringMap_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.StringMap.class, org.mojolang.mojo.core.StringMap.Builder.class);
  }

  public static final int VALS_FIELD_NUMBER = 1;
  private static final class ValsDefaultEntryHolder {
    static final com.google.protobuf.MapEntry<
        java.lang.String, java.lang.String> defaultEntry =
            com.google.protobuf.MapEntry
            .<java.lang.String, java.lang.String>newDefaultInstance(
                org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringMap_ValsEntry_descriptor, 
                com.google.protobuf.WireFormat.FieldType.STRING,
                "",
                com.google.protobuf.WireFormat.FieldType.STRING,
                "");
  }
  private com.google.protobuf.MapField<
      java.lang.String, java.lang.String> vals_;
  private com.google.protobuf.MapField<java.lang.String, java.lang.String>
  internalGetVals() {
    if (vals_ == null) {
      return com.google.protobuf.MapField.emptyMapField(
          ValsDefaultEntryHolder.defaultEntry);
    }
    return vals_;
  }

  public int getValsCount() {
    return internalGetVals().getMap().size();
  }
  /**
   * <code>map&lt;string, string&gt; vals = 1;</code>
   */

  @java.lang.Override
  public boolean containsVals(
      java.lang.String key) {
    if (key == null) { throw new NullPointerException("map key"); }
    return internalGetVals().getMap().containsKey(key);
  }
  /**
   * Use {@link #getValsMap()} instead.
   */
  @java.lang.Override
  @java.lang.Deprecated
  public java.util.Map<java.lang.String, java.lang.String> getVals() {
    return getValsMap();
  }
  /**
   * <code>map&lt;string, string&gt; vals = 1;</code>
   */
  @java.lang.Override

  public java.util.Map<java.lang.String, java.lang.String> getValsMap() {
    return internalGetVals().getMap();
  }
  /**
   * <code>map&lt;string, string&gt; vals = 1;</code>
   */
  @java.lang.Override

  public java.lang.String getValsOrDefault(
      java.lang.String key,
      java.lang.String defaultValue) {
    if (key == null) { throw new NullPointerException("map key"); }
    java.util.Map<java.lang.String, java.lang.String> map =
        internalGetVals().getMap();
    return map.containsKey(key) ? map.get(key) : defaultValue;
  }
  /**
   * <code>map&lt;string, string&gt; vals = 1;</code>
   */
  @java.lang.Override

  public java.lang.String getValsOrThrow(
      java.lang.String key) {
    if (key == null) { throw new NullPointerException("map key"); }
    java.util.Map<java.lang.String, java.lang.String> map =
        internalGetVals().getMap();
    if (!map.containsKey(key)) {
      throw new java.lang.IllegalArgumentException();
    }
    return map.get(key);
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
    com.google.protobuf.GeneratedMessageV3
      .serializeStringMapTo(
        output,
        internalGetVals(),
        ValsDefaultEntryHolder.defaultEntry,
        1);
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    for (java.util.Map.Entry<java.lang.String, java.lang.String> entry
         : internalGetVals().getMap().entrySet()) {
      com.google.protobuf.MapEntry<java.lang.String, java.lang.String>
      vals__ = ValsDefaultEntryHolder.defaultEntry.newBuilderForType()
          .setKey(entry.getKey())
          .setValue(entry.getValue())
          .build();
      size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(1, vals__);
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
    if (!(obj instanceof org.mojolang.mojo.core.StringMap)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.StringMap other = (org.mojolang.mojo.core.StringMap) obj;

    if (!internalGetVals().equals(
        other.internalGetVals())) return false;
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
    if (!internalGetVals().getMap().isEmpty()) {
      hash = (37 * hash) + VALS_FIELD_NUMBER;
      hash = (53 * hash) + internalGetVals().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.StringMap parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringMap parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.StringMap parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.StringMap parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.StringMap prototype) {
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
   * Protobuf type {@code mojo.core.StringMap}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.StringMap)
      org.mojolang.mojo.core.StringMapOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringMap_descriptor;
    }

    @SuppressWarnings({"rawtypes"})
    protected com.google.protobuf.MapField internalGetMapField(
        int number) {
      switch (number) {
        case 1:
          return internalGetVals();
        default:
          throw new RuntimeException(
              "Invalid map field number: " + number);
      }
    }
    @SuppressWarnings({"rawtypes"})
    protected com.google.protobuf.MapField internalGetMutableMapField(
        int number) {
      switch (number) {
        case 1:
          return internalGetMutableVals();
        default:
          throw new RuntimeException(
              "Invalid map field number: " + number);
      }
    }
    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringMap_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.StringMap.class, org.mojolang.mojo.core.StringMap.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.StringMap.newBuilder()
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
      internalGetMutableVals().clear();
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.BoxedProto.internal_static_mojo_core_StringMap_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.StringMap getDefaultInstanceForType() {
      return org.mojolang.mojo.core.StringMap.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.StringMap build() {
      org.mojolang.mojo.core.StringMap result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.StringMap buildPartial() {
      org.mojolang.mojo.core.StringMap result = new org.mojolang.mojo.core.StringMap(this);
      int from_bitField0_ = bitField0_;
      result.vals_ = internalGetVals();
      result.vals_.makeImmutable();
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
      if (other instanceof org.mojolang.mojo.core.StringMap) {
        return mergeFrom((org.mojolang.mojo.core.StringMap)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.StringMap other) {
      if (other == org.mojolang.mojo.core.StringMap.getDefaultInstance()) return this;
      internalGetMutableVals().mergeFrom(
          other.internalGetVals());
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
      org.mojolang.mojo.core.StringMap parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.StringMap) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private com.google.protobuf.MapField<
        java.lang.String, java.lang.String> vals_;
    private com.google.protobuf.MapField<java.lang.String, java.lang.String>
    internalGetVals() {
      if (vals_ == null) {
        return com.google.protobuf.MapField.emptyMapField(
            ValsDefaultEntryHolder.defaultEntry);
      }
      return vals_;
    }
    private com.google.protobuf.MapField<java.lang.String, java.lang.String>
    internalGetMutableVals() {
      onChanged();;
      if (vals_ == null) {
        vals_ = com.google.protobuf.MapField.newMapField(
            ValsDefaultEntryHolder.defaultEntry);
      }
      if (!vals_.isMutable()) {
        vals_ = vals_.copy();
      }
      return vals_;
    }

    public int getValsCount() {
      return internalGetVals().getMap().size();
    }
    /**
     * <code>map&lt;string, string&gt; vals = 1;</code>
     */

    @java.lang.Override
    public boolean containsVals(
        java.lang.String key) {
      if (key == null) { throw new NullPointerException("map key"); }
      return internalGetVals().getMap().containsKey(key);
    }
    /**
     * Use {@link #getValsMap()} instead.
     */
    @java.lang.Override
    @java.lang.Deprecated
    public java.util.Map<java.lang.String, java.lang.String> getVals() {
      return getValsMap();
    }
    /**
     * <code>map&lt;string, string&gt; vals = 1;</code>
     */
    @java.lang.Override

    public java.util.Map<java.lang.String, java.lang.String> getValsMap() {
      return internalGetVals().getMap();
    }
    /**
     * <code>map&lt;string, string&gt; vals = 1;</code>
     */
    @java.lang.Override

    public java.lang.String getValsOrDefault(
        java.lang.String key,
        java.lang.String defaultValue) {
      if (key == null) { throw new NullPointerException("map key"); }
      java.util.Map<java.lang.String, java.lang.String> map =
          internalGetVals().getMap();
      return map.containsKey(key) ? map.get(key) : defaultValue;
    }
    /**
     * <code>map&lt;string, string&gt; vals = 1;</code>
     */
    @java.lang.Override

    public java.lang.String getValsOrThrow(
        java.lang.String key) {
      if (key == null) { throw new NullPointerException("map key"); }
      java.util.Map<java.lang.String, java.lang.String> map =
          internalGetVals().getMap();
      if (!map.containsKey(key)) {
        throw new java.lang.IllegalArgumentException();
      }
      return map.get(key);
    }

    public Builder clearVals() {
      internalGetMutableVals().getMutableMap()
          .clear();
      return this;
    }
    /**
     * <code>map&lt;string, string&gt; vals = 1;</code>
     */

    public Builder removeVals(
        java.lang.String key) {
      if (key == null) { throw new NullPointerException("map key"); }
      internalGetMutableVals().getMutableMap()
          .remove(key);
      return this;
    }
    /**
     * Use alternate mutation accessors instead.
     */
    @java.lang.Deprecated
    public java.util.Map<java.lang.String, java.lang.String>
    getMutableVals() {
      return internalGetMutableVals().getMutableMap();
    }
    /**
     * <code>map&lt;string, string&gt; vals = 1;</code>
     */
    public Builder putVals(
        java.lang.String key,
        java.lang.String value) {
      if (key == null) { throw new NullPointerException("map key"); }
      if (value == null) {
  throw new NullPointerException("map value");
}

      internalGetMutableVals().getMutableMap()
          .put(key, value);
      return this;
    }
    /**
     * <code>map&lt;string, string&gt; vals = 1;</code>
     */

    public Builder putAllVals(
        java.util.Map<java.lang.String, java.lang.String> values) {
      internalGetMutableVals().getMutableMap()
          .putAll(values);
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


    // @@protoc_insertion_point(builder_scope:mojo.core.StringMap)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.StringMap)
  private static final org.mojolang.mojo.core.StringMap DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.StringMap();
  }

  public static org.mojolang.mojo.core.StringMap getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<StringMap>
      PARSER = new com.google.protobuf.AbstractParser<StringMap>() {
    @java.lang.Override
    public StringMap parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new StringMap(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<StringMap> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<StringMap> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.StringMap getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

