/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $post_message_request = {
    properties: {
        $schema: {
            type: 'string',
            description: `An optional URL to a JSON Schema document describing this resource`,
            format: 'uri',
        },
        fields: {
            type: 'dictionary',
            contains: {
                properties: {
                },
            },
        },
        message: {
            type: 'string',
            description: `Message to be Sent`,
            isRequired: true,
        },
        severity: {
            type: 'string',
            description: `Severity of Message`,
        },
        shortmessage: {
            type: 'string',
            description: `Short Message to be Sent`,
        },
        timestamp: {
            type: 'string',
            description: `Timestamp of Message`,
            format: 'date-time',
        },
        topic: {
            type: 'string',
            description: `Topic of Message`,
        },
    },
} as const;
