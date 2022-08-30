package com.github.exbotanical.resource.controllers.advice;

import com.github.exbotanical.resource.errors.GouacheException;
import com.github.exbotanical.resource.meta.Constants;
import com.github.exbotanical.resource.models.GouacheResponse;
import org.springframework.beans.ConversionNotSupportedException;
import org.springframework.beans.TypeMismatchException;
import org.springframework.core.annotation.AnnotationUtils;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.http.converter.HttpMessageNotReadableException;
import org.springframework.http.converter.HttpMessageNotWritableException;
import org.springframework.validation.BindException;
import org.springframework.web.HttpMediaTypeNotAcceptableException;
import org.springframework.web.HttpMediaTypeNotSupportedException;
import org.springframework.web.HttpRequestMethodNotSupportedException;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.MissingPathVariableException;
import org.springframework.web.bind.MissingServletRequestParameterException;
import org.springframework.web.bind.ServletRequestBindingException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.client.HttpServerErrorException;
import org.springframework.web.context.request.ServletWebRequest;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.context.request.async.AsyncRequestTimeoutException;
import org.springframework.web.multipart.support.MissingServletRequestPartException;
import org.springframework.web.servlet.NoHandlerFoundException;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.util.Objects;

/**
 * Global exception handler. Normalizes all exceptions to a GouacheResponse.
 */
@EnableWebMvc
@ControllerAdvice
@ResponseStatus
public class GouacheExceptionHandler extends ResponseEntityExceptionHandler {

