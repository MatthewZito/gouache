package com.github.exbotanical.resource.config;

import com.github.exbotanical.resource.controllers.interceptor.AuthInterceptor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.http.converter.json.MappingJackson2HttpMessageConverter;
import org.springframework.web.servlet.config.annotation.CorsRegistry;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

import java.util.List;

@Configuration
@EnableWebMvc
public class WebConfig implements WebMvcConfigurer {

  private static final Long MAX_AGE = 3600L;

  @Autowired
  AuthInterceptor authInterceptor;

  @Override
  public void addCorsMappings(CorsRegistry registry) {
    registry.addMapping("/**")
        .allowedHeaders(
            HttpHeaders.AUTHORIZATION,
            HttpHeaders.CONTENT_TYPE,
            HttpHeaders.ACCEPT)
        .allowedMethods(
            HttpMethod.GET.name(),
            HttpMethod.POST.name(),
            HttpMethod.PUT.name(),
            HttpMethod.DELETE.name())
        .maxAge(MAX_AGE)
        .allowedOrigins("http://localhost:3000")
        .allowCredentials(true);
  }

  @Override
  public void addInterceptors(InterceptorRegistry registry) {
    registry.addInterceptor(authInterceptor).addPathPatterns("/**");
  }

  @Override
  public void configureMessageConverters(List<HttpMessageConverter<?>> converters) {
    WebMvcConfigurer.super.configureMessageConverters(converters);
    // Instruct Spring to use the following converter in lieu of StringHttpMessageConverter; the
    // latter cannot convert String return types
    // as they will have been modified to a Response object by the time they're processed in the
    // underlying ByteArrayHttpMessageConverter.
    // @see https://stackoverflow.com/a/65015720/15159240
    converters.add(0, new MappingJackson2HttpMessageConverter());
  }
}
