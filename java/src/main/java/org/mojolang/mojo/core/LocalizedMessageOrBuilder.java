// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/error_details.proto

package org.mojolang.mojo.core;

public interface LocalizedMessageOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.LocalizedMessage)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string locale = 1;</code>
   * @return The locale.
   */
  java.lang.String getLocale();
  /**
   * <code>string locale = 1;</code>
   * @return The bytes for locale.
   */
  com.google.protobuf.ByteString
      getLocaleBytes();

  /**
   * <code>string message = 2;</code>
   * @return The message.
   */
  java.lang.String getMessage();
  /**
   * <code>string message = 2;</code>
   * @return The bytes for message.
   */
  com.google.protobuf.ByteString
      getMessageBytes();
}