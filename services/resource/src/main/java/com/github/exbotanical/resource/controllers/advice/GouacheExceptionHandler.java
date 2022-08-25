package com.github.exbotanical.resource.controllers.advice;

import com.github.exbotanical.resource.errors.GouacheException;
import com.github.exbotanical.resource.models.GouacheResponse;
import org.springframework.core.annotation.AnnotationUtils;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

/**
 * Global exception handler. Normalizes all exceptions to a GouacheResponse.
 *
 * @todo Test with other exception types e.g. HttpRequestMethodNotSupportedException.
 */
@EnableWebMvc
@ControllerAdvice
@ResponseStatus
public class GouacheExceptionHandler extends ResponseEntityExceptionHandler {

  @ExceptionHandler({GouacheException.class})
  public ResponseEntity<GouacheResponse> gouacheExceptionHandler(GouacheException e,
      WebRequest req) {
    System.out.println("XX: " + e);
    // Derive the message data and build a GouacheResponse object.
    GouacheResponse ret = GouacheResponse.builder()
        .friendly(e.getFriendly())
        .internal(e.getInternal())
        .build();

    HttpStatus status;

    ResponseStatus responseStatusAnnotation = e.getClass().getAnnotation(ResponseStatus.class);
    if (responseStatusAnnotation != null) {
      status = responseStatusAnnotation.value();
    } else {
      // Fallback to a 400.
      status = HttpStatus.BAD_REQUEST;
    }

    return ResponseEntity
        .status(status)
        .body(ret);
  }

  @ExceptionHandler({Exception.class})
  public ResponseEntity<GouacheResponse> defaultExceptionHandler(
      Exception e,
      WebRequest req) throws Exception {
    System.out.println("XX: " + e);

    // If the exception is annotated with @ResponseStatus rethrow it and let
    // the framework handle it - like the OrderNotFoundException example
    // at the start of this post.
    // @todo test
    if (AnnotationUtils.findAnnotation(e.getClass(), ResponseStatus.class) != null) {
      throw e;
    }

    // Build the fallback GouacheResponse object.
    GouacheResponse ret = GouacheResponse.builder()
        .friendly("An unknown exception occurred. @todo const")
        .internal(e.getMessage())
        .build();

    return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(ret);
  }
}
