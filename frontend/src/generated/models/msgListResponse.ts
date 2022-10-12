/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type msgListResponse = {
    count: number;
    data: Array<{
        appid: number;
        fields: Record<string, string>;
        id: string;
        message: string;
        severity: number;
        shortmsg: string;
        timestamp: string;
        topic: string;
    }>;
};

