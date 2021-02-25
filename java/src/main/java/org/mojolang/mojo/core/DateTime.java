// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/time.proto

package org.mojolang.mojo.core;

/**
 * Protobuf type {@code mojo.core.DateTime}
 */
public final class DateTime extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:mojo.core.DateTime)
    DateTimeOrBuilder {
private static final long serialVersionUID = 0L;
  // Use DateTime.newBuilder() to construct.
  private DateTime(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private DateTime() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new DateTime();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private DateTime(
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

            year_ = input.readInt32();
            break;
          }
          case 16: {

            month_ = input.readInt32();
            break;
          }
          case 24: {

            day_ = input.readInt32();
            break;
          }
          case 32: {

            hour_ = input.readInt32();
            break;
          }
          case 40: {

            minute_ = input.readInt32();
            break;
          }
          case 48: {

            seconds_ = input.readInt64();
            break;
          }
          case 56: {

            nanoseconds_ = input.readInt32();
            break;
          }
          case 122: {
            org.mojolang.mojo.core.TimeZone.Builder subBuilder = null;
            if (timeZone_ != null) {
              subBuilder = timeZone_.toBuilder();
            }
            timeZone_ = input.readMessage(org.mojolang.mojo.core.TimeZone.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(timeZone_);
              timeZone_ = subBuilder.buildPartial();
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
    return org.mojolang.mojo.core.TimeProto.internal_static_mojo_core_DateTime_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return org.mojolang.mojo.core.TimeProto.internal_static_mojo_core_DateTime_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            org.mojolang.mojo.core.DateTime.class, org.mojolang.mojo.core.DateTime.Builder.class);
  }

  public static final int YEAR_FIELD_NUMBER = 1;
  private int year_;
  /**
   * <code>int32 year = 1;</code>
   * @return The year.
   */
  @java.lang.Override
  public int getYear() {
    return year_;
  }

  public static final int MONTH_FIELD_NUMBER = 2;
  private int month_;
  /**
   * <code>int32 month = 2;</code>
   * @return The month.
   */
  @java.lang.Override
  public int getMonth() {
    return month_;
  }

  public static final int DAY_FIELD_NUMBER = 3;
  private int day_;
  /**
   * <code>int32 day = 3;</code>
   * @return The day.
   */
  @java.lang.Override
  public int getDay() {
    return day_;
  }

  public static final int HOUR_FIELD_NUMBER = 4;
  private int hour_;
  /**
   * <code>int32 hour = 4;</code>
   * @return The hour.
   */
  @java.lang.Override
  public int getHour() {
    return hour_;
  }

  public static final int MINUTE_FIELD_NUMBER = 5;
  private int minute_;
  /**
   * <code>int32 minute = 5;</code>
   * @return The minute.
   */
  @java.lang.Override
  public int getMinute() {
    return minute_;
  }

  public static final int SECONDS_FIELD_NUMBER = 6;
  private long seconds_;
  /**
   * <code>int64 seconds = 6;</code>
   * @return The seconds.
   */
  @java.lang.Override
  public long getSeconds() {
    return seconds_;
  }

  public static final int NANOSECONDS_FIELD_NUMBER = 7;
  private int nanoseconds_;
  /**
   * <code>int32 nanoseconds = 7;</code>
   * @return The nanoseconds.
   */
  @java.lang.Override
  public int getNanoseconds() {
    return nanoseconds_;
  }

  public static final int TIME_ZONE_FIELD_NUMBER = 15;
  private org.mojolang.mojo.core.TimeZone timeZone_;
  /**
   * <code>.mojo.core.TimeZone time_zone = 15;</code>
   * @return Whether the timeZone field is set.
   */
  @java.lang.Override
  public boolean hasTimeZone() {
    return timeZone_ != null;
  }
  /**
   * <code>.mojo.core.TimeZone time_zone = 15;</code>
   * @return The timeZone.
   */
  @java.lang.Override
  public org.mojolang.mojo.core.TimeZone getTimeZone() {
    return timeZone_ == null ? org.mojolang.mojo.core.TimeZone.getDefaultInstance() : timeZone_;
  }
  /**
   * <code>.mojo.core.TimeZone time_zone = 15;</code>
   */
  @java.lang.Override
  public org.mojolang.mojo.core.TimeZoneOrBuilder getTimeZoneOrBuilder() {
    return getTimeZone();
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
    if (year_ != 0) {
      output.writeInt32(1, year_);
    }
    if (month_ != 0) {
      output.writeInt32(2, month_);
    }
    if (day_ != 0) {
      output.writeInt32(3, day_);
    }
    if (hour_ != 0) {
      output.writeInt32(4, hour_);
    }
    if (minute_ != 0) {
      output.writeInt32(5, minute_);
    }
    if (seconds_ != 0L) {
      output.writeInt64(6, seconds_);
    }
    if (nanoseconds_ != 0) {
      output.writeInt32(7, nanoseconds_);
    }
    if (timeZone_ != null) {
      output.writeMessage(15, getTimeZone());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (year_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(1, year_);
    }
    if (month_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, month_);
    }
    if (day_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(3, day_);
    }
    if (hour_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(4, hour_);
    }
    if (minute_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(5, minute_);
    }
    if (seconds_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(6, seconds_);
    }
    if (nanoseconds_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(7, nanoseconds_);
    }
    if (timeZone_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(15, getTimeZone());
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
    if (!(obj instanceof org.mojolang.mojo.core.DateTime)) {
      return super.equals(obj);
    }
    org.mojolang.mojo.core.DateTime other = (org.mojolang.mojo.core.DateTime) obj;

    if (getYear()
        != other.getYear()) return false;
    if (getMonth()
        != other.getMonth()) return false;
    if (getDay()
        != other.getDay()) return false;
    if (getHour()
        != other.getHour()) return false;
    if (getMinute()
        != other.getMinute()) return false;
    if (getSeconds()
        != other.getSeconds()) return false;
    if (getNanoseconds()
        != other.getNanoseconds()) return false;
    if (hasTimeZone() != other.hasTimeZone()) return false;
    if (hasTimeZone()) {
      if (!getTimeZone()
          .equals(other.getTimeZone())) return false;
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
    hash = (37 * hash) + YEAR_FIELD_NUMBER;
    hash = (53 * hash) + getYear();
    hash = (37 * hash) + MONTH_FIELD_NUMBER;
    hash = (53 * hash) + getMonth();
    hash = (37 * hash) + DAY_FIELD_NUMBER;
    hash = (53 * hash) + getDay();
    hash = (37 * hash) + HOUR_FIELD_NUMBER;
    hash = (53 * hash) + getHour();
    hash = (37 * hash) + MINUTE_FIELD_NUMBER;
    hash = (53 * hash) + getMinute();
    hash = (37 * hash) + SECONDS_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getSeconds());
    hash = (37 * hash) + NANOSECONDS_FIELD_NUMBER;
    hash = (53 * hash) + getNanoseconds();
    if (hasTimeZone()) {
      hash = (37 * hash) + TIME_ZONE_FIELD_NUMBER;
      hash = (53 * hash) + getTimeZone().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static org.mojolang.mojo.core.DateTime parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.DateTime parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.DateTime parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static org.mojolang.mojo.core.DateTime parseFrom(
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
  public static Builder newBuilder(org.mojolang.mojo.core.DateTime prototype) {
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
   * Protobuf type {@code mojo.core.DateTime}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:mojo.core.DateTime)
      org.mojolang.mojo.core.DateTimeOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return org.mojolang.mojo.core.TimeProto.internal_static_mojo_core_DateTime_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return org.mojolang.mojo.core.TimeProto.internal_static_mojo_core_DateTime_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              org.mojolang.mojo.core.DateTime.class, org.mojolang.mojo.core.DateTime.Builder.class);
    }

    // Construct using org.mojolang.mojo.core.DateTime.newBuilder()
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
      year_ = 0;

      month_ = 0;

      day_ = 0;

      hour_ = 0;

      minute_ = 0;

      seconds_ = 0L;

      nanoseconds_ = 0;

      if (timeZoneBuilder_ == null) {
        timeZone_ = null;
      } else {
        timeZone_ = null;
        timeZoneBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return org.mojolang.mojo.core.TimeProto.internal_static_mojo_core_DateTime_descriptor;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.DateTime getDefaultInstanceForType() {
      return org.mojolang.mojo.core.DateTime.getDefaultInstance();
    }

    @java.lang.Override
    public org.mojolang.mojo.core.DateTime build() {
      org.mojolang.mojo.core.DateTime result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public org.mojolang.mojo.core.DateTime buildPartial() {
      org.mojolang.mojo.core.DateTime result = new org.mojolang.mojo.core.DateTime(this);
      result.year_ = year_;
      result.month_ = month_;
      result.day_ = day_;
      result.hour_ = hour_;
      result.minute_ = minute_;
      result.seconds_ = seconds_;
      result.nanoseconds_ = nanoseconds_;
      if (timeZoneBuilder_ == null) {
        result.timeZone_ = timeZone_;
      } else {
        result.timeZone_ = timeZoneBuilder_.build();
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
      if (other instanceof org.mojolang.mojo.core.DateTime) {
        return mergeFrom((org.mojolang.mojo.core.DateTime)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(org.mojolang.mojo.core.DateTime other) {
      if (other == org.mojolang.mojo.core.DateTime.getDefaultInstance()) return this;
      if (other.getYear() != 0) {
        setYear(other.getYear());
      }
      if (other.getMonth() != 0) {
        setMonth(other.getMonth());
      }
      if (other.getDay() != 0) {
        setDay(other.getDay());
      }
      if (other.getHour() != 0) {
        setHour(other.getHour());
      }
      if (other.getMinute() != 0) {
        setMinute(other.getMinute());
      }
      if (other.getSeconds() != 0L) {
        setSeconds(other.getSeconds());
      }
      if (other.getNanoseconds() != 0) {
        setNanoseconds(other.getNanoseconds());
      }
      if (other.hasTimeZone()) {
        mergeTimeZone(other.getTimeZone());
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
      org.mojolang.mojo.core.DateTime parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (org.mojolang.mojo.core.DateTime) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private int year_ ;
    /**
     * <code>int32 year = 1;</code>
     * @return The year.
     */
    @java.lang.Override
    public int getYear() {
      return year_;
    }
    /**
     * <code>int32 year = 1;</code>
     * @param value The year to set.
     * @return This builder for chaining.
     */
    public Builder setYear(int value) {
      
      year_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 year = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearYear() {
      
      year_ = 0;
      onChanged();
      return this;
    }

    private int month_ ;
    /**
     * <code>int32 month = 2;</code>
     * @return The month.
     */
    @java.lang.Override
    public int getMonth() {
      return month_;
    }
    /**
     * <code>int32 month = 2;</code>
     * @param value The month to set.
     * @return This builder for chaining.
     */
    public Builder setMonth(int value) {
      
      month_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 month = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearMonth() {
      
      month_ = 0;
      onChanged();
      return this;
    }

    private int day_ ;
    /**
     * <code>int32 day = 3;</code>
     * @return The day.
     */
    @java.lang.Override
    public int getDay() {
      return day_;
    }
    /**
     * <code>int32 day = 3;</code>
     * @param value The day to set.
     * @return This builder for chaining.
     */
    public Builder setDay(int value) {
      
      day_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 day = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearDay() {
      
      day_ = 0;
      onChanged();
      return this;
    }

    private int hour_ ;
    /**
     * <code>int32 hour = 4;</code>
     * @return The hour.
     */
    @java.lang.Override
    public int getHour() {
      return hour_;
    }
    /**
     * <code>int32 hour = 4;</code>
     * @param value The hour to set.
     * @return This builder for chaining.
     */
    public Builder setHour(int value) {
      
      hour_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 hour = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearHour() {
      
      hour_ = 0;
      onChanged();
      return this;
    }

    private int minute_ ;
    /**
     * <code>int32 minute = 5;</code>
     * @return The minute.
     */
    @java.lang.Override
    public int getMinute() {
      return minute_;
    }
    /**
     * <code>int32 minute = 5;</code>
     * @param value The minute to set.
     * @return This builder for chaining.
     */
    public Builder setMinute(int value) {
      
      minute_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 minute = 5;</code>
     * @return This builder for chaining.
     */
    public Builder clearMinute() {
      
      minute_ = 0;
      onChanged();
      return this;
    }

    private long seconds_ ;
    /**
     * <code>int64 seconds = 6;</code>
     * @return The seconds.
     */
    @java.lang.Override
    public long getSeconds() {
      return seconds_;
    }
    /**
     * <code>int64 seconds = 6;</code>
     * @param value The seconds to set.
     * @return This builder for chaining.
     */
    public Builder setSeconds(long value) {
      
      seconds_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 seconds = 6;</code>
     * @return This builder for chaining.
     */
    public Builder clearSeconds() {
      
      seconds_ = 0L;
      onChanged();
      return this;
    }

    private int nanoseconds_ ;
    /**
     * <code>int32 nanoseconds = 7;</code>
     * @return The nanoseconds.
     */
    @java.lang.Override
    public int getNanoseconds() {
      return nanoseconds_;
    }
    /**
     * <code>int32 nanoseconds = 7;</code>
     * @param value The nanoseconds to set.
     * @return This builder for chaining.
     */
    public Builder setNanoseconds(int value) {
      
      nanoseconds_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 nanoseconds = 7;</code>
     * @return This builder for chaining.
     */
    public Builder clearNanoseconds() {
      
      nanoseconds_ = 0;
      onChanged();
      return this;
    }

    private org.mojolang.mojo.core.TimeZone timeZone_;
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.TimeZone, org.mojolang.mojo.core.TimeZone.Builder, org.mojolang.mojo.core.TimeZoneOrBuilder> timeZoneBuilder_;
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     * @return Whether the timeZone field is set.
     */
    public boolean hasTimeZone() {
      return timeZoneBuilder_ != null || timeZone_ != null;
    }
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     * @return The timeZone.
     */
    public org.mojolang.mojo.core.TimeZone getTimeZone() {
      if (timeZoneBuilder_ == null) {
        return timeZone_ == null ? org.mojolang.mojo.core.TimeZone.getDefaultInstance() : timeZone_;
      } else {
        return timeZoneBuilder_.getMessage();
      }
    }
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     */
    public Builder setTimeZone(org.mojolang.mojo.core.TimeZone value) {
      if (timeZoneBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        timeZone_ = value;
        onChanged();
      } else {
        timeZoneBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     */
    public Builder setTimeZone(
        org.mojolang.mojo.core.TimeZone.Builder builderForValue) {
      if (timeZoneBuilder_ == null) {
        timeZone_ = builderForValue.build();
        onChanged();
      } else {
        timeZoneBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     */
    public Builder mergeTimeZone(org.mojolang.mojo.core.TimeZone value) {
      if (timeZoneBuilder_ == null) {
        if (timeZone_ != null) {
          timeZone_ =
            org.mojolang.mojo.core.TimeZone.newBuilder(timeZone_).mergeFrom(value).buildPartial();
        } else {
          timeZone_ = value;
        }
        onChanged();
      } else {
        timeZoneBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     */
    public Builder clearTimeZone() {
      if (timeZoneBuilder_ == null) {
        timeZone_ = null;
        onChanged();
      } else {
        timeZone_ = null;
        timeZoneBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     */
    public org.mojolang.mojo.core.TimeZone.Builder getTimeZoneBuilder() {
      
      onChanged();
      return getTimeZoneFieldBuilder().getBuilder();
    }
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     */
    public org.mojolang.mojo.core.TimeZoneOrBuilder getTimeZoneOrBuilder() {
      if (timeZoneBuilder_ != null) {
        return timeZoneBuilder_.getMessageOrBuilder();
      } else {
        return timeZone_ == null ?
            org.mojolang.mojo.core.TimeZone.getDefaultInstance() : timeZone_;
      }
    }
    /**
     * <code>.mojo.core.TimeZone time_zone = 15;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        org.mojolang.mojo.core.TimeZone, org.mojolang.mojo.core.TimeZone.Builder, org.mojolang.mojo.core.TimeZoneOrBuilder> 
        getTimeZoneFieldBuilder() {
      if (timeZoneBuilder_ == null) {
        timeZoneBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            org.mojolang.mojo.core.TimeZone, org.mojolang.mojo.core.TimeZone.Builder, org.mojolang.mojo.core.TimeZoneOrBuilder>(
                getTimeZone(),
                getParentForChildren(),
                isClean());
        timeZone_ = null;
      }
      return timeZoneBuilder_;
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


    // @@protoc_insertion_point(builder_scope:mojo.core.DateTime)
  }

  // @@protoc_insertion_point(class_scope:mojo.core.DateTime)
  private static final org.mojolang.mojo.core.DateTime DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new org.mojolang.mojo.core.DateTime();
  }

  public static org.mojolang.mojo.core.DateTime getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<DateTime>
      PARSER = new com.google.protobuf.AbstractParser<DateTime>() {
    @java.lang.Override
    public DateTime parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new DateTime(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<DateTime> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<DateTime> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public org.mojolang.mojo.core.DateTime getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

