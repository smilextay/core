// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/error_details.proto

package org.mojolang.mojo.core;

public interface RetryInfoOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.RetryInfo)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.core.Duration retry_delay = 1;</code>
   * @return Whether the retryDelay field is set.
   */
  boolean hasRetryDelay();
  /**
   * <code>.mojo.core.Duration retry_delay = 1;</code>
   * @return The retryDelay.
   */
  org.mojolang.mojo.core.Duration getRetryDelay();
  /**
   * <code>.mojo.core.Duration retry_delay = 1;</code>
   */
  org.mojolang.mojo.core.DurationOrBuilder getRetryDelayOrBuilder();
}