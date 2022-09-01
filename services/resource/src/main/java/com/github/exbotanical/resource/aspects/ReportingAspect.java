package com.github.exbotanical.resource.aspects;

import com.amazonaws.util.json.Jackson;
import com.github.exbotanical.resource.models.reporting.GouacheReport;
import com.github.exbotanical.resource.models.reporting.GouacheReportName;
import com.github.exbotanical.resource.models.reporting.RequestReport;
import com.github.exbotanical.resource.services.QueueSenderService;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.AfterThrowing;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.CodeSignature;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

/**
 * An aspect utilized for sending reports pertinent to system events.
 */
@Aspect
@Component
public class ReportingAspect {
  @Autowired
  private QueueSenderService queueSenderService;

  @Pointcut("within(@org.springframework.web.bind.annotation.RestController *)")
  private void restControllerPointcut() {
    // no-op
  }

  @AfterThrowing(pointcut = "restControllerPointcut()", throwing = "ex")
  private void logRestControllerException(JoinPoint joinPoint, Throwable ex) {
    final HttpServletRequest req =
        ((ServletRequestAttributes) RequestContextHolder.currentRequestAttributes()).getRequest();

    final Map<String, Object> parameters = deriveRequestParams(joinPoint);

    final String errorMessage = String.format(
        "Exception at %s.%s()",
        joinPoint.getSignature().getDeclaringTypeName(),
        joinPoint.getSignature().getName());

    final String cause = ex.getCause() != null ? ex.getCause().toString() : "NULL";

    RequestReport rl = RequestReport
        .builder()
        .path(req.getRequestURI())
        .method(req.getMethod())
        .parameters(parameters)
        .error(String.format("%s %s", errorMessage, cause))
        .build();

    queueSenderService.sendMessage(rl, GouacheReportName.HTTP_HANDLER_EX);
  }

  @Before("restControllerPointcut()")
  private void logInboundRequest(JoinPoint joinPoint) {
    final HttpServletRequest req =
        ((ServletRequestAttributes) RequestContextHolder.currentRequestAttributes()).getRequest();

    final Map<String, Object> parameters = deriveRequestParams(joinPoint);

    RequestReport rl = RequestReport
        .builder()
        .path(req.getRequestURI())
        .method(req.getMethod())
        .parameters(parameters)
        .build();


    queueSenderService.sendMessage(rl, GouacheReportName.HTTP_REQUEST_RECV);
  }

  private Map<String, Object> deriveRequestParams(JoinPoint joinPoint) {
    HashMap<String, Object> map = new HashMap<>();
    CodeSignature signature = (CodeSignature) joinPoint.getSignature();

    String[] parameterNames = signature.getParameterNames();

    for (int i = 0; i < parameterNames.length; i++) {
      map.put(parameterNames[i], joinPoint.getArgs()[i].toString());
    }

    return map;
  }
}
