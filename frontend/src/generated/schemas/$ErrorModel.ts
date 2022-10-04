/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $ErrorModel = {
    properties: {
        detail: {
            type: 'string',
            description: `A human-readable explanation specific to this occurrence of the problem.`,
        },
        errors: {
            type: 'array',
            contains: {
                properties: {
                    location: {
                        type: 'string',
                        description: `Where the error occured, e.g. 'body.items[3].tags' or 'path.thing-id'`,
                    },
                    message: {
                        type: 'string',
                        description: `Error message text`,
                    },
                    value: {
                        description: `The value at the given location`,
                        properties: {
                        },
                    },
                },
            },
        },
        instance: {
            type: 'string',
            description: `A URI reference that identifies the specific occurence of the problem.`,
            format: 'uri',
        },
        status: {
            type: 'number',
            description: `HTTP status code`,
            format: 'int32',
        },
        title: {
            type: 'string',
            description: `A short, human-readable summary of the problem type. This value should not change between occurances of the error.`,
        },
        type: {
            type: 'string',
            description: `A URI reference to human-readable documentation for the error.`,
            format: 'uri',
        },
    },
} as const;
