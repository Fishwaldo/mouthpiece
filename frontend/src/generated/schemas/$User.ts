/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $User = {
    properties: {
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
        email: {
            type: 'string',
            isRequired: true,
        },
        firstname: {
            type: 'string',
            isRequired: true,
        },
        id: {
            type: 'number',
            isRequired: true,
            format: 'int32',
        },
        lastname: {
            type: 'string',
            isRequired: true,
        },
        transports: {
            type: 'array',
            contains: {
                properties: {
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
            },
        },
        updatedat: {
            type: 'string',
            isRequired: true,
            format: 'date-time',
        },
    },
} as const;
