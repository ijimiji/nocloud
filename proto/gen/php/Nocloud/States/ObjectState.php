<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/states/proto/states.proto

namespace Nocloud\States;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>nocloud.states.ObjectState</code>
 */
class ObjectState extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string uuid = 1 [json_name = "uuid"];</code>
     */
    protected $uuid = '';
    /**
     * Generated from protobuf field <code>.nocloud.states.State state = 2 [json_name = "state"];</code>
     */
    protected $state = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $uuid
     *     @type \Nocloud\States\State $state
     * }
     */
    public function __construct($data = NULL) {
        \Nocloud\States\GPBMetadata\States::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string uuid = 1 [json_name = "uuid"];</code>
     * @return string
     */
    public function getUuid()
    {
        return $this->uuid;
    }

    /**
     * Generated from protobuf field <code>string uuid = 1 [json_name = "uuid"];</code>
     * @param string $var
     * @return $this
     */
    public function setUuid($var)
    {
        GPBUtil::checkString($var, True);
        $this->uuid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.nocloud.states.State state = 2 [json_name = "state"];</code>
     * @return \Nocloud\States\State|null
     */
    public function getState()
    {
        return $this->state;
    }

    public function hasState()
    {
        return isset($this->state);
    }

    public function clearState()
    {
        unset($this->state);
    }

    /**
     * Generated from protobuf field <code>.nocloud.states.State state = 2 [json_name = "state"];</code>
     * @param \Nocloud\States\State $var
     * @return $this
     */
    public function setState($var)
    {
        GPBUtil::checkMessage($var, \Nocloud\States\State::class);
        $this->state = $var;

        return $this;
    }

}

