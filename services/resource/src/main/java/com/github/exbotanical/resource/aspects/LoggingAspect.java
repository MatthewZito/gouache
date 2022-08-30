package com.github.exbotanical.resource.aspects;

import com.amazonaws.util.json.Jackson;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.exbotanical.resource.models.ReportName;
import com.github.exbotanical.resource.models.logs.ErrorLog;
import com.github.exbotanical.resource.models.logs.RequestLog;
import com.github.exbotanical.resource.services.QueueSenderService;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.AfterThrowing;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.CodeSignature;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

@Aspect
@Component
public class LoggingAspect {
  @Autowired
  private QueueSenderService queueSenderService;

  @Autowired
  private ObjectMapper mapper;

  private final Logger logger = LoggerFactory.getLogger(this.getClass());

  @Pointcut("within(@org.springframework.web.bind.annotation.RestController *)")
  private void restControllerPointcut() {
    // no-op
  }

  @AfterThrowing(pointcut = "restControllerPointcut()", throwing = "ex")
  private void logException(JoinPoint joinPoint, Throwable ex) {
    final HttpServletRequest req =
      ((ServletRequestAttributes) RequestContextHolder.currentRequestAttributes()).getRequest();

    final Map<String, Object> parameters = deriveRequestParams(joinPoint);

    final String errorMessage = String.format(
      "Exception at %s.%s()",
      joinPoint.getSignature().getDeclaringTypeName(),
      joinPoint.getSignature().getName());

    final String cause = ex.getCause() != null ? ex.getCause().toString() : "NULL";

    ErrorLog el = ErrorLog
      .builder()
      .path(req.getRequestURI())
      .method(req.getMethod())
      .parameters(parameters)
      .message(errorMessage)
      .cause(cause)
      .build();

    String data = Jackson.toJsonString(el);

    logger.info(data);

    queueSenderService.sendMessage(data, ReportName.HTTP_HANDLER_EX);
  }

  @Before("restControllerPointcut()")
  private void logInboundRequest(JoinPoint joinPoint) {
    final HttpServletRequest req =
      ((ServletRequestAttributes) RequestContextHolder.currentRequestAttributes()).getRequest();

    final Map<String, Object> parameters = deriveRequestParams(joinPoint);

    RequestLog rl = RequestLog
      .builder()
      .path(req.getRequestURI())
      .method(req.getMethod())
      .parameters(parameters)
      .build();

    String data = Jackson.toJsonString(rl);

    logger.info(data);

    queueSenderService.sendMessage(data, ReportName.HTTP_REQUEST_RECV);
  }

  private Map<String, Object> deriveRequestParams(JoinPoint joinPoint) {
    HashMap<String, Object> map = new HashMap<>();
    CodeSignature signature = (CodeSignature) joinPoint.getSignature();

    String[] parameterNames = signature.getParameterNames();
    for (int i = 0; i < parameterNames.length; i++) {
      map.put(parameterNames[i], joinPoint.getArgs()[i]);
    }

    return map;
  }
}
