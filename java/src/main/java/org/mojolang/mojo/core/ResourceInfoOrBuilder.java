// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/error_details.proto

package org.mojolang.mojo.core;

public interface ResourceInfoOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.ResourceInfo)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string resource_type = 1;</code>
   * @return The resourceType.
   */
  java.lang.String getResourceType();
  /**
   * <code>string resource_type = 1;</code>
   * @return The bytes for resourceType.
   */
  com.google.protobuf.ByteString
      getResourceTypeBytes();

  /**
   * <code>string resource_name = 2;</code>
   * @return The resourceName.
   */
  java.lang.String getResourceName();
  /**
   * <code>string resource_name = 2;</code>
   * @return The bytes for resourceName.
   */
  com.google.protobuf.ByteString
      getResourceNameBytes();

  /**
   * <code>string owner = 3;</code>
   * @return The owner.
   */
  java.lang.String getOwner();
  /**
   * <code>string owner = 3;</code>
   * @return The bytes for owner.
   */
  com.google.protobuf.ByteString
      getOwnerBytes();

  /**
   * <code>string description = 4;</code>
   * @return The description.
   */
  java.lang.String getDescription();
  /**
   * <code>string description = 4;</code>
   * @return The bytes for description.
   */
  com.google.protobuf.ByteString
      getDescriptionBytes();
}