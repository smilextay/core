// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/value.proto

package org.mojolang.mojo.core;

public interface ValueOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.Value)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.core.Values values_val = 1;</code>
   * @return Whether the valuesVal field is set.
   */
  boolean hasValuesVal();
  /**
   * <code>.mojo.core.Values values_val = 1;</code>
   * @return The valuesVal.
   */
  org.mojolang.mojo.core.Values getValuesVal();
  /**
   * <code>.mojo.core.Values values_val = 1;</code>
   */
  org.mojolang.mojo.core.ValuesOrBuilder getValuesValOrBuilder();

  /**
   * <code>.mojo.core.Object object_val = 2;</code>
   * @return Whether the objectVal field is set.
   */
  boolean hasObjectVal();
  /**
   * <code>.mojo.core.Object object_val = 2;</code>
   * @return The objectVal.
   */
  org.mojolang.mojo.core.Object getObjectVal();
  /**
   * <code>.mojo.core.Object object_val = 2;</code>
   */
  org.mojolang.mojo.core.ObjectOrBuilder getObjectValOrBuilder();

  /**
   * <code>bool bool_val = 3;</code>
   * @return Whether the boolVal field is set.
   */
  boolean hasBoolVal();
  /**
   * <code>bool bool_val = 3;</code>
   * @return The boolVal.
   */
  boolean getBoolVal();

  /**
   * <code>int64 int64_val = 4;</code>
   * @return Whether the int64Val field is set.
   */
  boolean hasInt64Val();
  /**
   * <code>int64 int64_val = 4;</code>
   * @return The int64Val.
   */
  long getInt64Val();

  /**
   * <code>double double_val = 5;</code>
   * @return Whether the doubleVal field is set.
   */
  boolean hasDoubleVal();
  /**
   * <code>double double_val = 5;</code>
   * @return The doubleVal.
   */
  double getDoubleVal();

  /**
   * <code>string string_val = 7;</code>
   * @return Whether the stringVal field is set.
   */
  boolean hasStringVal();
  /**
   * <code>string string_val = 7;</code>
   * @return The stringVal.
   */
  java.lang.String getStringVal();
  /**
   * <code>string string_val = 7;</code>
   * @return The bytes for stringVal.
   */
  com.google.protobuf.ByteString
      getStringValBytes();

  public org.mojolang.mojo.core.Value.ValueCase getValueCase();
}
