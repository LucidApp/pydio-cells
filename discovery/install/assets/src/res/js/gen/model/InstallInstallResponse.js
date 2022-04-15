/**
 * Pydio Cells Rest API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 4.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The InstallInstallResponse model module.
 * @module model/InstallInstallResponse
 * @version 4.0
 */
class InstallInstallResponse {
    /**
     * Constructs a new <code>InstallInstallResponse</code>.
     * @alias module:model/InstallInstallResponse
     */
    constructor() { 
        
        InstallInstallResponse.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>InstallInstallResponse</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstallInstallResponse} obj Optional instance to populate.
     * @return {module:model/InstallInstallResponse} The populated <code>InstallInstallResponse</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new InstallInstallResponse();

            if (data.hasOwnProperty('success')) {
                obj['success'] = ApiClient.convertToType(data['success'], 'Boolean');
            }
        }
        return obj;
    }


}

/**
 * @member {Boolean} success
 */
InstallInstallResponse.prototype['success'] = undefined;






export default InstallInstallResponse;

