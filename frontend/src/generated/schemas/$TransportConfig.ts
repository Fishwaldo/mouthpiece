/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $TransportConfig = {
    properties: {
        $schema: {
            type: 'string',
            description: `An optional URL to a JSON Schema document describing this resource`,
            format: 'uri',
        },
        config: {
            type: 'string',
            isRequired: true,
        },
        createdat: {
            type: 'string',
            isRequired: true,
            format: 'date-time',
        },
        deletedat: {
            properties: {
                time: {
                    type: 'string',
                    isRequired: true,
                    format: 'date-time',
                },
                valid: {
                    type: 'boolean',
                    isRequired: true,
                },
            },
            isRequired: true,
        },
        id: {
            type: 'number',
            isRequired: true,
            format: 'int32',
        },
        transport: {
            type: 'string',
            isRequired: true,
        },
        updatedat: {
            type: 'string',
            isRequired: true,
            format: 'date-time',
        },
    },
} as const;
