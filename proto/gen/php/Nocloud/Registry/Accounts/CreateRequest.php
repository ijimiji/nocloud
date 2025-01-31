<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/registry/proto/accounts/accounts.proto

namespace Nocloud\Registry\Accounts;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>nocloud.registry.accounts.CreateRequest</code>
 */
class CreateRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * New Account title
     *
     * Generated from protobuf field <code>string title = 1 [json_name = "title"];</code>
     */
    protected $title = '';
    /**
     * Credentials for New Account
     *
     * Generated from protobuf field <code>.nocloud.registry.accounts.Credentials auth = 2 [json_name = "auth"];</code>
     */
    protected $auth = null;
    /**
     * Namespace to put Account under
     *
     * Generated from protobuf field <code>string namespace = 3 [json_name = "namespace"];</code>
     */
    protected $namespace = '';
    /**
     * Account access level to parent namespace
     *
     * Generated from protobuf field <code>optional int32 access = 4 [json_name = "access"];</code>
     */
    protected $access = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $title
     *           New Account title
     *     @type \Nocloud\Registry\Accounts\Credentials $auth
     *           Credentials for New Account
     *     @type string $namespace
     *           Namespace to put Account under
     *     @type int $access
     *           Account access level to parent namespace
     * }
     */
    public function __construct($data = NULL) {
        \Nocloud\Registry\Accounts\GPBMetadata\Accounts::initOnce();
        parent::__construct($data);
    }

    /**
     * New Account title
     *
     * Generated from protobuf field <code>string title = 1 [json_name = "title"];</code>
     * @return string
     */
    public function getTitle()
    {
        return $this->title;
    }

    /**
     * New Account title
     *
     * Generated from protobuf field <code>string title = 1 [json_name = "title"];</code>
     * @param string $var
     * @return $this
     */
    public function setTitle($var)
    {
        GPBUtil::checkString($var, True);
        $this->title = $var;

        return $this;
    }

    /**
     * Credentials for New Account
     *
     * Generated from protobuf field <code>.nocloud.registry.accounts.Credentials auth = 2 [json_name = "auth"];</code>
     * @return \Nocloud\Registry\Accounts\Credentials|null
     */
    public function getAuth()
    {
        return $this->auth;
    }

    public function hasAuth()
    {
        return isset($this->auth);
    }

    public function clearAuth()
    {
        unset($this->auth);
    }

    /**
     * Credentials for New Account
     *
     * Generated from protobuf field <code>.nocloud.registry.accounts.Credentials auth = 2 [json_name = "auth"];</code>
     * @param \Nocloud\Registry\Accounts\Credentials $var
     * @return $this
     */
    public function setAuth($var)
    {
        GPBUtil::checkMessage($var, \Nocloud\Registry\Accounts\Credentials::class);
        $this->auth = $var;

        return $this;
    }

    /**
     * Namespace to put Account under
     *
     * Generated from protobuf field <code>string namespace = 3 [json_name = "namespace"];</code>
     * @return string
     */
    public function getNamespace()
    {
        return $this->namespace;
    }

    /**
     * Namespace to put Account under
     *
     * Generated from protobuf field <code>string namespace = 3 [json_name = "namespace"];</code>
     * @param string $var
     * @return $this
     */
    public function setNamespace($var)
    {
        GPBUtil::checkString($var, True);
        $this->namespace = $var;

        return $this;
    }

    /**
     * Account access level to parent namespace
     *
     * Generated from protobuf field <code>optional int32 access = 4 [json_name = "access"];</code>
     * @return int
     */
    public function getAccess()
    {
        return isset($this->access) ? $this->access : 0;
    }

    public function hasAccess()
    {
        return isset($this->access);
    }

    public function clearAccess()
    {
        unset($this->access);
    }

    /**
     * Account access level to parent namespace
     *
     * Generated from protobuf field <code>optional int32 access = 4 [json_name = "access"];</code>
     * @param int $var
     * @return $this
     */
    public function setAccess($var)
    {
        GPBUtil::checkInt32($var);
        $this->access = $var;

        return $this;
    }

}

