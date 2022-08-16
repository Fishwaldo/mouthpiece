/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $MessageResult = {
    properties: {
        $schema: {
            type: 'string',
            description: `An optional URL to a JSON Schema document describing this resource`,
            format: 'uri',
        },
        message_id: {
            type: 'number',
            description: `Message ID`,
            isRequired: true,
            format: 'int32',
        },
        status: {
            type: 'string',
            description: `Status of Message`,
            isRequired: true,
        },
    },
} as const;