  /**
   * An application-wide exception handler for GouacheException exceptions and subclasses thereof.
   *
   * @param ex  The exception.
   * @param req The current request.
   *
   * @return A normalized GouacheResponse.
   */
  @ExceptionHandler({ GouacheException.class })
  public ResponseEntity<GouacheResponse> gouacheExceptionHandler(GouacheException ex,
                                                                 WebRequest req) {
    // Derive the message data and build a GouacheResponse object.
    GouacheResponse ret = GouacheResponse.builder()
      .friendly(ex.getFriendly())
      .internal(ex.getInternal())
      .build();

    HttpStatus status;

    ResponseStatus responseStatusAnnotation = ex.getClass().getAnnotation(ResponseStatus.class);
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

  /**
   * An application-wide exception handler for uncaught Exception exceptions.
   *
   * @param ex  The exception.
   * @param req The current request.
   *
   * @return A normalized GouacheResponse.
   *
   * @throws Exception Forwards the Exception to Spring internals.
   */
  @ExceptionHandler({ Exception.class })
  public ResponseEntity<GouacheResponse> defaultExceptionHandler(
    Exception ex,
    WebRequest req) throws Exception {

    // If the exception is annotated with @ResponseStatus rethrow it and allow the framework to handle it.
    // @todo test
    if (AnnotationUtils.findAnnotation(ex.getClass(), ResponseStatus.class) != null) {
      throw ex;
    }

    if (ex instanceof HttpServerErrorException.InternalServerError) {
      return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(GouacheResponse.builder()
        .friendly(Constants.E_INTERNAL_SERVER_ERROR)
        .internal(ex.getMessage())
        .build());
    }

    // Build the fallback GouacheResponse object.
    GouacheResponse ret = GouacheResponse.builder()
      .friendly(Constants.E_GENERIC)
      .internal(ex.getMessage())
      .build();

    return ResponseEntity.status(resolveStatus(ex)).body(ret);
  }

  /**
   * Handle 404 NoHandlerFoundException exceptions.
   *
   * @param ex      The exception.
   * @param headers The headers to be written to the response.
   * @param status  The selected response status.
   * @param req     The current request.
   *
   * @return A normalized GouacheResponse.
   */
  @Override
  protected ResponseEntity<Object> handleNoHandlerFoundException(NoHandlerFoundException ex, HttpHeaders headers, HttpStatus status, WebRequest req) {
    String reqPath = ((ServletWebRequest) req).getRequest().getRequestURI().toString();
    System.out.println(reqPath);
    // Build the fallback GouacheResponse object.
    GouacheResponse ret = GouacheResponse.builder()
      .friendly(String.format(Constants.E_ROUTE_NOT_FOUND_FMT, reqPath))
      .internal(ex.getMessage())
      .build();

    return ResponseEntity.status(HttpStatus.NOT_FOUND).body(ret);
  }

  /**
   * Handle 405 HttpRequestMethodNotSupportedException exceptions.
   *
   * @param ex      The exception.
   * @param headers The headers to be written to the response.
   * @param status  The selected response status.
   * @param req     The current request.
   *
   * @return A normalized GouacheResponse.
   */
  @Override
  protected ResponseEntity<Object> handleHttpRequestMethodNotSupported(HttpRequestMethodNotSupportedException ex, HttpHeaders headers, HttpStatus status, WebRequest req) {
    String reqPath = ((ServletWebRequest) req).getRequest().getRequestURI().toString();
    String reqMethod = Objects.toString(((ServletWebRequest) req).getHttpMethod(), "UNKNOWN");

    // Build the fallback GouacheResponse object.
    GouacheResponse ret = GouacheResponse.builder()
      .friendly(String.format(Constants.E_METHOD_NOT_ALLOWED_FMT, reqMethod, reqPath))
      .internal(ex.getMessage())
      .build();

    return ResponseEntity.status(HttpStatus.METHOD_NOT_ALLOWED).body(ret);
  }

  /**
   * Handle 500 InternalServerError exceptions.
   *
   * @param ex      The exception.
   * @param body    The body for the response.
   * @param headers The headers for the response.
   * @param status  The response status.
   * @param req     The current request.
   *
   * @return A normalized GouacheResponse.
   */
  @Override
  protected ResponseEntity<Object> handleExceptionInternal(Exception ex, Object body, HttpHeaders headers, HttpStatus status, WebRequest req) {
    // Build the fallback GouacheResponse object.
    GouacheResponse ret = GouacheResponse.builder()
      .friendly(String.format(Constants.E_INTERNAL_SERVER_ERROR))
      .internal(ex.getMessage())
      .build();

    return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(ret);
  }

  /**
   * Resolve the correct HTTP status code given the exception type, falling back to 400 Bad Request.
   *
   * @param ex The provided exception from which to derive the HTTP status code.
   *
   * @return The resolved HTTP status code, with 400 Bad Request as the fallback value.
   */
  private HttpStatus resolveStatus(Exception ex) {

    if (ex instanceof HttpRequestMethodNotSupportedException) {
      return HttpStatus.METHOD_NOT_ALLOWED;
    } else if (ex instanceof HttpMediaTypeNotSupportedException) {
      return HttpStatus.UNSUPPORTED_MEDIA_TYPE;
    } else if (ex instanceof HttpMediaTypeNotAcceptableException) {
      return HttpStatus.NOT_ACCEPTABLE;
    } else if (ex instanceof MissingPathVariableException) {
      return HttpStatus.INTERNAL_SERVER_ERROR;
    } else if (ex instanceof MissingServletRequestParameterException) {
      return HttpStatus.BAD_REQUEST;
    } else if (ex instanceof ServletRequestBindingException) {
      return HttpStatus.BAD_REQUEST;
      // Check InternalServerError despite handling it explicitly in case we somehow miss it - at least we'll have the correct status code.
    } else if (ex instanceof ConversionNotSupportedException || ex instanceof HttpServerErrorException.InternalServerError) {
      return HttpStatus.INTERNAL_SERVER_ERROR;
    } else if (ex instanceof TypeMismatchException) {
      return HttpStatus.BAD_REQUEST;
    } else if (ex instanceof HttpMessageNotReadableException) {
      return HttpStatus.BAD_REQUEST;
    } else if (ex instanceof HttpMessageNotWritableException) {
      return HttpStatus.INTERNAL_SERVER_ERROR;
    } else if (ex instanceof MethodArgumentNotValidException) {
      return HttpStatus.BAD_REQUEST;
    } else if (ex instanceof MissingServletRequestPartException) {
      return HttpStatus.BAD_REQUEST;
    } else if (ex instanceof BindException) {
      return HttpStatus.BAD_REQUEST;
    } else if (ex instanceof NoHandlerFoundException) {
      return HttpStatus.NOT_FOUND;
    } else if (ex instanceof AsyncRequestTimeoutException) {
      return HttpStatus.SERVICE_UNAVAILABLE;
    }

    return HttpStatus.BAD_REQUEST;
  }
}
