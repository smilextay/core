// Copyright 2020 Google LLC
///
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
///
//     http://www.apache.org/licenses/LICENSE-2.0
///
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


/// Describes when the clients can retry a failed request. Clients could ignore
/// the recommendation here or retry when this information is missing from error
/// responses.
///
/// It's always recommended that clients should use exponential backoff when
/// retrying.
///
/// Clients should wait until `retry_delay` amount of time has passed since
/// receiving the error response before retrying.  If retrying requests also
/// fail, clients should use an exponential backoff scheme to gradually increase
/// the delay between retries based on `retry_delay`, until either a maximum
/// number of retries have been reached or a maximum retry delay cap has been
/// reached.
type RetryInfo {
    /// Clients should wait at least this long between retrying the same request.
    retry_delay: Duration @1
}

/// Describes additional debugging info.
type DebugInfo {
    /// The stack trace entries indicating where the error occurred.
    stack_entries: [String] @1

    /// Additional debugging information provided by the server.
    detail: String  @2
}

/// Describes how a quota check failed.
///
/// For example if a daily limit was exceeded for the calling project,
/// a service could respond with a QuotaFailure detail containing the project
/// id and the description of the quota limit that was exceeded.  If the
/// calling project hasn't enabled the service in the developer console, then
/// a service could respond with the project id and set `service_disabled`
/// to true.
///
/// Also see RetryInfo and Help types for other details about handling a
/// quota failure.
type QuotaFailure {
    /// A type type used to describe a single quota violation.  For example, a
    /// daily quota or a custom quota that was exceeded.
    type Violation {
        /// The subject on which the quota check failed.
        /// For example, "clientip:<ip address of client>" or "project:<Google
        /// developer project id>".
        subject: String  @1

        /// A description of how the quota check failed. Clients can use this
        /// description to find more about the quota configuration in the service's
        /// public documentation, or find the relevant quota limit to adjust through
        /// developer console.
        ///
        /// For example: "Service disabled" or "Daily Limit for read operations
        /// exceeded".
        description: String  @2
    }

    /// Describes all quota violations.
    violations: [Violation] @1
}

/// Describes the cause of the error with structured details.
///
/// Example of an error when contacting the "pubsub.googleapis.com" API when it
/// is not enabled:
///
///     { "reason": "API_DISABLED"
///       "domain": "googleapis.com"
///       "metadata": {
///         "resource": "projects/123",
///         "service": "pubsub.googleapis.com"
///       }
///     }
///
/// This response indicates that the pubsub.googleapis.com API is not enabled.
///
/// Example of an error that is returned when attempting to create a Spanner
/// instance in a region that is out of stock:
///
///     { "reason": "STOCKOUT"
///       "domain": "spanner.googleapis.com",
///       "metadata": {
///         "availableRegions": "us-central1,us-east2"
///       }
///     }
type ErrorInfo {
    /// The reason of the error. This is a constant value that identifies the
    /// proximate cause of the error. Error reasons are unique within a particular
    /// domain of errors. This should be at most 63 characters and match
    /// /[A-Z0-9_]+/.
    reason: String @1

    /// The logical grouping to which the "reason" belongs. The error domain
    /// is typically the registered service name of the tool or product that
    /// generates the error. Example: "pubsub.googleapis.com". If the error is
    /// generated by some common infrastructure, the error domain must be a
    /// globally unique value that identifies the infrastructure. For Google API
    /// infrastructure, the error domain is "googleapis.com".
    domain: String  @2

    /// Additional structured details about this error.
    ///
    /// Keys should match /[a-zA-Z0-9-_]/ and be limited to 64 characters in
    /// length. When identifying the current value of an exceeded limit, the units
    /// should be contained in the key, not the value.  For example, rather than
    /// {"instanceLimit": "100/request"}, should be returned as,
    /// {"instanceLimitPerRequest": "100"}, if the client exceeds the number of
    /// instances that can be created in a single (batch) request.
    metadata: {String: String} @3
}

/// Describes what preconditions have failed.
////
/// For example, if an RPC failed because it required the Terms of Service to be
/// acknowledged, it could list the terms of service violation in the
/// PreconditionFailure type.
type PreconditionFailure {
    /// A type type used to describe a single precondition failure.
    type Violation {
        /// The type of PreconditionFailure. We recommend using a service-specific
        /// enum type to define the supported precondition violation subjects. For
        /// example, "TOS" for "Terms of Service violation".
        type: String  @1

        /// The subject, relative to the type, that failed.
        /// For example, "google.com/cloud" relative to the "TOS" type would indicate
        /// which terms of service is being referenced.
        subject: String  @2

        /// A description of how the precondition failed. Developers can use this
        /// description to understand how to fix the failure.
        ///
        /// For example: "Terms of service not accepted".
        description: String  @3
    }

    /// Describes all precondition violations.
    violations: [Violation] @1
}

/// Describes violations in a client request. This error type focuses on the
/// syntactic aspects of the request.
type BadRequest {
    /// A type type used to describe a single bad request field.
    type FieldViolation {
        /// A path leading to a field in the request body. The value will be a
        /// sequence of dot-separated identifiers that identify a protocol buffer
        /// field. E.g., "field_violations.field" would identify this field.
        field: String  @1

        /// A description of why the request element is bad.
        description: String  @2
    }

    /// Describes all violations in a client request.
    field_violations: [FieldViolation] @1
}

/// Contains metadata about the request that clients can attach when filing a bug
/// or providing other forms of feedback.
type RequestInfo {
    /// An opaque that: String  should only be interpreted by the service generating
    /// it. For example, it can be used to identify requests in the service's logs.
    request_id: String @1

    /// Any data that was used to serve this request. For example, an encrypted
    /// stack trace that can be sent back to the service provider for debugging.
    serving_data: String @2
}

/// Describes the resource that is being accessed.
type ResourceInfo {
    /// A name for the type of resource being accessed, e.g. "sql table",
    /// "cloud storage bucket", "file", "Google calendar"; or the type URL
    /// of the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic".
    resource_type: String @1

    /// The name of the resource being accessed.  For example, a shared calendar
    /// name: "example.com_4fghdhgsrgh@group.calendar.google.com", if the current
    /// error is [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED].
    resource_name: String @2

    /// The owner of the resource (optional).
    /// For example, "user:<owner email>" or "project:<Google developer project
    /// id>".
    owner: String  @3

    /// Describes what error is encountered when accessing this resource.
    /// For example, updating a cloud project may require the `writer` permission
    /// on the developer console project.
    description: String  @4
}

/// Provides links to documentation or for performing an out of band action.
///
/// For example, if a quota check failed with an error indicating the calling
/// project hasn't enabled the accessed service, this can contain a URL pointing
/// directly to the right place in the developer console to flip the bit.
type Help {
    /// Describes a URL link.
    type Link {
        /// Describes what the link offers.
        description: String  @1

        /// The URL of the link.
        url: Url  @2
    }

    /// URL(s) pointing to additional information on handling the current error.
    links: [Link] @1
}

/// Provides a localized error type that is safe to return to the user
/// which can be attached to an RPC error.
type LocalizedMessage {
    /// The locale used following the specification defined at
    /// http://www.rfc-editor.org/rfc/bcp/bcp47.txt.
    /// Examples are: "en-US", "fr-CH", "es-MX"
    locale: String  @1

    /// The localized error type in the above locale.
    message: String  @2
}
