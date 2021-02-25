// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: mojo/core/phone_number.proto

package org.mojolang.mojo.core;

public interface PhoneNumberOrBuilder extends
    // @@protoc_insertion_point(interface_extends:mojo.core.PhoneNumber)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int32 country_code = 1;</code>
   * @return The countryCode.
   */
  int getCountryCode();

  /**
   * <code>uint64 national_number = 2;</code>
   * @return The nationalNumber.
   */
  long getNationalNumber();

  /**
   * <code>string extension = 3;</code>
   * @return The extension.
   */
  java.lang.String getExtension();
  /**
   * <code>string extension = 3;</code>
   * @return The bytes for extension.
   */
  com.google.protobuf.ByteString
      getExtensionBytes();

  /**
   * <code>bool italian_leading_zero = 4;</code>
   * @return The italianLeadingZero.
   */
  boolean getItalianLeadingZero();

  /**
   * <code>int32 number_of_leading_zeros = 8;</code>
   * @return The numberOfLeadingZeros.
   */
  int getNumberOfLeadingZeros();

  /**
   * <code>string raw_input = 5;</code>
   * @return The rawInput.
   */
  java.lang.String getRawInput();
  /**
   * <code>string raw_input = 5;</code>
   * @return The bytes for rawInput.
   */
  com.google.protobuf.ByteString
      getRawInputBytes();

  /**
   * <code>.mojo.core.PhoneNumber.CountryCodeSource country_code_source = 6;</code>
   * @return The enum numeric value on the wire for countryCodeSource.
   */
  int getCountryCodeSourceValue();
  /**
   * <code>.mojo.core.PhoneNumber.CountryCodeSource country_code_source = 6;</code>
   * @return The countryCodeSource.
   */
  org.mojolang.mojo.core.PhoneNumber.CountryCodeSource getCountryCodeSource();

  /**
   * <code>string preferred_domestic_carrier_code = 7;</code>
   * @return The preferredDomesticCarrierCode.
   */
  java.lang.String getPreferredDomesticCarrierCode();
  /**
   * <code>string preferred_domestic_carrier_code = 7;</code>
   * @return The bytes for preferredDomesticCarrierCode.
   */
  com.google.protobuf.ByteString
      getPreferredDomesticCarrierCodeBytes();
}
