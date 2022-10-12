/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $meResponse = {
    properties: {
        email: {
            type: 'string',
            isRequired: true,
        },
        id: {
            type: 'number',
            isRequired: true,
            format: 'int32',
        },
        name: {
            type: 'string',
            isRequired: true,
        },
        status: {
            type: 'string',
            isRequired: true,
        },
    },
} as const;
