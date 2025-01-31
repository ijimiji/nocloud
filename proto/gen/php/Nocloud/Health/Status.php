<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pkg/health/proto/health.proto

namespace Nocloud\Health;

use UnexpectedValueException;

/**
 * Protobuf type <code>nocloud.health.Status</code>
 */
class Status
{
    /**
     * Generated from protobuf enum <code>UNKNOWN = 0;</code>
     */
    const UNKNOWN = 0;
    /**
     * Service is up and running
     *
     * Generated from protobuf enum <code>SERVING = 1;</code>
     */
    const SERVING = 1;
    /**
     * Service is offline(down)
     *
     * Generated from protobuf enum <code>OFFLINE = 2;</code>
     */
    const OFFLINE = 2;
    /**
     * Routine is running
     *
     * Generated from protobuf enum <code>RUNNING = 3;</code>
     */
    const RUNNING = 3;
    /**
     * Routine is stopped
     *
     * Generated from protobuf enum <code>STOPPED = 4;</code>
     */
    const STOPPED = 4;
    /**
     * Internal error while making status
     *
     * Generated from protobuf enum <code>INTERNAL = 5;</code>
     */
    const INTERNAL = 5;
    /**
     * Check has errors
     *
     * Generated from protobuf enum <code>HASERRS = 6;</code>
     */
    const HASERRS = 6;
    /**
     * Service has no Routines
     *
     * Generated from protobuf enum <code>NOEXIST = 7;</code>
     */
    const NOEXIST = 7;

    private static $valueToName = [
        self::UNKNOWN => 'UNKNOWN',
        self::SERVING => 'SERVING',
        self::OFFLINE => 'OFFLINE',
        self::RUNNING => 'RUNNING',
        self::STOPPED => 'STOPPED',
        self::INTERNAL => 'INTERNAL',
        self::HASERRS => 'HASERRS',
        self::NOEXIST => 'NOEXIST',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}

