/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type post_message_request = {
    /**
     * An optional URL to a JSON Schema document describing this resource
     */
    $schema?: string;
    /**
     * Additional Fields
     */
    fields?: Record<string, any>;
    /**
     * Message to be Sent
     */
    message: string;
    /**
     * Severity of Message
     */
    severity?: string;
    /**
     * Short Message to be Sent
     */
    shortmessage?: string;
    /**
     * Timestamp of Message
     */
    timestamp?: string;
    /**
     * Topic of Message
     */
    topic?: string;
};

