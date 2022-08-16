/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $CheckerResult = {
    properties: {
        $schema: {
            type: 'string',
            description: `An optional URL to a JSON Schema document describing this resource`,
            format: 'uri',
        },
        details: {
            type: 'dictionary',
            contains: {
                properties: {
                    error: {
                        type: 'string',
                    },
                    status: {
                        type: 'string',
                        isRequired: true,
                    },
                    timestamp: {
                        type: 'string',
                        format: 'date-time',
                    },
                },
            },
        },
        status: {
            type: 'string',
            isRequired: true,
        },
    },
} as const;
