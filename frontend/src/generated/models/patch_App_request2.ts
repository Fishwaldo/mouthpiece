/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type patch_App_request2 = Array<{
    /**
     * JSON Pointer for the source of a move or copy
     */
    from?: string;
    /**
     * Operation name
     */
    op: 'add' | 'remove' | 'replace' | 'move' | 'copy' | 'test';
    /**
     * JSON Pointer to the field being operated on, or the destination of a move/copy operation
     */
    path: string;
    /**
     * The value to set
     */
    value?: any;
}>;
