// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/error.proto

package org.mojolang.mojo.core;

public interface ErrorOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.Error)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.mojo.core.ErrorCode code = 1;</code>
   * @return Whether the code field is set.
   */
  boolean hasCode();
  /**
   * <code>.mojo.core.ErrorCode code = 1;</code>
   * @return The code.
   */
  org.mojolang.mojo.core.ErrorCode getCode();
  /**
   * <code>.mojo.core.ErrorCode code = 1;</code>
   */
  org.mojolang.mojo.core.ErrorCodeOrBuilder getCodeOrBuilder();

  /**
   * <code>string message = 4;</code>
   * @return The message.
   */
  java.lang.String getMessage();
  /**
   * <code>string message = 4;</code>
   * @return The bytes for message.
   */
  com.google.protobuf.ByteString
      getMessageBytes();

  /**
   * <code>.mojo.core.Url document = 5;</code>
   * @return Whether the document field is set.
   */
  boolean hasDocument();
  /**
   * <code>.mojo.core.Url document = 5;</code>
   * @return The document.
   */
  org.mojolang.mojo.core.Url getDocument();
  /**
   * <code>.mojo.core.Url document = 5;</code>
   */
  org.mojolang.mojo.core.UrlOrBuilder getDocumentOrBuilder();

  /**
   * <code>repeated .mojo.core.Error causes = 10;</code>
   */
  java.util.List<org.mojolang.mojo.core.Error> 
      getCausesList();
  /**
   * <code>repeated .mojo.core.Error causes = 10;</code>
   */
  org.mojolang.mojo.core.Error getCauses(int index);
  /**
   * <code>repeated .mojo.core.Error causes = 10;</code>
   */
  int getCausesCount();
  /**
   * <code>repeated .mojo.core.Error causes = 10;</code>
   */
  java.util.List<? extends org.mojolang.mojo.core.ErrorOrBuilder> 
      getCausesOrBuilderList();
  /**
   * <code>repeated .mojo.core.Error causes = 10;</code>
   */
  org.mojolang.mojo.core.ErrorOrBuilder getCausesOrBuilder(
      int index);
}
