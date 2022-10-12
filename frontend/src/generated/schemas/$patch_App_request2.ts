/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $patch_App_request2 = {
    type: 'array',
    contains: {
        properties: {
            from: {
                type: 'string',
                description: `JSON Pointer for the source of a move or copy`,
            },
            op: {
                type: 'Enum',
                isRequired: true,
            },
            path: {
                type: 'string',
                description: `JSON Pointer to the field being operated on, or the destination of a move/copy operation`,
                isRequired: true,
            },
            value: {
                description: `The value to set`,
                properties: {
                },
            },
        },
    },
} as const;
