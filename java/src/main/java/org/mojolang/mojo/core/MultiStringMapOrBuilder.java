// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/boxed.proto

package org.mojolang.mojo.core;

public interface MultiStringMapOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.MultiStringMap)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>map&lt;string, .mojo.core.MultiStringMap.StringValues&gt; vals = 1;</code>
   */
  int getValsCount();
  /**
   * <code>map&lt;string, .mojo.core.MultiStringMap.StringValues&gt; vals = 1;</code>
   */
  boolean containsVals(
      java.lang.String key);
  /**
   * Use {@link #getValsMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.String, org.mojolang.mojo.core.MultiStringMap.StringValues>
  getVals();
  /**
   * <code>map&lt;string, .mojo.core.MultiStringMap.StringValues&gt; vals = 1;</code>
   */
  java.util.Map<java.lang.String, org.mojolang.mojo.core.MultiStringMap.StringValues>
  getValsMap();
  /**
   * <code>map&lt;string, .mojo.core.MultiStringMap.StringValues&gt; vals = 1;</code>
   */

  org.mojolang.mojo.core.MultiStringMap.StringValues getValsOrDefault(
      java.lang.String key,
      org.mojolang.mojo.core.MultiStringMap.StringValues defaultValue);
  /**
   * <code>map&lt;string, .mojo.core.MultiStringMap.StringValues&gt; vals = 1;</code>
   */

  org.mojolang.mojo.core.MultiStringMap.StringValues getValsOrThrow(
      java.lang.String key);
}