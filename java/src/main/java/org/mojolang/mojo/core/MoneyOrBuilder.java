// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/money.proto

package org.mojolang.mojo.core;

public interface MoneyOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.Money)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string currency_code = 1;</code>
   * @return The currencyCode.
   */
  java.lang.String getCurrencyCode();
  /**
   * <code>string currency_code = 1;</code>
   * @return The bytes for currencyCode.
   */
  com.google.protobuf.ByteString
      getCurrencyCodeBytes();

  /**
   * <code>int64 units = 2;</code>
   * @return The units.
   */
  long getUnits();

  /**
   * <code>int32 nanos = 3;</code>
   * @return The nanos.
   */
  int getNanos();
}
