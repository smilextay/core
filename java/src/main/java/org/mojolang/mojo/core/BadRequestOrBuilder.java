// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/error_details.proto

package org.mojolang.mojo.core;

public interface BadRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.BadRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>repeated .mojo.core.BadRequest.FieldViolation field_violations = 1;</code>
   */
  java.util.List<org.mojolang.mojo.core.BadRequest.FieldViolation> 
      getFieldViolationsList();
  /**
   * <code>repeated .mojo.core.BadRequest.FieldViolation field_violations = 1;</code>
   */
  org.mojolang.mojo.core.BadRequest.FieldViolation getFieldViolations(int index);
  /**
   * <code>repeated .mojo.core.BadRequest.FieldViolation field_violations = 1;</code>
   */
  int getFieldViolationsCount();
  /**
   * <code>repeated .mojo.core.BadRequest.FieldViolation field_violations = 1;</code>
   */
  java.util.List<? extends org.mojolang.mojo.core.BadRequest.FieldViolationOrBuilder> 
      getFieldViolationsOrBuilderList();
  /**
   * <code>repeated .mojo.core.BadRequest.FieldViolation field_violations = 1;</code>
   */
  org.mojolang.mojo.core.BadRequest.FieldViolationOrBuilder getFieldViolationsOrBuilder(
      int index);
}
