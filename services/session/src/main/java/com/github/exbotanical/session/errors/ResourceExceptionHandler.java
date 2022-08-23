package com.github.exbotanical.session.errors;

import com.github.exbotanical.session.entities.ErrorMessage;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@ControllerAdvice
@ResponseStatus
public class ResourceExceptionHandler extends ResponseEntityExceptionHandler {

  @ExceptionHandler(ResourceNotFoundException.class)
  public ResponseEntity<ErrorMessage> resourceNotFoundException(ResourceNotFoundException e, WebRequest req) {
    HttpStatus status = HttpStatus.NOT_FOUND;
    ErrorMessage message = new ErrorMessage(status, e.getMessage());

    return ResponseEntity
            .status(status)
            .body(message);
  }
}

/*

@Composed(Set.class)
public class InstrumentedSet {
  ...
}

Also, note that the special documentation required for inheritance clutters up
normal documentation, which is designed for programmers who create instances
of your class and invoke methods on them. **As of this writing, there is little in the
way of tools to separate ordinary API documentation from information of interest
only to programmers implementing subclasses.**
*/
