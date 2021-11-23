// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/error_details.proto

package org.mojolang.mojo.core;

public interface DebugInfoOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.DebugInfo)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated string stack_entries = 1;</code>
   * @return A list containing the stackEntries.
   */
  java.util.List<java.lang.String>
      getStackEntriesList();
  /**
   * <code>repeated string stack_entries = 1;</code>
   * @return The count of stackEntries.
   */
  int getStackEntriesCount();
  /**
   * <code>repeated string stack_entries = 1;</code>
   * @param index The index of the element to return.
   * @return The stackEntries at the given index.
   */
  java.lang.String getStackEntries(int index);
  /**
   * <code>repeated string stack_entries = 1;</code>
   * @param index The index of the value to return.
   * @return The bytes of the stackEntries at the given index.
   */
  com.google.protobuf.ByteString
      getStackEntriesBytes(int index);

  /**
   * <code>string detail = 2;</code>
   * @return The detail.
   */
  java.lang.String getDetail();
  /**
   * <code>string detail = 2;</code>
   * @return The bytes for detail.
   */
  com.google.protobuf.ByteString
      getDetailBytes();
}
