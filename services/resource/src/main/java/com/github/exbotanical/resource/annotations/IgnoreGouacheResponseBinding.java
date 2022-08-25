package com.github.exbotanical.resource.annotations;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * IgnoreGouacheResponseBinding instructs the system to bypass the normalization of the given
 * endpoint's return values into a GouacheResponse, instead returning an HTTP response containing
 * the as-is return value.
 */
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
public @interface IgnoreGouacheResponseBinding {

}
