/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
export const $App = {
    properties: {
        appname: {
            type: 'string',
            description: `Application Name`,
            isRequired: true,
            pattern: '^[a-z0-9]+$',
        },
        associatedusers: {
            type: 'array',
            contains: {
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
            },
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
        description: {
            type: 'string',
            description: `Description of Application`,
            isRequired: true,
        },
        filters: {
            type: 'array',
            contains: {
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
                    id: {
                        type: 'number',
                        isRequired: true,
                        format: 'int32',
                    },
                    name: {
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
            isRequired: true,
        },
        icon: {
            type: 'string',
            description: `Icon of Application`,
            isRequired: true,
        },
        id: {
            type: 'number',
            isRequired: true,
            format: 'int32',
        },
        status: {
            type: 'Enum',
            isRequired: true,
        },
        updatedat: {
            type: 'string',
            isRequired: true,
            format: 'date-time',
        },
        url: {
            type: 'string',
            description: `URL of Application`,
            isRequired: true,
        },
    },
} as const;
