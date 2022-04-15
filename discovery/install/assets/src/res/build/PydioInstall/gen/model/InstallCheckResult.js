'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }(); /**
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

var _ApiClient = require('../ApiClient');

var _ApiClient2 = _interopRequireDefault(_ApiClient);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

/**
 * The InstallCheckResult model module.
 * @module model/InstallCheckResult
 * @version 4.0
 */
var InstallCheckResult = function () {
  /**
   * Constructs a new <code>InstallCheckResult</code>.
   * @alias module:model/InstallCheckResult
   */
  function InstallCheckResult() {
    _classCallCheck(this, InstallCheckResult);

    InstallCheckResult.initialize(this);
  }

  /**
   * Initializes the fields of this object.
   * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
   * Only for internal use.
   */


  _createClass(InstallCheckResult, null, [{
    key: 'initialize',
    value: function initialize(obj) {}

    /**
     * Constructs a <code>InstallCheckResult</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/InstallCheckResult} obj Optional instance to populate.
     * @return {module:model/InstallCheckResult} The populated <code>InstallCheckResult</code> instance.
     */

  }, {
    key: 'constructFromObject',
    value: function constructFromObject(data, obj) {
      if (data) {
        obj = obj || new InstallCheckResult();

        if (data.hasOwnProperty('JsonResult')) {
          obj['JsonResult'] = _ApiClient2.default.convertToType(data['JsonResult'], 'String');
        }
        if (data.hasOwnProperty('Name')) {
          obj['Name'] = _ApiClient2.default.convertToType(data['Name'], 'String');
        }
        if (data.hasOwnProperty('Success')) {
          obj['Success'] = _ApiClient2.default.convertToType(data['Success'], 'Boolean');
        }
      }
      return obj;
    }
  }]);

  return InstallCheckResult;
}();

/**
 * @member {String} JsonResult
 */


InstallCheckResult.prototype['JsonResult'] = undefined;

/**
 * @member {String} Name
 */
InstallCheckResult.prototype['Name'] = undefined;

/**
 * @member {Boolean} Success
 */
InstallCheckResult.prototype['Success'] = undefined;

exports.default = InstallCheckResult;
